package main

import "fmt"

func main() {
	var scores [5]int
	scores[0] = 95
	scores[1] = 91
	scores[2] = 39
	scores[3] = 60
	scores[4] = 21

	sum := 0
	for i := 0; i < len(scores); i++ {
		sum += scores[i]
	}

	avg := sum / len(scores)

	fmt.Println(sum, avg)

	var arr [3]int16
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	fmt.Printf("%p\n", &arr[1])
	fmt.Printf("%p\n", &arr[2])

	var arr1 [3]int = [3]int{3, 6, 9}
	fmt.Println(arr1)

	var arr2 = [3]int{1, 4, 7}
	fmt.Println(arr2)

	var arr3 = [...]int{1, 4, 7}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 1, 0: 4, 2: 7}
	fmt.Println(arr4)

	fmt.Printf("%T\n", arr4)
}
