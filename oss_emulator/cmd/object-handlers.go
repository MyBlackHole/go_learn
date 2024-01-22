package emulator

import (
	"io"
	"net/http"
	"strconv"

	"github.com/minio/mux"
)

func (api objectAPIHandlers) PutObjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutObject")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]
	object, err := unescapePath(vars["object"])
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}
	var (
		rd        io.Reader = r.Body
		putObject           = objectAPI.PutObject
		size                = r.ContentLength
	)

	if size == -1 {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrMissingContentLength), r.URL)
		return
	}

	objInfo, err := putObject(ctx, bucket, object, size, rd)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}
	setPutObjHeaders(w, objInfo, false)

	writeSuccessResponseHeadersOnly(w)
}


func (api objectAPIHandlers) HeadObjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "HeadObject")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponseHeadersOnly(w, errorCodes.ToAPIErr(ErrServerNotInitialized))
		return
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]
    // 取变量对象
	object, err := unescapePath(vars["object"])
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}

	getObjectInfo := objectAPI.GetObjectInfo

	objInfo, err := getObjectInfo(ctx, bucket, object)
	if err != nil {
        writeErrorResponseHeadersOnly(w, toAPIError(ctx, err))
		return
	}

	if err = setObjectHeaders(w, objInfo); err != nil {
		writeErrorResponseHeadersOnly(w, toAPIError(ctx, err))
		return
	}

    w.WriteHeader(http.StatusOK)
}

func setObjectHeaders(w http.ResponseWriter, objInfo ObjectInfo) (err error) {
	setCommonHeaders(w)

	lastModified := objInfo.ModTime.UTC().Format(http.TimeFormat)
	w.Header().Set(LastModified, lastModified)

	if objInfo.ETag != "" {
		w.Header()[ETag] = []string{"\"" + objInfo.ETag + "\""}
	}

    w.Header().Set(ContentLength, strconv.FormatInt(objInfo.Size, 10))

	return nil
}

// 暂不提供支持
func (api objectAPIHandlers) NewMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "NewMultipartUpload")
    writeErrorResponseHeadersOnly(w, toAPIError(ctx, nil))
}
