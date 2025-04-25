package backtracking

import (
	"fmt"
)

func balanced_parentheses_improved(n int) [][]string {
	var result [][]string
	var current []string
	generate(n, 0, 0, current, &result)
	return result
}

func generate(n , open, close int, current []string, result *[][]string) {
	if len(current) == 2*n {
		combination := make([]string, len(current))
		copy(combination, current)
		*result = append(*result, combination)
		return
	}

	if open<n {
		generate(n, open+1, close, append(current,"("),result)
	}
	if close < open {
		generate(n,open,close+1,append(current,")"),result)
	}
}

/*
TestGenerateParentheses tests solution(s) with the following signature and problem description:

	GenerateParentheses(n int) []string

Given an integer n produce all valid variations of arranging
n pair of parentheses. e.g. for `2` it should return `()()` and `(())`.
*/
func TestGenerateParenthesesImproved() {
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
		got := balanced_parentheses_improved(test.n)
		if len(got) != test.count {
			fmt.Printf("Test #%d failed: got %d combinations, want %d\n", i, len(got), test.count)
		}
	}
}