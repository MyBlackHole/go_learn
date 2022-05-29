package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	x, y := robotgo.GetMousePos()
	println("x:", x, "y:", y)
}
