package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	uuidValue := uuid.New().String()
	fmt.Printf("%s", uuidValue)
}
