package arrays

import "fmt"

/*
input: list []int
output: int

Conditions assumed:
-> Unsorted slice of n positive integers where each number < n
-> There is at most one duplicate

Example:
-> 2, 1 = 1, 2-1 = 1
-> 3, 1,2 = 3, 3-3 = 0
-> 4, 1,2,3 = 6, 6-4 = 2 > 1,3,2,3 = 9 9-6 = 3
-> 5, 1,2,3,4 = 10, 10-5 = 5
*/


func findDuplucateInArray(list []int) int {
	// {2,4,1,3,4} , n = 5
	// find the sum of list, know the sum of n-1 elements
	var summation int
	var actual_sum int
	for i:=0; i<len(list); i++ {
		summation = summation + i
		actual_sum = actual_sum + list[i]
	}
	var diff int = actual_sum - summation
	if(diff<=0||diff>=len(list)){
		return -1
	}
	return diff
}

func TestFindDuplicate() {
	tests := []struct {
		list	[]int
		duplicate int
	}{
		{[]int{}, -1},
		{[]int{1, 2, 2}, 2},
		{[]int{1, 2, 3}, -1},
		{[]int{1, 1, 2, 3}, 1},
		{[]int{1, 2, 2, 3}, 2},
		{[]int{1, 2, 3, 2, 4, 5}, 2},
		{[]int{3, 2, 1, 4, 5, 4}, 4},
	}
	
	for i, test := range tests {
		if got := findDuplucateInArray(test.list); got != test.duplicate{
			fmt.Println("Failed test case #%d. Want %d got %d", i, test.duplicate, got) 
				
		}
	}
}


