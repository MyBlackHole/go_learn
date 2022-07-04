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

	engine.Sync2(&User{})

	if err != nil {
		fmt.Println(err.Error())
	}

	// 插入数据
	engine.Insert(&User{Name: "dj", Age: 18})

	// 更新为 0
	// affected, err := engine.Id(id).Cols("age").Update(&user)

	engine.Update(&User{Name: "ldj"})
	// engine.ID(1).Update(&User{Name: "ldj"})
	// engine.ID(2).Cols("name", "age").Update(&User{Name: "dj"})

	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 18})
}
