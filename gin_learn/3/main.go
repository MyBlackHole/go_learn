package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func func1(c *gin.Context) {
	time.Sleep(100 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"msg":    "time",
		"status": "no",
	})
}

func func2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":    "",
		"status": "ok",
	})
}

func main() {
	print("black")
	r := gin.Default()
	r.GET("/time", func1)
	r.GET("/name", func2)
	r.Run()
}
