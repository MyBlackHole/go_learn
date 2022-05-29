package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	// 必须大写开头
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"user": username,
			"age":  age,
			"page": page,
		})
	})

	// curl -X GET 127.0.0.1:8080/user\?username=1\&password=kjjj
	// curl -X GET 127.0.0.1:8080/user -F "username=aaaa" -F "password=aaaaaaa"
	r.GET("/user", func(ctx *gin.Context) {
		var userinfo User
		if err := ctx.ShouldBind(&userinfo); err == nil {
			ctx.JSON(http.StatusOK, userinfo)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.POST("/post", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// curl -X POST 127.0.0.1:8080/user\?username=1\&password=kjjj
	// curl -X POST 127.0.0.1:8080/user -F "username=aaaa" -F "password=aaaaaaa"
	r.POST("/user", func(ctx *gin.Context) {
		var userinfo User
		if err := ctx.ShouldBind(&userinfo); err == nil {
			ctx.JSON(http.StatusOK, userinfo)
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"error": err.Error(),
			})
		}
	})

	// curl -X GET 127.0.0.1:8080/list/123
	r.GET("/list/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(http.StatusOK, "%v", cid)
	})

	defaultRouters := r.Group("/v1")
	{
		// curl -X GET 127.0.0.1:8080/v1/\?username=1\&password=kjjj
		defaultRouters.GET("/", func(ctx *gin.Context) {
			var userinfo User
			if err := ctx.ShouldBind(&userinfo); err == nil {
				ctx.JSON(http.StatusOK, userinfo)
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			}
		})
	}

	r.Run()
}
