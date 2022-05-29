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
	user := &User{
		Name: "lzy",
		Age:  50,
	}

	affected, err := engine.Insert(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)

	users := make([]*User, 2)
	users[0] = &User{Name: "xhq", Age: 41}
	users[1] = &User{Name: "lhy", Age: 12}

	affected, err = engine.Insert(&users)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d records inserted, id1:%d, id2:%d\n", affected, users[0].Id, users[1].Id)
}
