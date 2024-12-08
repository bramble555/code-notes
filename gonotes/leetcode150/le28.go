package leetcode150

func strStr(haystack string, needle string) int {
	n1 := len(haystack)
	n2 := len(needle)
	if n2 > n1 {
		return -1
	}
	// haystack: a b c
	// needle:	   a b
	// 			a b 不成立
	// n1 = 3 n2 = 2
	for i := 0; i <= n1-n2; i++ {
		for j := 0; j < n2; j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if haystack[i+j] == needle[j] && j == n2-1 {
				return i
			}
		}
	}
	return -1
}
