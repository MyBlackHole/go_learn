package main

import "fmt"

func main() {
	fmt.Println(2 == 4)
	var count int = 40
	if count < 30 {
		fmt.Println("ok")
	}

	if count1 := 20; count1 < 30 {
		fmt.Println("o")
	} else {
		fmt.Println("k")
	}

	if count1 := 20; count1 < 30 {
		fmt.Println("o")
	} else if count1 > 100 {
		fmt.Println("k")
	} else {
		fmt.Println("?")
	}

	var i int

	switch i {
	// 10 与 i 类型必须一致
	case 10:
		fmt.Println(10)
		// 不用 bread 默认执行一个 case 退出
	case 20, 40, 00:
		fmt.Println(20)
	case 30:
		fmt.Println(30)
		// 可以放在任意位置，放在那都是先执行完 case 才到 default
	default:
		fmt.Println("???")
	}

	fmt.Println("---")

	switch {
	// 10 与 i 类型必须一致
	case 10 == 20:
		fmt.Println(10)
		// 不用 bread 默认执行一个 case 退出
	case 20 == 20:
		fmt.Println(20)
	case 30 == 20:
		fmt.Println(30)
		// 可以放在任意位置，放在那都是先执行完 case 才到 default
	default:
		fmt.Println("???")
	}

	switch b := 10; {
	case b > 5:
		fmt.Println("大于5")
	case b < 5:
		fmt.Println("大于5")
	}

	switch b := 10; {
	case b > 5:
		fmt.Println("大于5")
		// 穿透一层
		fallthrough
	case b > 8:
		fmt.Println("大于8")
	case b < 5:
		fmt.Println("大于5")
	}

	var sum int
	for i := 0; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

	ii := 0
	for ii < 5 {
		fmt.Println(ii)
		ii++
	}

	// for {
	// 	// 死循环
	//
	// }

	var str string = "black hole"
	for _, v := range str {
		fmt.Printf("%c\n", v)
	}

	var sum1 int = 0
	for i := 1; i <= 100; i++ {
		sum1 += i
		fmt.Println(sum1)
		if sum >= 300 {
			// 退出 for
			break
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("i:%v, j:%v\n", i, j)
			if i == 2 && j == 2 {
				// 结束最近的循环
				break
			}
		}
	}

lable2:
	for i := 0; i < 5; i++ {
		// lable1:
		for j := 0; j < 5; j++ {
			fmt.Printf("i:%v, j:%v\n", i, j)
			if i == 2 && j == 2 {
				// 结束 lable2 循环
				break lable2
			}
		}
	}

	for i := 0; i < 101; i++ {
		if i%6 != 0 {
			// 结束本次循环
			continue
		}
		fmt.Println(i)
	}

lable3:
	for i := 0; i < 5; i++ {
		// lable1:
		for j := 0; j < 5; j++ {
			fmt.Printf("i:%v, j:%v\n", i, j)
			if i == 2 && j == 2 {
				// 结束 lable2 循环
				continue lable3
			}
		}
	}

	fmt.Println("1")
	fmt.Println("2")
	if true {
		goto lable4
	}
	fmt.Println("3")
	fmt.Println("4")
lable4:
	fmt.Println("5")

	for i := 0; i < 101; i++ {
		fmt.Println(i)
		if i == 15 {
			return
		}
	}

	fmt.Println("return")
}
