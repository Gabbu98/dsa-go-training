package backtracking

import "strings"

var dirs = []string{"l", "u", "r", "d"}

func maze(m int, n int, walls [][2]int, start [2]int, finish [2]int, current [2]int, visited[][4]bool, path []string, result *string) {
	if current[0] == finish[0] && current[1] == finish[1] {
		strings.Join(path, *result)
		return
	}

	if current[0] >= m || current[1] >= n{
		return
	}

	for i:=0; i<len(walls); i++ {
		if walls[i][0] == current[0] && walls[i][1] == current[1] {
			return
		}
	}

	visited = append(visited, [4]bool{false,false,false,false})

	for i:=0; i<len(dirs); i++ {

		if !visited[len(visited)-1][i] {
			visited[len(visited)-1][i] = true
			path = append(path, dirs[i])
			newCurrent := updateCurrent(current, dirs[i])
			maze(m, n, walls, start, finish, newCurrent, visited, path, result)
			path = path[:len(path)-1]
			visited[len(visited)-1][i] = false // this might need improvement as it does not actually globally save the visited areas
		}
	}

	visited = visited[:len(visited)-1]
	
}

func updateCurrent(current [2]int, dir string) [2]int {
	newCurrent := current
	switch dir {
		case "l": newCurrent[1]=current[1]-1
		case "r": newCurrent[1]=current[1]+1
		case "d": newCurrent[0]=current[0]-1
		case "u": newCurrent[0]=current[0]+1
	}
	return newCurrent
}