package backtracking

func sudoku(board *[][]int, i *int, perm []int, used []bool, done *bool) {

	if len(*board) == 9 {
		*done = true
		return
	}

	if len(perm) == 9 {
		copyPerm := make([]int, len(perm))
		copy(copyPerm, perm)
		*board = append(*board, copyPerm)
		*i++
		return
	}

	for x:=0; x<len(used); x++ {

		if !used[x] && !isPresent() {
			used[x]=true
			perm = append(perm, x)
			sudoku(board, i, perm, used, done)
			perm = perm[:len(perm)-1]
			used[x]=false
		}
	}
}

func isPresent() bool {
	
}