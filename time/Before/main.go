package main

import (
	"fmt"
	"time"
)

func main() {
	t1str := "2026-12-08"
	t1time, _ := time.ParseInLocation("2006-01-02", t1str, time.Local)
	fmt.Println(t1time)
	if t1time.Before(time.Now()) {
		fmt.Println("t1time has arrived")
	} else {
		fmt.Println("t1time hasn't come yet")
	}
}
