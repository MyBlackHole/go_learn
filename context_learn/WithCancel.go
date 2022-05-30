package main

import (
	"context"
	"log"
	"os"
	"time"
)

var (
	logg *log.Logger
)

func work(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			logg.Println(`下班!`)
			ch <- true
			return
		default:
			logg.Println(`上班!`)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ch := make(chan bool)

	logg = log.New(os.Stdout, "", log.Ltime)
	ctx, cancel := context.WithCancel(context.Background())
	go work(ctx, ch)
	time.Sleep(10 * time.Second)
	//取消函数：当cancel被调用时,WithCancel遍历Done以执行关闭；
	cancel()

	// 这个chan是为了保证子的goroutine执行完,当然也可以不用chan用time.Sleep停止几秒
	<-ch
	logg.Println(`无脑发呆中!`)
}

/* outfile:
17:27:52 上班!
17:27:54 上班!
17:27:56 上班!
17:27:58 上班!
17:28:00 上班!
17:28:02 下班!
17:28:02 无脑发呆中!
*/
