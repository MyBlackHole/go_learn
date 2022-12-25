package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//实例化一个命令行程序
	oApp := cli.NewApp()
	//程序名称
	oApp.Name = "BlackHole"
	//程序的用途描述
	oApp.Usage = "cli 测试"
	//程序的版本号
	oApp.Version = "1.0.0"
	//该程序执行的代码
	oApp.Action = func(c *cli.Context) error {
		fmt.Println("Test")
		return nil
	}
	//启动
	if err := oApp.Run(os.Args); err != nil {
		log.Fatal(err)
	}
	/*
	 result:
	 [root@localhost cli]# go run main.go help

	 NAME:
	 GoTool - To save the world

	 USAGE:
	 main [global options] command [command options] [arguments...]

	 VERSION:
	 1.0.0

	 COMMANDS:
	 help, h Shows a list of commands or help for one command

	 GLOBAL OPTIONS:
	 --help, -h  show help
	 --version, -v print the version

	 [root@localhost cli]# go run main.go
	 Test
	*/
}
