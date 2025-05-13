package backtracking

import (
	"fmt"
	"reflect"
)

func balanced_parentheses(n int) [][]string {
	if n==0 {
		return [][]string{{""}}
	}

	var input []string = []string{}

	for i:=1; i<=n*2; i++ {
		if i % 2 != 0 {
			input = append(input, "(")
		} else {
			input = append(input, ")")
		}
	}

	var result [][]string = [][]string{}

	backtrack(input, &result, []string{}, 0, 0, make([]bool, len(input)))

	return result
}

func backtrack(input []string, result *[][]string, permutation []string, num_closing int, num_open int, used []bool) {

	if len(permutation) == len(input) {
		for i:=0; i<len(*result); i++ {
			if reflect.DeepEqual(permutation, (*result)[i]) {
				return
			}
		}

		copyPerm := make([]string, len(permutation))
		copy(copyPerm, permutation)
		*result = append(*result, copyPerm)
		return
	}

	for i:=0; i<len(input); i++ {
		isClosing := (i+1) % 2 == 0
		
		if !used[i] && isClosing && (num_closing < num_open) {
			num_closing++
			used[i] = true
			permutation = append(permutation, input[i])
			backtrack(input, result, permutation, num_closing, num_open, used)
			used[i] = false
			num_closing--
			permutation = permutation[:len(permutation)-1]
			
		} else if !used[i] && !isClosing && (num_open <= len(input)/2) {
			num_open++
			used[i] = true
			permutation = append(permutation, input[i])
			backtrack(input, result, permutation, num_closing, num_open, used)
			used[i] = false
			num_open--
			permutation = permutation[:len(permutation)-1]
		}

		
	}

}

/*
TestGenerateParentheses tests solution(s) with the following signature and problem description:

	GenerateParentheses(n int) []string

Given an integer n produce all valid variations of arranging
n pair of parentheses. e.g. for `2` it should return `()()` and `(())`.
*/
func TestGenerateParentheses() {
	tests := []struct {
		n     int
		count int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 5},
		{4, 14},
		{5, 42},
	}
	
	for i, test := range tests {
		got := balanced_parentheses(test.n)
		if len(got) != test.count {
			fmt.Printf("Test #%d failed: got %d combinations, want %d\n", i, len(got), test.count)
		}
	}
}