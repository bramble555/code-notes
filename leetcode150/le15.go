package leetcode150

import "slices"

func threeSum(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	// 排序方便去重
	slices.Sort(nums)
	// 思路 固定 i ， j 和 k 向中间移动
	// 加入结果的之后去重(改变i 和 j)
	for i := 0; i < n-2; i++ {
		// 每次重置 i 和 j
		j := i + 1
		k := n - 1
		if i > 0 && nums[i] == nums[i-1] {
			// 让 i 继续 + 1，但是不能写 i++ ， 为什么呢？
			// 写 i++，又让 j < k 这个循环启动了，可能 i此时已经不满足 i < n-2 了
			continue
		}
		for j < k {
			// 太小了
			if nums[i]+nums[j]+nums[k] < 0 {
				j++
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--
			} else {
				// 记录下来 nums[j], nums[k]
				l := nums[j]
				r := nums[k]
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == l {
					j++
				}
				for j < k && nums[k] == r {
					k--
				}
			}
		}
	}
	return res
}
