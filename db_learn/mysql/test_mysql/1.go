package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID  int64
	A1  string
	A2  string
	A3  string
	A4  string
	A5  string
	A6  string
	A7  string
	A8  string
	A9  string
	A10 string
}

type TimeTmp struct {
	Time int64  `json:"time"`
	Err  string `json:"err"`
	Code string `json:"code"`
	Resp string `json:"resp"`
}

type Resp struct {
	Err      string    `json:"err"`
	TimeSum  int64     `json:"timeSum"`
	TimeMax  int64     `json:"timeMax"`
	TimeMin  int64     `json:"timeMin"`
	// TimeMin  int64     `json:"timeMin"`
	// DataList []TimeTmp `json:"dataList"`
}

type MysqlInfo struct {
	Username string `json:"userNmae"`
	Password string `json:"passWord"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"dataBase"`
}

type Config struct {
	MysqlInfo MysqlInfo `json:"mysqlInfo"`
	Count     int       `json:"count"`
	DelTable  bool      `json:"delTable"`
}

func RunTime(userList []User) []TimeTmp {
	var timeTmpList []TimeTmp
	for _, item := range userList {
		var timeTmp TimeTmp
		timeTmp.Code = fmt.Sprintf("%v", item)

		t := time.Now()
		_, err := engine.Insert(item)
		if err != nil {
			timeTmp.Err = err.Error()
		}
		duration := time.Since(t)
		timeTmp.Time = duration.Nanoseconds()
		timeTmpList = append(timeTmpList, timeTmp)
	}
	return timeTmpList
}

func structByReflect(inStructPtr interface{}) {
	rType := reflect.TypeOf(inStructPtr)
	rVal := reflect.ValueOf(inStructPtr)
	if rType.Kind() == reflect.Ptr {
		// 传入的inStructPtr是指针，需要.Elem()取得指针指向的value
		rType = rType.Elem()
		rVal = rVal.Elem()
	} else {
		panic("inStructPtr must be ptr to struct")
	}

	// 遍历结构体
	for i := 0; i < rType.NumField(); i++ {
		name := rType.Field(i).Name
		if name == "ID" {
			continue
		}
		f := rVal.Field(i)
		a := rand.Intn(100)
		f.Set(reflect.ValueOf(GetRandomString2(a)))
	}
}

func NewData(count int) []User {
	var userList []User
	for ; count > 0; count-- {
		user := new(User)
		structByReflect(user)
		userList = append(userList, *user)
	}
	return userList
}

var engine *xorm.Engine

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomString2(n int) string {
	randBytes := make([]byte, n/2)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func main() {

	// 配置
	var config Config

	// 返回值
	var resp Resp

	bytAll, err := ioutil.ReadFile("config.json")
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	err = json.Unmarshal(bytAll, &config)
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	userList := NewData(config.Count)

	engine, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=3s&parseTime=true&loc=Local&charset=utf8", config.MysqlInfo.Username, config.MysqlInfo.Password, config.MysqlInfo.Host, config.MysqlInfo.Port, config.MysqlInfo.Database))
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	err = engine.Ping()
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	err = engine.CreateTables(new(User)) // 表创建
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	DataList := RunTime(userList)

	bytAll, err = json.Marshal(DataList)
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", DataList)
		return
	}
	err = ioutil.WriteFile("DataList.json", bytAll, 0755)
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", bytAll)
		return
	}

	resp.TimeMin = 99999999999999
	for _, item := range DataList {
		if item.Time > resp.TimeMax {
			resp.TimeMax = item.Time
		}
		if item.Time < resp.TimeMin {
			resp.TimeMin = item.Time
		}
		resp.TimeSum += item.Time
	}

	bytAll, err = json.Marshal(resp)
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}


	err = ioutil.WriteFile("resp.json", bytAll, 0755)
	if err != nil {
		resp.Err = err.Error()
		fmt.Printf("%v", resp)
		return
	}

	if config.DelTable {
		err = engine.DropTables(new(User)) // 删除表
		if err != nil {
			resp.Err = err.Error()
			fmt.Printf("%v", resp)
			return
		}
	}

	defer engine.Close() // 退出后关闭
}
