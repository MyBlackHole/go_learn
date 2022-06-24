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

	querySql := "select * from user limit 1"
	reuslts, err := engine.Query(querySql)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, record := range reuslts {
		for key, val := range record {
			fmt.Println(key, string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, err := engine.Exec(updateSql, "dj", 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res.RowsAffected())
}
