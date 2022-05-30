package main

import (
	"fmt"
	"time"
)

func CurrentDateNotFormat() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("20060102")
}

func CurrentDate() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02")
}

func CurrentTimeNotFormat() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("150405")
}

func CurrentTime() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("15:04:05")
}

func CurrentDateTime() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("2006-01-02 15:04:05")
}

func CurrentDateTimeNotFormat() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("20060102_150405")
}

func TestDateTime() {
	fmt.Println(CurrentDate())              //2020-03-28
	fmt.Println(CurrentDateNotFormat())     //20200328
	fmt.Println(CurrentTime())              //11:11:26
	fmt.Println(CurrentTimeNotFormat())     //111126
	fmt.Println(CurrentDateTime())          //2020-03-28 11:11:26
	fmt.Println(CurrentDateTimeNotFormat()) //20200328_111126
}

func main() {
	TestDateTime()
	now := int(time.Now().Unix())
	fmt.Print(now)

}
