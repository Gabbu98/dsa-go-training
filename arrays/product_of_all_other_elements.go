package arrays

import (
	"fmt"
)

/*
Input: array of integers A
Process: a new array B where B[i] = product of all items in A except A[i] excluding division
*/

func Product(A []int) []int {
	var m []int = make([]int, len(A))

	for i:=0; i<len(A); i++ {
		value := A[i]
		
		for j:=0; j<len(A); j++ {
			if j!=i {
				if m[j]==0{
					m[j] = 1 * value
				} else {
					m[j] = m[j] * value
				}
			}
		}

	}
	fmt.Print(m)
	return m 
}