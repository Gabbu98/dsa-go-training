package backtracking

import (
	"fmt"
	"strings"
)

var dirs = []string{"l", "u", "r", "d"}
type Key [2]int

func maze_driver(m int, n int, wall [][2]int, start [2]int, finish [2]int) string {
	result := ""
	visited := map[Key]bool{}

	for i:= 0; i<m; i++ {
		for j:=0; j<n; j++ {
			visited[Key{i,j}] = false
		}
	}

	maze(m, n, wall, start, finish, start, &visited, []string{}, &result)

	return result
}


func maze(m int, n int, walls [][2]int, start [2]int, finish [2]int, current [2]int, visited *map[Key]bool, path []string, result *string) {
	if current[0] == finish[0] && current[1] == finish[1] {
		*result = strings.Join(path, "")
		return
	}

	if current[0] >= m || current[1] >= n || current[0] < 0 || current[1] < 0{
		return
	}

	for i:=0; i<len(walls); i++ {
		if walls[i][0] == current[0] && walls[i][1] == current[1] {
			return
		}
	}

	for i:=0; i<len(dirs); i++ {

		nextCurrent := updateCurrent(current, dirs[i])

		if !isVisited(visited, nextCurrent) && len(*result) == 0 {
			(*visited)[Key{nextCurrent[0], nextCurrent[1]}]=true
			path = append(path, dirs[i])
			maze(m, n, walls, start, finish, nextCurrent, visited, path, result)
			path = path[:len(path)-1]
		}
	}
	
}

func isVisited(visited *map[Key]bool, nextCurrent [2]int) bool {
	return (*visited)[Key{nextCurrent[0], nextCurrent[1]}]
}

func updateCurrent(current [2]int, dir string) [2]int {
	newCurrent := current
	switch dir {
		case "l": newCurrent[1]=current[1]-1
		case "u": newCurrent[0]=current[0]-1
		case "r": newCurrent[1]=current[1]+1
		case "d": newCurrent[0]=current[0]+1
	}
	return newCurrent
}

/*
TestMaze tests solution(s) with the following signature and problem description:

	func Maze(walls [][2]int, start, finish [2]int,m,n int) string

Given the coordinates of walls in a m x n maze in tuples of {row, col} format, and start
and finish coordinates in the same format, find a path from start to finish. return
a string of directions like `lrud` (left, right, up, down) to get a robot from
start to finish.

The robot can only move in the four left, right, down and up directions and not
through walls. The robot does not know where the finish line is so it has to
explore every possible cell in the order of directions given.
*/
func TestMaze() {
	tests := []struct {
		m      int
		n      int
		walls  [][2]int
		start  [2]int
		finish [2]int
		moves  string
	}{
		{5, 5, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"},
		{10, 10, [][2]int{}, [2]int{0, 0}, [2]int{0, 1}, "r"},
		{5, 5, [][2]int{}, [2]int{0, 0}, [2]int{4, 4}, "rldrrurrdldllldrrrrd"},
		{5, 5, [][2]int{{1, 1}, {1, 2}, {1, 3}, {2, 3}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{2, 4}, "rrrrdd"},
		{5, 5, [][2]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 3}, {2, 1}, {2, 2}, {2, 3}, {3, 1}, {4, 1}}, [2]int{4, 0}, [2]int{1, 2}, "uuurr"},
		{5, 5, [][2]int{{1, 0}, {1, 1}, {1, 2}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrrrddllllddrrrr"},
		{5, 5, [][2]int{{1, 0}, {1, 1}, {1, 4}, {1, 3}, {3, 1}, {3, 2}, {3, 3}, {3, 4}}, [2]int{0, 0}, [2]int{4, 4}, "rrddllddrrrr"},
	}

	for i, test := range tests {
		if got := maze_driver(test.m, test.n, test.walls, test.start, test.finish); test.moves != got {
			fmt.Printf("Failed test case #%d. Want %s got %s", i, test.moves, got)
		}
	}
}