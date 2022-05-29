package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.baidu.com/"
	// url := "http://product.auto.163.com/picture/ckindex/4F6K0008/000CBYNW.html"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	// req.Header.Add("Accept-Encoding", "gzip")
	// req.Header.Add("X-SS-REQ-TICKET", "1611635376309")
	// req.Header.Add("passport-sdk-version", "30")
	// req.Header.Add("sdk-version", "2")
	// req.Header.Add("Cookie", "qh[360]=1; odin_tt=1cacf31943c52e803d36b6c29c33efcb32d9179ec0b850f3508c03cf88ce4ccd5d94daf962268b0a94fe59bc19dd69b5525c0d67fd63aadb7908c397af05172f; install_id=2058700123549015; ttreq=1$7857f888dfce4b0fb75057e506d3896fd65de16e; WIN_WH=481_694; PIXIEL_RATIO=1.4562500715255737; FRM=new")
	// req.Header.Add("X-Tyhon", "DRktIW86IyBCTjZrCW0/aGpEByJ6SiEQZn1ZEvU=")
	// req.Header.Add("X-Khronos", "1611635376")
	// req.Header.Add("X-Gorgon", "0404809d4001ae5ca252cb9a89bebe3d057da9c75d01ecb026e1")
	// req.Header.Add("Host", "product.auto.163.com")
	// req.Header.Add("Connection", "Keep-Alive")
	// req.Header.Add("User-Agent", "okhttp/3.10.0.1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	ioutil.WriteFile("1.html", body, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.StatusCode)
}
