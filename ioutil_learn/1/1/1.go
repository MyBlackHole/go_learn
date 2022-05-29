package main

import (
	"fmt"
	"os"
)

func FileExist(path string) bool {
  _, err := os.Lstat(path)
  return !os.IsNotExist(err)
}

func main()  {
	fmt.Print(FileExist("2.go"))
}
