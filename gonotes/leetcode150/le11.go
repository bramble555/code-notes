package leetcode150

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	res := 0
	for i < j {
		var temp int
		if height[i] <= height[j] {
			temp = height[i] * (j - i)
			i++
		} else {
			temp = height[j] * (j - i)
			j--
		}
		res = max(res, temp)
	}
	return res
}
