package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	logg *log.Logger
)

func work(ctx context.Context, ch chan bool) {
	for {
		/*
		   Deadline:是获取设置的超时时间:
		   第一个返回值:设置的超时时间，到超时时间Context会自动发起取消请求
		   第二个返回值ok==false时表示没有设置截止时间，如果需要取消的话需要调用取消函数进行取消。
		*/
		if deadline, ok := ctx.Deadline(); ok {
			fmt.Println(deadline)
			if time.Now().After(deadline) {
				logg.Println(`超时退出!`)
				//这里是为了演示,Context中的Err()输出:context deadline exceeded
				logg.Printf(ctx.Err().Error())
				ch <- true
				return
			}

		}
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			ch <- true
			return
		default:
			logg.Println(`上班!`)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan bool)
	logg = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go work(ctx, ch)
	time.Sleep(10 * time.Second)
	//取消函数：当cancel被调用时,context.WithTimeout设置的时间超过后,关闭ctx.Done通道；
	cancel()
	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	<-ch
	logg.Println(`无脑发呆中!`)
}

/* outfile:
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:42 上班!
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:43 上班!
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:44 上班!
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:45 上班!
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:46 上班!
2019-01-30 18:23:47.851042 +0800 CST m=+5.000360703
18:23:47 超时退出!
18:23:47 context deadline exceeded //这里就是ctx超时的时候ctx.Err的错误消息。
18:23:52 无脑发呆中!
*/
