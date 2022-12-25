package main

import (
	"io/ioutil"
	"log"
	"regexp"
)

func main() {
	bytAll, err := ioutil.ReadFile("c.log")
	if err != nil {
		log.Print(err.Error())
	}
	text := string(bytAll)

	r, err := regexp.Compile("http.*")
	if err != nil {
		log.Print(err.Error())
	}
	strList := r.FindAllString(text, -1)
	for _, v := range strList {
		log.Print(" --- " + v + "\n")
	}
}
