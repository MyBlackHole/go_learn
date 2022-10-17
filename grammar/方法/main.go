package main

import "fmt"

type Worker struct {
	name string
	age  int
}

func (w Worker) work() {
	fmt.Println(w)
}

func (w *Worker) rest() {
	fmt.Println(w)
}

func main() {
	w1 := Worker{name: "吴", age: 26}
	w1.work()

	w2 := &Worker{name: "吴", age: 26}
	w2.work()

	w1.rest()
	w2.rest()
}
