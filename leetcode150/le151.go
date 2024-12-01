package leetcode150

import (
	"strings"
)

func reverseWords(s string) string {
	// 删除俩边字符串
	s = strings.Trim(s, " ")
	sArrTemp := strings.Split(s, " ")
	sArr := make([]string, 0)
	for i := 0; i < len(sArrTemp); i++ {
		if sArrTemp[i] != "" {
			sArr = append(sArr, sArrTemp[i])
		}
	}
	// 4   0 1  反转俩次
	// 5   0 1  反转俩次
	for i := 0; i <= len(sArr)/2; i++ {
		swap(&sArr, i, len(sArr)-1-i)
	}
	return strings.Join(sArr, " ")
}
func swap(s *[]string, i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
