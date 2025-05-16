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

func closingAndMatch(r rune, top rune) bool {
	if top != '{' && r == '}' {
		return false
	}
	if top != '[' && r == ']' {
		return false
	}
	if top != '(' && r == ')' {
		return false
	}

	return true
}

func isExpressionBalancedRecursive(stack *Stack, r []rune) bool {

	if len(r) == 0 {
		return true
	}

	top, err := stack.Pop()

	if err != nil {
		stack.Push(r[0])
	} else {
		if !closingAndMatch(r[0], top) {
			return false
		} 
	}

	r = r[1:]

	return isExpressionBalancedRecursive(stack, r)


}


func isExpressionBalanced(s string) bool {
	stack:=new(Stack)
	r := []rune(s)

	if len(r) == 0 {
		return true
	}

	if len(r) == 1 {
		return false
	}

	isExpressionBalancedRecursive(stack, r)

	return true
}

func TestIsExpressionBalanced() {
	tests := []struct {
		expression string
		isValid bool
	}{
		{"", true},
		{"()", true},
		// {"(){", false},
		// {"(){}", true},
		// {"(){}]", false},
		// {"(){}][", false},
		// {"(){}[]", true},
		// {"({}[])", true},
		// {"({[]})", true},
		// {"({[])", false},
		// {"(({[]))", false},
		// {")({[])", false},
	}

	for i, test := range tests {
		fmt.Println(i)
		if got := isExpressionBalanced(test.expression); got != test.isValid {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.isValid, got)
		}
	}
}