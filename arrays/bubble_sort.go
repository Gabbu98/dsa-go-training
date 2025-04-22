package arrays

import (
	"fmt"
	"slices"
)

/*
input: slice of unsorted integers
conditions: sort without creating a new array

bubble sort algorithm:
1. Go through each item
2. If item at n[i] is larger than n[i+1],swap
3. If swap go back to 1, else list is sorted
*/

func bubbleSort(a []int) []int {
	
	if len(a) == 0 {
		return a
	}

	var isSorted bool = false
	var index int = 0

	for !isSorted {
		if len(a)-1==index {
			isSorted = true
			continue
		}

		if a[index] > a[index+1] {
			a[index], a[index+1] = a[index+1], a[index]
			index = 0
			continue
		}

		index++
	}
	return a
}

func TestBubbleSort() {
	tests := []struct {
		input, sorted []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
	}
	for i, test := range tests {
		bubbleSort(test.input)
		if !slices.Equal(test.input, test.sorted) {
			fmt.Println("Failed test case #%d. Want %v got %v", i, test.sorted, test.input)
		}
	}
}