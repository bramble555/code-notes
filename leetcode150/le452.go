package leetcode150

import (
	"slices"
)

func findMinArrowShots(points [][]int) int {
	n := len(points)

	if n == 0 {
		return 0
	}
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})
	res := 1
	axis := points[0][1]
	for i := 1; i < n; i++ {
		// 查看当前 i 与后面的是有交集,没有交集 res + 1
		if axis < points[i][0] {
			res++
			axis = points[i][1]
		}
	}
	return res
}
