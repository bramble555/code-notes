package leetcode150

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	// 转换小写
	s = strings.ToLower(s)
	// 只保留字母和数字
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	s = strings.Join(strings.FieldsFunc(s, f), "")
	return flag(s)
}

// 判断是否为回文字符串
func flag(s string) bool {
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
