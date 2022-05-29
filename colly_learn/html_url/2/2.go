package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	// url := "https://httpbin.org/delay/1"

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// create a request queue with 2 consumer threads
	// 设置消费者 线程数  缓存大小
	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 1}, // Use default queue storage
	)

	// c.OnRequest(func(r *colly.Request) {
	// })

	c.OnError(func(response *colly.Response, err error) {
		q.AddRequest(response.Request)
		fmt.Printf("%s\n", response.Request.URL.String())
	})

	for i := 0; i < 5; i++ {
		// Add URLs to the queue
		// 1、添加待访问url到 队列中
		q.AddURL("https://www.google.com/search?q=ij&oq=ij&aqs=edge..69i57.38992j0j1&sourceid=chrome&ie=UTF-8")
		q.AddURL("https://baike.com/item/%E5%8F%8C%E5%90%91%E9%93%BE%E8%A1%A8/2968731?fr=aladdin")
	}
	// Consume URLs
	// 2、开始消费 访问
	q.Run(c)

}

