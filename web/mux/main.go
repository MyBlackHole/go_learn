package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// IMPORTANT: you must specify an OPTIONS method matcher for the middleware to set CORS headers
	// r.HandleFunc("/foo", fooHandler).Methods(http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodOptions)
	r.Methods(http.MethodPost).Path("/{bucket}/{object:.+}").HandlerFunc(fooHandler).Queries("append", "").Queries("position", "{position:[0-9]+}")
	r.Use(mux.CORSMethodMiddleware(r))

	r.NotFoundHandler = ErrorResponseHandler()

    log.Printf("start listen\n")
	err := http.ListenAndServe(":80", r)
	if err != nil {
		log.Panic(err)
	}
}

func ErrorResponseHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			return
		}
        log.Printf("errorResponseHandler\n")

		_, err := w.Write([]byte("errorResponseHandler"))
		if err != nil {
			log.Panic(err)
		}
	}
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
    log.Printf("fooHandler\n")

	_, err := w.Write([]byte("foo"))
	if err != nil {
		log.Panic(err)
	}
}
