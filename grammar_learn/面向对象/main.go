package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct {
	Person
	school string
}

func main() {

	p1 := Person{name: "吴", age: 26}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	s1 := Student{Person{name: "吴", age: 26}, "小学"}
	fmt.Println(s1)

	var s2 Student
	s2.Person.name = "吴"
	s2.Person.age = 18
	s2.school = "小学"
	fmt.Println(s2)

	var s3 Student
	s3.name = "吴"
	s3.age = 18
	s3.school = "小学"
	fmt.Println(s3)
}
