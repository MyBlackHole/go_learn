package main

import (
	"fmt"
	"time"
)

func main() {
	t1str := "2026-12-08 12:00:00"
	t1time, _ := time.ParseInLocation("2006-01-02 15:04:05", t1str, time.Local)
	if t1time.Before(time.Now()) {
		fmt.Println("t1time has arrived")
	} else {
		fmt.Println("t1time hasn't come yet")
	}
}
