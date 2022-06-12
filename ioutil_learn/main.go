package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fileName := "./1/1.go"
	data, err := ioutil.ReadFile(fileName)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	fileName1 := "1.txt"
	s1 := "test文件"
	err = ioutil.WriteFile(fileName1, []byte(s1), os.ModePerm)
	fmt.Println(err)

	s2 := "水边八"
	r1 := strings.NewReader(s2)
	data, err = ioutil.ReadAll(r1)
	fmt.Println(err)
	fmt.Println(data)
	fmt.Println(string(data))

	// 目录读取, 一层
	dirName := "./1"
	fileInfos, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		fmt.Printf("%s --- %t \n", fileInfos[i].Name(), fileInfos[i].IsDir())
	}

	// 临时文件与临时目录
	dir, err := ioutil.TempDir("1", "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(dir)
	fmt.Println(dir)

	file, err := ioutil.TempFile(dir, "Test")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(file.Name())
	fmt.Println(file.Name())
}
