package emulator

import (
    "net/http"
	"github.com/minio/mux"
)


type objectAPIHandlers struct {
	ObjectAPI func() ObjectLayer
}

func registerAPIRouter(router *mux.Router) {
	router.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)

}
