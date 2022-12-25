package main

// import (
// 	"bytes"
// 	"database/sql"
// 	"fmt"
// 	"strings"
// 	"time"

// 	_ "github.com/go-sql-drivkjer/mysql"
// )

// const (
// 	USERNAME     = "yunrun"
// 	PASSWORD     = "yr2020!QAZ"
// 	NETWORK      = "tcp"
// 	SERVER       = "192.168.1.219"
// 	PORT         = 3384
// 	DATABASE     = "urun_car"
// 	TIMEOUT      = 5
// 	READTIMEOUT  = 5
// 	WRITETIMEOUT = 5
// )

// // const (
// // 	USERNAME     = "root"
// // 	PASSWORD     = "123456"
// // 	NETWORK      = "tcp"
// // 	SERVER       = "127.0.0.1"
// // 	PORT         = 3306
// // 	DATABASE     = "urun_car"
// // 	TIMEOUT      = 5
// // 	READTIMEOUT  = 5
// // 	WRITETIMEOUT = 5
// // )

// var MysqlDB *sql.DB

// func init() {
// 	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?timeout=%ds&readTimeout=%ds&writeTimeout=%ds&charset=utf8", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE, TIMEOUT, READTIMEOUT, WRITETIMEOUT)
// 	var err error
// 	MysqlDB, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		fmt.Print(err)
// 		return
// 	}
// 	MysqlDB.SetConnMaxLifetime(100 * time.Second)
// }

// func getIBBMS(brand_pro string) (carList []Car) {
// 	car := new(Car)

// 	sql := fmt.Sprintf("select ID, brand_pro, Brand, Model, style, source from new_car where brand_pro ='%s'", brand_pro)
// 	rows, err := MysqlDB.Query(sql)
// 	defer func() {
// 		if rows != nil {
// 			rows.Close()
// 		}
// 	}()
// 	if err != nil {
// 		fmt.Printf("Insert failed,err:%v", err)
// 		return
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&car.ID, &car.BrandPro, &car.Brand, &car.Model, &car.Style, &car.Source)

// 		car.Brand = strings.ReplaceAll(car.Brand, " ", "")
// 		car.Model = strings.ReplaceAll(car.Model, " ", "")
// 		car.Style = strings.ReplaceAll(car.Style, " ", "")

// 		if err != nil {
// 			fmt.Printf("Scan failed,err:%v", err)
// 			return
// 		}
// 		// fmt.Print(*car)
// 		carList = append(carList, *car)
// 	}
// 	return
// }

// func getGB() (brandList []string) {
// 	rows, err := MysqlDB.Query("select brand_pro from new_car group by brand_pro")
// 	defer func() {
// 		if rows != nil {
// 			rows.Close()
// 		}
// 	}()
// 	if err != nil {
// 		fmt.Printf("Insert failed,err:%v", err)
// 		return
// 	}

// 	var brand string

// 	for rows.Next() {
// 		err = rows.Scan(&brand)
// 		if err != nil {
// 			fmt.Printf("Scan failed,err:%v", err)
// 			return
// 		}
// 		// fmt.Println(brand)
// 		brandList = append(brandList, brand)
// 	}
// 	return
// }

// func batch(carGroupList []CarGroup2) {
// 	if len(carGroupList) == 0 {
// 		fmt.Print("空消息")
// 		return
// 	}

// 	var buffer bytes.Buffer
// 	sql := "INSERT INTO urun_car.new_car_group3 (brand_pro, year, a_id, b_id, brand_equal_str, model_equal_str, style_equal_str, brand_equal_count, model_equal_count, style_equal_count, ab_brand_nes, ba_brand_nes, ab_model_nes, ba_model_nes, ab_style_nes, ba_style_nes, ab_brand_nec, ba_brand_nec, ab_model_nec, ba_model_nec, ab_style_nec, ba_style_nec, ab_brand_flag, ba_brand_flag, ab_model_flag, ba_model_flag, ab_style_flag, ba_style_flag, status) VALUES "
// 	if _, err := buffer.WriteString(sql); err != nil {
// 		fmt.Print(err.Error())
// 	}

// 	for index, value := range carGroupList {
// 		if index == len(carGroupList)-1 {
// 			buffer.WriteString(fmt.Sprintf("('%s', '%d', '%d', '%d', '%s', '%s', '%s', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d');", value.BrandPro, value.Year, value.AID, value.BID, value.BrandEqualStr, value.ModelEqualStr, value.StyleEqualStr, value.BrandEqualCount, value.ModelEqualCount, value.StyleEqualCount, value.AbBrandNes, value.BaBrandNes, value.AbModelNes, value.BaModelNes, value.AbStyleNes, value.BaStyleNes, value.AbBrandNec, value.BaBrandNec, value.AbModelNec, value.BaModelNec, value.AbStyleNec, value.BaStyleNec, value.AbBrandFlag, value.BaBrandFlag, value.AbModelFlag, value.BaModelFlag, value.AbStyleFlag, value.BaStyleFlag, value.Status))
// 		} else {
// 			buffer.WriteString(fmt.Sprintf("('%s', '%d', '%d', '%d', '%s', '%s', '%s', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d'),", value.BrandPro, value.Year, value.AID, value.BID, value.BrandEqualStr, value.ModelEqualStr, value.StyleEqualStr, value.BrandEqualCount, value.ModelEqualCount, value.StyleEqualCount, value.AbBrandNes, value.BaBrandNes, value.AbModelNes, value.BaModelNes, value.AbStyleNes, value.BaStyleNes, value.AbBrandNec, value.BaBrandNec, value.AbModelNec, value.BaModelNec, value.AbStyleNec, value.BaStyleNec, value.AbBrandFlag, value.BaBrandFlag, value.AbModelFlag, value.BaModelFlag, value.AbStyleFlag, value.BaStyleFlag, value.Status))
// 		}
// 	}

// 	_, err := MysqlDB.Exec(buffer.String())
// 	if err != nil {
// 		fmt.Print("插入消息失败：", err.Error())
// 		batch(carGroupList)
// 	}
// 	return
// }