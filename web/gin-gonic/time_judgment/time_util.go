package main

import "time"

// 获取10位时间戳
func getTimeNow10()(i int64){
	i = time.Now().Unix()
	return
}

// 获取13位时间戳
func getTimeNow13()(i int64){
	i = getTimeNow10() * 1000
	return
}

func CurrentDate() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02")
}

func CurrentDateTime() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}
