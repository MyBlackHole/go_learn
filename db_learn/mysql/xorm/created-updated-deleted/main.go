package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type A struct {
	ProductId int64 `json:"product_id" xorm:"not null default 0 comment('商品 id') INT(11)"`
}

type Model struct {
	Id                  int64   `json:"id" xorm:"pk autoincr comment('auto_increment') BIGINT(20)"`
	ProductId           int64   `json:"product_id" xorm:"not null default 0 comment('商品 id') INT(11)"`
	ProductMainImageUrl string  `json:"product_main_image_url" xorm:"not null default '' comment('商品图片') VARCHAR(255)"`
	ProductSku          string  `json:"sku" xorm:"not null default '' comment('商品 sku') VARCHAR(255)"`
	ProductNameCn       string  `json:"product_name_cn" xorm:"not null default '' comment('商品中文名') VARCHAR(255)"`
	ProductSpec         string  `json:"product_spec" xorm:"not null default '' comment('商品规格') VARCHAR(255)"`
	OperatorId          int64   `json:"operator_id" xorm:"not null default 0 comment('商品创建人 id') INT(11)"`
	OperatorName        string  `json:"operator_name" xorm:"not null default '' comment('商品创建人名') VARCHAR(255)"`
	PurchaseId          int64   `json:"purchase_id" xorm:"not null default 0 comment('最后一次采购单的采购员 id') INT(11)"`
	PurchaseUserName    string  `json:"purchase_user_name" xorm:"not null default '' comment('该商品最后采购单的采购员名') VARCHAR(255)"`
	AccountId           int64   `json:"account_id" xorm:"not null default 0 comment('最后一次销售订单的店铺 id') INT(11)"`
	AccountName         string  `json:"account_name" xorm:"not null default '' comment('该商品最后销售订单的店铺名') VARCHAR(255)"`
	WarehouseId         int64   `json:"warehouse_id" xorm:"not null default 0 comment('仓库 id') INT(11)"`
	WarehouseName       string  `json:"warehouse_name" xorm:"not null default '' comment('仓库名') VARCHAR(255)"`
	InventoryList       []*A    `json:"inventory_list" xorm:"comment('库位 id 与对应库存数 列表') JSON"`
	QuantityTotal       int     `json:"quantity_total" xorm:"not null default 0 comment('库存总数量') INT(11)"`
	ProductAveragePrice string  `json:"product_average_price" xorm:"not null default '' comment('平均采购价') VARCHAR(255)"`
	UnsalablePriceTotal float64 `json:"unsalable_price_total" xorm:"comment('滞销总金额') FLOAT"`
	OutOrderState       bool    `json:"out_order_state" xorm:"not null default 0 comment('出货状态') TINYINT(1)"`
	UnsalableDay        int64   `json:"unsalable_day" xorm:"comment('滞销天数') INT(11)"`
	Sales7Day           int64   `json:"sales_7day" xorm:"sales_7day comment('近 7 天销量') INT(11)"`
	Sales15Day          int64   `json:"sales_15day" xorm:"sales_15day comment('近 15 天销量') INT(11)"`
	Sales30Day          int64   `json:"sales_30day" xorm:"sales_30day comment('近 30 天销量') INT(11)"`
	Sales60Day          int64   `json:"sales_60day" xorm:"sales_60day comment('近 60 天销量') INT(11)"`
	Sales90Day          int64   `json:"sales_90day" xorm:"sales_90day comment('近 90 天销量') INT(11)"`
	Sales180Day         int64   `json:"sales_180day" xorm:"sales_180day comment('近 180 天销量') INT(11)"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:123456@/test?charset=utf8")

	// 自动检测和创建表，这个检测是根据表的名字
	// 自动检测和新增表中的字段，这个检测是根据字段名，同时对表中多余的字段给出警告信息
	// 自动检测，创建和删除索引和唯一索引，这个检测是根据索引的一个或多个字段名，而不根据索引名称。因此这里需要注意，如果在一个有大量数据的表中引入新的索引，数据库可能需要一定的时间来建立索引。
	// 自动转换varchar字段类型到text字段类型，自动警告其它字段类型在模型和数据库之间不一致的情况。
	// 自动警告字段的默认值，是否为空信息在模型和数据库之间不匹配的情况
	err := engine.Sync2(&Model{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 插入数据
	ret, err := engine.Insert(&Model{ProductId: 100, InventoryList: []*A{&A{ProductId: 10000}}})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ret)

	p1 := &Model{}
	engine.Where("product_id = ?", 100).Get(p1)
	fmt.Printf("after insert:%v", p1.InventoryList[0])
	time.Sleep(5 * time.Second)

	// engine.Table(&Model{}).ID(p1.Id).Update(map[string]interface{}{"age": 30})

	// p2 := &Player{}
	// engine.Where("name = ?", "dj").Get(p2)
	// fmt.Println("after update:", p2)
	// time.Sleep(5 * time.Second)

	// engine.ID(p1.Id).Delete(&Player{})

	// p3 := &Player{}
	// engine.Where("name = ?", "dj").Unscoped().Get(p3)
	// fmt.Println("after delete:", p3)
}
