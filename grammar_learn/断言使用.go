package main

import "fmt"

func main() {
	var i interface{} = 10
	t1, ok := i.(int)
	fmt.Printf("%d-%t\n", t1, ok)

	fmt.Println("=====分隔线1=====")

	t2, ok := i.(string)
	fmt.Printf("%s-%t\n", t2, ok)

	fmt.Println("=====分隔线2=====")

	var k interface{} // nil
	t3, ok := k.(interface{})
	fmt.Println(t3, "-", ok)

	fmt.Println("=====分隔线3=====")
	k = 10
	t4, ok := k.(interface{})
	fmt.Printf("%d-%t\n", t4, ok)

	t5, ok := k.(int)
	fmt.Printf("%d-%t\n", t5, ok)
}
