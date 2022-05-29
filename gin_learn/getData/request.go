package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func xgRequest(xgperson *XGPerson) (xgresp XGResp, err error) {
	data, err := json.Marshal(&xgperson)
	if err != nil {
		return
	}

	fmt.Print(string(data))

	response, err := postJson("https://xgsj.istarshine.com/v3/xsearch?token=ac94e046-485c-4e8e-a42c-059a7d855c3b", data)
	if err != nil {
		return
	}

	ioutil.WriteFile("1.json", response.Body, 0755)

	err = json.Unmarshal(response.Body, &xgresp)
	if err != nil {
		return
	}

	if xgresp.Mes != "success" {
		err = errors.New(xgresp.Mes)
		return
	}
	return
}

func request(person *Person) (resp Resp, err error) {
	data, err := json.Marshal(&person)
	if err != nil {
		return
	}

	fmt.Print(string(data))

	response, err := postJson("http://47.112.188.43:8081/itf/newslist", data)
	if err != nil {
		return
	}

	err = json.Unmarshal(response.Body, &resp)
	if err != nil {
		return
	}

	if resp.Message != "success" {
		err = errors.New(resp.Message)
		return
	}

	return
}
