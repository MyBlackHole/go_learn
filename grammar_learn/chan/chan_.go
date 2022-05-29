package main

import "fmt"
import "time"

type CH struct {
	limitChan chan bool
}

func read(ch CH) {
	for true {
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch.limitChan)
	}
}

func main() {
	var ch CH
	var count int
	ch.limitChan = make(chan bool, 10)
  go read(ch)
	for true {
		ch.limitChan <- true
    count += 1
		fmt.Println(count)
	}
}
