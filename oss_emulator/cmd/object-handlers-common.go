package emulator

import (
	"net/http"
)

const (
	ETag = "ETag"
)

func setPutObjHeaders(w http.ResponseWriter, objInfo ObjectInfo, delete bool) {
	if objInfo.ETag != "" && !delete {
		w.Header()[ETag] = []string{`"` + objInfo.ETag + `"`}
	}
}
