package emulator

import (
	"context"
)

type ObjectLayer interface {
	MakeBucket(ctx context.Context, bucket string) error
	GetBucketInfo(ctx context.Context, bucket string) (bucketInfo BucketInfo, err error)
}
