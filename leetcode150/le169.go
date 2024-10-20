package leetcode150

func majorityElement(nums []int) int {
	n := len(nums)
	numsMap := make(map[int]int, n)
	// for i := 0; i < n; i++ {
	// 	numsMap[nums[i]] = 0
	// }
	// for i := 0; i < n; i++ {
	// 	numsMap[nums[i]]++
	// }
	// for num, v := range numsMap {
	// 	if v > n/2 {
	// 		return num
	// 	}
	// }

	// 上面太冗余了

	for _, v := range nums {
		numsMap[v]++
		if numsMap[v] > n/2 {
			return v
		}
	}

	return nums[0]
}
