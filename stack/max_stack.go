package stack

import (
	"errors"
	"fmt"
)

/*
Implement a stack that can return the max of element it contains
*/

type MaxStack struct {
	stack []int
	maxStack []int
}

func (maxStack *MaxStack) Push(i int) {
	maxStack.stack = append(maxStack.stack, i)

	if len(maxStack.maxStack) != 0 {
		var max int = maxStack.maxStack[len(maxStack.maxStack)-1]

		if max <= i {
			maxStack.maxStack = append(maxStack.maxStack, i)
		} else {
			maxStack.maxStack = append(maxStack.maxStack, max)
		}

	} else {
		maxStack.maxStack = append(maxStack.maxStack, i)
	}
}

// Define the custom error
var ErrEmptyStack = errors.New("stack is empty")

func (maxStack *MaxStack) Pop() (int, error) {
	if len(maxStack.stack) == 0 {
		return -1,ErrEmptyStack
	}
	var value int = maxStack.stack[len(maxStack.stack)-1]
	maxStack.stack = maxStack.stack[:len(maxStack.stack)-1]
	maxStack.maxStack = maxStack.maxStack[:len(maxStack.maxStack)-1]
	return value,nil
}

func (maxStack *MaxStack) Max() (int,error) {
	if len(maxStack.maxStack)==0 {
		return -1,ErrEmptyStack
	}
	return maxStack.maxStack[len(maxStack.maxStack)-1],nil
}


/*
TestMaxStack tests solution(s) with the following signature and problem description:

	func (maxStack *MaxStack) Push(i int)
	func (maxStack *MaxStack) Max() int

Implement a stack that can return the max of element it contains.

For example if we push {2,4,5} to the stack, max should return 5.
*/
func TestMaxStack() {
	tests := []struct {
		push []int
		pop  int
		max  int
	}{
		{[]int{}, 0, -1},
		{[]int{1, 2, 3, 4, 5}, 2, 3},
		{[]int{1, 2, 3, 4, 5}, 0, 5},
		{[]int{5, 4, 3, 2, 1}, 2, 5},
		{[]int{1, 5, 3, 4}, 1, 5},
	}

	for i, test := range tests {
		stack := new(MaxStack)
		for _, p := range test.push {
			stack.Push(p)
		}

		for range test.pop {
			_, err := stack.Pop()
			if err != nil {
				fmt.Printf("unexpected error: %s", err)
			}
		}

		got,_ := stack.Max()
		if got != test.max {
		fmt.Printf("Failed test case #%d. Want %#v got %#v", i, test.max, got)
		}
	}
}

func TestEmptyStackError() {
	stack := new(MaxStack)
	if _, err := stack.Pop(); !errors.Is(err, ErrEmptyStack) {
		fmt.Printf("Calling Pop on an empty stack did not result in empty stack error. Want %#v got %#v", ErrEmptyStack, err)
	}
}
