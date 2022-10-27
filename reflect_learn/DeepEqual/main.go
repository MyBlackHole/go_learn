package main

import (
	"fmt"

	"reflect"
)

type tt struct {
	Code int
	a    *int
}

func main() {
	var i int
	var j int
	i = 2
	j = 2
	m1 := map[string]int{"a": 1, "b": 2, "c": 3}
	m2 := map[string]int{"a": 1, "c": 3, "b": 2}

	fmt.Println("m1 == nil?", m1 == nil)
	fmt.Println("m2 != nil?", m2 != nil)
	//fmt.Println("m1==m2",m1==m2)
	fmt.Println("cmpMap(m1,m2) = ", cmpMap(m1, m2))
	fmt.Println("reflect.DeepEqual(m1,m2) = ", reflect.DeepEqual(m1, m2))
	fmt.Println()
	m3 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 1}
	fmt.Println("cmpMap(m1,m3)=", cmpMap(m1, m3))
	fmt.Println("reflect.DeepEqual(m1,m3) = ", reflect.DeepEqual(m1, m3))

	fmt.Println("reflect.DeepEqual(m1,m3) = ----", reflect.DeepEqual(&tt{Code: 1, a: &i}, &tt{Code: 1, a: &j}))
}
func cmpMap(m1, m2 map[string]int) bool {
	for k1, v1 := range m1 {
		if v2, has := m2[k1]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	for k2, v2 := range m2 {
		if v1, has := m1[k2]; has {
			if v1 != v2 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}