/*
# 什么是uuid? #

	uuid是Universally Unique Identifier的缩写，即通用唯一识别码。

	uuid的目的是让分布式系统中的所有元素，都能有唯一的辨识资讯，而不需要透过中央控制端来做辨识资讯的指定。如此一来，
	每个人都可以建立不与其它人冲突的 uuid。

	A universally unique identifier (UUID) is a 128-bit number used to identify information in computer systems.
	go生成uuid：

	目前，golang中的uuid还没有纳入标准库，我们使用github上的开源库：

			-----------------------------------------------------------------
			go get -u github.com/satori/go.uuid
			-----------------------------------------------------------------
*/
package main

import (
	"github.com/satori/go.uuid"
)
func main() {
	// 创建 UUID v4
	u1 := uuid.NewV4()
	println(`生成的UUID v4：`)
	println(u1.String())

	// 创建可以进行错误处理的 UUID v4
	u2 := uuid.NewV4()

	println(`生成的UUID v4：`)
	println(u2.String())

	// 解析 字符串 到 UUID
	u2 = uuid.Must(uuid.FromString(`9b08dc66-a56f-4b64-b9eb-93efed77eb2b`))
	println(`解析 字符串 到 UUID 成功！解析到的 UUID 如下：`)
	println(u2.String())
}
