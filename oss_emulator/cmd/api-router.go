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

    // 未实现
    bucketRouter.Methods(http.MethodPost).Path("/{object:.+}").HandlerFunc(httpTraceAll(api.NewMultipartUploadHandler)).Queries("uploads", "")

    bucketRouter.Methods(http.MethodHead).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.HeadObjectHandler))

    bucketRouter.Methods(http.MethodPut).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.PutObjectHandler))

    bucketRouter.Methods(http.MethodPost).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.AppendObjectHandler)).Queries("append", "").Queries("position", "{position:.*}")

    bucketRouter.Methods(http.MethodGet).Path("/{object:.+}").HandlerFunc(httpTraceHdrs(api.GetObjectHandler))

    bucketRouter.Methods(http.MethodGet).HandlerFunc((httpTraceAll(api.ListObjectsHandler)))

	bucketRouter.Methods(http.MethodHead).HandlerFunc(httpTraceAll(api.HeadBucketHandler))

	bucketRouter.Methods(http.MethodPut).HandlerFunc(httpTraceAll(api.PutBucketHandler))

	apiRouter.Methods(http.MethodGet).Path(SlashSeparator).HandlerFunc(httpTraceAll(api.ListBucketsHandler))

	apiRouter.MethodNotAllowedHandler = httpTraceAll(methodNotAllowedHandler("OSS"))
}
