package leetcode150

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if major != nums[i] {
			if count > 0 {
				count--
				continue
			}
			// count = 0 的情况
			major = nums[i]
			count = 1
		} else {
			count++
		}
	}
	return major
}
