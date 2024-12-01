package leetcode150

func isValidSudoku(board [][]byte) bool {
	// 下标为 0-8
	// 填充的是 1-9
	flag := make([]bool, 9)
	// 判断每一行是否有重复
	for i := 0; i < 9; i++ {
		flag = make([]bool, 9)
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				t := board[i][j] - '1'
				if flag[t] {
					return false
				} else {
					flag[t] = true
				}
			}
		}
	}

	// 判断每一列是否有重复
	for i := 0; i < 9; i++ {
		flag = make([]bool, 9)
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' {
				t := board[j][i] - '1'
				if flag[t] {
					return false
				} else {
					flag[t] = true
				}
			}
		}
	}

	// 判断 3*3 是否有重复
	// i j 代表 每一个 3*3的左上角元素
	// k m 代表 每一个 3*3矩阵里面的行和列
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			flag = make([]bool, 9)
			for k := 0; k < 3; k++ {
				for m := 0; m < 3; m++ {
					if board[i+k][j+m] != '.' {
						t := board[i+k][j+m] - '1'
						if flag[t] {
							return false
						} else {
							flag[t] = true
						}
					}
				}
			}
		}
	}
	return true
}
