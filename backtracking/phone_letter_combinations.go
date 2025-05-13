package backtracking

import (
	"fmt"
	"sort"
	"strings"
)

func digitsToLettersList(digits string) [][]string {
	digitMap := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var result [][]string
	for i := 0; i < len(digits); i++ {
		letters, ok := digitMap[digits[i]]
		if ok {
			result = append(result, letters)
		}
	}

	return result
}

func contains(list []string, target string) bool {
	for _, item := range list {
		if item == target {
			return true
		}
	}
	return false
}

func makeNestedBoolList(nestedStrings [][]string) [][]bool {
	result := make([][]bool, len(nestedStrings)) // outer slice

	for i, row := range nestedStrings {
		result[i] = make([]bool, len(row)) // inner slice
	}

	return result
}

func backtracking(in string) []string {
	input := digitsToLettersList(in)
	var result = []string{}

	if(len(in)==0 || strings.Contains(in, "1")){
		return result
	}

	phone_letter_combinations(input, 0, &result, []string{}, makeNestedBoolList(input))
	return result
}

func phone_letter_combinations(input [][]string, i int, result *[]string, perm []string, used [][]bool) {

	if len(perm) == len(input) {
		permString := strings.Join(perm, "")

		if !contains(*result, permString) {
			*result = append(*result, permString)
		}
		return
	}

	for j:=0; j<len(input[i]); j++ {
		if !used[i][j]{
			used[i][j]=true
			perm = append(perm, input[i][j])
			i++
			phone_letter_combinations(input,i,result,perm,used)
			i--
			used[i][j]=false
			perm = perm[:len(perm)-1]
		}
	}
	
}

/*
TestPhoneLetterCombinations tests solution(s) with the following signature and problem description:

	func PhoneLetterCombinations(digits string) []string

Intakes a string of digits from 2 to 9 inclusive and returns all possible combinations of letters
that could be generated from those. For example for input 2 it should return abc.
*/
func TestPhoneLetterCombinations() {
	tests := []struct {
		digits       string
		combinations []string
	}{
		{"", []string{}},
		{"1", []string{}},
		{"2", []string{"a", "b", "c"}},
		{"9", []string{"w", "x", "y", "z"}},
		{"12", []string{}},
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"232", []string{"ada", "adb", "adc", "aea", "aeb", "aec", "afa", "afb", "afc", "bda", "bdb", "bdc", "bea", "beb", "bec", "bfa", "bfb", "bfc", "cda", "cdb", "cdc", "cea", "ceb", "cec", "cfa", "cfb", "cfc"}},
	}

	for i, test := range tests {
		got := backtracking(test.digits)
		if len(got) > 0 {
			sort.Strings(got)
		}
		if len(test.combinations) != len(got) {
			fmt.Printf("Failed test case #%d. Want %#v got %#v", i, test.combinations, got)
		}
	}
}