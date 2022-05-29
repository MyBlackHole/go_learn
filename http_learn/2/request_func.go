package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
    Statuses int
    Body []byte
}

func httpPost() {
    resp, err := http.Post("https://www.abcd123.top/api/v1/login",
        "application/x-www-form-urlencoded",
        strings.NewReader("username=test&password=ab123123"))
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}

func httpPostForm() {
    resp, err := http.PostForm("https://www.denlery.top/api/v1/login",
        url.Values{"username": {"auto"}, "password": {"auto123123"}})
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    fmt.Println(string(body))
}

func httpPostJson() {
    jsonStr :=[]byte(`{ "username": "auto", "password": "auto123123" }`)
    url:= "https://www.denlery.top/api/v1/login"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        // handle error
    }
    defer resp.Body.Close()

    statuscode := resp.StatusCode
    hea := resp.Header
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))
    fmt.Println(statuscode)
    fmt.Println(hea)
}

func postJson(url string, jsonByt []byte) (response Response,err error) {
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