package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	ss "github.com/shadowsocks/shadowsocks-go/shadowsocks"
)

func getMd5(str string) (ret string) {
	data := []byte(str)
	has := md5.Sum(data)
	ret = fmt.Sprintf("%x", has) //将[]byte转成16进制
	return
}

type ConfigList struct {
	Configs []ss.Config `json:"configs"`
}

func getConfig(ch chan string, name string) (configList ConfigList, err error) {
	params := url.Values{}
	Url, err := url.Parse("https://q60py.91rong.com.cn:38010/ProxyConfig/WinFormProxyConfig/ProxyServers")
	if err != nil {
		return
	}
	params.Set("uname", name)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)
	if err != nil {
		ch <- err.Error()
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	_ = json.Unmarshal(body, &configList)
	return
}

type result struct {
	StatusCode int    `json:"statusCode"`
	LoginState string `json:"LoginState"`
	UserId     int    `json:"UserId"`
}

func login(ch chan string, name string, pass string, ukey string) (statusCode int, err error) {
	params := url.Values{}
	Url, err := url.Parse("https://q60py.91rong.com.cn:38010/ProxyConfigUKey/WinFormProxyConfig/ProxyServers")

	if err != nil {
		return
	}

	pass = getMd5(pass)
	params.Set("username", name)
	params.Set("passwd", pass)
	params.Set("ukey", ukey)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()

	resp, err := http.Get(urlPath)

	if err != nil {
		ch <- err.Error()
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res result
	_ = json.Unmarshal(body, &res)
	statusCode = res.StatusCode
	return
}

func mainLoginVPN(ch chan string, name string, pass string) {
	// ukey := get_ukey()

	// statusCode, err := login(ch, name, pass, ukey)
	// if err != nil {
	// 	ch <- err.Error()
	//	return
	// }
	// fmt.Println(statusCode)

	configList, err := getConfig(ch, name)
	if err != nil {
		ch <- err.Error()
	}

	fmt.Println(configList)
	runSS(configList.Configs[2])
}
