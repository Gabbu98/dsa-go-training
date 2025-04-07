package arrays

import "fmt"

/*
Input: Array of numbers
Output: Unique triplets with sum == 0
Example:
{1,2,-4,6,3} -> {-4,1,3} because -4+1+3 = 0
*/

func ExtractZeroSumTriplets(list []int) []int {
  result := make([]int, 3)
  // if length is 0 return a list with three zeros
  if len(list) == 0 {
    return result
  }

  list = sort(list)
  
  // The two pointer system


  // 0+1+2=3-4=-1
  // -4+1+2+3=3
  // -4+3+1 = 0

  // {7,3,5,-9,6,-4,4}
  // 7,3,5,-9
  // {-20,-4,3,4,5,6,7,1}
  // 7,
  return []int{}
} 

func sort(unsorted[]int) []int {
  var sorted []int
  
  for i:=0; i<len(unsorted); i++ {
    current := unsorted[i]
    var tempList []int
    if len(sorted) == 0 {
      tempList = append(sorted, current)
    }

    for j:=0; j < len(sorted); j++ {
      if current == sorted[j] {
        break
      }
      if current < sorted[j]{
        tempList = append([]int{current}, sorted[j:]...)
        tempList = append(sorted[:j],tempList...)      
        break 
      }

      if j == len(sorted) - 1 {
        tempList = append(sorted, current)
        break
      }
      
    }
    sorted = tempList
  }
  return sorted
}