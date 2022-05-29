package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)
	fmt.Println(age)

	var name string
	fmt.Scanf("%d %s", &age, &name)
	fmt.Println(age, name)
}
