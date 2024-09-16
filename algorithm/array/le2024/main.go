package main

import "fmt"

func maxConsecutiveAnswers(answerKey string, k int) int {
	// 把F换为T
	keyBytes := []byte(answerKey)
	n := len(keyBytes)
	if n == 1 {
		return 1
	}
	// 当前翻转次数
	flipNum := 0
	res := 0
	// 滑动窗口
	low := 0
	hi := 0
	for ; hi < n; hi++ {
		if keyBytes[hi] == 'F' {
			flipNum++
		}
		// 左窗口向右滑动
		if flipNum > k {
			for keyBytes[low] == 'T' {
				low++
			}
			low++
			flipNum--

		}
		res = max(res, hi-low+1)
	}
	// 把T换为F
	flipNum = 0
	low = 0
	hi = 0
	for ; hi < n; hi++ {
		if keyBytes[hi] == 'T' {
			flipNum++
		}
		if flipNum > k {
			for keyBytes[low] == 'F' {
				low++
			}
			low++
			flipNum--
		}
		res = max(res, hi-low+1)
	}
	return res

}
func main() {
	fmt.Println(maxConsecutiveAnswers("TFFT", 1))

}
