package main

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engineA := gin.Default()
	engineA.Use(hasXiaohong)
	engineA.Run(":12345")
}

func hasXiaohong(c *gin.Context) {
	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}
	defer c.Request.Body.Close()

	if bytes.Contains(data, []byte("小红")) {
		// 发送给服务 B
		// 数据封装
		http.Post("服务 B 的地址", "contentType", bytes.NewBuffer(data))
	}

	// 注意这时 c.Request.Body 已经读完了，需要重新将读出来的值给放回去，之后的处理就依然可以使用 c.Request.Body 了。
	bufReader := bytes.NewBuffer(data)           // 这只是一个 Reader
	c.Request.Body = io.NopCloser(bufReader) // 这边通过 NopCloser 包装成 ReadCloser
}
