package backtracking

import "math"

// input: an integer representing a n by n chessboard
// output: all possible arrangements of placing n queens such that the queens do not attack each other

// checker: cross and x check
// used 2d array: n by n; when a check is performed it marks all the boxes made unavailable
// iterate box by box
func resetBoard(n int) [][]bool {
	board := make([][]bool, n)
	for i:=0; i<n; i++ {
		board[i] = make([]bool, n)
	}
	return board
}

func nqueens_drive(n int) bool {
	useds := [][][]bool{}
	board := resetBoard(n)
	used := make([]bool, n)
	for i:=0; i<n; i++ {
		board[i] = make([]bool, n)
	}

	for i:=0; i<n; i++ {
		board[0][i]=true
		if !nqueens(n, &board,used) {
			return false
		} else {
			useds = append(useds, board)
		}
		board=resetBoard(n)
	}

	return true
}

func nqueens(n int, board *[][]bool, used []bool) bool {
	for row:=0; row<n; row++ {
		if !used[row] { 
			used[row] = true

			for col:=0; col<n; col++ {
				if !(*board)[row][col] && !isUsed(row, col, board) {
					(*board)[row][col] = true
					if nqueens(n, board, used){
						return true
					}
				}
			}

			used[row] = false
			return false
		}
	}

	return false
}

func isUsed(currentRow int, currentCol int, used *[][]bool) bool {

	// check cross sectionaly
	for row:=0; row<len(*used); row++ {
		for col:=0; col<len((*used)[row]); col++ {

			if col==currentCol {
				return (*used)[row][col]
			}

			if row==currentRow {
				return (*used)[row][col]
			}
			
			colDiff:=col-currentCol
			rowDiff:=row-currentRow

			if math.Abs(float64(colDiff)) == math.Abs(float64(rowDiff)) {
				return (*used)[row][col]
			}

		}
	}
	return false
}

