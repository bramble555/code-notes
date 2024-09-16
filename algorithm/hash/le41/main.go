package main

import (
	"fmt"
)

// 哈希解法
func firstMissingPositiveHash(nums []int) int {
	n := len(nums)
	numsHash := make(map[int]bool, n)
	// 初始化numsHash，当然也可以不初始化
	for i := 1; i <= n; i++ {
		numsHash[i] = false
	}
	for _, v := range nums {
		if v >= 1 && v <= n {
			numsHash[v] = true
		}
	}
	for i := 1; i <= n; i++ {
		if !numsHash[i] {
			return i
		}
	}
	// 都未出现
	return n + 1

}

// 标记法似乎实现不了
// 标记法 核心思想，把符合条件的数字(1-n) 例如2  把nums[1]=-nums[1]
// func firstMissingPositive(nums []int) int {
// 	n := len(nums)

// 	// 把不符合条件的数组都变为n+1，或者其他值
// 	for i := 0; i < n; i++ {
// 		if nums[i] < 1 || nums[i] > n {
// 			nums[i] = n + 1
// 		}
// 	}

// 	// 核心标记: 细节，一定是把正数变为负数，但是如果不+abs，也有可能实现负数变为了正数
// 	for i := 0; i < n; i++ {
// 		val := int(math.Abs(float64(nums[i])))
// 		if val <= n {
// 			nums[val-1] = -int(math.Abs(float64(nums[val-1])))
// 		}
// 	}
// 	fmt.Println(nums)
// 	// 查找第一个非负数的位置，即为所求的最小未出现的正整数
// 	for i := 0; i < n; i++ {
// 		if nums[i] > 0 {
// 			return i + 1
// 		}
// 	}
// 	// 如果所有位置都被标记，则说明1到n都已经存在
// 	return n + 1

// }
func main() {
	// fmt.Println(firstMissingPositive([]int{1, 1}))

	fmt.Println(firstMissingPositiveHash([]int{1, 1}))
}
