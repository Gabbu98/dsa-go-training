package backtracking

import "strings"

func phone_letter_combinations(input []string, result *[]string, perm []string, mapInput [][]string, skip bool) {

	if len(perm) == len(input) {
		permString := strings.Join(perm, "")
		*result = append(*result, permString)
		return
	}

	for i:=0; i<len(mapInput); i++ {
		if skip {
			skip = false
			continue
		} 

		input = mapInput[i]
		
		for j:=0; j<len(input); j++ {
			skip = true
			perm = append(perm, input[j])
			phone_letter_combinations(input, result, perm, mapInput, skip)
			skip = false
			perm = perm[:len(perm)-1]
		}
	}
}