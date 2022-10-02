package main

import "fmt"
import "time"

// import "log/gotool"

var iii int

func test() (i int, j int) {
	i = 1
	j = 2
	return
}

func main() {

	var i int = 100
	// const i int = 100
	i = 200
	// var i64 int64 = int64(i)
	fmt.Println(i)

	var i8 int8 = -1

	var i81 int8 = -127 / i8

	fmt.Println(i81)
	var in1 interface{}
	var in2 interface{}

	in1 = "ok"
	in2 = 6

	if in1 == in2 {
		fmt.Println(in1)
	} else {
		fmt.Println(in2)
	}

	// var (
	// 	i = 0
	// )
	// for {

	// 	fmt.Printf("http://product.cheshi.com/static/selectcar/%s.html\n", string(rune('A'+i)))
	// 	i++
	// 	if i > 25 {
	// 		break
	// 	}
	// }
	var s []string
	// var b []string
	if len(s) == 0 {
		fmt.Print(s)
	}

	fmt.Print(60 * 60 * 24 * 7 * time.Second)
	i1, j1 := test()
	fmt.Println(i1, j1)
	fmt.Println(iii)

	var arr = []int{1, 2, 3}

	for i := 0; i < len(arr)-1; {

		j := i + 1
		if j > len(arr)-1 {
			j = len(arr) - 1
		}
		fmt.Println(arr[i:j])
		i = j
	}
}
