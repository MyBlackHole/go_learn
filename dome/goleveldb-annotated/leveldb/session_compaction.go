// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package leveldb

import (
	"sync/atomic"

	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/memdb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func (s *session) pickMemdbLevel(umin, umax []byte, maxLevel int) int {
	v := s.version()
	defer v.release()
	return v.pickMemdbLevel(umin, umax, maxLevel)
}

func (s *session) flushMemdb(rec *sessionRecord, mdb *memdb.DB, maxLevel int) (int, error) {
	// Create sorted table.
	iter := mdb.NewIterator(nil)
	defer iter.Release()
	// 使用iterator，将kv/meta data写入sst文件，并创建tFile，记录fd/size/imin/imax/seekLeft
	t, n, err := s.tops.createFrom(iter)
	if err != nil {
		return 0, err
	}

	// Pick level other than zero can cause compaction issue with large
	// bulk insert and delete on strictly incrementing key-space. The
	// problem is that the small deletion markers trapped at lower level,
	// while key/value entries keep growing at higher level. Since the
	// key-space is strictly incrementing it will not overlaps with
	// higher level, thus maximum possible level is always picked, while
	// overlapping deletion marker pushed into lower level.
	// See: https://github.com/syndtr/goleveldb/issues/127.
	// 选择一个level，将frozen table直接推到该level
	// 现在不支持往高level推，原因见注释
	flushLevel := s.pickMemdbLevel(t.imin.ukey(), t.imax.ukey(), maxLevel)
	// 在sessionRecord中记录在flushLevel中add新文件
	rec.addTableFile(flushLevel, t)

	s.logf("memdb@flush created L%d@%d N·%d S·%s %q:%q", flushLevel, t.fd.Num, n, shortenb(int(t.size)), t.imin, t.imax)
	return flushLevel, nil
}

// Pick a compaction based on current state; need external synchronization.
// 返回一个compaction结构，这个结构表示compaction状态
// 通过newCompaction，获取到level/level+1层参与compaction的文件
func (s *session) pickCompaction() *compaction {
	v := s.version()

	var sourceLevel int
	var t0 tFiles

	// 发起compaction的第一步就是查找参与compaction的source层文件
	// 由于触发compaction有两个条件：(1) 写 (2) 读
	// 下面的两个分支分别处理上述两个条件
	if v.cScore >= 1 { // cScore是在生成新的version(也就是compaction)时计算的
		// cLevel是此次compaction的source level，与最大的cScore一致
		sourceLevel = v.cLevel
		// TODO
		cptr := s.getCompPtr(sourceLevel)
		tables := v.levels[sourceLevel]
		// 从cLevel挑选出需要compaction的sst文件
		// cptr应该是上一次compaction的key
		for _, t := range tables {
			// 将第一个key大于上一次compaction key的文件加入到source的compaction队列中
			if cptr == nil || s.icmp.Compare(t.imax, cptr) > 0 {
				t0 = append(t0, t)
				break
			}
		}

		// 如果没有符合条件的，则将第一个文件加入到队列
		if len(t0) == 0 {
			t0 = append(t0, tables[0])
		}
	} else {
		// 非cSore>1的场景，则是seekLeft == 0触发
		// 这种场景下，参与compaction的source层文件就是seekLeft == 0的文件，记录在cSeek中
		if p := atomic.LoadPointer(&v.cSeek); p != nil {
			ts := (*tSet)(p)
			sourceLevel = ts.level
			t0 = append(t0, ts.table)
		} else {
			v.release()
			return nil
		}
	}

	return newCompaction(s, v, sourceLevel, t0)
}

// Create compaction from given level and range; need external synchronization.
func (s *session) getCompactionRange(sourceLevel int, umin, umax []byte, noLimit bool) *compaction {
	v := s.version()

	if sourceLevel >= len(v.levels) {
		v.release()
		return nil
	}

	t0 := v.levels[sourceLevel].getOverlaps(nil, s.icmp, umin, umax, sourceLevel == 0)
	if len(t0) == 0 {
		v.release()
		return nil
	}

	// Avoid compacting too much in one shot in case the range is large.
	// But we cannot do this for level-0 since level-0 files can overlap
	// and we must not pick one file and drop another older file if the
	// two files overlap.
	if !noLimit && sourceLevel > 0 {
		limit := int64(v.s.o.GetCompactionSourceLimit(sourceLevel))
		total := int64(0)
		for i, t := range t0 {
			total += t.size
			if total >= limit {
				s.logf("table@compaction limiting F·%d -> F·%d", len(t0), i+1)
				t0 = t0[:i+1]
				break
			}
		}
	}

	return newCompaction(s, v, sourceLevel, t0)
}

func newCompaction(s *session, v *version, sourceLevel int, t0 tFiles) *compaction {
	c := &compaction{
		s:           s,
		v:           v,
		sourceLevel: sourceLevel,
		levels:      [2]tFiles{t0, nil},
		// TODO
		maxGPOverlaps: int64(s.o.GetCompactionGPOverlaps(sourceLevel)),
		tPtrs:         make([]int, len(v.levels)),
	}
	// 根据source level，来expand需要compaction的文件，包含source level和source+1
	c.expand()
	c.save()
	return c
}

// compaction represent a compaction state.
type compaction struct {
	s *session
	v *version

	sourceLevel int
	// 2维数组，第一维是进行compaction的两个level，第二维是每个level参与compaction的文件
	levels        [2]tFiles
	maxGPOverlaps int64

	gp                tFiles
	gpi               int
	seenKey           bool
	gpOverlappedBytes int64
	imin, imax        internalKey
	tPtrs             []int
	released          bool

	snapGPI               int
	snapSeenKey           bool
	snapGPOverlappedBytes int64
	snapTPtrs             []int
}

func (c *compaction) save() {
	c.snapGPI = c.gpi
	c.snapSeenKey = c.seenKey
	c.snapGPOverlappedBytes = c.gpOverlappedBytes
	c.snapTPtrs = append(c.snapTPtrs[:0], c.tPtrs...)
}

func (c *compaction) restore() {
	c.gpi = c.snapGPI
	c.seenKey = c.snapSeenKey
	c.gpOverlappedBytes = c.snapGPOverlappedBytes
	c.tPtrs = append(c.tPtrs[:0], c.snapTPtrs...)
}

func (c *compaction) release() {
	if !c.released {
		c.released = true
		c.v.release()
	}
}

// Expand compacted tables; need external synchronization.
// 基于source level的待compaction文件进行expand，查找参加compaction的文件
func (c *compaction) expand() {
	// 参与compaction的文件数的上限
	limit := int64(c.s.o.GetCompactionExpandLimit(c.sourceLevel))
	// source level的所有文件
	vt0 := c.v.levels[c.sourceLevel]
	vt1 := tFiles{}
	if level := c.sourceLevel + 1; level < len(c.v.levels) {
		// sourcelevel + 1的全部文件
		vt1 = c.v.levels[level]
	}

	// c.levels是source，source+1层参与compaction的文件
	t0, t1 := c.levels[0], c.levels[1]
	// compaction源文件的[imin, imax]
	imin, imax := t0.getRange(c.s.icmp)
	// We expand t0 here just incase ukey hop across tables.
	// 扩展source level参与compaction的文件,原则很简单，就是查找与source文件的[imin, imax]有overlap的文件
	t0 = vt0.getOverlaps(t0, c.s.icmp, imin.ukey(), imax.ukey(), c.sourceLevel == 0)
	if len(t0) != len(c.levels[0]) {
		// 更新imin,imax
		imin, imax = t0.getRange(c.s.icmp)
	}
	// 使用imin,imax去level+1层查找overlap的文件
	t1 = vt1.getOverlaps(t1, c.s.icmp, imin.ukey(), imax.ukey(), false)
	// Get entire range covered by compaction.
	// level N/level N+1的[min,max]并集
	amin, amax := append(t0, t1...).getRange(c.s.icmp)

	// See if we can grow the number of inputs in "sourceLevel" without
	// changing the number of "sourceLevel+1" files we pick up.
	// 在不修改level+1的前提下，尽量扩大level层参与compaction的文件数，这么处理的原因是：
	/*
                        a        b         c       d
      source level   <-----><--------><--------><------>
                       A    B     C    D     E     F
      level +1       <---><---><----><----><----><---->
      上面的例子中，比如发起compaction的源文件是b，查找compaction文件的流程如下：
      1. 在source层进行expand，由于不是level 0，不存在overlap。结果还是b
      2. 使用b文件的key范围[min,max]到level+1层进行比较，得到overlap的文件B、C、D
      3. 此时b、B、C、D文件为compaction的输入文件

      但是还有另外一种情况：
                        a      b         c       d
      source level   <---><--------><--------><------>
                         A      B     C    D     E
      level +1       <-------><----><----><----><---->
      这种情况下的compaction流程为：
      1. 在source层进行expand，由于不是level 0，不存在overlap。结果还是b
      2. 使用b文件的key范围[min,max]到level+1层进行比较，得到overlap的文件A、B
      3. 这时候compaction的输入文件是否就是b、A、B了呢？观察一下可以发现，source level的a文件
         key的范围完全在A+B范围内。这时候把a进入到compaction的目标文件，完全不会影响level+1的key的范围
      4. 因此，原则上：在不修改level+1的前提下，尽量扩大level层参与compaction的文件数

      在第一个例子中，能不能把与B、C、D有overlap的a加进来呢？答案是否定的。因为leveldb中，除了level 0
      以外，其他level不能存在key的overlap。如果把a加入到compaction的源文件，生成的结果文件会与A存在overlap。
	*/
	if len(t1) > 0 {
		exp0 := vt0.getOverlaps(nil, c.s.icmp, amin.ukey(), amax.ukey(), c.sourceLevel == 0)
		if len(exp0) > len(t0) && t1.size()+exp0.size() < limit {
			xmin, xmax := exp0.getRange(c.s.icmp)
			exp1 := vt1.getOverlaps(nil, c.s.icmp, xmin.ukey(), xmax.ukey(), false)
			if len(exp1) == len(t1) {
				c.s.logf("table@compaction expanding L%d+L%d (F·%d S·%s)+(F·%d S·%s) -> (F·%d S·%s)+(F·%d S·%s)",
					c.sourceLevel, c.sourceLevel+1, len(t0), shortenb(int(t0.size())), len(t1), shortenb(int(t1.size())),
					len(exp0), shortenb(int(exp0.size())), len(exp1), shortenb(int(exp1.size())))
				imin, imax = xmin, xmax
				t0, t1 = exp0, exp1
				amin, amax = append(t0, t1...).getRange(c.s.icmp)
			}
		}
	}

	// Compute the set of grandparent files that overlap this compaction
	// (parent == sourceLevel+1; grandparent == sourceLevel+2)
	// level+2层与此次compaction有重叠的文件
	if level := c.sourceLevel + 2; level < len(c.v.levels) {
		c.gp = c.v.levels[level].getOverlaps(c.gp, c.s.icmp, amin.ukey(), amax.ukey(), false)
	}

	c.levels[0], c.levels[1] = t0, t1
	c.imin, c.imax = imin, imax
}

// Check whether compaction is trivial.
func (c *compaction) trivial() bool {
	return len(c.levels[0]) == 1 && len(c.levels[1]) == 0 && c.gp.size() <= c.maxGPOverlaps
}

func (c *compaction) baseLevelForKey(ukey []byte) bool {
	for level := c.sourceLevel + 2; level < len(c.v.levels); level++ {
		tables := c.v.levels[level]
		for c.tPtrs[level] < len(tables) {
			t := tables[c.tPtrs[level]]
			if c.s.icmp.uCompare(ukey, t.imax.ukey()) <= 0 {
				// We've advanced far enough.
				if c.s.icmp.uCompare(ukey, t.imin.ukey()) >= 0 {
					// Key falls in this file's range, so definitely not base level.
					return false
				}
				break
			}
			c.tPtrs[level]++
		}
	}
	return true
}

func (c *compaction) shouldStopBefore(ikey internalKey) bool {
	for ; c.gpi < len(c.gp); c.gpi++ {
		gp := c.gp[c.gpi]
		if c.s.icmp.Compare(ikey, gp.imax) <= 0 {
			break
		}
		if c.seenKey {
			c.gpOverlappedBytes += gp.size
		}
	}
	c.seenKey = true

	if c.gpOverlappedBytes > c.maxGPOverlaps {
		// Too much overlap for current output; start new output.
		c.gpOverlappedBytes = 0
		return true
	}
	return false
}

// Creates an iterator.
func (c *compaction) newIterator() iterator.Iterator {
	// Creates iterator slice.
	icap := len(c.levels)
	if c.sourceLevel == 0 {
		// Special case for level-0.
		icap = len(c.levels[0]) + 1
	}
	its := make([]iterator.Iterator, 0, icap)

	// Options.
	ro := &opt.ReadOptions{
		DontFillCache: true,
		Strict:        opt.StrictOverride,
	}
	strict := c.s.o.GetStrict(opt.StrictCompaction)
	if strict {
		ro.Strict |= opt.StrictReader
	}

	for i, tables := range c.levels {
		if len(tables) == 0 {
			continue
		}

		// Level-0 is not sorted and may overlaps each other.
		if c.sourceLevel+i == 0 {
			for _, t := range tables {
				its = append(its, c.s.tops.newIterator(t, nil, ro))
			}
		} else {
			it := iterator.NewIndexedIterator(tables.newIndexIterator(c.s.tops, c.s.icmp, nil, ro), strict)
			its = append(its, it)
		}
	}

	return iterator.NewMergedIterator(its, c.s.icmp, strict)
}
