package leetcode150

func removeDuplicates(nums []int) int {
	n := len(nums)
	// 返回值
	if n == 1 || n == 0 {
		return n
	}
	slow := 1
	fast := 1
	for ; fast < n; fast++ {
		if nums[fast] != nums[slow-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
