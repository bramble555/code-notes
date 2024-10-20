package leetcode150

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pre := make([]int, n)
	suf := make([]int, n)
	res := make([]int, n)
	for {
		for {

		}
	}
	// 	[1,2,3,4]
	// pre	1  2  6 24
	// suf	24  24 12 4
	// res  24  12 8 6
	// 发现规律！！
	// 注意第一个数字和最后一个数字
	pre[0] = nums[0]
	for i := 1; i < n; i++ {
		pre[i] = nums[i] * pre[i-1]
	}
	suf[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suf[i] = nums[i] * suf[i+1]
	}
	res[0] = suf[1]
	res[n-1] = pre[n-2]
	for i := 1; i <= n-2; i++ {
		res[i] = pre[i-1] * suf[i+1]
	}
	return res
}
