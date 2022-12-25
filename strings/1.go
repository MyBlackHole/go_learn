package main

import (
	"fmt"
	"strings"
)

func main() {
	func4()
}

func func4() {
	var s = []string{
		"ok",
		"kkk",
	}
	var sss *[]string = &s
	ss := strings.Join(*sss, ",")
	fmt.Print(ss)
}

func func3() {
	s := "/a/b"
	i := strings.HasPrefix(s, "/a")
	print(i)
}

type OK struct {
	a int
	*int
}

func func2() {
	ok := OK{a: 1}
	i := 0
	ok.int = &i
	print(ok.a, ok.int)
}

func func1() {
	s := "/black_hole/user/app"
	sl := strings.Split(s, "/")
	for _, s := range sl {
		fmt.Println(s)
	}
	fmt.Println(s[0])
	fmt.Println(len(sl))
	var ss string
	if "" == ss {
		fmt.Println(111)
	}
	fmt.Println(ss)
	fmt.Println(222)
	fmt.Printf("\u2225")
}
