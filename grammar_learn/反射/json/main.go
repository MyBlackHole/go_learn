package main

import "fmt"
import "encoding/json"
import "reflect"

type Settings struct {
	A int
	B [][]int
}

func main() {

	// 准备模拟json数据
	var arr [][]int
	arr = append(arr, []int{1, 2})
	s, _ := json.Marshal(&arr)

	// 模拟动态设置B项
	var t Settings
	key := "B"

	tt := reflect.ValueOf(&t)
	tv := tt.Elem().FieldByName(key)

	fmt.Println("type:", tv.Type().String())

	newObj := reflect.New(tv.Type())
	err := json.Unmarshal(s, newObj.Interface())
	if err != nil {
		fmt.Println(err.Error())
	}

	tv.Set(newObj.Elem())
	fmt.Println(t)
}
