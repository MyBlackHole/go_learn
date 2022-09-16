package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //执行一些初始化操作
	"github.com/jmoiron/sqlx"
)

func main() {

	//1 链接方式一
	DB, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		fmt.Println("链接出错：", err)
	}
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		fmt.Println("通信出错", err)
	}
	fmt.Printf("ok\n")
	// 2 链接方式二
	//DB:=sqlx.MustOpen("mysql","root:123@tcp(127.0.0.1:3306)/lqz?charset=utf8")
	//defer DB.Close()
	//err:=DB.Ping()
	//if err!=nil {
	//	fmt.Println("通信出错",err)
	//}

}
