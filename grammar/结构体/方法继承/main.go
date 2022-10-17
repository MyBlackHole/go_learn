package main

import (
	"fmt"
)

type Animal struct {
	Name string
}

func (a *Animal) Eat() {
	fmt.Printf("%v is eating", a.Name)
	fmt.Println()
}

type Cat struct {
	Animal
}

func main() {
	cat := &Cat{
		Animal: Animal{
			Name: "cat",
		},
	}
	cat.Name = ""
	cat.Eat() // cat is eating
}
