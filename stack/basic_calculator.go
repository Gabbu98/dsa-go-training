package stack

import (
	"fmt"
	"strconv"
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

// function evaluate(expression):
//     stack = empty stack
//     index = 0

//     while index < length of expression:
//         token = expression[index]

//         if token is a digit:
//             parse full number and push to stack

//         else if token is an operator:
//             while stack is not empty and precedence(token) <= precedence(top of stack):
//                 pop operator and operands, compute, and push result
//             push operator to stack

//         else if token is '(':
//             result, new_index = evaluate(sub-expression starting at index + 1)
//             push result to stack
//             index = new_index

//         else if token is ')':
//             while stack contains operators:
//                 pop operator and operands, compute, and push result
//             return top of stack and current index

//         index += 1

//     while stack contains operators:
//         pop operator and operands, compute, and push result

//     return top of stack


var operands = []string{".","(",")","-","+","/","*"}

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

	for i < len(operands) {

		if operands[i] == top {
			topIndex = i
		} else if operands[i] == current {
			currentIndex = i
		}
		i++
	}

	if currentIndex > topIndex {
		return true
	}

	return false
}

func count(stackNumber *Stack, stackOperands *Stack, tok string) (float64, error) {
    var result float64 = 0

	top, emptyStackErr := stackOperands.PopString()

	for emptyStackErr == nil {

		right, rightErr := stackNumber.PopFloat64()
		left, leftErr := stackNumber.PopFloat64()

		if leftErr != nil {
			return -1, leftErr
		} else if rightErr != nil {
			return -1, rightErr
		} else {
			switch top {
				case "*": result = left * right
				case "/": result = left / right
				case "+": result = left + right
				case "-": result = left - right
			}
			stackNumber.PushFloat64(result)
		}

		top, emptyStackErr = stackOperands.PopString()
		
		if tok == "" && emptyStackErr != nil  && isCurrentHigher(top, tok) {
			stackOperands.PushString(top)
			return result, nil
		}
	}

	return result, emptyStackErr

}

func BasicCalculatorRecursive(stackNumber *Stack, stackOps *Stack, input string, index int) (float64, int, error) {
	for index < len(input) {
		tok := string(input[index])

		if !isOperation(string(tok)) {
			val , err := strconv.ParseFloat(tok, 64)

			if err != nil {
				return -1, index, err
			}

			stackNumber.PushFloat64(val)
		} else {
			if tok == "(" {
				result, newIndex, err := BasicCalculatorRecursive(new(Stack), new(Stack), input, index+1)
				
				if err != nil {
					return -1, newIndex, err
				}
				
				stackNumber.PushFloat64(result)
				index = newIndex
			} else if tok == ")" {
				result, _ := count(stackNumber, stackOps, tok)

				return result, index, nil
			} else if tok == "." {
				index++
				if !isOperation(string(input[index])) {
					val, err := stackNumber.PopFloat64()

					if err != nil {

						decimal,_ := strconv.ParseFloat(strconv.FormatFloat(val, 'f', 2, 64) + "." + string(input[index]),64)
						stackNumber.PushFloat64(decimal)
					} else {
						return -1, index, nil
					}
				} else {
					return -1, index, nil				}
			} else {
				top, _ := stackOps.PopString()

				if isCurrentHigher(top, tok) {
					if top!="" {
						stackOps.PushString(top)
					}
					stackOps.PushString(tok)
				} else  {
					stackOps.PushString(top)
					result, _ := count(stackNumber, stackOps, tok)

					stackNumber.PushFloat64(result)
					stackOps.PushString(tok)
				}
			}
		}
		index++
	}

	for len(stackOps.stackString) != 0 {
		result, err := count(stackNumber, stackOps, "")

		if err != nil {
			return result, index, err
		}

		stackNumber.PushFloat64(result)
	}

	result, err := stackNumber.PopFloat64()

	if err != nil {
		return result, index, err
	}

	return result, index, nil
}

func BasicCalculator(input string) (float64, error) {
	if input == "" {
		return -1, nil
	}

	stackNumber := new(Stack)
	stackOps := new(Stack)

	result, _, errorMessage := BasicCalculatorRecursive(stackNumber, stackOps, input, 0)

	return result, errorMessage
}

func TestBasicCalculator() {
	tests := []struct {
		expression string
		expectErr  bool
		outcome    float64
	}{
		// {"", true, -1},
		// {"1++", true, -1},
		// {"1+2", false, 3},
		// {"1*(2+3)", false, 5},
		// {"1+2+3", false, 6},
		// {"1+(3-2)", false, 2},
		// {"9/3", false, 3},
		// {"3-9/3", false, 0},
		// {"1*(2+3+(4*5))", false, 25},
		// {"1*(2+3)+(4*5)", false, 25},
		// {"5.10/2", false, 2.55}, // to solve this i would first split by the delimeters (/,*,+,-)
	}

	for i, test := range tests {
		got, _ := BasicCalculator(test.expression)
		// if err != nil && !test.expectErr {
		// 	fmt.Println("Failed test case #%d. Did not expect an error. Error : %s", i, err)
		// }

		if got != test.outcome {
			fmt.Println("Failed test case #%d. Want %#v got %#v", i, test.outcome, got)
		}
	}
}
