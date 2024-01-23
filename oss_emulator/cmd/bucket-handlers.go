package emulator

import (
	"github.com/minio/mux"
	"net/http"
)

const (
	xOSSIOErrCodeHeader = "x-oss-error-code"
	xOSSErrDescHeader   = "x-oss-error-desc"
)

func (api objectAPIHandlers) PutBucketHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "PutBucket")

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
	ctx := newContext(r, w, "HeadBucket")

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

func (api objectAPIHandlers) ListBucketsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "ListBuckets")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
	}

	listBuckets := objectAPI.ListBuckets

	bucketsInfo, err := listBuckets(ctx)
	if err != nil {
		writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
		return
	}
	response := generateListBucketsResponse(bucketsInfo)
	encodedSuccessResponse := encodeResponse(response)
	writeSuccessResponseXML(w, encodedSuccessResponse)
}

func (api objectAPIHandlers) ListObjectsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(r, w, "ListObjects")

	objectAPI := api.ObjectAPI()
	if objectAPI == nil {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(ErrServerNotInitialized), r.URL)
	}

	// vars := mux.Vars(r)
	// bucket := vars["bucket"]

	prefix, marker, delimiter, maxKeys, encodingType, s3Error := getListObjectsArgs(r.Form)
	if s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL)
		return
	}

	if s3Error := validateListObjectsArgs(prefix, marker, delimiter, encodingType, maxKeys); s3Error != ErrNone {
		writeErrorResponse(ctx, w, errorCodes.ToAPIErr(s3Error), r.URL)
		return
	}

	// listObjects := objectAPI.ListObjects

	// _, err := listObjects(ctx, bucket, prefix, marker, delimiter, maxKeys)
	// if err != nil {
	// 	writeErrorResponse(ctx, w, toAPIError(ctx, err), r.URL)
	// 	return
	// }
}
