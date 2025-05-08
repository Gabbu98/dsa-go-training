package backtracking

import "math"

// input: an integer representing a n by n chessboard
// output: all possible arrangements of placing n queens such that the queens do not attack each other

// checker: cross and x check
// used 2d array: n by n; when a check is performed it marks all the boxes made unavailable
// iterate box by box
func nqueens_drive(n int) {
	useds := [][][]bool{}
	board := make([][]string, n)
	used := make([][]bool, n)
	for i:=0; i<n; i++ {
		board[i] = make([]string, n)
		used[i] = make([]bool, n)
		for j:=0; j<n; j++ {
			used [i][j] = false
			board [i][j] = "empty"
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			used[i][j] = true
			board[i][j] = "queen"
			if nqueens(n,0,&used,&board){
				useds = append(useds, used)
			}
			used[i][j] = false
			board[i][j] = "empty"
		}
	}
}

func nqueens(n int, nQueensCounter int, used *[][]bool, result *[][]string) bool{

	
	for row:=0; row<n; row++ {
		for col:=0; col<n; col++ {
			
			if !isUsed(row, col, used) && !(*used)[row][col] {
				(*used)[row][col] = true
				(*result)[row][col] = "queen"
				nQueensCounter++

				if nQueensCounter==n{
					return true
				}

				if nqueens(n,nQueensCounter,  used, result) {
					return true
				} else {
					nQueensCounter--
					(*used)[row][col] = false
					(*result)[row][col] = "empty"
					return false
				}
				
			}

		}
	}

	return true

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