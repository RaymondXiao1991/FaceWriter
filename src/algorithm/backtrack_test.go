package algorithm

import (
	"testing"
)

func TestDFS(t *testing.T) {
	DFS(0)
	t.Log(SolutionCounts)
}

func TestSolve(t *testing.T) {
	SolveEightQueen(0)
}
