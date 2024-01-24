package emulator

import (
	"github.com/nutsdb/nutsdb"
)

const (
	// 1MB
	SegmentSize = 10 * 1024 * 1024
)

func InitMeta(metaDir string) {
	var err error
	globalMetaDb, err = nutsdb.Open(
		nutsdb.DefaultOptions,
		nutsdb.WithDir(metaDir),
		// nutsdb.WithSegmentSize(SegmentSize),
	)
	if err != nil {
		panic(err)
	}

}

func (fi FileInfo) ToObjectInfo(bucket, object string) ObjectInfo {
	objInfo := ObjectInfo{
		IsDir:        HasSuffix(object, SlashSeparator),
		Bucket:       bucket,
		Name:         object,
		Size:         fi.Size,
		ModTime:      fi.ModTime,
		StorageClass: "Standard",
		Type:         "Multipart",
		ETag:         "00000000000000000",
	}
	return objInfo
}
