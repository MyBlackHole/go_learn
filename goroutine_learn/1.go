package main

import (
	"fmt"
	"time"
)

func run(done chan int) {
	for {
		select {
		case <-done:
			fmt.Println("exiting...")
			done <- 1
			break
		default:
		}

		time.Sleep(time.Second * 1)
		fmt.Println("do something")
	}
}

func main() {
	c := make(chan int)

	go run(c)

	fmt.Println("wait")
	time.Sleep(time.Second * 5)

	c <- 1
	<-c

	fmt.Println("main exited")
}
