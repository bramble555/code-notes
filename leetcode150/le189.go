package leetcode150

func rotate(nums []int, k int) {
	// 有一种特殊情况， k > len(nums) ,所以一开始要对 k 取余操作
	// 以下方法属于常规做法，当 len(nums) 超级大会超时
	k %= len(nums)
	temps := make([]int, k)
	tempsIdx := 0
	n := len(nums)
	// 初始化 temps
	// n = 7  k = 3 i = 3 的 时候 break
	// n = 4  k = 2 i = 1 的 时候 break
	// 条件也可以改为 i > n-1-k
	for i := n - 1; tempsIdx < k; i-- {
		temps[tempsIdx] = nums[i]
		tempsIdx++
	}
	tempsIdx = 0
	// 移动的次数
	for i := 0; i < k; i++ {
		// 均向后移动一个元素
		for j := n - 2; j >= 0; j-- {
			nums[j+1] = nums[j]
		}
		nums[0] = temps[tempsIdx]
		tempsIdx++
	}
}
