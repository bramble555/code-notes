package leetcode150

func removeDuplicates2(nums []int) int {
	n := len(nums)
	// slow 指的是下一个要存的数组的索引 (也就是数组的长度，因为索引从0开始)
	slow := 2
	fast := 2
	if n <= 2 {
		return n
	}
	for ; fast < n; fast++ {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}

	}
	return slow
}
