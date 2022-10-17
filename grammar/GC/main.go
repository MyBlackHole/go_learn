package main

import (
	"fmt"
	"runtime"
	"time"
)

func INTPrtTest() {
	p := make([]*int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
		// out
		// GC took 9.205332484s
		// GC took 5.772939114s
		// GC took 5.626722723s
		// GC took 6.222295537s
		// GC took 6.081612918s
		// GC took 6.05748775s
		// GC took 5.86972785s
		// GC took 5.895571828s
		// GC took 5.994007784s
		// GC took 6.447978242s
	}

	runtime.KeepAlive(p)

}

func INTTest() {
	p := make([]int, 1e9)

	for i := 0; i < 10; i++ {
		start := time.Now()
		runtime.GC()
		fmt.Printf("GC took %s \n", time.Since(start))
		// out
		// GC took 762.213µs
		// GC took 255.852µs
		// GC took 238.934µs
		// GC took 234.915µs
		// GC took 193.474µs
		// GC took 208.56µs
		// GC took 177.434µs
		// GC took 199.421µs
		// GC took 191.172µs
		// GC took 177.833µs
	}

	runtime.KeepAlive(p)

}
func main() {
	INTTest()
}
