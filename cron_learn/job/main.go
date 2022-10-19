package main

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Hello ", g.Name)

	b := 0
	j := 1 / b
	print(j)
}

func main() {
	// c := cron.New()
	// c.AddJob("@every 1s", GreetingJob{"dj"})
	// // c.AddFunc()
	// c.Start()
	test()

	time.Sleep(50 * time.Second)
}

func test() {
	go func() {

		// cron-v1时间写法示例timeout := "秒 分 时 日 月 年"
		// 每分钟的第5s、25s、45s各执行一次：5,25,45 * * * * *
		// 每12s执行一次：*/12 * * * * *
		// 每隔1分钟的第0秒执行一次：0 */1 * * * *
		// 每天23:30:00执行一次：0 30 23 * * *
		// 每天凌晨1:00:00执行一次：0 0 1 * * *
		// 每月1号早上6:00:00执行一次：0 0 6 1 * *
		// 在每小时的26分、29分、33分各执行一次：0 26,29,33 * * *
		// 每天的0点、13点、18点、21点都执行一次：0 0 0,13,18,21 * *
		// 每天十点到十二点每五秒执行一次：*/5 * 10-12 * * *

		num := 0 // 运行次数

		local, _ := time.LoadLocation("Asia/Shanghai")
		interval := cron.New(cron.WithLocation(local), cron.WithSeconds())
		// 周六到周日
		timeout := "TZ=US/Pacific * * * * * 3"
		_, err := interval.AddFunc(timeout, func() {
			num++
			log.Println("全局定时器已开启=num=", num)
		})

		log.Println("全局定时器运行结果，error=", err)
		interval.Start()
	}()
}
