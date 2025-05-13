package stack

import "errors"

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

func (maxStack *MaxStack) Max() int {
	return maxStack.maxStack[len(maxStack.maxStack)-1]
}

