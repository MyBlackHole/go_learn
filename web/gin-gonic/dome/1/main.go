package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "Black%v", "Hole")
	})

	r.GET("/new", func(ctx *gin.Context) {
		ctx.String(200, "Black%v", "Hole new")
	})

	r.POST("/add", func(ctx *gin.Context) {
		ctx.String(200, "Black%v", "Hole add")
	})

	r.Run()
}
