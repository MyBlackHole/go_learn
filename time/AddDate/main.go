package main

import (
	"fmt"
	"time"
)

func main() {
	nowTime := time.Now()
	getTime := nowTime.AddDate(0, 0, -1)                //年，月，日   获取一天前的时间
	resTime := getTime.Format("2006-01-02 15:04:05+08") //获取的时间的格式
	fmt.Println(resTime)

	getTime = nowTime.AddDate(0, -1, 0)             //年，月，日   获取一个月前的时间
	resTime = getTime.Format("2006-01-02 15:04:05") //获取的时间的格式
	fmt.Println(resTime)

	getTime = nowTime.AddDate(-2, 0, 0)  //年，月，日   获取两年前的时间
	resTime = getTime.Format("20060102") //获取的时间的格式
	fmt.Println(resTime)

}
