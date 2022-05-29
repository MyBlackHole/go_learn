package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main1() {
	ls := exec.Command("ls")
	out, _ := ls.CombinedOutput()
	fmt.Printf("%s\n", out)
	// ls.Run()
}

func main() {
	fmt.Println(os.Args[0])
	fmt.Println(len(os.Args))
	// ls.Run()
}
