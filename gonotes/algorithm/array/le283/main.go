package main

import "fmt"

// 思路：和remove-element一样，快慢指针，只是把最后面值赋值为0
func moveZeroes(nums []int) []int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
	return nums

}
func main() {
	fmt.Println(moveZeroes([]int{0}))
}
