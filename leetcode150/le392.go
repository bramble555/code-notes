package leetcode150

func isSubsequence(s string, t string) bool {
	// 	// 判断 t 是否包含了 s
	n1 := len(s)
	n2 := len(t)
	if n2 < n1 {
		return false
	}
	if n1 == 0 {
		return true
	}
	dp := make([]int, n2)
	// i 指向 s
	i := 0
	// 初始化 dp[1]
	if t[0] == s[0] {
		dp[0] = 1
		i++
	}
	for j := 1; j < n2; j++ {
		if t[j] == s[i] {
			dp[j] = dp[j-1] + 1
			i++
		}
		dp[j] = dp[j-1]
		if i == n1 || dp[j] == n1 {
			return true
		}
	}
	return false
}

// func isSubsequence(s string, t string) bool {
// 	// 判断 t 是否包含了 s
// 	n1 := len(s)
// 	n2 := len(t)
// 	if n2 < n1 {
// 		return false
// 	}
// 	if n1 == 0 {
// 		return true
// 	}
// 	// 双指针 i 指向 s ，j 指向 t
// 	i := 0
// 	j := 0
// 	for i < n1 && j < n2 {
// 		if s[i] == t[j] {
// 			i++
// 			j++
// 		} else {
// 			j++
// 		}
// 		// i == n1 的时候，说明包含了
// 		if i == n1 {
// 			return true
// 		}
// 		if j == n2 {
// 			return false
// 		}
// 	}
// 	return false
// }
