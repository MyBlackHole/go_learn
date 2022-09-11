package main

import (
	"encoding/json"
	"fmt"
)

type Test struct {
	A string
	B int
	C string
}

func (t *Test) UnmarshalJSON(data []byte) error {
	type testAlias Test
	test := &testAlias{
		A: "default A",
		B: -2,
	}
	_ = json.Unmarshal(data, test)
	*t = Test(*test)
	return nil
}

// var example []byte = []byte(`[{"A": "1", "C": "3"}, {"A": "4", "B": 2}, {"C": "5"}]`)
var example []byte = []byte(`{"A": "1", "C": "3"}`)

func main() {
	var out *Test
	_ = json.Unmarshal(example, out)
	fmt.Print(out)
}
