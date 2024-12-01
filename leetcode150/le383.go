package leetcode150

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	// ransomNote 范围小，magazine 范围大
	n := len(ransomNote)
	m := len(magazine)
	if n > m {
		return false
	}
	mapStr := make(map[uint8]int)
	// 把范围大的增加到 mapStr 里面
	for i := 0; i < m; i++ {
		_, ok := mapStr[magazine[i]]
		if ok {
			mapStr[magazine[i]]++
		}
		if !ok {
			mapStr[magazine[i]] = 1
		}
	}
	fmt.Println(mapStr)
	// 把范围小的 --
	for i := 0; i < n; i++ {
		_, ok := mapStr[ransomNote[i]]
		if ok {
			mapStr[ransomNote[i]]--
			if mapStr[ransomNote[i]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	fmt.Println(mapStr)
	return true
}
