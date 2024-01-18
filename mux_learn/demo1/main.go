package main

import (
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
    router.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
    router.Use(mux.CORSMethodMiddleware(router))
    
    http.ListenAndServe(":8080", router)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    if r.Method == http.MethodOptions {
        return
    }

    w.Write([]byte("foo"))
}
