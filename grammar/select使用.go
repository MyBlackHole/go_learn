package main

import (
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Black Hole 1"
}

func server2(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "Black Hole 2"
}

func main() {
	func2()
}

func func2() {
	ch := make(chan string)
	go server1(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			print(v)
		default:
			print("no value\n")
		}
	}
}

func func1() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		print(s1)
	case s2 := <-output2:
		print(s2)
	}
}
