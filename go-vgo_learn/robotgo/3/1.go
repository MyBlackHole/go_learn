package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {
	robotgo.MoveMouse(800, 400)

	robotgo.MoveMouseSmooth(800, 400)

	robotgo.MoveMouseSmooth(800, 400, 20.0, 200.0)
}
