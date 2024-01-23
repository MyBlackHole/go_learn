package emulator

import (
	"time"
)

const metacacheStreamVersion = 2

type VolInfo struct {
	Name string

	Created time.Time
}


type FileInfo struct {
	Volume string `msg:"v,omitempty"`

	Name string `msg:"n,omitempty"`

	Size int64 `msg:"sz"`

	ModTime time.Time `msg:"mt"`
}

func newFileInfo(object string) (fi FileInfo) {
    fi.Volume = object
	return
}

type WalkDirOptions struct {
	Bucket string

	BaseDir string

	Recursive bool

	ReportNotFound bool

	FilterPrefix string

	ForwardTo string

	Limit int
}
