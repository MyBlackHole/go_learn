package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
)

func main() {
	// resp, err := http.Get("http://i.imgur.com/Peq1U1u.jpg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer resp.Body.Close()

	bytAll, err := ioutil.ReadFile("1.jpg")
	if err != nil {
		fmt.Println(err)
	}

	m, name, err := image.Decode(bytes.NewReader(bytAll))
	if err != nil {
		log.Fatal(err)
	}
	g := m.Bounds()
	// Get height and width
	height := g.Dy()
	width := g.Dx()
	// The resolution is height x width
	resolution := height * width
	// Print results
	fmt.Println(resolution, "pixels")
	fmt.Printf("%dx%d -- %s\n", width, height, name)
}


func GetImgWH(filename string) (int, int, error) {
	bytAll, err := ioutil.ReadFile("1.jpg")
	if err != nil {
		return 0, 0, err
	}

	m, _, err := image.Decode(bytes.NewReader(bytAll))
	if err != nil {
		log.Fatal(err)
	}
	g := m.Bounds()
	return g.Dx(), g.Dy(), nil
}