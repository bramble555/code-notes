package leetcode150

import "regexp"

func lengthOfLastWord(s string) int {
	n := len(s)
	var i int = n - 1
	res := 0
	// 去除末尾空格
	for s[i] == ' ' {
		i--
	}
	// 数数
	for i >= 0 {
		if s[i] != ' ' {
			res++
			i--
		} else {
			return res
		}
	}
	return res
}
func lengthOfLastWord2(s string) int {
	// 去除末尾空格
	n := len(s)
	var i int = n - 1
	// 去除末尾空格
	for s[i] == ' ' {
		i--
	}
	rep := regexp.MustCompile(`\s+$`)
	match := rep.FindString(s[:i])
	return len(match)
}
