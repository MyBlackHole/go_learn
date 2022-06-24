package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}

	affected, err := engine.Where("name = ?", "lzy").Delete(&User{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d records deleted\n", affected)
}
