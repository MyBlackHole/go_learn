package emulator

import (
	"context"
    "io"
)

type ObjectLayer interface {
	MakeBucket(ctx context.Context, bucket string) error
	GetBucketInfo(ctx context.Context, bucket string) (bucketInfo BucketInfo, err error)
    PutObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error)
}
