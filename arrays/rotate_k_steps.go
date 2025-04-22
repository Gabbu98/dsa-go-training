package arrays

import (
	"fmt"
	"slices"
)

/*
input: list of integers and number k
process: rotate k times
output: rotated list of integers

[0,1,2],3 => [3,4,5] 3%3=0, 4%3=1, 5%3=2
[0,1,2],4 => [4,5,6] 4%3=1, 5%3=2, 6%3=0
*/

func rotateKSteps(a []int, k int) []int {
	
	if k<=0 {
		return a
	}
	var rotated_a = make([]int, len(a))

	for i := 0; i<len(a); i++ {
		var index int = i+k;
		
		if index >= len(a) {
			index = index % len(a)
		}

		rotated_a[index]= a[i]

	}

	return rotated_a
}

func TestRotateKSteps() {
	tests := []struct {
		list        []int
		steps       int
		rotatedList []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1, 2, 3}, 1, []int{3, 1, 2}},
		{[]int{1, 2, 3}, 2, []int{2, 3, 1}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
	}

	for i, test := range tests {
		list := rotateKSteps(test.list, test.steps)
		if !slices.Equal(test.rotatedList, list) {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.rotatedList, test.list)
		}
	}
}
