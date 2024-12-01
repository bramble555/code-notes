package leetcode150

func longestConsecutive(nums []int) int {
	n := len(nums)
	setNums := make(map[int]struct{}, n)
	for i := 0; i < n; i++ {
		setNums[nums[i]] = struct{}{}
	}
	maxLen := 0
	for num, _ := range setNums {
		curLen := 1

		// this num is the first of the group
		if _, ok := setNums[num-1]; !ok {
			searchNum := num + 1
			// search the sequence num of the num
			for {
				if _, ok := setNums[searchNum]; ok {
					curLen++
					searchNum++
				} else {
					break
				}
			}
		}
		maxLen = max(curLen, maxLen)
	}
	return maxLen
}
