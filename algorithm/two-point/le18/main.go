package main

import (
	"fmt"
	"sort"
)

// 和三数之和一样，只是多加一层for循环
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := make([][]int, 0)
	if n < 4 {
		return res
	}
	sort.Ints(nums)

	// i后面至少要有三个数字，n-3,下标从0开始再-1
	for i := 0; i <= n-4; i++ {
		// i去重
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j <= n-3; j++ {
			// j去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 1
			right := n - 1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					// 加入结果
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					l, r := nums[left], nums[right]
					// left和right继续互相靠拢,去重

					for left < right && l == nums[left] {
						left++
					}
					for left < right && r == nums[right] {
						right--
					}
				}
			}
		}
	}
	return res
}
func main() {
	nums := []int{1, 0, -1, 0, 2, -2}
	// 排序后结果为 -2 -1 0 0 1 2
	target := 0
	fmt.Println(fourSum(nums, target))
}
