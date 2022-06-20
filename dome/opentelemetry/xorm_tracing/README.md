# Golang XORM分布式链路追踪（源码分析）
使用`XORM`和`Opentracing`，让你彻彻底底摆脱繁琐的`CRUD`的阴影，将工作重心转移至业务逻辑

> 系统环境  
> go version go1.14.3 windows/amd64  
> xorm.io/xorm 1.0.3  


# 一、 ⚠️提醒

- XORM版本`1.0.2`及以上（支持以`Hook钩子函数`方式侵入XORM执行过程）  

- [项目Gitee源码](https://gitee.com/avtion/xormWithTracing)  

- 一定要看`Q&A`，看`Q&A`，看`Q&A`！

- 在之前的 [Golang实战 XORM搭配OpenTracing+Jaeger链路监控让SQL执行一览无遗](https://www.avtion.cn/post/golang%E5%AE%9E%E6%88%98-xorm%E6%90%AD%E9%85%8Dopentracing-jaeger%E9%93%BE%E8%B7%AF%E7%9B%91%E6%8E%A7%E8%AE%A9sql%E6%89%A7%E8%A1%8C%E4%B8%80%E8%A7%88%E6%97%A0%E9%81%97/) 文章中使用的侵入日志模块的方式是`1.0.2`版本以下侵入的方式，但是这种方式存在并发问题，会丢失`span`，如果升级到了`1.0.2`版本或以上就应该使用`Hook钩子`的方式侵入

# 二、过程

## 1. 使用Docker启动Opentracing + jaeger服务

**使用的镜像：`jaegertracing/all-in-one:1.18`**

Docker命令
```
docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:1.18
```

浏览器访问`localhost:16686`，可以看到`JaegerUI`界面，如下所示：

![输入图片说明](https://images.gitee.com/uploads/images/2020/0813/224040_42d51718_2051718.png "clipboard.png")

至此，以`内存`作为数据寄存方式的`OpenTracing+Jaeger`服务成功运行。

> 这里不详细讲解 `Opentracing` 的使用，有兴趣的同学可以查看官方技术文档  
> 同时这里搭配的 `jaeger` 进行使用，方便我们调试和线上小规模部署



## 2. 安装Xorm、OpenTracing和Jaeger

Xorm - 需要 `1.0.2` 版本及以上才能支持`Hook钩子函数`
```
go get xorm.io/xorm
```

OpenTracing和Jaeger - 只需要安装`Jaeger-Client`就会依赖`Opentracing`
```
go get github.com/uber/jaeger-client-go
```

## 3. 初始化Opentracing --> 通常由服务统一初始化
> main.go

```go
package main

func initJaeger() (closer io.Closer, err error) {
	// 根据配置初始化Tracer 返回Closer
	tracer, closer, err := (&config.Configuration{
		ServiceName: "xormWithTracing",
		Disabled:    false,
		Sampler: &config.SamplerConfig{
			Type: jaeger.SamplerTypeConst,
			// param的值在0到1之间，设置为1则将所有的Operation输出到Reporter
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
		    // 请注意，如果logSpans没开启就一定没有结果
			LogSpans:           true,
            // 请一定要注意，LocalAgentHostPort填错了就一定没有结果
			LocalAgentHostPort: "localhost:6831",
		},
	}).NewTracer()
	if err != nil {
		return
	}

	// 设置全局Tracer - 如果不设置将会导致上下文无法生成正确的Span
	opentracing.SetGlobalTracer(tracer)
	return
}
```

## 3. 定义Hook钩子结构体
> custom_hook.go

```go
package main

type TracingHook struct {
	// 注意Hook伴随DB实例的生命周期，所以我们不能在Hook里面寄存span变量
	// 否则就会发生并发问题
	before func(c *contexts.ContextHook) (context.Context, error)
	after  func(c *contexts.ContextHook) error
}

// xorm的hook接口需要满足BeforeProcess和AfterProcess函数
func (h *TracingHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return h.before(c)
}

func (h *TracingHook) AfterProcess(c *contexts.ContextHook) error {
	return h.after(c)
}

// 让编译器知道这个是xorm的Hook，防止编译器无法检查到异常
var _ contexts.Hook = &TracingHook{}
```

## 4. 实现钩子函数
> custom_hook.go

```go
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

    // 记录
	sp.LogFields(tracerLog.String("SQL", sql))
	sp.LogFields(tracerLog.Object("args", c.Args))
	sp.SetTag("execute_time", c.ExecuteTime)

	return nil
}
```

## 5. 把钩子函数挂载在钩子Hook上
> custom_hook.go

```go
func NewTracingHook() *TracingHook {
	return &TracingHook{
		before: before,
		after:  after,
	}
}
```

## 6. 在xorm引擎初始化时挂载自定义的Hook
> main.go

```go
func NewEngineForHook() (engine *xorm.Engine, err error) {
	// XORM创建引擎
	engine, err = xorm.NewEngine("mysql", "root:root@(localhost:3306)/test?charset=utf8mb4")
	if err != nil {
		return
	}

	// 使用我们的钩子函数 <---- 重要
	engine.AddHook(NewTracingHook())
	return
}

```


# 四、 单元测试 ---> 不会如何使用的请认真看一下
> main_test.go

```go
// XORM技术文档范例
type User struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
}

// 新方式进行上下文注入，要求 xorm 1.0.2版本
func TestNewEngineForHook(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	// 初始化XORM引擎
	engine, err := NewEngineForHook()
	if err != nil {
		t.Fatal(err)
	}

	// 初始化Tracer
	closer, err := initJaeger()
	if err != nil {
		t.Fatal(err)
	}
	defer closer.Close()

	// 生成新的Span - 注意将span结束掉，不然无法发送对应的结果
	span := opentracing.StartSpan("xorm sync")
	defer span.Finish()

	// 把生成的Root Span写入到Context上下文，获取一个子Context
	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// 将子上下文传入Session
	session := engine.Context(ctx)

	// Sync2同步表结构
	if err := session.Sync2(&User{}); err != nil {
		t.Fatal(err)
	}

	// 插入一条数据
	if _, err := session.InsertOne(&User{Name: fmt.Sprintf("test-%d", rand.Intn(1<<10))}); err != nil {
		t.Fatal()
	}
}
```

1. 再次查看jaeger可以看到一共有三个记录，其中最下面有3个span的trace是我们主动发起数据库操作的记录，上面两个trace是xorm自行发起的数据库操作

![输入图片说明](https://images.gitee.com/uploads/images/2020/0813/224120_c43720e5_2051718.png "clipboard.png")

2. 点击进入trace，展开子span，我们可以看到xorm执行的SQL语句和参数、执行时间以及所占的比例

![输入图片说明](https://images.gitee.com/uploads/images/2020/0813/224131_035930d3_2051718.png "clipboard.png")

![输入图片说明](https://images.gitee.com/uploads/images/2020/0813/224141_9d441a3e_2051718.png "clipboard.png")

# 五、 Q&A

## 1. jaeger没有记录？
- 如果访问`jaeger`发现压根没有服务请求`jaeger`请务必检查初始化`opentracing`是否成功（特别检查一下`LocalAgentHostPort`参数是否填写正确）
- 子`span`没有`finish`

## 2. 为什么不使用GORM？
- GORM不支持上下文侵入，至少GORM目前版本不支持，仅支持引擎全局上下文，这也导致了无法单独追踪每个请求，而只能追踪整个GORM引擎，GORM估计在2.0版本支持上下文侵入

## 3. 为什么通过日志模块侵入不能实现并发安全？
- xorm的会话在调用日志使用的是值传递去传递ContextHook，所以不能把span寄存到ContextHook中的上下文Ctx，而是只能暂时寄存在全局日志Logger中，这也导致了日志实例中的span是并发不安全的，高并发情况下会丢失span
- 日志模块侵入依旧是xorm`1.0.2`以下的版本最好的侵入方式，但是不推荐

## 4. 为什么有一些莫名其妙的SQL查询？
- xorm有时候会根据特殊情况对SQL进行优化，我们可以看到有时候xorm会查询一下表结构，以方便后续的查询

## 5. 有什么实际用处？
- 方便业务上线后进行追踪，实际Opentracing的用处就是如此

# 六、 源码分析（没兴趣可以不看）

## 1. 注册钩子结构体

> xorm -> engine.go

```go
func (engine *Engine) AddHook(hook contexts.Hook) {
	engine.db.AddHook(hook)
}
```

> xorm -> db.go

```go
func (db *DB) AddHook(h ...contexts.Hook) {
	db.hooks.AddHook(h...)
}
```

> xorm -> hook.go

```go
type Hook interface {
	BeforeProcess(c *ContextHook) (context.Context, error)
	AfterProcess(c *ContextHook) error
}

type Hooks struct {
	hooks []Hook
}

func (h *Hooks) AddHook(hooks ...Hook) {
	h.hooks = append(h.hooks, hooks...)
}
```

我们需要传入一个`contexts.Hook`，是一个接口类型，我们只需要实现两个方法就能实现这个接口

## 2. 钩子上下文结构体

> xorm -> hook.go

```go
// ContextHook represents a hook context
type ContextHook struct {
    // 开始时间
	start       time.Time
	// 上下文
	Ctx         context.Context
	// SQL语句
	SQL         string        // log content or SQL
	// SQL参数
	Args        []interface{} // if it's a SQL, it's the arguments
	// 查询结果
	Result      sql.Result
	// 执行时间
	ExecuteTime time.Duration
	// 如果发生错误，会赋值
	Err         error // SQL executed error
}
```

`ContextHook`结构体作为Hook接口调用函数的入参，包括了我们需要的基本数据

## 3. 执行SQL语句时调用逻辑

> xorm -> db.go

```go
func (db *DB) beforeProcess(c *contexts.ContextHook) (context.Context, error) {
	if db.NeedLogSQL(c.Ctx) {
	    // <-- 重要，这里是将日志上下文转化成值传递
	    // 所以不能修改context.Context的内容
		db.Logger.BeforeSQL(log.LogContext(*c))
	}
	// Hook是指针传递，所以可以修改context.Context的内容
	ctx, err := db.hooks.BeforeProcess(c)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (db *DB) afterProcess(c *contexts.ContextHook) error {
    // 和beforeProcess同理，日志上下文不能修改context.Context的内容
    // 而hook可以
	err := db.hooks.AfterProcess(c)
	if db.NeedLogSQL(c.Ctx) {
		db.Logger.AfterSQL(log.LogContext(*c))
	}
	return err
}
```

- 这一段就是实际SQL查询过程中调用日志和Hook的过程，从这里可以非常明显的看到日志模块传入的是值而不是指针，从而导致了我们无法修改日志模块中的上下文实现span的传递，只能利用全局日志实例来传递span，这直接出现了并发安全问题
- 而Hook的传递使用的是指针传递，将`contexts.ContextHook`的指针传入钩子函数执行流程，允许我们直接操作`Ctx`