package main

import (
	"fmt"
)

// 空间复杂度O(1)
// 双指针消除中间多余的空格，只留一个空格
// 类比消除数组中多余的重复元素
func reverseWords(s string) string {
	sByt := []byte(s)
	// 消除开头的空格
	slow, fast := 0, 0
	for ; fast < len(sByt); fast++ {
		if sByt[fast] != ' ' {
			break
		}
	}
	// 消除中间多余的空格，只留一个空格
	for ; fast < len(sByt); fast++ {
		if fast > 0 && sByt[fast] == ' ' && sByt[fast] == sByt[fast-1] {
			continue
		}
		sByt[slow] = sByt[fast]
		slow++
	}
	// 不能一次性消除所有多余的空格，因为末尾空格总有一个(如果原string末尾有空格)
	// 此时可能末尾有一个空格，也可能没有,消除末尾空格
	if sByt[slow-1] == ' ' {
		slow--
	}
	sByt = sByt[:slow]
	// 翻转整个有效string
	reverse(sByt)
	// 翻转每个单词
	start, end := 0, 0
	for ; end < len(sByt); end++ {
		if sByt[end] == ' ' {
			// 左闭右开
			reverse(sByt[start:end])
			start = end + 1
			end = end + 1
		}
	}
	// 结束的时候end=len(sByt)，前面代码都没错误的话，end是无效，并且sByt[end-1]是一个字符
	// 翻转最后一个单词
	reverse(sByt[start:end])
	return string(sByt)

}
func reverse(sByt []byte) {
	for i, j := 0, len(sByt)-1; i < j; i, j = i+1, j-1 {
		sByt[i], sByt[j] = sByt[j], sByt[i]
	}
}
func main() {

	s := " hello    world  "
	fmt.Println(reverseWords(s))
}
