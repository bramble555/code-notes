package leetcode150

import (
	"slices"
)

func hIndex(citations []int) int {
	n := len(citations)
	slices.Sort(citations)
	var res int = 0
	// 3, 0, 6, 1, 5
	// 0 1 3 5 6
	//     3 2 1
	for i := n - 1; i >= 0; i-- {
		if citations[i] > res {
			res++
		}
	}
	return res

}
