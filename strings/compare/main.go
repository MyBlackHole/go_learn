package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("GO", "go"))
	fmt.Println(strings.Compare("go", "go"))
	fmt.Println(strings.Compare("PO0546000070", "PO0546000080"))
}
