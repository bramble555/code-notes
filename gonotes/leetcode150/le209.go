package leetcode150

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)

	temp := 0
	res := n + 1
	i := 0
	for j := 0; j < n; j++ {
		// 一直向右滑动
		temp += nums[j]

		for temp >= target {
			res = min(res, j-i+1)
			temp -= nums[i]
			i++
		}
	}
	// 判断是否有满足条件(如果 n 个数字加起来都不大于 target)
	if res == n+1 {
		return 0
	}
	return res
}
