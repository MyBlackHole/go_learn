package main

import (
	"fmt"
	"reflect"
)

//定义struct
type T struct {
	Age      int
	Name     string
	Children []int
}

func structByReflect(data map[string]interface{}, inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}

	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		t := rType.Field(i)
		f := rVal.Field(i)
		if v, ok := data[t.Name]; ok {
			f.Set(reflect.ValueOf(v))
		} else {
			panic(t.Name + " not found")
		}
	}
}

func main() {
	// 初始化测试用例
	t := T{12, "someone-life", nil}
	fmt.Printf("%v", t)

	// 反射获取测试对象对应的struct枚举类型
	s := reflect.ValueOf(&t).Elem()

	// 内置常用类型的设值方法，利用Field序号get
	s.Field(0).SetInt(123)

	// 这里将slice转成reflect.Value类型; 用于下一步赋值
	sliceValue := reflect.ValueOf([]int{1, 2, 3})

	// 使用字段名称get，并赋值。赋值类型必须为reflect.Value
	s.FieldByName("Children").Set(sliceValue)
}
