package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	XDGCONFIGHOME := os.Getenv("XDG_CONFIG_HOME")
	var PACKAGE string = path.Join(XDGCONFIGHOME, "cn.com.urundata.shadowsocks", "config.json")
	fmt.Println(PACKAGE)
}
