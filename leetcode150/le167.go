package leetcode150

func twoSum(numbers []int, target int) []int {
	n := len(numbers)
	var i int = 0
	var j int = n - 1
	for i < j {
		// 数字太小
		if numbers[i]+numbers[j] < target {
			i++
		}
		// 数字太大
		if numbers[i]+numbers[j] > target {
			j--
		}
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		}

	}
	return []int{0, 0}
}
