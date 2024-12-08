package leetcode150

import "strings"

func wordPattern(pattern string, s string) bool {
	sArr := strings.Split(s, " ")
	n := len(pattern)
	m := len(sArr)
	if n != m {
		return false
	}

	// p 映射到 s 上
	mapStr := make(map[uint8]string, 0)
	// 存放到 s
	arr := make(map[string]struct{}, 0)
	for i := 0; i < n; i++ {
		v, ok := mapStr[pattern[i]]
		// p[i] 不存在 先看 s[i] 是否存在
		if !ok {
			// s[i] 存在就 false
			if _, ok := arr[sArr[i]]; ok {
				return false
			}
			// s[i] 不存在建立映射关系
			mapStr[pattern[i]] = sArr[i]
			arr[sArr[i]] = struct{}{}
		} else {
			if v != sArr[i] {
				return false
			}
		}
	}
	return true
}
