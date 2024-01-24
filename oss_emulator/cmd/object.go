package emulator

import (
	"bytes"
	"context"
	"github.com/nutsdb/nutsdb"
	"io"
)

type Objects struct {
	disk StorageAPI
}

func (o *Objects) MakeBucket(ctx context.Context, bucket string) (err error) {
	// disk := o.disk

	var exist bool

	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if err != nil {
		err = toObjectErr(err)
		return
	}

	if exist {
		return toObjectErr(errVolumeExists)
	}

	err = globalMetaDb.Update(func(tx *nutsdb.Tx) error {
		err = tx.NewBucket(nutsdb.DataStructureBTree, bucket)
		err = toObjectErr(err)
		return err
	})

	if err != nil {
		err = toObjectErr(err)
		return
	}

	// 预计方些桶信息 (不 put bucket 有问题)
	err = globalMetaDb.Update(
		func(tx *nutsdb.Tx) error {
			key := []byte(bucket)
			val := []byte(bucket)
			return tx.Put(bucket, key, val, 0)
		})

	if err != nil {
		err = toObjectErr(err)
		return
	}

	// // 创建桶
	// err := disk.MakeVol(ctx, bucket)
	return nil
}

func (o *Objects) GetBucketInfo(ctx context.Context, bucket string) (bucketInfo BucketInfo, err error) {
	// disk := o.disk
	var exist bool

	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if !exist {
		err = toObjectErr(errVolumeNotFound)
		return
	}

	// volInfo, err := disk.StatVol(ctx, bucket)
	// if err != nil {
	// 	return
	// }
	return BucketInfo{Name: bucket}, nil
}

func (o *Objects) PutObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error) {
	disk := o.disk

	var exist bool

	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if !exist {
		err = toObjectErr(errVolumeNotFound)
		return
	}

	fi := newFileInfo(pathJoin(bucket, object))

	partName := "part.1"
	written, err := disk.CreateFile(ctx, fi.Volume, partName, size, r)
	if err != nil {
		err = toObjectErr(err)
		return
	}
	fi.Size = written
	fi.ModTime = UTCNow()

	disk.WriteMetadata(ctx, bucket, object, fi)

	objInfo = fi.ToObjectInfo(bucket, object)
	return
}

func (o *Objects) AppendObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error) {
	disk := o.disk

	var exist bool

	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if !exist {
		err = toObjectErr(errVolumeNotFound)
		return
	}

	fi, err := disk.ReadMetadata(ctx, bucket, object)
	if err != nil {
		fi = newFileInfo(pathJoin(bucket, object))
	}

	partName := "part.1"
	written, err := disk.AppendFile(ctx, fi.Volume, partName, size, r)
	if err != nil {
		err = toObjectErr(err)
		return
	}
	fi.Size += written
	fi.ModTime = UTCNow()

	disk.WriteMetadata(ctx, bucket, object, fi)

	objInfo = fi.ToObjectInfo(bucket, object)
	return
}

func (o *Objects) GetObjectInfo(ctx context.Context, bucket, object string) (info ObjectInfo, err error) {
	disk := o.disk

	fi, err := disk.ReadMetadata(ctx, bucket, object)
	if err != nil {
		err = toObjectErr(err)
		return
	}

	info = fi.ToObjectInfo(bucket, object)

	return
}

func (o *Objects) ListBuckets(ctx context.Context) (buckets []BucketInfo, err error) {
	// disk := o.disk

	buckets = make([]BucketInfo, 0, 32)
	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		return tx.IterateBuckets(nutsdb.DataStructureBTree, "*", func(bucket string) bool {
			bi := BucketInfo{
				Name: bucket,
				// Created: v.Created,
			}
			buckets = append(buckets, bi)
			return true
		})
	})
	if err != nil {
		err = toObjectErr(err)
		return
	}

	// vols, err := disk.ListVols(ctx)
	// if err != nil {
	// 	return
	// }

	return
}

func maxKeysPlusOne(maxKeys int, addOne bool) int {
	if maxKeys < 0 || maxKeys > maxObjectList {
		maxKeys = maxObjectList
	}
	if addOne {
		maxKeys++
	}
	return maxKeys
}

func (o *Objects) ListObjects(ctx context.Context, bucket, prefix, marker, delimiter string, maxKeys int) (result ListObjectsInfo, err error) {
	var exist bool

	// 检查桶是否存在
	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if !exist {
		err = toObjectErr(errVolumeNotFound)
		return
	}

	var markerFound bool
	var existMarker bool
	var existPrefix bool
	var count int

	if len(prefix) > 0 {
		existPrefix = true
	}

	if len(marker) > 0 {
		existMarker = true
	} else {
		markerFound = true
	}

	var keys [][]byte
	var values [][]byte
	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		keys, values, err = tx.GetAll(bucket)
		return err
	})
	if err != nil {
		return
	}

	for index, key := range keys {
		// 存在下一次开始位置
		if existMarker && bytes.Compare([]byte(marker), key) <= 0 {
			markerFound = true
		}

		if !markerFound {
			// 未找到开始位置
			continue
		}

		// 找到开始位置了
		// 存在前缀检查
		// 检查前缀是否匹配
		if existPrefix && !bytes.HasPrefix(key, []byte(prefix)) {
			continue
		}

		// 跳过跟桶名一致对象
		if bytes.Equal([]byte(bucket), key) {
			continue
		}

		// 校验都通过了
		count += 1
		if count <= maxKeys {
			var fi FileInfo
			_, err = fi.UnmarshalMsg(values[index])
			if err != nil {
				return
			}

			result.Objects = append(result.Objects, fi.ToObjectInfo(bucket, string(key)))
			continue
		}
		// 超出 maxKeys 范围
		// 记录被截断了
		result.IsTruncated = true
		result.NextMarker = string(key)
	}

	return
}

func (o *Objects) GetObject(ctx context.Context, bucket string, object string, w io.Writer, writeHeadCall func(objInfo ObjectInfo) error) (err error) {
	disk := o.disk

	var exist bool

	err = globalMetaDb.View(func(tx *nutsdb.Tx) error {
		exist = tx.ExistBucket(nutsdb.DataStructureBTree, bucket)
		return nil
	})

	if !exist {
		err = toObjectErr(errVolumeNotFound)
		return
	}

	fi, err := disk.ReadMetadata(ctx, bucket, object)
	if err != nil {
		err = toObjectErr(err)
		return
	}

    objInfo := fi.ToObjectInfo(bucket, object)
    err = writeHeadCall(objInfo)
	if err != nil {
		return
	}

	partName := "part.1"
	err = disk.ReadFile(ctx, fi.Volume, partName, w)
	if err != nil {
		err = toObjectErr(err)
		return
	}

	return
}
