package main

import (
	"fmt"
	"gotool"
	"strings"
)

func main() {
	gotool.New()
	func1()
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
	ss := strings.ReplaceAll(s, "/", "")
	fmt.Print(ss)
}
