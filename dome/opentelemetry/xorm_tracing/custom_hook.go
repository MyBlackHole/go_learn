package main

import (
	"context"
	"github.com/opentracing/opentracing-go"
	tracerLog "github.com/opentracing/opentracing-go/log"
	"xorm.io/builder"
	"xorm.io/xorm/contexts"
)

// 使用新定义类型
// context.WithValue方法中注释写到
// 提供的键需要可比性，而且不能是字符串或者任意内建类型，避免不同包之间
// 调用到相同的上下文Key发生碰撞，context的key具体类型通常为struct{}，
// 或者作为外部静态变量（即开头字母为大写）,类型应该是一个指针或者interface类型
type xormHookSpan struct{}

var xormHookSpanKey = &xormHookSpan{}

type TracingHook struct {
	// 注意Hook伴随DB实例的生命周期，所以我们不能在Hook里面寄存span变量
	// 否则就会发生并发问题
	before func(c *contexts.ContextHook) (context.Context, error)
	after  func(c *contexts.ContextHook) error
}

// 让编译器知道这个是xorm的Hook，防止编译器无法检查到异常
var _ contexts.Hook = &TracingHook{}

func (h *TracingHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return h.before(c)
}

func (h *TracingHook) AfterProcess(c *contexts.ContextHook) error {
	return h.after(c)
}

func NewTracingHook() *TracingHook {
	return &TracingHook{
		before: before,
		after:  after,
	}
}

func before(c *contexts.ContextHook) (context.Context, error) {
	// 这里一定要注意，不要拿第二个返回值作为上下文进行替换，而是用自己的key
	span, _ := opentracing.StartSpanFromContext(c.Ctx, "xorm sql execute")
	c.Ctx = context.WithValue(c.Ctx, xormHookSpanKey, span)
	return c.Ctx, nil
}

func after(c *contexts.ContextHook) error {
	// 自己实现opentracing的SpanFromContext方法，断言将interface{}转换成opentracing的span
	sp, ok := c.Ctx.Value(xormHookSpanKey).(opentracing.Span)
	if !ok {
		// 没有则说明没有span
		return nil
	}
	defer sp.Finish()

	// 记录我们需要的内容
	if c.Err != nil {
		sp.LogFields(tracerLog.Object("err", c.Err))
	}

	// 使用xorm的builder将查询语句和参数结合，方便后期调试
	sql, _ := builder.ConvertToBoundSQL(c.SQL, c.Args)

	sp.LogFields(tracerLog.String("SQL", sql))
	sp.LogFields(tracerLog.Object("args", c.Args))
	sp.SetTag("execute_time", c.ExecuteTime)

	return nil
}
