package leetcode150

import "slices"

func merge2(intervals [][]int) [][]int {
	n := len(intervals)
	res := make([][]int, 0, n)
	if n == 0 {
		return res
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0] // 按第一个元素排序
	})
	tempStart := intervals[0][0]
	tempEnd := intervals[0][1]

	for i := 1; i < n; i++ {
		// 区间有重合
		if intervals[i][0] <= tempEnd {
			tempEnd = max(tempEnd, intervals[i][1])
		} else {
			res = append(res, []int{tempStart, tempEnd})
			tempStart = intervals[i][0]
			tempEnd = intervals[i][1]
		}
	}
	// 加入最后一组元素
	res = append(res, []int{tempStart, tempEnd})
	return res
}
