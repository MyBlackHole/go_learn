package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	fileInfo, err := os.Stat("1.go")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("%T\n", fileInfo)
	// 名
	fmt.Println(fileInfo.Name())
	// 大小
	fmt.Println(fileInfo.Size())
	// 是否目录
	fmt.Println(fileInfo.IsDir())
	// 权限
	fmt.Println(fileInfo.Mode())

	fileName := "./1.go"
	// 是否绝对路径
	fmt.Println(filepath.IsAbs(fileName))
	// 获取绝对路径
	fmt.Println(filepath.Abs(fileName))

	// 获取父目录
	fmt.Println(path.Join(fileName, ".."))

	// 创建一层目录
	err = os.Mkdir("./test", os.ModePerm)
	// // 创建多层目录
	// err = os.MkdirAll("./test/test", os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}

	// 创建文件
	file1, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(file1)

	file2, err := os.OpenFile("test1.txt", os.O_RDONLY|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(file2)

	err = os.RemoveAll("test")
	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.Remove("test.txt")

	if err != nil {
		fmt.Println(err.Error())
	}

	err = os.Remove("test1.txt")

	if err != nil {
		fmt.Println(err.Error())
	}
}
