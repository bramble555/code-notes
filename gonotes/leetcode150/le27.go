package leetcode150

func removeElement(nums []int, val int) int {
	slow := 0
	fast := 0
	// 返回值
	count := 0
	n := len(nums)
	for fast < n {
		if nums[fast] == val {
			fast++
			continue

		}
		nums[slow] = nums[fast]
		count++
		slow++
		fast++
	}
	return count
}
