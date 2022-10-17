package main

import "fmt"

type A struct {
	AA string `json:"aa"`
}

type B struct {
	AAList []*A   `json:"aa_list"`
	BB     string `json:"bb"`
}

func main() {
	var f func()
	f = fun1
	f()

	arr1 := fun2()
	fmt.Printf("%T-- %p-- %v\n", arr1, &arr1, arr1)

	arr3 := fun3()
	fmt.Printf("%T-- %p-- %v\n", arr3, &arr3, arr3)

	b := new(B)
	aaMap := make(map[string][]*A)
	bbMap := make(map[string]*A)

	a := new(A)

	aaMap["a"] = []*A{}
	aaList := aaMap["a"]
	aaList = append(aaList, a)
	b.AAList = aaMap["a"]
	bbMap["b"] = a
	fmt.Println(b.AAList[0].AA)
	bbMap["b"].AA = "cc"
	fmt.Println(a)
	fmt.Println(b.AAList[0].AA)

}

func fun3() *[4]int {
	return &[4]int{1, 2, 3, 4}
}

func fun2() [4]int {
	return [4]int{1, 2, 3, 4}
}

func fun1() {
	fmt.Println("ok")
}
