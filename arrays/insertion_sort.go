package arrays

import (
	"fmt"
	"slices"
)

/*
input: array of unsorted integers
conditions: sort without creating a new array
algorith:
1. Create a sorted list and unsorted list, put first element in sorted, rest in unsorted
2. Compare the first element of unsorted with every element in sorted, from last, and swap as necessary
3. If elements remain in unsorted repeat 2
*/

// [2,1] bug
// iteration 0 sorted [2], iteration 1 problematic
func insertionSort(a []int) []int {
	sorted := []int {0}
	var isSorted bool = false

	if len(a)==0 {
		return sorted[1:]
	}

	index := 0

	for !isSorted {

		if index==0 {
			sorted = append(sorted, a[0])
			a = a[1:]
			index++

			if len(a) == 0 {
				isSorted = true
			} 

			continue
		}

		for j:= len(sorted)-1; j>=0; j-- {
			if a[0] > sorted[j] {
				newList := append([]int{a[0]}, sorted[j+1:]...)
				sorted = append(sorted[:j+1], newList...)
				break
			}
		}

		if len(a)>1 {
			a = a[1:]
		} else {
			a = []int {}
			isSorted = true
		}
		
	}

	return sorted[1:]
}

func TestInsertionSort() {
	tests := []struct {
		input, sorted []int
	}{
		// {[]int{}, []int{}},
		// {[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{4, 2, 3, 1, 5}, []int{1, 2, 3, 4, 5}},
	}
	for i, test := range tests {
		result := insertionSort(test.input)
		if !slices.Equal(result, test.sorted) {
			fmt.Println("Failed test case #%d. Want %v got %v", i, test.sorted, result)
		}
	}
}