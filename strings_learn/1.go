package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/black_hole/user/app"
	sl := strings.Split(s, "/")
	for _, s := range sl {
		fmt.Println(s)
	}

	fmt.Println(s[0])
}
