package main

import "io/ioutil"

func main() {
	ioutil.WriteFile("test.txt", []byte(`ok`), 0755)
}
