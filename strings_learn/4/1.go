package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "http://img.12365auto.com/g/202011/17/202011171245551519_sst.jpg"
	url = strings.ReplaceAll(url, "http://img.12365auto.com/g/", "http://img.12365auto.com/g/temp/")
	url = strings.ReplaceAll(url, "_sst", "")
	fmt.Printf("%s", url)
}

