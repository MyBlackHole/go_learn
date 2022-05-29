package main

import (
	"fmt"
	"strconv"
)

type Pint struct {
	x int
	y int
}

var z interface{}

func main() {
	// var a Pint = Pint{1, 1, 4}
	// b := Pint{1, 1, 5}
	// var s string = a.z.(string)
	z = 4.0
	ss := interface3String(z)
	fmt.Println(ss)
}

func interface2String(inter interface{}) {
	switch inter.(type) {

	case string:
		fmt.Println("string", inter.(string))
		break
	case int:
		fmt.Println("int", inter.(int))
		break
	case float64:
		fmt.Println("float64", inter.(float64))
		break
	}
}

func interface3String(inter interface{}) (s string) {
	switch inter.(type) {
		case string:
			s = inter.(string)
			break
		case int:
			s = strconv.Itoa(inter.(int)) 
			break
		case float64:
			s = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
			break
		default:
			s = ""
	}
	return
}
