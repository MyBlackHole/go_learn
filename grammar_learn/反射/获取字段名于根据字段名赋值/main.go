package main

import (
	"fmt"
	"reflect"
	"time"
)

type JtFinanceTransaction struct {
	Id              int64     `json:"id" xorm:"pk autoincr comment('auto_increment') BIGINT(20)"`
	UserId          int       `json:"user_id" xorm:"not null default 0 comment('用户ID') INT(11)"`
	AccountId       int64     `json:"account_id" xorm:"not null default 0 comment('jt_platform_account表ID') BIGINT(20)"`
	SiteId          string    `json:"site_id" xorm:"not null default '' comment('站点ID') CHAR(2)"`
	PostedDate      time.Time `json:"posted_date" xorm:"not null default '0000-00-00 00:00:00' comment('事件发生时间') TIMESTAMP"`
	PostedDateLocal time.Time `json:"posted_date_local" xorm:"not null default '0000-00-00 00:00:00' comment('事件发生时间的北京时间') TIMESTAMP"`
	Location        string    `json:"location" xorm:"not null default '' comment('时区') CHAR(15)"`
	TransactionType string    `json:"transaction_type" xorm:"not null default '' comment('事务类型') VARCHAR(50)"`
	TransactionId   string    `json:"transaction_id" xorm:"not null default '' comment('事务ID') CHAR(30)"`
	AmazonOrderId   string    `json:"amazon_order_id" xorm:"not null default '' comment('亚马逊订单ID') VARCHAR(30)"`
	SaleOrderId     int       `json:"sale_order_id" xorm:"not null default 0 comment('订单ID') INT(11)"`
	Sku             string    `json:"sku" xorm:"not null default '' comment('sku') VARCHAR(50)"`
	SaleOrderItemId int       `json:"sale_order_item_id" xorm:"not null default 0 comment('jt_sale_order_items表ID') INT(11)"`
	FeeType         string    `json:"fee_type" xorm:"not null default '' comment('费用类型') VARCHAR(255)"`
	CurrencyCode    string    `json:"currency_code" xorm:"not null default '' comment('货币代码') CHAR(3)"`
	ExchangeRate    string    `json:"exchange_rate" xorm:"not null default 0.000000 comment('汇率') DECIMAL(11,6)"`
	Amount          string    `json:"amount" xorm:"not null default 0.000000 comment('金额') DECIMAL(11,6)"`
	CreateTime      time.Time `json:"create_time" xorm:"not null default '0000-00-00 00:00:00' comment('创建时间') TIMESTAMP"`
	LastUpdateTime  time.Time `json:"last_update_time" xorm:"not null default 'CURRENT_TIMESTAMP' comment('最后更新时间') TIMESTAMP"`
}

func main() {
	typ := reflect.TypeOf(JtFinanceTransaction{})
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		fmt.Println("结构体里的字段名", field.Name)
	}

	// var financeTransaction JtFinanceTransaction
	// v := reflect.ValueOf(&financeTransaction)
	// fmt.Println(v.Elem().FieldByName("Amount").Kind())
	// var s string
	// s = "aa"
	// v.Elem().FieldByName("Amount").Set(reflect.ValueOf(s))
	// fmt.Printf("%+v", financeTransaction)

	financeTransaction := new(JtFinanceTransaction)
	v := reflect.ValueOf(financeTransaction)
	fmt.Println(v.Elem().FieldByName("Amount").Kind())
	var s string
	s = "aa"
	v.Elem().FieldByName("Amount").Set(reflect.ValueOf(s))
	fmt.Printf("%+v", financeTransaction)

	// for i := 0; i < v.NumField(); i++ {
	// 	field := v.Field(i)
	// }
}
