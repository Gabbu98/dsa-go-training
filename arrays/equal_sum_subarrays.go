package arrays

import (
	"fmt"
	"reflect"
)

/*
input: list of integers A
output: two sub arrays with equal sums without changing order of the elements in the list
example: [1,7,3,5] => [1,7] and [3,5] = 8
*/

func equalSumSubArrays(A []int) [][]int {
	var m = make(map[int]int)
	var sum int
	// iterate and sum -> store in a map
	for i:=0; i<len(A); i++ {
		sum += A[i]
		m[sum] = i
	}

	if sum%2!=0 || sum==0 {
		return [][]int{}
	}
	
	var div int = sum / 2

	v,exists:=m[div]

	if !exists {
		return [][]int{}
	}

	return [][]int{A[:v+1],A[v+1:]}
}

func TestEqualSumSubArrays() {
	tests := []struct {
		list      []int
		subArrays [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2, 2, 3}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{{1, 2}, {3}}},
		{[]int{2, 3, 5}, [][]int{{2, 3}, {5}}},
		{[]int{1, 7, 3, 5}, [][]int{{1, 7}, {3, 5}}},
		{[]int{-4, 1, 1, 1, 1}, [][]int{}},
		{[]int{4, 1, 1, 1, 1}, [][]int{{4}, {1, 1, 1, 1}}},
		{[]int{1, 1, 1, 1, 4}, [][]int{{1, 1, 1, 1}, {4}}},
	}

	for i, test := range tests {
		if got := equalSumSubArrays(test.list); !reflect.DeepEqual(got, test.subArrays) {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.subArrays, got)
		}
	}
}