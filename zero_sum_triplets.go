/*
Input: Array of numbers
Output: Unique triplets with sum == 0
Example:
{1,2,-4,6,3} -> {-4,1,3} because -4+1+3 = 0
*/

func extractZeroSumTriplets(list []int) [3]int {
  var result := make([]int, 3)
  // if length is 0 return a list with three zeros
  if len(list) {
    return result
  }

  // 0+1+2=3-4=-1
  // -4+1+2+3=3
  // -4+3+1 = 0

  // {7,3,5,-9,6,-4,4}
  // 7,3,5,-9
  // {-20,-4,3,4,5,6,7,1}
  // 7,
} 
