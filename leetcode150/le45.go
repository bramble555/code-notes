package leetcode150

import "math"

func jump(nums []int) int {
	n := len(nums)
	res := make([]int, n)
	// 给 res 切片初始值
	for j := 1; j < n; j++ {
		res[j] = math.MaxInt
	}
	res[0] = 0
	for i := 0; i < n; i++ {
		for j := i; j <= nums[i]+i && j < n; j++ {
			res[j] = min(res[i]+1, res[j])
		}
	}
	return res[n-1]
}
