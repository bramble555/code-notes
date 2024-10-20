package leetcode150

func maxProfit2(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	var res int
	for i := 1; i < n; i++ {
		res += max(0, prices[i]-prices[i-1])
	}
	return res
}
