package emulator

import (
	"context"
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
