package leetcode150

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	// mapS 记录 某个字符或者数字上一次出现的位置
	mapS := make(map[uint8]int, n)
	var l int = 0
	res := 0
	for r := 0; r < n; r++ {
		// 先检查是否存在; 还要保证 左指针不能向左移动
		if v, exist := mapS[s[r]]; exist && v >= l {
			// 更新左指针
			l = v + 1

		}
		res = max(res, r-l+1)
		// 更新 s[r] 上一次出现的位置
		mapS[s[r]] = r

	}
	return res
}
