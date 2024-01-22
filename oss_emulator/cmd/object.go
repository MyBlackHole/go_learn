package emulator

import (
	"context"
	"io"
)

type Objects struct {
	disk StorageAPI
}

func (o *Objects) MakeBucket(ctx context.Context, bucket string) error {
	disk := o.disk

	disk.MakeVol(ctx, bucket)

	return nil
}

func (o *Objects) GetBucketInfo(ctx context.Context, bucket string) (bucketInfo BucketInfo, err error) {
	disk := o.disk

	volInfo, err := disk.StatVol(ctx, bucket)
	if err != nil {
		return
	}
	return BucketInfo{Name: bucket, Created: volInfo.Created}, nil
}

func (o *Objects) PutObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error) {
	disk := o.disk

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
	disk := o.disk

	buckets = make([]BucketInfo, 0, 32)
	vols, err := disk.ListVols(ctx)
	if err != nil {
		return
	}

	for _, v := range vols {
		bi := BucketInfo{
			Name:    v.Name,
			Created: v.Created,
		}
		buckets = append(buckets, bi)
	}
	return
}
