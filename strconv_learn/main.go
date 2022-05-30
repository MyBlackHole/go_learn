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
}
