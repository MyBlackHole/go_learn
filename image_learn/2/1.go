package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"net/http"
)

func main() {
	resp, err := http.Get("http://i.imgur.com/Peq1U1u.jpg")
	if err != nil {
		return // handle error somehow
	}
	defer resp.Body.Close()
	img, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return // handle error somehow
	}
	fmt.Println(img.Width*img.Height, "pixels")
}
