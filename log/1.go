package main

import (
	"log"
	"os"
)

var logger *log.Logger

func init()  {
    logFile, err := os.OpenFile("./c.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        log.Panic("打开日志文件异常")
    }
    logger = log.New(logFile, "success ", log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
    logger.Println("自定义logger")
    http://img.12365auto.com/g/201605/11/201605111832507362.jpg
    http://img.12365auto.com/g/temp/201605/11/201605111832507362.jpg 
}
