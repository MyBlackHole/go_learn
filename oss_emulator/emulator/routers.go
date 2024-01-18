package emulator

import (
	"fmt"
	"net/http"

	"github.com/minio/mux"
)

func configureServerHandler(server serverCtxt) (http.Handler, error) {
	router := mux.NewRouter().SkipClean(true)

	registerAPIRouter(router)
	return router, nil
}

func listenAndServe(server serverCtxt, handler http.Handler) error {
	addr := fmt.Sprintf(":%d", server.Port)
	return http.ListenAndServe(addr, handler)
}
