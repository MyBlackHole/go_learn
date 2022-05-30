package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) eat() {
	fmt.Println("p Person")
}

type Student struct {
	Person
	school string
}

func (s Student) eat() {
	fmt.Println("s Student")
}

func (s Student) study() {
	fmt.Println("s Student study")
}

func main() {
	p1 := Person{name: "wu", age: 26}
	fmt.Println(p1)
	p1.eat()

	s1 := Student{Person{name: "wu", age: 26}, "sha"}
	fmt.Println(s1)
	s1.eat()
	s1.Person.eat()
}
