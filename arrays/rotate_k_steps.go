package arrays

import "fmt"
/*
input: list of integers and number k
process: rotate k times
output: rotated list of integers

[0,1,2],3 => [3,4,5] 3%3=0, 4%3=1, 5%3=2
[0,1,2],4 => [4,5,6] 4%3=1, 5%3=2, 6%3=0
*/

func RotateKSteps(a []int, k int): []int {
	
	if k<=0 {
		return a
	}

	for i := 0; i<len(a); i++ {
		var index int = i+k;

		if index >= len(a) {
			index = index % len(a)
		}

		a[index] a[i] = a[i] a[index]

	}

	fmt.Println(a)

	return a
}