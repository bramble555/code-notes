package main

import (
	"fmt"
)

// 如果排序，那么不能保证子数组是连续的，审题的时候忽略了...
// 213 [12,28,83,4,25,26,25,2,25,25,25,12] 不能通过

// 暴力解法
func minSubArrayLen2(target int, nums []int) int {
	var sum int
	var count int        // 用于记录每一次找到>=target所需要数字
	min := len(nums) + 1 // 用于记录结果
	for i := 0; i < len(nums); i++ {
		// 每次初始化count sum
		sum = 0
		count = 0
		for j := i; j < len(nums); j++ {
			count++
			sum += nums[j]
			if sum >= target {
				if count < min {
					min = count
				}
				break
			}
		}

	}
	if min == len(nums)+1 {
		min = 0
	}
	return min
}

// 正解是滑动窗口
// 双指针 i代表起始位置,j代表结束位置
// 一开始j一直增加直到满足条件
// 此时i一直减直到不满足条件
// 循环上述俩个步骤
// 注意i一直减是在j++的循环的里面，不是并列if,是for
func minSubArrayLen(target int, nums []int) int {
	i := 0
	j := 0
	sum := 0
	var res int = len(nums) + 1 // 要返回的满足条件的最小值
	for ; j < len(nums); j++ {
		// i向右滑动过程
		sum += nums[j]
		// 如果满足条件，记录下len,此时i一直减直到不满足条件
		for sum >= target {
			res = min(res, j-i+1)
			sum -= nums[i]
			i++
		}
	}
	// 没有找到满足条件的时候
	if res == len(nums)+1 {
		res = 0
	}
	return res

}
func minSubArrayLenc(target int, nums []int) int {
	i := 0
	j := 0
	sum := nums[0]
	var res int = len(nums) + 1 // 要返回的满足条件的最小值
	for i < j && j < len(nums) {
		// 不满足条件 j++
		if sum < target {
			j++
			sum += nums[j]
		} else {
			// 记录满足条件的时候min
			res = min(res, j-i+1)
			i++
			sum -= nums[i]
		}

	}
	// 没有找到满足条件的时候
	if res == len(nums)+1 {
		return 0
	}
	return res
}
func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, nums))
}
