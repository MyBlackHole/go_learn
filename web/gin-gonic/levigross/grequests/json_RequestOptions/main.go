package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/levigross/grequests"
)

type RequestOptions struct {

	// Data is a map of key values that will eventually convert into the
	// the body of a POST request.
	Data map[string]string

	// Params is a map of query strings that may be used within a GET request
	Params map[string]string

	// QueryStruct is a struct that encapsulates a set of URL query params
	// this paramter is mutually exclusive with `Params map[string]string` (they cannot be combined)
	// for more information please see https://godoc.org/github.com/google/go-querystring/query
	QueryStruct interface{}

	// Files is where you can include files to upload. The use of this data
	// structure is limited to POST requests
	Files []grequests.FileUpload

	// JSON can be used when you wish to send JSON within the request body
	JSON interface{}

	// Headers if you want to add custom HTTP headers to the request,
	// this is your friend
	Headers map[string]string
}

func main() {
	var err error
	var b []byte
	var s string = `{"Data":null,"Params":null,"QueryStruct":null,"Files":null,"JSON":null,"Headers":{"ok": "ok"}}`
	var requestoptions = grequests.RequestOptions{}
	var options = RequestOptions{}

	var testString string

	b, err = json.Marshal(options)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(b))

	err = json.Unmarshal([]byte(s), &requestoptions)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%v\n", requestoptions)

	fmt.Printf("\"%s\"\n", testString)

	os.Exit(0)
}
