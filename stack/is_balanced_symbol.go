package stack

import (
	"fmt"
)

// need to separate
func closingAndMatch(r string, top string) bool {

	// if both are open return false

	if top == "{" && r == "}" {
		return true
	}
	if top == "[" && r == "]" {
		return true
	}
	if top == "(" && r == ")" {
		return true
	}

	return false
}

func closing(r string) bool {

	if r == "}" {
		return true
	}
	if r == "]" {
		return true
	}
	if r == ")" {
		return true
	}

	return false
}

func isExpressionBalancedRecursive(stack *Stack, r []rune) bool {

	if len(r) == 0 && len(stack.stack) == 0 {
		return true
	} else if len(r) == 0 && len(stack.stack) != 0 {
		return false
	}

	top, err := stack.Pop()

	if err != nil {
		stack.Push(r[0])
	} else {
		if closing(string(r[0])) && !closing(string(top)){
			// if top is open and r is close but not of same type -> return false
			// esle
			if !closingAndMatch(string(r[0]), string(top)) {
				return false
			} 
		} else {
			stack.Push(top)
			stack.Push(r[0])
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

	return isExpressionBalancedRecursive(stack, r)

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
		fmt.Println(i)
		if got := isExpressionBalanced(test.expression); got != test.isValid {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.isValid, got)
		}
	}
}