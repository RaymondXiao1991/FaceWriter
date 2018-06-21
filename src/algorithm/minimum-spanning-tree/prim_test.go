package minimum_spanning_tree

import (
	"testing"
)

func TestMiniSpanTreePrim(t *testing.T) {
	var adjMat [vexCounts][vexCounts]UINT
	MiniSpanTreePrim(AdjMatrix(adjMat), 0) //Prim算法，从顶点0开始.
}
