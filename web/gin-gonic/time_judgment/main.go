package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.POST("/Task/WeiboAddTimeFilter", saveData)
	r.POST("/Task/WeiboTimeFilter", getData)
	r.POST("/Log/addChargeLog", addChargeLog)
	r.POST("/Log/findChargeLog", findChargeLog)
	r.GET("/Comment/SelectBlogger", readUser)
	// 209
	r.Run("0.0.0.0:8010")
	// // 206
	// r.Run("0.0.0.0:8012")
}
