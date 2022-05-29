package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "GoTool"
	//程序的用途描述
	oApp.Usage = "To save the world"
	//程序的版本号
	oApp.Version = "1.0.0"

	//设置多个命令处理函数
	oApp.Commands = []cli.Command{
		{
			//命令全称
			Name: "lang",
			//命令简写
			Aliases: []string{"l"},
			//命令详细描述
			Usage: "Setting language",
			//命令处理函数
			Action: func(c *cli.Context) {
				// 通过c.Args().First()获取命令行参数
				fmt.Printf("language=%v \n", c.Args().First())
			},
		},
		{
			Name:    "encode",
			Aliases: []string{"e"},
			Usage:   "Setting encoding",
			Action: func(c *cli.Context) {
				fmt.Printf("encoding=%v \n", c.Args().First())
			},
		},
	}

	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}

	/*
	 [root@localhost cli]# go run main.go l english
	 language=english

	 [root@localhost cli]# go run main.go e utf8
	 encoding=utf8

	 [root@localhost cli]# go run main.go help
	 NAME:
	 GoTool - To save the world

	 USAGE:
	 main [global options] command [command options] [arguments...]

	 VERSION:
	 1.0.0

	 COMMANDS:
	 lang, l Setting language
	 encode, e Setting encoding
	 help, h Shows a list of commands or help for one command

	 GLOBAL OPTIONS:
	 --help, -h  show help
	 --version, -v print the version
	*/
}
