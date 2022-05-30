package main

import (
	"errors"
	"fmt"
	// "os"
)

func main() {
	// f, err := os.Open("1.txt")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// fmt.Println(f)

	err1 := errors.New("error")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)
	err := checkAge(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func checkAge(age int) error {
	if age < 0 {
		return errors.New("----")

		// err := fmt.Errorf("jfskf %v", age)
		// return err
	}

	fmt.Println("---", age)
	return nil
}
