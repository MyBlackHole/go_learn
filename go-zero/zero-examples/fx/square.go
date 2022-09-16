package main

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/fx"
)

func main() {
	result, err := fx.From(func(source chan<- interface{}) {
		// 生成数据源
		for i := 0; i < 10; i++ {
			source <- i
		}
	}).Map(func(item interface{}) interface{} {
		// 修改每一个 item
		i := item.(int)
		return i * i
	}).Filter(func(item interface{}) bool {
		// 返回 true 保留
		i := item.(int)
		return i%2 == 0
	}).Distinct(func(item interface{}) interface{} {
		// 去重
		return item
	}).Reduce(func(pipe <-chan interface{}) (interface{}, error) {
		var result int
		// 求和
		for item := range pipe {
			i := item.(int)
			result += i
		}
		// 返回
		return result, nil
	})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
