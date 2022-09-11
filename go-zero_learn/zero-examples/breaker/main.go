package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/zeromicro/go-zero/core/breaker"
	"github.com/zeromicro/go-zero/core/lang"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/cheggaaa/pb.v1"
)

const (
	duration        = time.Minute * 5
	breakRange      = 20
	workRange       = 50
	requestInterval = time.Millisecond
	// multiply to make it visible in plot
	stateFator = float64(time.Second/requestInterval) / 2
)

type (
	server struct {
		state int32
	}

	metric struct {
		calls int64
	}
)

// 原子加一
func (m *metric) addCall() {
	atomic.AddInt64(&m.calls, 1)
}

// 原子重置零, 返回原值
func (m *metric) reset() int64 {
	return atomic.SwapInt64(&m.calls, 0)
}

func newServer() *server {
	return &server{}
}

func (s *server) serve(m *metric) bool {
	m.addCall()
	return atomic.LoadInt32(&s.state) == 1
}

func (s *server) start() {
	go func() {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		var state int32
		for {
			var v int32
			if state == 0 {
				v = r.Int31n(breakRange)
			} else {
				v = r.Int31n(workRange)
			}
			time.Sleep(time.Second * time.Duration(v+1))
			// 最后一位取反
			state ^= 1
			atomic.StoreInt32(&s.state, state)
		}
	}()
}

func runBreaker(s *server, br breaker.Breaker, duration time.Duration, m *metric) {
	ticker := time.NewTicker(requestInterval)
	defer ticker.Stop()
	done := make(chan lang.PlaceholderType)
	fp, err := os.Create("result.log")
	logx.Must(err)
	defer fp.Close()

	go func() {
		time.Sleep(duration)
		close(done)
	}()

	for {
		select {
		case <-ticker.C:
			var msg string
			var msgBat string
			err := br.Do(func() error {
				if s.serve(m) {
					msg = "nil"
					return nil
				} else {
					msg = "err"
					return breaker.ErrServiceUnavailable
				}
			})
			msg = fmt.Sprintf("[%s:%s]", msg, err)
			if msg != msgBat {
				fmt.Fprintf(fp, "%s\n", msg)
				msgBat = msg
			}
		case <-done:
			return
		}
	}
}

// 断路器
func main() {
	srv := newServer()
	srv.start()

	gb := breaker.NewBreaker()
	fp, err := os.Create("result.csv")
	logx.Must(err)
	defer fp.Close()
	fmt.Fprintln(fp, "seconds,state,googleCalls,netflixCalls")

	var gm, nm metric
	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		var seconds int
		for range ticker.C {
			seconds++
			gcalls := gm.reset()
			ncalls := nm.reset()
			fmt.Fprintf(fp, "%d,%.2f,%d,%d\n",
				seconds, float64(atomic.LoadInt32(&srv.state))*stateFator, gcalls, ncalls)
		}
	}()

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	go func() {
		runBreaker(srv, gb, duration, &gm)
		waitGroup.Done()
	}()

	go func() {
		bar := pb.New(int(duration / time.Second)).Start()
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for range ticker.C {
			bar.Increment()
		}
		bar.Finish()
	}()

	waitGroup.Wait()
}
