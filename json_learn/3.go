package main

import (
	"fmt"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
)

// func main() {
// 	const json = `{"city":"dhaka","type":"weekly","temperatures":[30,39.9,35.4,33.5,31.6,33.2,30.7]}`
// 	avg := gojsonq.New().FromString(json).From("temperatures").Avg()
// 	fmt.Printf("Average temperature: %.2f", avg) // 33.471428571428575
// }
func main() {
	jsonStr := `{
   "code":"0",
   "message":"success",
   "data":{
       "thumb_up":0,
       "system":0,
       "im":0,
       "avatarUrl":"https://profile.csdnimg.cn/F/4/B/helloworld",
       "invitation":0,
       "comment":0,
       "follow":0,
       "totalCount":0,
       "coupon_order":0,
       "pageNumbers":[
           1,
           2,
           3,
           4,
           5,
           6
       ]
   },
   "status":true
   }`

	// 保存 json 字符串中 pageNumbers 字段列表数据
	pageNumbers := make([]int, 0)
	result, err := gojsonq.New().FromString(jsonStr).FindR("data.pageNumbers")
	if err != nil {
		fmt.Println(err)
	}
	// 将提前的数据解析为自定义变量类型
	err = result.As(&pageNumbers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pageNumbers)
}
