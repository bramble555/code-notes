package leetcode150

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	res := make([][]string, 0)
	if n == 0 {
		return res
	}
	mapStr := make(map[string][]string, 0)
	for i := 0; i < n; i++ {
		bytes := []byte(strs[i])
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		_, ok := mapStr[string(bytes)]
		if !ok {
			mapStr[string(bytes)] = []string{strs[i]}
		} else {
			mapStr[string(bytes)] = append(mapStr[string(bytes)], strs[i])
		}
	}
	for _, v := range mapStr {
		res = append(res, v)
	}
	return res
}
