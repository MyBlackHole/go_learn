package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func handleSearch(ctx context.Context) {

	for {
		time.Sleep(time.Second)
		fmt.Println("working……")
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1) //添加

	// 如何实现go 子程序优雅的退出
	go handleSearch(ctx)
	// 案例监听5秒
	time.Sleep(time.Second * 5)
	// 连续5秒都没完成则自动退出
	cancel()  // 退出
	wg.Wait() // 等待结束
	fmt.Println("over")
}
