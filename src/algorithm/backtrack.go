package algorithm

import "fmt"

const N = 4

var col [N]int
var d1, d2 [N * 2]int
var SolutionCounts int

func DFS(x int) {
	if x == N {
		SolutionCounts++
	} else {
		for i := 0; i < N; i++ {
			if col[i] == 0 && d1[i+x] == 0 && d2[i-x+N] == 0 {
				col[i], d1[i+x], d2[i-x+N] = 1, 1, 1
				DFS(x + 1)
				col[i], d1[i+x], d2[i-x+N] = 0, 0, 0
			}
		}
	}
}

const Num = 4

var count = 1
var queens [Num][Num] int

func print() {
	fmt.Printf("第%d种解法:\n", count)
	for i := 0; i < Num; i++ {
		for j := 0; j < Num; j++ {
			if queens[i][j] == 1 {
				fmt.Printf("%s ", "■")
			} else {
				fmt.Printf("%s ", "□")
			}
		}
		fmt.Println()
	}
}

func setQueen(row, col int) bool {
	if row == 0 {
		queens[row][col] = 1
		return true
	}

	for i := 0; i < Num; i++ {
		if queens[row][i] == 1 {
			return false
		}
	}

	for i := 0; i < Num; i++ {
		if queens[i][col] == 1 {
			return false
		}
	}

	for i, j := row, col; i < Num && j < Num; i, j = i+1, j+1 {
		if queens[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if queens[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i < Num && j >= 0; i, j = i+1, j-1 {
		if queens[i][j] == 1 {
			return false
		}
	}
	for i, j := row, col; i >= 0 && j < Num; i, j = i-1, j+1 {
		if queens[i][j] == 1 {
			return false
		}
	}
	queens[row][col] = 1
	return true
}

func SolveEightQueen(row int) {
	if row == Num {
		print()
		count++
		return
	}
	for i := 0; i < Num; i++ {
		if setQueen(row, i) {
			SolveEightQueen(row + 1)
		}
		queens[row][i] = 0
	}
}
