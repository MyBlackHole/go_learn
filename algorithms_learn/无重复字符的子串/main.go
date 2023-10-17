package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    // left 左指针
	var left, res int
	usedChar := make(map[rune]int)
	for i, v := range s {
        // left <= k: 当前字符串, 上一次出现位置小于当前窗口的左开始位置, 避免往回移动
		if k, ok := usedChar[v]; ok && left <= k {
            // 移动窗口左位置
			left = k + 1
		} else {
            // i-left+1: 计算滑动窗口大小
			res = max(res, i-left+1)
		}
        // 修改记录当前字符最新出现位置
		usedChar[v] = i
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	var t string = "3456789x1234123"
	var m string = "sdkjfisf56789x1234123"
	fmt.Println(lengthOfLongestSubstring(t))
	fmt.Println(lengthOfLongestSubstring(m))
	fmt.Println(lengthOfLongestSubstring("1234321"))
}
