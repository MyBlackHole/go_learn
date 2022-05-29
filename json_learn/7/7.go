package main

import (
	"encoding/json"
	"fmt"
)

type Value struct {
}

func main() {

	var strList []string
	strList = append(strList, "ok")
	strList = append(strList, "ko")

	str, err := json.Marshal(strList)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", str)

	var strList1 []string
	err = json.Unmarshal(str, &strList1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strList1)

	var mapList map[string]Value
	mapList = make(map[string]Value)
	mapList["1"] = Value{}
	mapList["2"] = Value{}

	str, err = json.Marshal(mapList)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", str)

	var mapList1 map[string]Value
	err = json.Unmarshal(str, &mapList1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(mapList1)
	if s, err := mapList1["3"]; err {
		fmt.Println(s)
	}

	value, ok := mapList1["2"]
	if ok {
		fmt.Println(err)
	}
	fmt.Println(value)

}
