package main

import (
	"fmt"
	"strconv"
)

type Pint struct {
	x int
	y int
}

var z interface{}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {
	// var a Pint = Pint{1, 1, 4}
	// b := Pint{1, 1, 5}
	// var s string = a.z.(string)

	var p1 Person
	fmt.Println(p1)
	p1.name = "吴"
	p1.age = 26
	p1.sex = "男"
	p1.address = "广西"
	fmt.Println(p1)

	p2 := Person{}
	p2.name = "吴"
	p2.age = 26
	p2.sex = "男"
	p2.address = "广西"
	fmt.Println(p2)

	p3 := Person{"吴", 25, "男", "广西"}
	fmt.Println(p3)
	fmt.Printf("%p---%T\n", &p3, p3)

	p4 := p3
	p4.name = "定"
	fmt.Println(p4)
	fmt.Println(p3)
	fmt.Printf("%p---%T\n", &p4, p4)

	var pp1 *Person
	pp1 = &p4
	fmt.Println(pp1)
	fmt.Printf("%p---%T\n", &pp1, pp1)
	fmt.Printf("%v\n", *pp1)

	fmt.Println("---------------匿名结构体-----------------")

	type Student struct {
		name string
		age  int
	}

	s1 := Student{name: "吴", age: 26}
	fmt.Println(s1.name, s1.age)

	s2 := struct {
		name string
		age  int
	}{
		name: "吴",
		age:  26,
	}
	fmt.Println(s2.name, s2.age)

	fmt.Println("---------------匿名字段-----------------")
	type Worker struct {
		string
		int
	}

	w2 := Worker{"吴", 26}
	fmt.Println(w2)
	fmt.Println(w2.string)
	fmt.Println(w2.int)

	fmt.Println("--------------结构体嵌套------------------")

	// z = 4.0
	// ss := interface3String(z)
	// fmt.Println(ss)
}

func interface3String(inter interface{}) (s string) {
	switch inter.(type) {
	case string:
		s = inter.(string)
		break
	case int:
		s = strconv.Itoa(inter.(int))
		break
	case float64:
		s = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	default:
		s = ""
	}
	return
}

func interface2String(inter interface{}) {
	switch inter.(type) {

	case string:
		fmt.Println("string", inter.(string))
		break
	case int:
		fmt.Println("int", inter.(int))
		break
	case float64:
		fmt.Println("float64", inter.(float64))
		break
	}
}
