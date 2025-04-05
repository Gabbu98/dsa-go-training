package arrays

import (
	"fmt"
	"math"
	"slices"
)

/*
Given two slices of positive integers, sum the slices into one slice of integers.
For example {2,9} + {9,9,9} return {1,0,2,8}
*/

func sumSlices(a, b []int) []int {
	
	// get the maximum length from both slices
	var length_of_a int = len(a)
	var length_of_b int = len(b)

	if length_of_a==0 {
		return b
	} else if length_of_b == 0 {
		return a
	}

	var iteration_length int
	
	if length_of_a>=length_of_b {
		iteration_length = length_of_a
		number_of_zeros := iteration_length-length_of_b
		zeros := make(int[], number_of_zeros)
		b = append(zeros, b)
	} else if length_of_b > length_of_a {
		iteration_length = length_of_b
		number_of_zeros := iteration_length-length_of_a
		zeros := make(int[], number_of_zeros)
		a = append(zeros, a)
	}

	var result []int = []int{}
	var current int = 0
	for i := iteration_length-1; i>=0; i-- {
		var sum int
		if current<len(result) {
			sum = a[i] + b[i]
		}
		sum = a[i] + b[i] + sum[current]

		if sum>9 {
			var remainder int = sum % 10
			var digit int = sum/10
			if current<len(result) {
				result[current]=digit
			}else{
				result = append(result,digit)
			}
			result = append(result,remainder)
			
		} else {
                        if current<len(result) {
				result[current]=digit
			}else{
				result = append(result,digit)
			}
		}
		current++
	}
	

	for i:=0; i<=iteration_length-1; i++ {
		result[i], result[iteration_length-1-i] = result[iteration_length-1-i], result[i]
	}

	return result
}

func TestAddSliceOfTwoNumbers() {
	tests := []struct {
		num1, num2, sum []int
	}{
		{[]int{1}, []int{}, []int{1}},
		{[]int{1}, []int{0}, []int{1}},
		{[]int{1}, []int{1}, []int{2}},
		{[]int{1}, []int{9}, []int{1, 0}},
		{[]int{2, 5}, []int{3, 5}, []int{6, 0}},
		{[]int{2, 9}, []int{9, 9, 9}, []int{1, 0, 2, 8}},
		{[]int{9, 9, 9}, []int{9, 9, 9}, []int{1, 9, 9, 8}},
	}
	for i, test := range tests {
		
		if got := sumSlices(test.num1, test.num2); !slices.Equal(got, test.sum) {
			fmt.Println("Failed test case #%d. Want %v got %v", i, test.sum, got)
		}
	}
}
