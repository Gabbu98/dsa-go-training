package arrays
/*
input: list []int
output: int

Unsorted slice of n positive integers where each number < n
There is at most one duplicate, return it
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

	return actual_sum-summation
}

// 2, 1 = 1, 2-1 = 1
// 3, 1,2 = 3, 3-3 = 0 
// 4, 1,2,3 = 6, 6-4 = 2 > 1,3,2,3 = 9 9-6 = 3
// 5, 1,2,3,4 = 10, 10-5 = 5
