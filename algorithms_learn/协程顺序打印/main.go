package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	catCh := make(chan struct{})
	dogCh := make(chan struct{})
	fishCh := make(chan struct{})
	// 打印cat的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-catCh
			fmt.Println("cat")
			dogCh <- struct{}{}
		}
	}()
	// 打印dog的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-dogCh
			fmt.Println("dog")
			fishCh <- struct{}{}
		}
	}()
	// 打印fish的协程
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			<-fishCh
			fmt.Println("fish")
			if i < 9 {
				catCh <- struct{}{}
			} else {
				// 9 退出
				break
			}
		}
	}()
	// 启动第一个协程
	catCh <- struct{}{}
	wg.Wait()
}
