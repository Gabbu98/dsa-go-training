package backtracking

import (
	"fmt"
	"reflect"
)

// input: list of integers
// output: all possible combinations

func permutations(input []int)[][]int{
	var result [][]int = [][]int{}

	backtrack(input, &result, []int{}, make([]bool, len(input)))
	
	return result
}

func backtrack(input []int,result *[][]int, permutation []int, used []bool){
	if len(permutation) == len(input) {
		copyPerm := make([]int, len(permutation))
		copy(copyPerm, permutation)
		*result = append(*result, copyPerm)
		return
	}

	for i := 0; i < len(input); i++ {
		if !used[i] {
			used[i] = true
			permutation = append(permutation, input[i])
			backtrack(input, result, permutation, used)
			used[i] = false
			permutation = permutation[:len(permutation)-1]
		}
	}
}

/*
TestPermutations tests solution(s) with the following signature and problem description:

	func Permutations(input []int) [][]int

Given a list of integers like {1,2}, produce all possible combinations like {1,2},{2,1}.
*/
func TestPermutations() {
	tests := []struct {
		nums         []int
		permutations [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{{1}}},
		{[]int{1, 2}, [][]int{{1, 2}, {2, 1}}},
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{1, 2, 3, 4}, [][]int{{1, 2, 3, 4}, {1, 2, 4, 3}, {1, 3, 2, 4}, {1, 3, 4, 2}, {1, 4, 3, 2}, {1, 4, 2, 3}, {2, 1, 3, 4}, {2, 1, 4, 3}, {2, 3, 1, 4}, {2, 3, 4, 1}, {2, 4, 3, 1}, {2, 4, 1, 3}, {3, 2, 1, 4}, {3, 2, 4, 1}, {3, 1, 2, 4}, {3, 1, 4, 2}, {3, 4, 1, 2}, {3, 4, 2, 1}, {4, 2, 3, 1}, {4, 2, 1, 3}, {4, 3, 2, 1}, {4, 3, 1, 2}, {4, 1, 3, 2}, {4, 1, 2, 3}}},
	}

	for i, test := range tests {
		got := permutations(test.nums)
		if !reflect.DeepEqual(test.permutations, got) {
			fmt.Printf("Failed test case #%d.\nWant: %#v\nGot:  %#v\n", i, test.permutations, got)
		}
	}
}