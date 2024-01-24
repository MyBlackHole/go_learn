package emulator

import (
	"context"
    "io"
)

type ObjectLayer interface {
	MakeBucket(ctx context.Context, bucket string) error
	GetBucketInfo(ctx context.Context, bucket string) (bucketInfo BucketInfo, err error)
    PutObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error)
    GetObject(ctx context.Context, bucket string, object string, w io.Writer, writeHeadCall func(objInfo ObjectInfo) error) (err error)
    AppendObject(ctx context.Context, bucket string, object string, size int64, r io.Reader) (objInfo ObjectInfo, err error)
    GetObjectInfo(ctx context.Context, bucket, object string) (objInfo ObjectInfo, err error)
	ListBuckets(ctx context.Context) (buckets []BucketInfo, err error)
	ListObjects(ctx context.Context, bucket, prefix, marker, delimiter string, maxKeys int) (result ListObjectsInfo, err error)
}
