package arrays

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
	return m 
}

// [1,2,3,4,5] => [120,60,40,...] => [a,ab,abc,abcd,abcde]
// [abcde,abcd,abc,ab,a]

// left = [1,1,12,123,1234]
// right = [2345,345,45,1]

// 0: left=1 right=2*3*4*5
// 1: left=1 right=3*4*5
// 2: left=1*2 right=4*5


// traverse max to min and populate right
// traverse min to max and find the left product and populate result into list[i]]

func ProductOptimized(A []int) []int {
	// Create a slice to hold the product of all elements to the right of each index
	var right []int = make([]int, len(A))
	var product int = 1

	// Fill the 'right' slice with the running product of elements to the right of each index
	for i := len(A) - 1; i >= 0; i-- {
		if len(A)-1 == i {
			// The last element has no elements to the right, so product is 1
			right[i] = product
			continue
		}
		// Update product and assign to right[i]
		product *= A[i+1]
		right[i] = product
	}

	// Reset product for left-side processing
	product = 1
	value := 1

	// Traverse the array from left to right
	for i := 0; i <= len(A)-1; i++ {

		if i == 0 {
			// For the first element, there's no element to the left
			value = product * right[i]
			continue
		}

		// Update product with the previous element (left-side product)
		product *= A[i-1]

		// Overwrite A[i-1] with the final product of left * right
		A[i-1] = value

		// Calculate new value using updated product and right[i]
		value = product * right[i]

		// Special handling for last element, assign final value
		if i == len(A)-1 {
			A[i] = value
		}
	}

	return A
}
