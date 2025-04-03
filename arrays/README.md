# Arrays
- A very important data structure
- A fixed memory block
- Time Complexity of O(1) and Space Complexity of O(n) where n is they size of the array

## Implementation
- Arrays are values in Go, when passed to a functional it is an additional memory allocation
- Solution is pointers.
- The size of the array must be known during compile time.


```
package main

import "fmt"

func main() {
	var nums1 [2]int
	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums1, nums2) // Prints [0 0] [1 2 3]
}
```
# Slices
- more flexible abstraction of arrays
- more convenient access
- passed by reference
- `append` function provides dynamic resizing
- data manipulation can happen via `(inclusive) [low:high] (exclusive)` ranges


```
nums := []int {1,2,3}
nums = append([]int{0}, nums...) // add a new element to start
nums = nums[:len(nums)-1] // removes last element
```

## Make function
- creates a zeroed slice of a given length and capacity


```
nums := make([]int, 2)
nums[0], nums[1] = 0, 1
```

# Complexity
- accessing array O(1)
- searching, adding or deleting O(n)
- useful data structures:
    - linked lists
    - trees
    - hash tables