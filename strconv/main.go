package main

import "strconv"
import "fmt"

func main() {
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T -- %t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T -- %s\n", ss1, ss1)

	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T -- %d\n", i2, i2)

	i, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(i)
	}

	var a byte = 'A'
	b := 'A' + 1
	i1 := 1
	fmt.Println(a)
	fmt.Println(a + byte(i1))
	fmt.Println(string(b))
	fmt.Printf("%T\n", b)

	fmt.Printf("%c\n", 65)
}
