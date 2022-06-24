package main

import (
	"io/ioutil"
	"log"
)



func main()  {
	bytAll, err := ioutil.ReadFile("c.log")
	if err != nil {
		log.Print(err.Error())
	}
	print(string(bytAll))
}