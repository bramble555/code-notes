package leetcode150

func canJump(nums []int) bool {
	n := len(nums)
	flag := make([]bool, n)
	flag[0] = true
	// [2,3,1,1,4]
	// i = 3 temp = 1  j = 4 j <= 3+1=4

	// [1 0]
	// i = 0 temp =1

	// [2 0 0]
	// i = 1 temp = 0
	// i = 0 temp = 2

	for i := n - 2; i >= 0; i-- {
		temp := nums[i]
		for j := i + 1; j <= i+temp && j < n; j++ {
			flag[j] = true
		}
	}
	// 遍历整个 flag
	for i := 0; i < n; i++ {
		if !flag[i] {
			return false
		}
	}
	return true
}
