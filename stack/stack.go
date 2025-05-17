package stack

import "errors"

type Stack struct {
	stack []rune
}

func (stack *Stack) Push(i rune) {
	stack.stack = append(stack.stack, i)
}

// Define the custom error
var ErrorEmptyStack = errors.New("stack is empty")

func (stack *Stack) Pop() (rune, error) {
	if len(stack.stack) == 0 {
		return -1,ErrorEmptyStack
	}
	var value rune = stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return value,nil
}
