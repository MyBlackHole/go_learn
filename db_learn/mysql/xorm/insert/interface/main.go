package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Test struct {
	A string
	B string
}

type INTERFACE interface{}

func (c *Test) FromDB(bytes []byte) error {
	return json.Unmarshal(bytes, c)
}

func (c *Test) ToDB() (bytes []byte, err error) {
	bytes, err = json.Marshal(c)
	return
}

type User struct {
	Id    int64     `json:"id" xorm:"pk autoincr comment('auto_increment') BIGINT(20)"`
	Value INTERFACE `json:"value"  xorm:"not null comment('任务内容') json"`
}

// type UserTest struct {
// 	Id    int64 `json:"id" xorm:"pk autoincr comment('auto_increment') BIGINT(20)"`
// 	Value Test  `json:"value"  xorm:"not null comment('任务内容') json"`
// }

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")
	err := engine.Sync2(&User{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user := &User{}
	// user.Value = ""
	user.Value = Test{A: "a", B: "b"}

	affected, err := engine.Insert(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)
}
