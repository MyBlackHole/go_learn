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

	location, err = time.LoadLocation("Europe/Stockholm")
	if err != nil {
		fmt.Println(err.Error(), "----")
		return
	}

	fmt.Println(location.String())
	fmt.Println(t.In(location).String())

	location, err = time.LoadLocation("Europe/Warsaw")
	if err != nil {
		fmt.Println(err.Error(), "----")
		return
	}

	fmt.Println(location.String())
	fmt.Println(t.In(location).String())

	location, err = time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err.Error(), "----")
		return
	}

	fmt.Println(location.String())
	fmt.Println(t.In(location).String())

}
