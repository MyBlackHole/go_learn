package main

import (
	"fmt"
	"time"
)

// ANSIC       = "Mon Jan _2 15:04:05 2006"
// UnixDate    = "Mon Jan _2 15:04:05 MST 2006"

func CurrentDateNotFormat() string {
	timeUnix := time.Now().Unix()
	return time.Unix(timeUnix, 0).Format("20060102")
}

func CurrentDate() string {
	timeUnix := time.Now().Unix()
	fmt.Println(time.Unix(timeUnix, 0), "Unix")
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

func ParseYMDDate(dateTime string) time.Time {
	dt, _ := time.Parse("2006-01-02", dateTime)
	return dt
}

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())
	fmt.Println()

	fmt.Println(time.Now())
	TestDateTime()
	now := int(time.Now().Unix())
	fmt.Println(now)

	fmt.Println(time.Now())

	fmt.Println()
	time.Parse(time.Now().Format("2006-01-02"), "2006-01-02")
	t := time.Date(2022, 7, 26, 12, 25, 0, 0, time.UTC)
	fmt.Println(t.String())
	t1 := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
	fmt.Println(t1, "t1")

	fmt.Println(time.Date(2022, 7, 26, 12, 25, 0, 0, time.UTC).Location().String())

	fmt.Println(t1.Weekday())
	t, err := time.Parse(time.RFC3339, "2022-06-07T16:20:03+00:00")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(t)

	fmt.Printf("%s\n", ParseYMDDate("2022-09-27"))

	l, _ := time.LoadLocation("America/Los_Angeles")

	fmt.Printf("%d\n", int(time.Now().Weekday()))
	fmt.Printf("%d\n", int(time.Now().In(l).Weekday()))

	fmt.Println(ParseInLocation("24:00"))
}

func ParseInLocation(timeStr string) (time.Time, error) {
	return time.ParseInLocation("15:04", timeStr, time.Local)
}
