package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	byts, err := ioutil.ReadFile("1.json")
	if err != nil {
		fmt.Printf("%s", err.Error())
	}

	var xg = XGResp{}

	json.Unmarshal(byts, &xg)
	fmt.Print(xg)
}


func TestIntFloat(t *testing.T) {
	var f float64 = 0.1
	fmt.Print(int64(f))
}