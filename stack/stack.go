package stack

import "errors"

type Stack struct {
	stack []rune
	stackString string
}

func (stack *Stack) Push(i rune) {
	stack.stack = append(stack.stack, i)
}

func (stack *Stack) PushString(i string) {
	stack.stackString = stack.stackString + i
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

func (stack *Stack) PopString() (string, error) {
	if len(stack.stackString) == 0 {
		return "",ErrorEmptyStack
	}
	var value string = string(stack.stackString[len(stack.stackString)-1])
	stack.stackString = stack.stackString[:len(stack.stackString)-1]
	return value,nil
}
