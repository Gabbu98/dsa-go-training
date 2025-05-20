package stack

import (
	"fmt"
	"strconv"
	"strings"
)

// Scan the infix expression from left to right.
// If the scanned character is an operand, put it in the postfix expression.
// Otherwise, do the following
// If the precedence of the current scanned operator is higher than the precedence of the operator on top of the stack, or if the stack is empty, or if the stack contains a ‘(‘, then push the current operator onto the stack.
// Else, pop all operators from the stack that have precedence higher than or equal to that of the current operator. After that push the current operator onto the stack.
// If the scanned character is a ‘(‘, push it to the stack.
// If the scanned character is a ‘)’, pop the stack and output it until a ‘(‘is encountered, and discard both the parenthesis.
// Repeat steps 2-5 until the infix expression is scanned.
// Once the scanning is over, Pop the stack and add the operators in the postfix expression until it is not empty.
// Finally, print the postfix expression.

// func BasicCalculator(input string) (float64, error)

// input: string with numbers, parentheses and basic arithmetic ops in the following order */+- return result

var operands = []string{"(",")","-","+","/","*"}

func isOperation(r string) bool {

	for _, s := range operands {
		if s == r {
			return true
		}
	}
	return false
}

func isCurrentHigher(top string, current string) bool {

	var currentIndex int = -1
	var topIndex int = -1

	i := 0

	for (currentIndex == -1 && topIndex == -1) || i < len(operands) {

		if operands[i] == top {
			topIndex = i
		} else if operands[i] == current {
			currentIndex = i
		}
		
	}

	if currentIndex > topIndex {
		return true
	}

	return false
}

func BasicCalculatorRecursive(stackNumber *Stack, stackOps *Stack, input string) (float64, error) {

	for len(input) > 0 {
		var val string = string(input[0])

		if(isOperation(val)) {
			top, err := stackOps.PopString()
			isCurrentHigherResult := true
			if err == nil {
				isCurrentHigherResult = isCurrentHigher(top, val)
			} else {
				return -1, err
			}

			if isCurrentHigherResult {
				stackOps.PushString(val)
			} else {
				// calc
			}
		} else {
			value, err := strconv.ParseFloat(val, 64)
			
			if err == nil {
				stackNumber.PushFloat64(value)
			} else {
				return -1, err
			}
		}
	}

	return -1, nil
}

func BasicCalculator(input string) (float64, error) {
	if input == "" {
		return -1, nil
	}

	stackNumber := new(Stack)
	stackOps := new(Stack)

	result, errorMessage := BasicCalculatorRecursive(stackNumber, stackOps, input)
}

func TestBasicCalculator() {
	tests := []struct {
		expression string
		expectErr  bool
		outcome    float64
	}{
		{"", true, -1},
		{"1++", true, -1},
		{"1+2", false, 3},
		{"1*(2+3)", false, 5},
		{"1+2+3", false, 6},
		{"1+(3-2)", false, 2},
		{"9/3", false, 3},
		{"3-9/3", false, 0},
		{"1*(2+3+(4*5))", false, 25},
		{"1*(2+3)+(4*5)", false, 25},
		{"5.10/2", false, 2.55},
	}

	for i, test := range tests {
		got, err := BasicCalculator(test.expression)
		if err != nil && !test.expectErr {
			fmt.Println("Failed test case #%d. Did not expect an error. Error : %s", i, err)
		}

		if got != test.outcome {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.outcome, got)
		}
	}
}