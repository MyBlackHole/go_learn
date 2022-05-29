package main

import "fmt"

func main() {
	var countmap map[string]string
	countmap = make(map[string]string)
	countmap["a"] = "a"
	countmap["b"] = "b"
	for item := range countmap {
		fmt.Println(item)
	}
}
