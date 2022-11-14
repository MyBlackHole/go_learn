package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getRouteParams(c *gin.Context) {
	// 获取路由参数为name的值
	// http://127.0.0.1:8888/route/card  输出 card
	name := c.Param("name")
	c.String(http.StatusOK, name)
}

func getRoutePregParams(c *gin.Context) {

	// 获取路由参数为name的值 和job的值 job为无限长,忽略/,并且包含前一个/
	// http://127.0.0.1:8888/route/card/
	// 输出 card /

	// http://127.0.0.1:8888/route/card/name
	// 输出 card /name

	// http://127.0.0.1:8888/route/card/name/hello
	// 输出 card /name/hello

	name := c.Param("name")
	job := c.Param("job")
	c.String(http.StatusOK, name+" "+job)
}

/*
*

	获取路由参数
*/
func getParams(c *gin.Context) {

	// http://127.0.0.1:8888/?name=card
	// 输出 card card 程序员

	// http://127.0.0.1:8888/?name=card&job=coder
	// 输出 card card coder

	// 当有多个参数相同时,取第一个value
	// http://127.0.0.1:8888/?name=card&job=coder&name=周起
	// 输出 card card coder

	// http://127.0.0.1:8888/?name=周起&job=coder&name=card
	// 输出 周起 周起 coder

	// c.Query() 等同于 c.Request.URL.Query().Get()
	name := c.Query("name")
	name2 := c.Request.URL.Query().Get("name")

	// 设置默认参数,如果job参数不存在,默认为程序员
	name3 := c.DefaultQuery("job", "程序员")
	c.String(http.StatusOK, name+" "+name2+" "+name3)
}

func getRepeatParams(c *gin.Context) {
	name := c.QueryArray("name")
	// http://127.0.0.1:8888/repeat?name=周起&name=card
	// 输出  name:[周起 card], type:[]string
	fmt.Printf("name:%v, type:%T", name, name)

	// http://127.0.0.1:8888/repeat?job[a]=周起&job[b]=card
	// 输出  job:map[a:周起 b:card], type:map[string]string
	job := c.QueryMap("job")
	fmt.Printf("job:%v, type:%T", job, job)
}

func main() {
	router := gin.Default()

	// 获取路由参数
	router.GET("/route/:name", getRouteParams)

	// 获取正则路由参数
	router.GET("/route/:name/*job", getRoutePregParams)

	// 获取queryString
	router.GET("/", getParams)
	// 获取同名参数
	router.GET("/repeat", getRepeatParams)

	router.Run(":8888")
}
