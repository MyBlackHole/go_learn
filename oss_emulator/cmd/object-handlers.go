package emulator

import (
	"github.com/minio/mux"
	"io"
	"net/http"
)

func (api objectAPIHandlers) PutObjectHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutObjectHandler")

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
