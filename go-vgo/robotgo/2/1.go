package main

import (
	"github.com/go-vgo/robotgo"
)

func main() {

	// 获取屏幕上一个像素的颜色（十六进制颜色，不带 #）
	// 第1、2个参数：像素的座标
	color := robotgo.GetPixelColor(800, 400)
	println(`颜色：`, color)

	// 获取屏幕大小
	width, height := robotgo.GetScreenSize()
	println(`宽：`, width, ` 高：`, height)

	// 屏幕截图
	// 第1、2个参数：截图的座标
	// 第3、4个参数：截图的宽高
	bitmap := robotgo.CaptureScreen(0, 0, 800, 400)
	// 全屏截图
	// bitmap := robotgo.CaptureScreen(0, 0, width, height)

	// 保存截图
	// 第1个参数：截图数据
	// 第2个参数：保存的文件名，不支持UTF-8字符的文件名字（例如：中文）
	robotgo.SaveBitmap(bitmap, `screen.png`)
}
