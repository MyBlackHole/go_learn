package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

type Black struct {
	Phone
}

func (black Black) call() {
	fmt.Println("I am black, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	black := new(Black)
	black.call()

	fmt.Println("-----------------接口类型-------------------")
	fmt.Println("-----------------空接口-------------------")

	type A interface {
	}

	type Cat struct {
		color string
	}

	type Person struct {
		name string
	}

	test1 := func(a interface{}) {
		fmt.Println("---->", a)
	}

	var a1 A = Cat{color: "ok"}
	test1(a1)

	var p1 Person = Person{name: "sm"}
	test1(p1)

	map1 := make(map[string]interface{})
	map1["a"] = "ksfj"
	map1["b"] = 12
	map1["c"] = a1
	test1(map1)

	slica1 := make([]interface{}, 0, 10)
	slica1 = append(slica1, "ksfjl", 123, a1)
	// slica1 := make([]interface{}, 3, 10)
	// slica1[0] = "ksdjf"
	// slica1[1] = 123
	// slica1[2] = a1

	printSlica(slica1...)

	fmt.Println("-----------------接口嵌套-------------------")
	// var c C2 = Cat2{}
	var c C2
	c = Cat2{}
	c.test1()

	fmt.Println("-----------------接口断言-------------------")

	var c1 C2
	c1 = Cat2{name: "ok"}
	c1.test1()

	if v, ok := c1.(Cat3); ok {
		fmt.Println(ok)
		fmt.Println(v.name)
	} else {
		fmt.Println(ok)
		fmt.Println("---")
	}

	switch ins := c1.(type) {
	case Cat2:
		fmt.Println("-----switch--------")
		fmt.Println(ins.name)
	}
}

type A2 interface {
	test1()
}

type B2 interface {
	test2()
}

type C2 interface {
	A2
	B2
	test3()
}

type Cat3 struct {
	name string
}

func (c Cat3) test1() {
	fmt.Println("test1")

}

func (c Cat3) test2() {
	fmt.Println("test2")
}

func (c Cat3) test3() {
	fmt.Println("test2")
}

type Cat2 struct {
	name string
}

func (c Cat2) test1() {
	fmt.Println("test1")

}

func (c Cat2) test2() {
	fmt.Println("test2")
}

func (c Cat2) test3() {
	fmt.Println("test2")
}

// printSlica(slica1)
// func printSlica(list []interface{}) {
func printSlica(list ...interface{}) {
	for i := 0; i < len(list); i++ {
		fmt.Printf("%v\n", list[i])
	}
}
