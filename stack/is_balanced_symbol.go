package stack

import (
	"errors"
	"fmt"
)

// isExpressionBalanced(s string)

// given a set of symbols []{}{}, determins if input is balanced or not

// example {}{}{} is balanced

// iterate char by char and push to stack

// {{([])}}

// 3


// ({((}])

// (){}][

// [
// ]


// recursively( string, stack, isBalanced)

// 	- constraint: isBalance == false
// 		return false
		
// 	- baseCase: string == empty 
// 		return isBalanced

type Stack struct {
	stack []int
}

func (stack *Stack) Push(i int) {
	stack.stack = append(stack.stack, i)

	if len(stack.stack) != 0 {
		var max int = stack.stack[len(stack.stack)-1]

		if max <= i {
			stack.stack = append(stack.stack, i)
		} else {
			stack.stack = append(stack.stack, max)
		}

	} else {
		stack.stack = append(stack.stack, i)
	}
}

// Define the custom error
var ErrorEmptyStack = errors.New("stack is empty")

func (stack *Stack) Pop() (int, error) {
	if len(stack.stack) == 0 {
		return -1,ErrorEmptyStack
	}
	var value int = stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return value,nil
}


func isExpressionBalanced(s string) bool {
	stack:=new(Stack)
	return isExpressionBalancedRecursive(s, stack, true)
}

func isExpressionBalancedRecursive(s string, stack *Stack) bool {
	if len(s) == 0 {
		return true
	}

	if isClosing(s[0]) && !match(s[0], &stack)  {
		return false
	}
	
	if len(s) > 1 {
		s = s[1:]
	} else {
		s = ""
	}

	stack.Push(s[0])

	return isExpressionBalancedRecursive(s, stack)
}

func TestIsExpressionBalanced() {
	tests := []struct {
		expression string
		isValid bool
	}{
		{"", true},
		{"()", true},
		{"(){", false},
		{"(){}", true},
		{"(){}]", false},
		{"(){}][", false},
		{"(){}[]", true},
		{"({}[])", true},
		{"({[]})", true},
		{"({[])", false},
		{"(({[]))", false},
		{")({[])", false},
	}

	for i, test := range tests {
		if got := isExpressionBalanced(test.expression); got != test.isValid {
			fmt.Printf("Failed test case #%d. Want %#v got %#v", i, test.isValid, got)
		}
	}
}