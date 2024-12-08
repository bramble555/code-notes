package leetcode150

import "fmt"

func rotate2(matrix [][]int) {
	n := len(matrix)
	count := 0
	for i := 0; i < n; i++ {
		for j := count; j < n; j++ {
			if i != j {
				matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
			}

		}
		count++
	}
	// n = 3
	// 0 2
	//
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
	fmt.Println(matrix)
}
