package main

import (
	"fmt"
)

// 双指针算法
// 假设arr := []int{-10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 如果从小到大从前往后加入到新数组中，那么81，64..... 显然是错的
// 应该从大到小从后往前加入到新数组中，那么，从后往前依次是100,81,64,49...

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	i := 0
	j := len(nums) - 1
	k := len(nums) - 1 //从大到小从后往前加入到新数组的下标
	for i <= j {
		leftVal := nums[i] * nums[i]
		rightVal := nums[j] * nums[j]
		// 比较左边值大还是右边值大
		if leftVal >= rightVal {
			res[k] = leftVal
			k--
			i++
		} else {
			res[k] = rightVal
			k--
			j--
		}
	}
	return res
}
func main() {
	arr := []int{-10, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sortedSquares(arr))
}
