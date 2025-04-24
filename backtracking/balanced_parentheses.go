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

	backtrack(input, &result, []string{}, 0, n)

	return result
}

func backtrack(input []string, result *[][], permutation []string, num_closing int, remaining_open int) {

	

}