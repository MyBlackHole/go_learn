package emulator

import (

	"time"
)

type BucketInfo struct {
	Name string

	Created time.Time
}


type ObjectInfo struct {
	Bucket string

	Name string

	ModTime time.Time

	Size int64

	IsDir bool

	ETag string

	IsLatest bool
}
