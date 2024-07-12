package main

import (
	"fmt"
)

// 暴力双层for循环,假设每种输入只会对应一个答案,但是没考虑到一种输入对应多个答案的情况，leetcode上面也是对应一个答案，能通过
func twoSum2(nums []int, target int) []int {
	res := make([]int, 0)
	// flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res = append(res, i, j)
				// flag = true
			}
		}
		if len(res) == 2 {
			break
		}
	}
	return res
}

// 正确做法是hashmap,因为要查找target - nums[i] 的值是否遍历过，如果遍历过，那么就找到了
// 为什么使用map呢，因为要返回对应的下标，如果返回其值，用set
// 为什么map[v][i] 这样存放呢？ 因为要查找的是v，返回的i
// 思路:遍历nums,然后添加到map里面，再查找target - nums[i]是否有， 有个小疑问，相同的数字不就覆盖了？[3,3] target = 6,那么就没有值？？
// 小疑问解答：并不是每次先把nums[i]添加到map里面，而是先查找map是否target - nums[i] 的值，如果有则添加当前索引和之前的索引
func twoSum1(nums []int, target int) []int {
	res := make([]int, 0)
	mapNums := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 查找target - nums[i]是否有
		need := target - nums[i]
		if preIndex, ok := mapNums[need]; ok {
			res = append(res, i, preIndex)
		}
		// 是否有结果
		mapNums[nums[i]] = i
		if len(res) == 2 {
			break
		}
	}
	return res
}

// 小优化，不用建立res
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, 0)
	for i, v := range nums {
		if pre, ok := numsMap[target-v]; ok {
			return []int{pre, i}
		}
		numsMap[v] = i
	}
	return []int{}
}

// 不能用双指针法。如果用map保存v-i，v会有重复的，i被覆盖了
func main() {
	nums := []int{1, 2, 2, 3}
	fmt.Println(twoSum2(nums, 4))
	fmt.Println(twoSum1(nums, 4))
	fmt.Println(twoSum(nums, 4))

}
