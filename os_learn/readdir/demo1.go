package main

import (
	"fmt"
	"os"
)

func listDirByReadDir(path string) {
    // 获取 path 下一层的文件与目录
    // 一次性读完返回(文件量多会堵塞很久)
    lst, err := os.ReadDir(path)
    if err != nil {
        panic(err)
    }
    for _, val := range lst {
        if val.IsDir() {
            fmt.Printf("[%s]\n", val.Name())
        } else {
            fmt.Println(val.Name())
        }
    }
}

func main()  {
    listDirByReadDir("/media/black/Data/Documents/Go")
}

