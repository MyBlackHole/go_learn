package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
	_ "gopkg.in/mgo.v2/bson"
)

var Session *mgo.Session
var Databases *mgo.Database
var mongoErr error

const (
	MONGO_HOST = "127.0.0.1"
	MONGO_PORT = "27017"
	MONGO_DB   = "URun_time_judgment_log"
	MONGO_USER = ""
	MONGO_PWD  = ""
)

func init() {
	// 创建链接
	Session, mongoErr = mgo.Dial(fmt.Sprintf("%s:%s", MONGO_HOST, MONGO_PORT))
	if mongoErr != nil {
		fmt.Println("链接失败！")
		os.Exit(1)
	}
	// 选择DB
	Databases = Session.DB(MONGO_DB)

	if MONGO_USER == MONGO_PWD && MONGO_PWD == "" {
		return
	}

	// 登陆
	mongoErr = Databases.Login(MONGO_USER, MONGO_PWD)

	if mongoErr != nil {
	    log.Panic("登陆验证失败！")
	}

	// collection = Databases.C(c)
	// defer Session.Close()
}
