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

func setObjectHeaders(w http.ResponseWriter, objInfo ObjectInfo) (err error) {
	setCommonHeaders(w)

	lastModified := objInfo.ModTime.UTC().Format(http.TimeFormat)
	w.Header().Set(LastModified, lastModified)

	if objInfo.ETag != "" {
		w.Header()[ETag] = []string{"\"" + objInfo.ETag + "\""}
	}

    // w.Header().Set(ContentLength, strconv.FormatInt(objInfo.Size, 10))
	w.Header().Set(NextAppendPosition, strconv.FormatInt(objInfo.Size, 10))

	return nil
}
