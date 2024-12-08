package leetcode150

func candy(ratings []int) int {
	n := len(ratings)
	res := 0
	candyCount := make([]int, n)
	candyCount[0] = 1
	// 从前往后遍历寻找递增
	for i := 1; i < n; i++ {
		// 初始化 res[i]
		candyCount[i] = 1
		if ratings[i-1] < ratings[i] {
			candyCount[i] = max(candyCount[i-1]+1, candyCount[i])
		}
	}
	// 从后往前遍历寻找递减
	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candyCount[i] = max(candyCount[i+1]+1, candyCount[i])
		}
		res += candyCount[i]
	}
	res += candyCount[n-1]
	return res
}
