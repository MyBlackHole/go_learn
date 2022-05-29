package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Player struct {
	Id        int64
	Name      string
	Age       int
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")

	// 自动检测和创建表，这个检测是根据表的名字
	// 自动检测和新增表中的字段，这个检测是根据字段名，同时对表中多余的字段给出警告信息
	// 自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称。因此这里需要注意，如果在一个有大量数据的表中引入新的索引，数据库可能需要一定的时间来建立索引。
	// 自动转换varchar字段类型到text字段类型，自动警告其它字段类型在模型和数据库之间不一致的情况。
	// 自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	engine.Sync2(&Player{})

	// 插入数据
	engine.Insert(&Player{Name: "dj", Age: 18})

	p1 := &Player{}
	engine.Where("name = ?", "dj").Get(p1)
	fmt.Println("after insert:", p1)
	time.Sleep(5 * time.Second)

	engine.Table(&Player{}).ID(p1.Id).Update(map[string]interface{}{"age": 30})

	p2 := &Player{}
	engine.Where("name = ?", "dj").Get(p2)
	fmt.Println("after update:", p2)
	time.Sleep(5 * time.Second)

	engine.ID(p1.Id).Delete(&Player{})

	p3 := &Player{}
	engine.Where("name = ?", "dj").Unscoped().Get(p3)
	fmt.Println("after delete:", p3)
}
