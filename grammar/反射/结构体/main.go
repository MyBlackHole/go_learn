package main

import (
	"reflect"
)

type FromData struct {
	ProductId      int64
	CellId         int64
	ProductName    string
	ProductSku     string
	CellName       string
	QuantityLeft   int
	CheckInventory int
	Remark         string
}

// 获取结构体值列表
func (this *FromData) ToList() (list []interface{}) {
	retThis := reflect.ValueOf(this)
	if retThis.Kind() == reflect.Ptr {
		retThis = retThis.Elem()
		for i := 0; i < retThis.NumField(); i++ {
			list = append(list, retThis.Field(i).Interface())
		}
	}
	return
}

// 根据列表填充值
func (this *FromData) SetList(list []interface{}) {
	retHello := reflect.ValueOf(this)
	if retHello.Kind() == reflect.Ptr {
		retHello = retHello.Elem()
		for i := 0; i < retHello.NumField(); i++ {
			value := list[i]
			if value != nil && reflect.ValueOf(value).Kind() == retHello.Field(i).Kind() {
				retHello.Field(i).Set(reflect.ValueOf(value))
			}
		}
	}
}

func (this *FromData) DeepEqual(dst *FromData) bool {
	return reflect.DeepEqual(this, dst)
}
