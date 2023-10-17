package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "./1.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	bs := make([]byte, 4, 4)
	// for {
	// 	n, err := file.Read()
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// }
	n, err := file.Read(bs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(n)
	fmt.Println(bs)
}
