package backtracking

import (
	"fmt"
	"reflect"
)

func sudoku_driver(board [][]int) [][]int {
	sudoku(&board)

	return board
}

func sudoku(board *[][]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if (*board)[row][col] == 0 {
				for val := 1; val <= 9; val++ {
					if isUsable(board, row, col, val) {
						(*board)[row][col] = val
						if sudoku(board) {
							return true
						}
						(*board)[row][col] = 0 // backtrack
					}
				}
				return false // if no valid number found
			}
		}
	}
	return true // solved
}

func isUsable(board *[][]int, row int, col int, value int) bool {
	// Check row
	for i := 0; i < 9; i++ {
		if (*board)[row][i] == value {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if (*board)[i][col] == value {
			return false
		}
	}

	// Check 3x3 box
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if (*board)[startRow+i][startCol+j] == value {
				return false
			}
		}
	}

	return true
}


// TestSudoku tests solution(s) with the following signature and problem description:

// 	func Sudoku(board [][]int) bool

// Given a partially filled 9x9 grid with integers from 1 to 9 representing a Sudoku
// board and 0 representing an empty slot that needs to be filled, write a function
// that solves the board such that the values in each row, column and each of
// the 9 3x3 sub-grids are unique.
// */
func TestSudoku() {
	tests := []boardAndSolution{
		testBoard1(),
		testBoard2(),
		testBoard3(),
		testBoard4(),
	}

	for i, test := range tests {
		if sudoku_driver(test.board); !reflect.DeepEqual(test.solution, test.board) {
			fmt.Printf("Failed test case #%d. Want %#v got %#v", i, test.solution, test.board)
		}
	}
}

type boardAndSolution struct {
	board    [][]int
	solution [][]int
}

func testBoard1() boardAndSolution {
	return boardAndSolution{
		board: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		solution: [][]int{
			{1, 2, 3, 4, 5, 6, 7, 8, 9},
			{4, 5, 6, 7, 8, 9, 1, 2, 3},
			{7, 8, 9, 1, 2, 3, 4, 5, 6},
			{2, 1, 4, 3, 6, 5, 8, 9, 7},
			{3, 6, 5, 8, 9, 7, 2, 1, 4},
			{8, 9, 7, 2, 1, 4, 3, 6, 5},
			{5, 3, 1, 6, 4, 2, 9, 7, 8},
			{6, 4, 2, 9, 7, 8, 5, 3, 1},
			{9, 7, 8, 5, 3, 1, 6, 4, 2},
		},
	}
}

func testBoard2() boardAndSolution {
	return boardAndSolution{
		board: [][]int{
			{1, 0, 6, 0, 0, 2, 3, 0, 0},
			{0, 5, 0, 0, 0, 6, 0, 9, 1},
			{0, 0, 9, 5, 0, 1, 4, 6, 2},
			{0, 3, 7, 9, 0, 5, 0, 0, 0},
			{5, 8, 1, 0, 2, 7, 9, 0, 0},
			{0, 0, 0, 4, 0, 8, 1, 5, 7},
			{0, 0, 0, 2, 6, 0, 5, 4, 0},
			{0, 0, 4, 1, 5, 0, 6, 0, 9},
			{9, 0, 0, 8, 7, 4, 2, 1, 0},
		},
		solution: [][]int{
			{1, 4, 6, 7, 9, 2, 3, 8, 5},
			{2, 5, 8, 3, 4, 6, 7, 9, 1},
			{3, 7, 9, 5, 8, 1, 4, 6, 2},
			{4, 3, 7, 9, 1, 5, 8, 2, 6},
			{5, 8, 1, 6, 2, 7, 9, 3, 4},
			{6, 9, 2, 4, 3, 8, 1, 5, 7},
			{7, 1, 3, 2, 6, 9, 5, 4, 8},
			{8, 2, 4, 1, 5, 3, 6, 7, 9},
			{9, 6, 5, 8, 7, 4, 2, 1, 3},
		},
	}
}

func testBoard3() boardAndSolution {
	return boardAndSolution{
		board: [][]int{
			{3, 0, 0, 8, 0, 1, 0, 0, 6},
			{0, 0, 4, 2, 0, 5, 9, 0, 0},
			{0, 5, 0, 0, 0, 0, 0, 8, 0},
			{0, 7, 8, 1, 0, 9, 5, 6, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 6, 2, 7, 0, 4, 1, 3, 0},
			{0, 4, 0, 0, 0, 0, 0, 2, 0},
			{0, 0, 3, 5, 0, 7, 6, 0, 0},
			{0, 0, 0, 4, 0, 0, 0, 0, 1},
		},
		solution: [][]int{
			{3, 2, 7, 8, 9, 1, 4, 5, 6},
			{6, 8, 4, 2, 7, 5, 9, 1, 3},
			{1, 5, 9, 3, 4, 6, 2, 8, 7},
			{4, 7, 8, 1, 3, 9, 5, 6, 2},
			{5, 3, 1, 6, 2, 8, 7, 4, 9},
			{9, 6, 2, 7, 5, 4, 1, 3, 8},
			{7, 4, 6, 9, 1, 3, 8, 2, 5},
			{2, 1, 3, 5, 8, 7, 6, 9, 4},
			{8, 9, 5, 4, 6, 2, 3, 7, 1},
		},
	}
}

func testBoard4() boardAndSolution {
	return boardAndSolution{
		board: [][]int{
			{0, 0, 0, 8, 0, 1, 0, 5, 0},
			{6, 0, 4, 2, 0, 5, 9, 0, 0},
			{0, 5, 0, 0, 0, 0, 0, 0, 7},
			{0, 7, 8, 1, 0, 9, 5, 6, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 6, 2, 7, 0, 4, 1, 3, 0},
			{0, 0, 0, 0, 0, 0, 0, 2, 0},
			{0, 1, 3, 5, 8, 0, 0, 9, 0},
			{0, 0, 0, 4, 0, 0, 0, 0, 1},
		},
		solution: [][]int{
			{3, 2, 7, 8, 9, 1, 4, 5, 6},
			{6, 8, 4, 2, 7, 5, 9, 1, 3},
			{1, 5, 9, 3, 4, 6, 2, 8, 7},
			{4, 7, 8, 1, 3, 9, 5, 6, 2},
			{5, 3, 1, 6, 2, 8, 7, 4, 9},
			{9, 6, 2, 7, 5, 4, 1, 3, 8},
			{7, 4, 6, 9, 1, 3, 8, 2, 5},
			{2, 1, 3, 5, 8, 7, 6, 9, 4},
			{8, 9, 5, 4, 6, 2, 3, 7, 1},
		},
	}
}