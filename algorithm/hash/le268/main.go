package main

import (
	"fmt"
	"math"
)

// 找出 [0, n] 这个范围内没有出现在数组中的那个数。
// 哈希算法
func missingNumberHash(nums []int) int {
	n := len(nums)
	numsHash := make(map[int]bool, n+1)
	for _, v := range nums {
		if v >= 0 && v <= n {
			numsHash[v] = true
		}
	}
	for i := 0; i < n+1; i++ {
		if flag, _ := numsHash[i]; !flag {
			return i
		}

	}
	return n
}

// 标记算法
func missingNumber(nums []int) int {

	n := len(nums)
	if n == 1 && nums[0] == 0 {
		return 1
	}
	// 把不符合条件的数组都变为n，或者其他值
	for i := 0; i < n; i++ {
		if nums[i] < 0 || nums[i] > n-1 {
			nums[i] = n
		}
	}

	// 核心标记: 细节，一定是把正数变为负数，但是如果不+abs，也有可能实现负数变为了正数
	for i := 0; i < n; i++ {
		val := int(math.Abs(float64(nums[i])))
		if val < n {
			nums[val] = -int(math.Abs(float64(nums[val])))
		}
	}
	// 查找第一个非负数的位置，即为所求的最小未出现的正整数
	for i := 0; i < n; i++ {
		if i == n-1 && nums[i] == 0 {
			return i
		}
		if nums[i] > 0 {
			return i
		}
	}
	// 如果所有位置都被标记，则说明0到n-1都已经存在
	return n

}
func main() {
	// fmt.Println(missingNumberHash([]int{3, 0, 1, 2}))
	fmt.Println(missingNumber([]int{0}))
}
