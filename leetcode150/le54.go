package leetcode150

func spiralOrder(matrix [][]int) []int {
	// 行
	m := len(matrix)
	// 列
	n := len(matrix[0])
	res := make([]int, 0, m*n)
	left := 0
	right := n - 1
	top := 0
	bottom := m - 1
	for left <= right && top <= bottom {
		// 最上面元素
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		// 只有一行的情况
		if top > bottom {
			break
		}
		// 最右边元素
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--
		// 只有一列情况
		if left > right {
			break
		}
		// 最底下面元素
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--
		if top > bottom {
			break
		}
		// 最左边元素
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}
