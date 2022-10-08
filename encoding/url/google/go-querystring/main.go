package main

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

// go get -u github.com/google/go-querystring/query

type Info struct {
	Name    string `json:"name" url:"name"`
	Age     int    `json:"age" url:"age"`
	Comment string `url:"comments"`
}

type SubNested struct {
	Value string `url:"value"`
}

type Nested struct {
	A   SubNested  `url:"a"`
	B   *SubNested `url:"b"`
	Ptr *SubNested `url:"ptr,omitempty"`
	C   string     `url:"c"`
	D   string     `url:"-"`
	E   *SubNested `url:"e"`
}

func main() {
	in := Nested{
		A: SubNested{Value: "ab"},
		B: &SubNested{Value: "b"},
	}
	v, err := query.Values(in)   // 会自动a-z排序、转义
	fmt.Println(v.Encode(), err) // age=2&comments=%E6%B5%8B%E8%AF%95&name=aq <nil>
}
