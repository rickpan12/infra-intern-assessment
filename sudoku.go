package main

func isValid(board [][]int, row, col int, ch int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == ch || board[row][i] == ch || board[3*(row/3)+i/3][3*(col/3)+i%3] == ch {
			return false
		}
	}
	return true
}

func solve(board [][]int) [][]int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				for ch := 0; ch <= 9; ch++ {
					if isValid(board, i, j, ch) {
						board[i][j] = ch
						if result := solve(board); result != nil {
							return result
						}
						board[i][j] = 0
					}
				}
				return nil
			}
		}
	}
	return board
}

func SolveSudoku(board [][]int) [][]int {
	return solve(board)
}