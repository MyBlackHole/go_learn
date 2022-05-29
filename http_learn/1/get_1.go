package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getMd5(str string) (ret string) {
	data := []byte(str)
	has := md5.Sum(data)
	ret = fmt.Sprintf("%x", has) //将[]byte转成16进制
	//fmt.Println(ret)
	return
}

type Config struct {
	Server       interface{} `json:"server"`
	ServerPort   int         `json:"server_port"`
	LocalPort    int         `json:"local_port"`
	LocalAddress string      `json:"local_address"`
	Password     string      `json:"password"`
	Method       string      `json:"method"`
}

func (c *Config) UnmarshalJSON(data []byte) error {
	type config Config
	test := &config{
		LocalPort:    1080,
		LocalAddress: "127.0.0.1",
	}
	_ = json.Unmarshal(data, test)
	*c = Config(*test)
	return nil
}

type ConfigList struct {
	Configs []Config `json:"configs"`
}

func getConfig(name string) (configList ConfigList, err error) {
	params := url.Values{}
	Url, err := url.Parse("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		return
	}
	params.Set("uname", name)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//_json = string(body)
	_ = json.Unmarshal(body, &configList)
	return
}

type result struct {
	StatusCode int    `json:"statusCode"`
	LoginState string `json:"LoginState"`
	UserId     int    `json:"UserId"`
}

func login(name string, pass string, ukey string) (statusCode int, err error) {
	params := url.Values{}
	Url, err := url.Parse("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	if err != nil {
		return
	}
	pass = getMd5(pass)
	params.Set("username", name)
	params.Set("passwd", pass)
	params.Set("ukey", ukey)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	fmt.Println(urlPath)
	resp, err := http.Get(urlPath)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	var res result
	_ = json.Unmarshal(body, &res)
	statusCode = res.StatusCode
	return
}
func main() {
	//ret := getMd5("333333333333")
	//fmt.Println(ret)

	//statusCode, err := login("333333", "333333333333", "3333333333333333333333333333333")
	//if err != nil {
	//	return
	//}
	//fmt.Println(statusCode)
	configList, err := getConfig("jstest")
	if err != nil {
		return
	}
	fmt.Println(configList)
}
