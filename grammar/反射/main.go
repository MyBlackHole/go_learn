package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Hello struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	h := Hello{
		Name: "小王",
		Age:  15,
	}
	retH := reflect.TypeOf(h)
	//获取结构体里的名称
	for i := 0; i < retH.NumField(); i++ {
		field := retH.Field(i)
		fmt.Println("结构体里的字段名", field.Name)
		// fmt.Println("结构体里的字段属性:", field.Type)
		fmt.Printf("结构体里的字段属性:%T---%v\n", field.Type, field.Type)
		fmt.Printf("结构体里的字段类型:%T---%v\n", field.Type.Kind(), field.Type.Kind())
		fmt.Println("结构体里面的字段的tag标签", field.Tag)
		// reflect.Value 转 Interface
		fmt.Printf("%v %T\n", reflect.ValueOf(h).FieldByName(field.Name).Interface(), reflect.ValueOf(h).FieldByName(field.Name).Interface())
		fmt.Println("---------------")
	}
	//提取tag标签里的信息
	name, bool := retH.FieldByName("Name")
	if bool {
		fmt.Println("tag标签的信息为", name.Tag.Get("json"))
	}

	var list []interface{} = []interface{}{112, nil}
	// 通过反射填充结构体
	// hello := new(Hello)
	// retHello := reflect.ValueOf(hello)
	var hello Hello
	retHello := reflect.ValueOf(&hello)
	if retHello.Kind() == reflect.Ptr {
		retHello = retHello.Elem()
		for i := 0; i < retHello.NumField(); i++ {
			var value interface{}
			switch retHello.Field(i).Kind() {
			case reflect.String:
				switch v := list[i].(type) {
				case int:
					value = fmt.Sprintf("%d", v)
				case string:
					value = v
				}
			case reflect.Int:
				switch v := list[i].(type) {
				case int:
					value = v
				case string:
					var err error
					value, err = strconv.Atoi(v)
					if err != nil {
						fmt.Println(err.Error())
					}
				}
			}
			fmt.Printf("%T--%v\n", value, value)
			if value != nil && reflect.ValueOf(value).Kind() == retHello.Field(i).Kind() {
				retHello.Field(i).Set(reflect.ValueOf(value))
			}
		}
	}
	fmt.Println(hello)

	retHello = reflect.ValueOf(&hello)
	var list2 []interface{}
	if retHello.Kind() == reflect.Ptr {
		retHello = retHello.Elem()
		for i := 0; i < retHello.NumField(); i++ {
			list2 = append(list2, retHello.Field(i).Interface())
		}
	}

	fmt.Println(list2)
}
