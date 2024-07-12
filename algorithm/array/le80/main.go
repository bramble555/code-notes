package main

import "fmt"

// 快慢指针
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	// 0和1俩个下标的元素值一定不会超过俩次
	slow := 2
	fast := 2
	for ; fast < len(nums); fast++ {
		// 判断条件是nums[fast] != nums[slow-2]
		// 因为当nums[fast] == nums[slow-2] 的时候，nums[fast] == nums[slow-1]
		// 此时不满足条件，其他的时候，均满足条件
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
func main() {
	nums := []int{0, 0, 1, 1, 1, 2}
	fmt.Println(removeDuplicates(nums))
}
