package main

import (
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/executors"
)

func main() {
	executor := executors.NewBulkExecutor(
		func(items []interface{}) {
			fmt.Println(len(items))
			fmt.Printf("%+v\n", items)
		},
		executors.WithBulkTasks(11),
	)
	for {
		if err := executor.Add(3); err != nil {
			fmt.Println(err)
			return
		}

		time.Sleep(time.Millisecond * 90)
	}
}
