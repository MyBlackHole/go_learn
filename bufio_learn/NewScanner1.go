// 接收用户输入的账号密码
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(Login())
}

func Login() (map[string]string, map[string]string) {
	user := map[string]string{}
	pass := map[string]string{}
	fmt.Println("欢迎光临!")

	fmt.Println("请输入用户名:")
	LoginUser := bufio.NewScanner(os.Stdin)
	if LoginUser.Scan() {
		user["userName"] = LoginUser.Text()
	} else {
		user["userName"] = "输入错误"
	}

	fmt.Println("请输入登录密码:")
	LoginPassword := bufio.NewScanner(os.Stdin)
	if LoginPassword.Scan() {
		pass["PassWord"] = LoginPassword.Text()
	} else {
		pass["PassWord"] = "您输入错误"
	}

	return user,pass
}
