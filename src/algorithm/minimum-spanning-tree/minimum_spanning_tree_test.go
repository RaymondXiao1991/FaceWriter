package minimum_spanning_tree

import (
	"testing"
)

func TestAdjMatrix(t *testing.T) {
	var adjMat [vexCounts][vexCounts]UINT
	AdjMatrix(adjMat) //邻接矩阵
	t.Log(AdjMatrix(adjMat))
}
