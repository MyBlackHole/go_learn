package emulator

import (
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
		return
	}

	if exist {
		return errVolumeExists
	}

	err = globalMetaDb.Update(func(tx *nutsdb.Tx) error {
		err = tx.NewBucket(nutsdb.DataStructureBTree, bucket)
		return err
	})

	if err != nil {
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
		err = errVolumeNotFound
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
		err = errVolumeNotFound
		return
	}

	fi := newFileInfo(pathJoin(bucket, object))

	partName := "part.1"
	err = disk.CreateFile(ctx, fi.Volume, partName, size, r)
	if err != nil {
		return
	}
	fi.Size = size
	fi.ModTime = UTCNow()

	disk.WriteMetadata(ctx, bucket, object, fi)

	objInfo = fi.ToObjectInfo(bucket, object)
	return
}

func (o *Objects) GetObjectInfo(ctx context.Context, bucket, object string) (info ObjectInfo, err error) {
	disk := o.disk

	fi, err := disk.ReadMetadata(ctx, bucket, object)
	if err != nil {
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
