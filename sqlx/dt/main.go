package main

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	userName = "test"
	password = "P@ssw0rd"
	ip       = "192.168.90.207"
	port     = "8888"
	dbName   = "TEST"
)

// // mysql
// const (
// 	userName = "black"
// 	password = "1358"
// 	ip       = "127.0.0.1"
// 	port     = "3306"
// 	dbName   = "test_xa"
// )

// Db数据库连接池
var DB *sql.DB

type User struct {
	id    int64
	score int8
	name  string
}

func init() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		panic("open database fail")
	}
	fmt.Println("connnect success")
}

func TeleDBDTInsertUser() error {
	var err error
	var sql string
	var cxt context.Context
	cxt = context.Background()

	conn, err := DB.Conn(cxt)
	if err != nil {
		return err
	}
	defer conn.Close()

	sql = fmt.Sprintf("UDAL DT START;")
	println(sql)
	_, err = conn.ExecContext(cxt, sql)
	if err != nil {
		println(err.Error())
		return err
	}

	for i := 0; i < 10; i++ {
		sql = fmt.Sprintf("insert into user (score, name) values(RAND() * 100, 'test1');")
		println(sql)
		_, err = conn.ExecContext(cxt, sql)
		if err != nil {
			println(err.Error())
			return err
		}
	}

	// sql = fmt.Sprintf("insert into user (score, name) values(RAND() * 100, 'test1');")
	// println(sql)
	// _, err = conn.ExecContext(cxt, sql)
	// if err != nil {
	// 	println(err.Error())
	// 	return err
	// }

	sql = fmt.Sprintf("COMMIT;")
	println(sql)
	_, err = conn.ExecContext(cxt, sql)
	if err != nil {
		println(err.Error())
		return err
	}

	return err
}

func main() {
	defer DB.Close()

	var concurrency = 1
	var count = 100

	var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(count int) {
			defer waitGroup.Done()
			for i := 0; i < count; i++ {
				TeleDBDTInsertUser()
			}
		}(count)
	}

	waitGroup.Wait()
}
