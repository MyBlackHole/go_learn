package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 基本数据类型转 string
	var n1 int = 19
	var n2 float32 = 1.01
	var n3 bool = false
	var n4 byte = 'a'

	var s1 string = fmt.Sprintf("%d", n1)
	fmt.Printf("%T--%q\n", s1, s1)

	var s2 string = fmt.Sprintf("%f", n2)
	fmt.Printf("%T--%q\n", s2, s2)

	var s3 string = fmt.Sprintf("%t", n3)
	fmt.Printf("%T--%q\n", s3, s3)

	var s4 string = fmt.Sprintf("%c", n4)
	fmt.Printf("%T--%q\n", s4, s4)

	s1 = strconv.FormatInt(int64(n1), 10)
	fmt.Printf("%T--%q\n", s1, s1)

	// f dddd.ddd
	// 9 保留小数点后面 9 位
	// 64 表示 float64 类型
	s2 = strconv.FormatFloat(float64(n2), 'f', 9, 64)
	fmt.Printf("%T--%q\n", s2, s2)

	s3 = strconv.FormatBool(n3)
	fmt.Printf("%T--%q\n", s3, s3)

	// string --> bool
	s1 = "true"
	var b bool
	b, _ = strconv.ParseBool(s1)
	fmt.Printf("%T--%v\n", b, b)

	// string --> int64
	s2 = "19"
	var i int64
	i, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Printf("%T--%v\n", i, i)

	s3 = "1.111"
	var f float64
	f, _ = strconv.ParseFloat(s3, 64)
	fmt.Printf("%T--%v\n", f, f)

	// 转换失败会返回对应类型默认值
}
