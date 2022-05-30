package main

import "strings"
import "bufio"
import "fmt"
import "io"

func main() {

	s := strings.NewReader("a\r\nb")
	r := bufio.NewReader(s)

	for {
		token, _, err := r.ReadLine()
		if len(token) > 0 {
			fmt.Printf("Token (ReadLine): %q\n", token)
		}
		if err == io.EOF {
			fmt.Println(err.Error())
		}
		if err != nil {
			break
		}
	}
	s.Seek(0, io.SeekStart)
	r.Reset(s)

	for {
		token, err := r.ReadBytes('\n')
		fmt.Printf("Token (ReadBytes): %q\n", token)
		if err == io.EOF {
			fmt.Println(err.Error())
		}
		if err != nil {
			break
		}
	}

	s.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(s)
	for scanner.Scan() {
		fmt.Printf("Token (Scanner): %q\n", scanner.Text())
	}
}
