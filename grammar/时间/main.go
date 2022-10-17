package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%v -- %T\n", now, now)
	fmt.Println(now)

	fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	fmt.Println("---------")

	// 指定格式
	fmt.Println(now.Format("2006/01/02 15/04/04"))
	fmt.Println(now.Format("2006-01-02 15:04:04"))
}
