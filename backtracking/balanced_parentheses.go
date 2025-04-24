package backtracking

func balanced_parentheses(n int) [][]string {
	var input []string = make([]string, n)

	for i:=1; i<(n+1)*2; i++ {
		if i % 2 == 0 {
			input = append(input, "{")
		} else {
			input = append(input, "}")
		}
	}

	var result [][]string = [][]string{}

	backtrack(input, &result, []string{}, 0, 0)

	return result
}

func backtrack(input []string, result *[][], permutation []string, num_closing int, num_open int) {

	if len(permutation) == len(input) {
		copyPerm := make([]string, len(permutation))
		copy(copyPerm, permutation)
		*result = append(*result, copyPerm)
		return
	}

	for i:=1; i<=len(input); i++ {
		isClosing := i % 2 == 0
		
		if isClosing && (num_closing < num_open) {

		} else if !isClosing && (num_open <= len(input)/2) {

		} else {
			*result = [][]string{}
			break
		}
	}

	

}