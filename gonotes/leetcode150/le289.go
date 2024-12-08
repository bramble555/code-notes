package leetcode150

// live -> dead -2
// dead -> live 2
// neighbor
var neighbor = [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

func gameOfLife(board [][]int) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 活着
			if board[i][j] == 1 {
				num := count(board, i, j)
				if num > 3 || num < 2 {
					board[i][j] = -2
				}
			} else {
				num := count(board, i, j)
				if num == 3 {
					board[i][j] = 2
				}
			}
		}
	}
	update(board)
}
func count(board [][]int, row, col int) int {
	m := len(board)
	n := len(board[0])
	aliveCount := 0
	for _, offset := range neighbor {
		newRow, newCol := row+offset[0], col+offset[1]
		if newRow >= 0 && newRow < m && newCol >= 0 && newCol < n {
			if board[newRow][newCol] == 1 || board[newRow][newCol] == -2 {
				aliveCount++
			}
		}
	}
	return aliveCount

}
func update(board [][]int) {
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 活着
			if board[i][j] == 2 {
				board[i][j] = 1
			}
			if board[i][j] == -2 {
				board[i][j] = 0
			}
		}
	}
}
