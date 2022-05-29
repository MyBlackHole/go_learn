package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
    _ "gopkg.in/mgo.v2/bson"
    "os"
)

var Session *mgo.Session
var Databases *mgo.Database
var MgoError error

const (
    MONGO_HOST  = "172.17.0.4"
    MONGO_PORT  = "27017"
    MONGO_DB    = ""
    MONGO_USER  = ""
    MONGO_PWD   = ""
)

func init() {
    // 创建链接
    Session, MgoError = mgo.Dial(fmt.Sprintf("%s:%s", MONGO_HOST, MONGO_PORT))
    if MgoError != nil {
        fmt.Println("链接失败！")
        os.Exit(1)
    }
    // 选择DB
    Databases = Session.DB(MONGO_DB)
    // // 登陆
    // MgoError = Databases.Login(MONGO_USER, MONGO_PWD)
    // if MgoError != nil {
    //     fmt.Println("登陆验证失败！")
    //     os.Exit(1)
    // }
    // defer Session.Close()
}

func Demos() {
    // 选择一个要操作的Collection
    c := Databases.C("demo")
    // count
    fmt.Println(c.Count())
    // insert
    c.Insert(map[string]string{"name": "马超"})
    // update
    c.Update(map[string]string{"name": "马超"}, map[string]string{"name": "黄忠"})
    // remove
    c.Remove(map[string]string{"name": "马超"})
    // find - count
    num, MgoError := c.Find(map[string]string{"name": "黄忠"}).Count()
    if MgoError != nil {
        fmt.Println(MgoError.Error())
    } else {
        fmt.Println(num)
    }
    // find - one
    var one map[string]interface{}
    MgoError = c.Find(map[string]string{"name": "黄忠"}).One(&one)
    if MgoError != nil {
        fmt.Println(MgoError.Error())
    } else {
        fmt.Println(one)
    }
    // find - all
    var all []map[string]interface{}
    MgoError = c.Find(map[string]string{"name": "黄忠"}).All(&all)
    if MgoError != nil {
        fmt.Println(MgoError.Error())
    } else {
        fmt.Println(all)
    }
}


func main()  {
	Demos()
}