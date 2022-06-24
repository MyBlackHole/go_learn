package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	// 必须大写开头
	Name string
	Age  int `json:"age"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Black%v", "Hole")
	})

	r.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]interface{}{
			"success": true,
			"msg":     "black",
		})
	})

	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"success": true,
			"msg":     "black",
		})
	})

	r.GET("/json3", func(ctx *gin.Context) {
		ctx.JSON(200, &User{})
	})

	// 127.0.0.1:8080/jsonp\?callback=sjfs
	// sjfs({"Name":"","age":0})
	r.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(200, &User{})
	})

	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, gin.H{
			"success": true,
			"msg":     "black",
		})
	})

	r.GET("/html", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "----",
		})
	})
	r.Run()
}
