package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(100) // 因为有两个动作，所以增加2个计数
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println("Goroutine ", i)
			wg.Done() // 操作完成，减少一个计数
		}(i)
	}

	wg.Wait() // 等待，直到计数为0
}
