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

    StorageClass string

    Type string

	Name string

	ModTime time.Time

	Size int64

	IsDir bool

	ETag string

	IsLatest bool
}

type ListObjectsInfo struct {
    // 截断状态(超出单次返回最大数量)
	IsTruncated bool

	NextMarker string

	Objects []ObjectInfo

	Prefixes []string
}

type listPathOptions struct {
	ID string

	Bucket string

	BaseDir string

	Prefix string

	FilterPrefix string

	Marker string

	Limit int

	InclDeleted bool

	Recursive bool

	Separator string

	Create bool

	IncludeDirectories bool

	Transient bool
}
