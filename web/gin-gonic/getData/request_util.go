package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Statuses int
	Body     []byte
}

func postJson(url string, jsonByt []byte) (response Response, err error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonByt))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	response.Statuses = resp.StatusCode
	response.Body, err = ioutil.ReadAll(resp.Body)
	return
}
