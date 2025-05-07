package backtracking

import "math"

// input: an integer representing a n by n chessboard
// output: all possible arrangements of placing n queens such that the queens do not attack each other

// checker: cross and x check
// used 2d array: n by n; when a check is performed it marks all the boxes made unavailable
// iterate box by box


func nqueens(n int, used *[][]bool, result *[][]string) bool{
	
	for row:=0; row<n; row++ {
		for col:=0; col<n; col++ {
			
			if !isUsed(row, col, used) {
				(*used)[row][col] = true
				(*result)[row][col] = "queen"

				if nqueens(n, used, result) {
					return true
				} else {
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
	if((*used)[row][col]){
		return true
	}

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
}