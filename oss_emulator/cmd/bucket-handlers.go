package emulator

import (
    "net/http"
	"github.com/minio/mux"
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

