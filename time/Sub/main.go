package main

import (
	"fmt"
	"time"
)

func main() {
	t1str := "2022-06-08 12:00:00"
	t1time, _ := time.ParseInLocation("2006-01-02 15:04:05", t1str, time.Local)

	// 计算天数差
	fmt.Println(int(t1time.Sub(time.Now()).Hours() / 24))
}
