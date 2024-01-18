package emulator

import (
	"github.com/minio/mux"
	"net/http"
)

func setObjectLayer(o ObjectLayer) {
	globalObjectAPI = o
}

func newObjectLayerFn() ObjectLayer {
	return globalObjectAPI
}

type objectAPIHandlers struct {
	ObjectAPI func() ObjectLayer
}

func registerAPIRouter(router *mux.Router) {
	api := objectAPIHandlers{
		ObjectAPI: newObjectLayerFn,
	}

	apiRouter := router.PathPrefix(SlashSeparator).Subrouter()

    bucketRouter := apiRouter.PathPrefix("/{bucket}").Subrouter()

	bucketRouter.Methods(http.MethodPut).HandlerFunc(httpTraceAll(api.PutBucketHandler))

	apiRouter.MethodNotAllowedHandler = httpTraceAll(methodNotAllowedHandler("OSS"))
}
