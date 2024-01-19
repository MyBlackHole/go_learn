package emulator

import (

	"time"
)

type BucketInfo struct {
	Name string

	Created time.Time
	Deleted time.Time
}
