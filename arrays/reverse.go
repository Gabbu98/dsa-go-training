package arrays

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