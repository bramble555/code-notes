package leetcode150

func isHappy(n int) bool {
	seen := make(map[int]struct{}) // 用于记录出现过的数字

	for n != 1 {
		if _, ok := seen[n]; ok { // 如果当前数字已出现过，说明进入了循环
			return false
		}
		seen[n] = struct{}{}
		// 计算数字的平方和
		temp := 0
		for n > 0 {
			digit := n % 10
			temp += digit * digit
			n /= 10
		}
		n = temp
	}
	return true
}
