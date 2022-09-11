package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type panicJob struct {
	count int
}

func (p *panicJob) Run() {
	p.count++
	if p.count == 1 {
		panic("oooooooooooooops!!!")
	}

	fmt.Println("hello world")
}

func test() {
	panic("oooooooooooooops!!!")
}

func main() {
	c := cron.New()
	// cron.Recover()()
	// c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(test))
	// c.AddJob("@every 1s", cron.NewChain(cron.Recover(cron.DefaultLogger)).Then(&panicJob{}))
	c.Start()

	time.Sleep(5 * time.Second)
}
