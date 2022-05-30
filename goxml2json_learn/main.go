package main

import (
	"fmt"
	"strings"

	xj "github.com/basgys/goxml2json"
)

func main() {
	// xml is an io.Reader
	xml := strings.NewReader(`<?xml version="1.0" encoding="UTF-8"?><hello>world</hello>`)
	json, err := xj.Convert(xml)
	if err != nil {
		panic("That's embarrassing...")
	}

	fmt.Println(json.String())
	// {"hello": "world"}
}
