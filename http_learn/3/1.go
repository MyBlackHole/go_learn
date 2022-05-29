package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func main() {
	reg, _ := regexp.Compile("您的IP地址是：\\[.+?\\]")

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://202.105.233.40:80") //根据定义Proxy func(*Request) (*url.URL, error)这里要返回url.URL
	}

	transport := &http.Transport{Proxy: proxy}
	client := &http.Client{Transport: transport}

	resp, err := client.Get("http://ip168.com/") //请求并获取到对象,使用代理
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get("http://ip168.com/")   //请求并获取到对象
	dataproxy, err := ioutil.ReadAll(resp.Body) //取出主体的内容
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body) //取出主体的内容
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s",data) //打印
	sproxy := reg.FindString(string(dataproxy))
	s := reg.FindString(string(data))
	res.Body.Close()
	resp.Body.Close()
	fmt.Printf("不使用代理:%s", s)     //打印
	fmt.Printf("使用代理:%s", sproxy) //打印
}
