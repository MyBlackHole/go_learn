package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//单个文件上传
func main() {
	r := gin.Default()
	//8<<20 即 8*2^20=8M
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传文件出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename+"上传成功")
	})
	r.Run()
}
