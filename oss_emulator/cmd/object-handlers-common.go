package emulator

import (
	"net/http"
	"strconv"
)

const (
	ETag               = "ETag"
	LastModified       = "Last-Modified"
	NextAppendPosition = "x-oss-next-append-position"
)

func setPutObjHeaders(w http.ResponseWriter, objInfo ObjectInfo, delete bool) {
	if objInfo.ETag != "" && !delete {
		w.Header()[ETag] = []string{`"` + objInfo.ETag + `"`}
	}

	w.Header().Set(NextAppendPosition, strconv.FormatInt(objInfo.Size, 10))
}
