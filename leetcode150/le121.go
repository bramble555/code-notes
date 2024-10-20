package leetcode150

import "math"

// func maxProfit2(prices []int) int {
// 	var min int
// 	var max int
// 	n := len(prices)
// 	if n == 1 {
// 		return 0
// 	}
// 	for i := 0; i < n-1; i++ {
// 		min = prices[i]
// 		for j := i + 1; j < n; j++ {
// 			max = maxNums(prices[j]-min, max)
// 		}
// 	}
// 	if max < 0 {
// 		max = 0
// 	}
// 	return max
// }
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 1 {
		return 0
	}
	maxPurchase := math.MaxInt64
	maxProfit := 0
	for i := 0; i < n; i++ {
		maxPurchase = minNums(maxPurchase, prices[i])
		// 计算此时最大利润
		maxProfit = maxNums(maxProfit, prices[i]-maxPurchase)
	}
	return maxProfit
}
func minNums(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxNums(a, b int) int {
	if a > b {
		return a
	}
	return b
}
