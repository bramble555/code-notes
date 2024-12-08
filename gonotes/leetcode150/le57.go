package leetcode150

func insert(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	if n == 0 {
		return [][]int{newInterval}
	}
	res := make([][]int, 0)
	for i := 0; i < n; i++ {
		if intervals[i][0] > newInterval[1] {
			res = append(res, newInterval)
			for j := i; j < n; j++ {
				res = append(res, intervals[j])

			}
			return res
		} else if intervals[i][1] < newInterval[0] {
			res = append(res, intervals[i])
		} else {
			newInterval[0] = min(intervals[i][0], newInterval[0])
			newInterval[1] = max(intervals[i][1], newInterval[1])
		}
	}
	res = append(res, newInterval)
	return res
}
