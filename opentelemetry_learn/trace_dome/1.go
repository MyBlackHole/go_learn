package main

import (
	"fmt"
	"io"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
)

// InitJaeger ...
func InitJaeger(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := jaegercfg.FromEnv()

	cfg.Sampler.Type = "const"
	cfg.Sampler.Param = 1
	cfg.Reporter.LocalAgentHostPort = "127.0.0.1:6831"
	cfg.Reporter.LogSpans = true

	tracer, closer, err := cfg.New(service, jaegercfg.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}

func main() {
	tracer, closer := InitJaeger("hello-world")
	defer closer.Close()
	opentracing.InitGlobalTracer(tracer)

	helloStr := "hello jaeger"
	span := tracer.StartSpan("say-hello")
	time.Sleep(time.Duration(2) * time.Millisecond)
	println(helloStr)
	span.Finish()
}
