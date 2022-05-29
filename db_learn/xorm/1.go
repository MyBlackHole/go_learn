package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

const (
	USERNAME     = "******"
	PASSWORD     = "8888888888"
	NETWORK      = "tcp"
	SERVER       = "8888888888888"
	PORT         = 3306
	DATABASE     = "urun_car"
	TIMEOUT      = 50
	READTIMEOUT  = 50
	WRITETIMEOUT = 50
)

type new_car_img_marked struct {
	ImgURL       string      `xorm:"'img_url'"`
	FdfsURL      string      `xorm:"'fdfs_url'"`
	ImgHash      string      `xorm:"'img_hash'"`
	ImgPpi       string      `xorm:"'img_ppi'"`
	Color        string      `xorm:"'color'"`
	Source       int         `xorm:"'source'"`
	DataID       int         `xorm:"'data_id'"`
	ImgErrStatus string      `xorm:"'img_err_status'"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?timeout=3s&parseTime=true&loc=Local&charset=utf8", USERNAME, PASSWORD, SERVER, PORT, DATABASE))
	if err != nil {
		fmt.Printf("%v", err.Error())
	} else {
		var item new_car_img_marked
		item.Source = 1
		var i int64
		i, err = engine.Insert(item)
		if err == nil {
			fmt.Println(i)
			return
		}
	}
}
