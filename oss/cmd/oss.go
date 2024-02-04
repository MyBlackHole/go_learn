package oss

import (
	"fmt"
    "net/http"

	"github.com/minio/mux"
	cli "github.com/urfave/cli/v2"
)

var serverCmd = &cli.Command{
	Name:  "server",
	Usage: "start object storage server",
	Action: serverMain,
    OnUsageError: func(cCtx *cli.Context, err error, isSubcommand bool) error {
        panic(err)
    },
}

func registerAPIRouter(router *mux.Router) {
    api := OssServer{}
	router.Methods(http.MethodHead).HandlerFunc(httpTrace(api.HeadHandler))
    router.Methods(http.MethodGet).HandlerFunc(httpTrace(api.GetHandler))
	router.Methods(http.MethodPut).HandlerFunc(httpTrace(api.PutHandler))
    router.Methods(http.MethodPost).HandlerFunc(httpTrace(api.PostHandler))
    router.Methods(http.MethodOptions).HandlerFunc(httpTrace(api.OptionsHandler))
    router.Methods(http.MethodDelete).HandlerFunc(httpTrace(api.DeleteHandler))
}


func serverMain(ctx *cli.Context) error {
	fmt.Println("start server main")

    GConfig = NewConfig()

	router := mux.NewRouter().SkipClean(true)
    registerAPIRouter(router)

	return listenAndServe(router)
}


func listenAndServe(handler http.Handler) error {
	addr := fmt.Sprintf(":%d", 80)
	return http.ListenAndServe(addr, handler)
}
