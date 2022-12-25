package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewUUID()

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Print(id.String())
}
