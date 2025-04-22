package arrays

import (
	"fmt"
	"reflect"
)

/*
Input: Array of numbers
Output: Unique triplets with sum == 0
Example:
{1,2,-4,6,3} -> {-4,1,3} because -4+1+3 = 0
*/

func extractZeroSumTriplets(list []int) [][]int {
  result := [][]int{}
  // if length is 0 return a list with three zeros
  if len(list) < 3  {
    return result
  }

  sortedList := sort(list)

  isZero := false
  fixedPointer := 0
  jPointer := 1
  kPointer := len(sortedList)-1
  
  var triplets [][]int
  if len(sortedList) == 0 {
    return append(triplets, list[0:3])
  }

  // The two pointer system
  for !isZero {
    if kPointer==jPointer || fixedPointer==jPointer {
      isZero = true
      break
    }

    sum := sortedList[fixedPointer]+sortedList[jPointer]+sortedList[kPointer]
    
    if sum == 0 {
      triplets = append(triplets, []int{sortedList[fixedPointer], sortedList[jPointer], sortedList[kPointer]})
    }

    if sum > 0 {
      kPointer--
    }

    if sum < 0 {
      jPointer++
    }
  }

  return triplets
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

func ZeroSumTriplets() {
  tests := []struct {
		list     []int
		triplets [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{1}, [][]int{}},
		{[]int{1, 2}, [][]int{}},
		{[]int{1, 2, 3}, [][]int{}},
		{[]int{1, -4, 3}, [][]int{{-4, 1, 3}}},
		{[]int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{1, 2, -4, 6, 3}, [][]int{{-4, 1, 3}}},
		{[]int{-1, -2, -8, 6, 2, 3}, [][]int{{-8, 2, 6}, {-2, -1, 3}}},
		{[]int{1, -2, -4, 5, -2, 4, 1, 3}, [][]int{{-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}}},
	}
	
	for i, test := range tests {
		if got := extractZeroSumTriplets(test.list); !reflect.DeepEqual(got, test.triplets) {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.triplets, got)
		}
	}
}