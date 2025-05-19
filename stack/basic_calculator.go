package stack

import "fmt"

// func BasicCalculator(input string) (float64, error)

// input: string with numbers, parentheses and basic arithmetic ops in the following order */+- return result


func BasicCalculatorRecursive(numberStack *Stack, operationStack *Stack) (float64, error) {
	return -1,nil
}

func isOperation(r rune) (bool, error) {
	var s string = string(r)

	if s == " " {
		return false, nil
	}

	if s == "{" || s == "}" || s == "*" || s == "/" || s == "+" || s == "-" {
		return true, nil
	}

	return false, nil
}

func BasicCalculator(input string) (float64, error) {
	if input == "" {
		return -1,nil
	}

	stackNumber := new(Stack)
	stackOps := new(Stack)

	for _, element := range input {
		isOps, _ := isOperation(element)
		if !isOps {
			stackNumber.Push(element)
			continue
		}
		stackOps.Push(element)
	}

	return BasicCalculatorRecursive(stackNumber, stackOps)
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