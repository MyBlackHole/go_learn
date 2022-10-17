package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/zeromicro/go-zero/core/collection"
)

const interval = time.Minute

type Func func() error

var traditional = flag.Bool("traditional", false, "enable traditional mode")

func main() {
	flag.Parse()

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Printf("goroutines: %d\n", runtime.NumGoroutine())
			}
		}
	}()

	timingWheelMode()
}

func timingWheelMode() {
	tw, err := collection.NewTimingWheel(time.Second, 600, func(key, value interface{}) {
		job(value.(Func))
	})
	if err != nil {
		log.Fatal(err)
	}

	defer tw.Stop()

	ticker := time.NewTicker(time.Second * 5)

	for {
		select {
		case <-ticker.C:

			for k, v := range map[string]Func{
				"A": A,
				"B": B,
			} {
				tw.SetTimer(k, v, time.Second*1)
			}
		}
	}
}

func job(f Func) {
	f()
}

func A() error {
	fmt.Println("A")
	return nil
}

func B() error {
	fmt.Println("B")
	return nil
}
