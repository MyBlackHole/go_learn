package main

import (
	"log"
	"os"
)

func main() {
	bytAll, err := os.ReadFile("c.log")
	if err != nil {
		log.Print(err.Error())
	}
	print(string(bytAll))
}
