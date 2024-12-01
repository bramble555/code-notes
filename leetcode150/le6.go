package leetcode150

import "fmt"

func convert(s string, numRows int) string {
	n := len(s)
	flag := make([]int, 0)

	for {
		for j := 1; j <= numRows; j++ {
			flag = append(flag, j)
			if len(flag) == n {
				goto EndLoop
			}
		}
		for j := numRows - 1; j >= 2; j-- {
			flag = append(flag, j)
			if len(flag) == n {
				goto EndLoop
			}
		}
	}
EndLoop:
	// 	 	a b c d e f
	// 		1 2 3 2 1 2
	// 输出	a e b d f c
	res := make([]rune, 0)
	for j := 1; j <= numRows; j++ {
		for i := 0; i < n; i++ {
			if flag[i] == j {
				res = append(res, rune(s[i]))
			}
		}
	}
	fmt.Println(string(res))
	return string(res)

}
