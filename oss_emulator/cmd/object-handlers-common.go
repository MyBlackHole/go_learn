package emulator

import (
	"net/http"
)

const (
	ETag         = "ETag"
	LastModified = "Last-Modified"
)

func setPutObjHeaders(w http.ResponseWriter, objInfo ObjectInfo, delete bool) {
	if objInfo.ETag != "" && !delete {
		w.Header()[ETag] = []string{`"` + objInfo.ETag + `"`}
	}
}
