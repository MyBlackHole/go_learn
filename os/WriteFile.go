package main

import (
	"os"
)

func main() {
	// ioutil.WriteFile("test.txt", []byte(`ok`), 0755)
    os.WriteFile("test.txt", []byte(`ok`), 0755)
}
