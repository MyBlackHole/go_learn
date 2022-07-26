package main

import (
	"fmt"
	"time"
)

func main() {
	// string转化为时间，layout必须为 "2006-01-02 15:04:05"
	var timeLayout string = "2006-01-02 15:04:05"
	// 有时间区间8小时问题
	times, _ := time.Parse(timeLayout, "2014-06-15 08:37:18")
	timeUnix := times.Unix()
	fmt.Printf("times is %+v \n, timeUnix is %+v", times, timeUnix)
	// 建议
	times, _ = time.ParseInLocation("2006-01-02 15:04:05", timeLayout, time.Local)
	fmt.Printf("times is %+v \n, timeUnix is %+v\n", times, timeUnix)
	var month time.Month
	month = 2

	startTime := time.Date(2022, month, 1, 0, 00, 0, 0, time.UTC)
	end := time.Unix(time.Date(2022, month, 1, 0, 00, 0, 0, time.UTC).Unix()-1, 0).In(time.UTC)
	fmt.Println(startTime, "----")
	fmt.Println(end, "----")
}
