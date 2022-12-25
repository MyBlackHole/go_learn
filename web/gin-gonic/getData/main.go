package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go send_dindin()

	r := gin.Default()
	r.POST("/get_total", getData)
	r.GET("/total", getRedisData)
	r.Run()
}

func send_dindin() {
	for true {
		var text, title string
		var key string = CurrentDate()

		count, err := redisdb.SCard(key).Result()
		if err != nil {
			log.Print(err)
		}

		resp, err := redisdb.SRandMemberN(key, 10).Result()
		if err != nil {
			log.Print(err)
		}

		title = "当天请求量 --- " + strconv.Itoa((int)(count))

		text = "### " + title + "\n" + "> [请求示例]"
		for _, redisDataStr := range resp {
			text += "> " + redisDataStr + "\n"
		}

		SendDingMsg(title, text)
		time.Sleep(10 * 60 * time.Second)
	}
}
