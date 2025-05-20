package stack

import "errors"

type Stack struct {
	stack []rune
	stackFloat64 []float64
	stackString []string
}

func (stack *Stack) Push(i rune) {
	stack.stack = append(stack.stack, i)
}

func (stack *Stack) PushFloat64(i float64) {
	stack.stackFloat64 = append(stack.stackFloat64, i)
}

func (stack *Stack) PushString(i string) {
	stack.stackString = append(stack.stackString, i)
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

func (stack *Stack) PopFloat64() (float64, error) {
	if len(stack.stack) == 0 {
		return -1,ErrorEmptyStack
	}
	var value float64 = stack.stackFloat64[len(stack.stackFloat64)-1]
	stack.stackFloat64 = stack.stackFloat64[:len(stack.stackFloat64)-1]
	return value,nil
}

func (stack *Stack) PopString() (string, error) {
	if len(stack.stack) == 0 {
		return "",ErrorEmptyStack
	}
	var value string = stack.stackString[len(stack.stackString)-1]
	stack.stackString = stack.stackString[:len(stack.stackString)-1]
	return value,nil
}
