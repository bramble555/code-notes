package main

import (
	"fmt"
	"strings"
)

// 判断是否为周期串
// 有一个结论：如果字符串在其掐头去尾的双倍字符串中，它就是周期串
func repeatedSubstringPattern(s string) bool {
	s2 := (s + s)
	// 掐头去尾
	s2 = s2[1 : len(s2)-1]
	// 这个Contains方法实际上就是KMP算法。KMP算法先搁置。
	return strings.Contains(s2, s)
}
func main() {
	fmt.Println(repeatedSubstringPattern("abcabc"))
}
