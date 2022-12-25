package foo

type Foo interface {
	Bar(x int) int
}

func Baz(foo Foo, x int) int {
	return foo.Bar(x) + 1
}

// 执行生成
// mockgen -source=foo.go -destination=mock_foo.go
