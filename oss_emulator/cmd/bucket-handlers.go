package emulator

import (
	"github.com/minio/mux"
	"net/http"
)

const (
	xOSSIOErrCodeHeader = "x-oss-error-code"
	xOSSErrDescHeader = "x-oss-error-desc"
)

func (api objectAPIHandlers) HoleWorld(w http.ResponseWriter, r *http.Request) {
	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}

	w.Write([]byte(bucket))
	w.Write([]byte("foo"))
}

func (api objectAPIHandlers) PutBucketHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutBucketHandler")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	if err := objectAPI.MakeBucket(ctx, bucket); err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}
	w.Header().Set(Location, pathJoin(SlashSeparator, bucket))

	writeSuccessResponseHeadersOnly(w)
}


func (api objectAPIHandlers) HeadBucketHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "HeadBucketHandler")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
	}

	vars := mux.Vars(r)
	bucket := vars["bucket"]

	getBucketInfo := objectAPI.GetBucketInfo

	if _, err := getBucketInfo(ctx, bucket); err != nil {
		writeErrorResponseHeadersOnly(w, toAPIError(ctx, err))
		return
	}

	writeResponse(w, http.StatusOK, nil, mimeXML)
}
