package le1004

// 滑动窗口  把0翻转为1，求1最长的数量
// 1004 -> 2024
func longestOnes(nums []int, k int) int {
	res := 0
	n := len(nums)
	flipNum := 0
	lo := 0
	hi := 0
	for ; hi < n; hi++ {
		if nums[hi] == 0 {
			flipNum++
		}
		if flipNum > k {
			for nums[lo] == 1 {
				lo++
			}
			lo++
			flipNum--
		}
		res = max(res, hi-lo+1)
	}
	return res
}
func main() {
	longestOnes([]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2)
}
