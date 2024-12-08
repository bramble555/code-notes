package leetcode150

// -1,-100,3,99 k = 2
//  3 99 -1  -100
// 发现规律了吗，反转前 K 个元素 ，再反转后 n-K 个元素
func rotate(nums []int, k int) {
	// 有一种特殊情况， k > len(nums) ,所以一开始要对 k 取余操作
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)

}
func reverse(nums []int, left, right int) {
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
