package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.String())

	location, err := time.LoadLocation("US/Pacific")
	if err != nil {
		fmt.Println(err.Error(), "----")
		return
	}

	fmt.Println(location.String())

	fmt.Println(t.In(location).String())
}
