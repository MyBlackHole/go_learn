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
	engine, _ := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	err := engine.Sync2(&User{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sess := engine.NewSession()

	_, err = sess.Insert(&([]*User{&User{Name: "xhq", Age: 41}, &User{Name: "lhy", Age: 12}}))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var users []*User
	err = sess.Where("name = ?", "xhq").Find(&users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, user := range users {
		fmt.Printf("%+v\n", user)
	}

}
