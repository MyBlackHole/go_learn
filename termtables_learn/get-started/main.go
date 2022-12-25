package main

import (
	"fmt"
	"github.com/scylladb/termtables"
)

func main() {
	t := termtables.CreateTable()
	t.AddHeaders("User", "Age")
	t.AddRow("dj", 18)
	t.AddRow("darjun", 30)
	fmt.Println(t.Render())
}
