package backtracking

func SolveNQueens(n int) [][]string {
	solutions := [][]string{}
	board := make([][]rune, n)
	for i := range board {
		board[i] = make([]rune, n)
		for j:=range board[i] {
			board[i][j] = '.'
		}
	}
	nqueens(n,0,board,&solutions)
	return solutions
}

func nqueens(n int, row int, board [][]rune, solutions *[][]string) {
	if row == n {
		solution := make([]string, n)
		for i:=range board {
			solution[i] = string(board[i])
		}
		*solutions = append(*solutions, solution)
		return
	}

	for col:=0; col<n; col++ {
		if isSafe(row, col, board, n) {
			board[row][col] = 'Q'
			nqueens(n, row+1,board,solutions)
			board[row][col]='.'
		}
	}
}

func isSafe(row int, col int, board [][]rune, n int) bool {
	// check column

	// chec upper left diagonal 

	// check upper right diagonal
}