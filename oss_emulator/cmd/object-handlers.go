package emulator

import (
	"io"
	"net/http"

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

 //    // 存在等于 -1 的情况
	// if size == -1 {
	// 	writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrMissingContentLength), r.URL)
	// 	return
	// }

	objInfo, err := putObject(ctx, bucket, object, size, rd)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}
	setPutObjHeaders(w, objInfo, false)

	writeSuccessResponseHeadersOnly(w)
}


func (api objectAPIHandlers) AppendObjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "AppendObject")

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
		appendObject           = objectAPI.AppendObject
		size                = r.ContentLength
	)

	// if size == -1 {
	// 	writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrMissingContentLength), r.URL)
	// 	return
	// }

	objInfo, err := appendObject(ctx, bucket, object, size, rd)
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

// 暂不提供支持
func (api objectAPIHandlers) NewMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "NewMultipartUpload")
    writeErrorResponseHeadersOnly(w, toAPIError(ctx, nil))
}
