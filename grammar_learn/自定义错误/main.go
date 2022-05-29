package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("ok")
}

func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		return errors.New("num2 == 0")
	} else {
		result := num1 / num2
		fmt.Println(result)
		return nil
	}
}
