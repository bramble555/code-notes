package leetcode150

func twoSum2(nums []int, target int) []int {
	n := len(nums)
	if n < 2 {
		return []int{}
	}
	mapNum := make(map[int]int)
	for i := 0; i < n; i++ {
		dif := target - nums[i]
		if v, ok := mapNum[dif]; ok {
			return []int{i, v}
		} else {
			mapNum[nums[i]] = i
		}
	}
	return []int{}
}
