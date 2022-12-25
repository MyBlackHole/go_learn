package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	XDGCONFIGHOME := os.Getenv("XDG_CONFIG_HOME")
	var PACKAGE string = path.Join(XDGCONFIGHOME, "cn.com.urundata.shadowsocks", "config.json")
	fmt.Println(PACKAGE)

	fileName := "./1.go"
	// 是否绝对路径
	fmt.Println(filepath.IsAbs(fileName))
	// 获取绝对路径
	fmt.Println(filepath.Abs(fileName))

	// 获取父目录
	fmt.Println(path.Join(fileName, ".."))
}
