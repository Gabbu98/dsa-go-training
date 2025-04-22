package arrays

import (
	"fmt"
	"slices"
)

/*
- input: slice of integers, start index and end index
- function: reverse integers in place without using extra memory
*/
func reverse(list []int, start, end int) {
    // input: 4,2,1,8,9,5 => 5,9,8,1,2,4        
	// >4,2,1,8,9,5< => 5,>9,8,1,2<,4 => 5,2,>8,1<,9,4
	// 
	for start<end {
		list[start], list[end] = list[end], list[start]
		start++
		end--
	}
	
}

func TestReverseInLine() {
	tests := []struct {
		list	[]int
		start	int
		end		int
		reversed	[]int
	}{
		{[]int{}, 0, 0, []int{}},
		{[]int{1, 2, 3}, 1, 2, []int{1, 3, 2}},
		{[]int{1, 2, 3}, 2, 1, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, 1, 2, []int{1, 3, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 0, 4, []int{5, 4, 3, 2, 1}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 4, []int{5, 4, 3, 2, 1, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, 0, 3, []int{4, 3, 2, 1, 5, 6}},
	}

	for i, test := range tests {
		reverse(test.list, test.start, test.end)
		if !slices.Equal(test.list, test.reversed) {
			fmt.Println("Failed test case #%d, expected %#v instead got %#v", i, test.reversed, test.list)
		}
	}
}

