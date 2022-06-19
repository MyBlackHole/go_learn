package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
)

func main() {
	buf := new(bytes.Buffer)
	pi := math.Pi

	err := binary.Write(buf, binary.LittleEndian, pi)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(buf.Bytes())
}
