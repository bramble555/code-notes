package leetcode150

import "fmt"

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	var minLen int = 200
	var minIndexStr int
	// 比较出来哪个字符串最小
	for i := 0; i < n; i++ {
		minLen = min(minLen, len(strs[i]))
		minIndexStr = i
	}
	maxCount := 200
	for i := 0; i < n; i++ {
		// 当前字符串和最小字符串的相同前缀的最大长度
		maxTempLen := minLen
		for j := 0; j < minLen; j++ {
			if strs[i][j] != strs[minIndexStr][j] {
				// 第 j 个 比较失败，那么 最大长度为 j-1 -0 +1
				maxTempLen = j
				break
			}
		}
		maxCount = min(maxTempLen, maxCount)
	}
	fmt.Println(maxCount)
	// if maxCount == -1 {
	// 	return ""
	// }
	return strs[minIndexStr][:maxCount]

}
