package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func request() {

    url := "http://192.168.78.212:8009/api/v3/policy_job"
	method := "PATCH"

	payload := strings.NewReader(`{
  "policy_id": 2,
  "backup_object_id": 2,
  "policy_job_id_list": ["jjjjjj"],
  "version": "v2",
  "enable": true
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImFkbWluIiwiaWF0IjoxNjg3OTcwMjQxfQ.NdtJRJcMtgh2iO7jdj64axhOm6zAWmz4dzyVshFUxsA")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}


func main() {
	for i := 0; i < 200; i++ {
        go request()
    }
    time.Sleep(100 * time.Second)
}
