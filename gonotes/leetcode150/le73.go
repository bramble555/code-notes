package leetcode150

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])
	// 记录 0 出现的行
	rows := make(map[int]struct{}, 0)
	// 记录 0 出现的列
	columns := make(map[int]struct{}, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = struct{}{}
				columns[j] = struct{}{}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果该行或该列包含 0，设置该位置为 0
			if _, ok1 := rows[i]; ok1 {
				matrix[i][j] = 0
				continue
			} else if _, ok2 := columns[j]; ok2 { // 如果行检查未通过，则检查列
				matrix[i][j] = 0
			}
		}
	}
}
