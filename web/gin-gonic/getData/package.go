package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func build_package(resp *MyResp) (packageList []Package) {
	for _, value := range resp.List {
		var p Package
		var headers Headers
		headers.Key = value.ID
		headers.Topic = "professionsupply"
		headers.Timestamp = getTimeNow13()
		p.Headers = headers
		data, err := json.Marshal(value)
		if err != nil {
			fmt.Printf("value: [%v], err: [%v]\n", value, err)
			continue
		}
		p.Body = string(data)
		packageList = append(packageList, p)
	}
	return
}

func send_package(packageList []Package) {
	if len(packageList) > 0 {
		data, err := json.Marshal(packageList)
		if err != nil {
			fmt.Printf("packageList: [%v], err: [%v]\n", packageList, err)
		}

		url := "http://119.3.212.162:38011"
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Printf("url: [%s], StatusCode: [%d], err: [%v]\n", url, resp.StatusCode, err)
		}
		defer resp.Body.Close()
	}
}

func package_main(resp MyResp) {
	packageList := build_package(&resp)
	send_package(packageList)
}
