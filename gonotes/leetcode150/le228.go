package leetcode150

import (
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {
	n := len(nums)
	res := make([]string, 0, n)
	if n == 0 {
		return res
	}
	for i := 0; i < n; {
		first := nums[i]

		// 判断当前元素与下一个元素是否连续
		for i < n-1 && nums[i+1] == nums[i]+1 {
			i++
		}
		// 此时不连续了， 1 2 3 5
		// 		  i: 0 1 2

		// 只有一个元素的情况
		if first == nums[i] {
			res = append(res, strconv.Itoa(first))
		} else {
			var sb strings.Builder
			sb.WriteString(strconv.Itoa(first))
			sb.WriteString("->")
			sb.WriteString(strconv.Itoa(nums[i]))
			res = append(res, sb.String())
		}
		i++
	}

	return res
}
