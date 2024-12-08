package leetcode150

func containsNearbyDuplicate(nums []int, k int) bool {
	n := len(nums)
	mapNum := make(map[int]int, n) // 初始化map的容量
	for i := 0; i < n; i++ {
		if j, ok := mapNum[nums[i]]; ok && i-j <= k {
			return true
		}
		mapNum[nums[i]] = i
	}
	return false
}
