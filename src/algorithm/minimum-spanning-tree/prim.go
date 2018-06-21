package minimum_spanning_tree

import "fmt"

type CloseEdge struct {
	data       VertexData // 表示u中顶点信息
	lowestCost UINT       // 最小代价
}

/*type node struct {
	data       VertexData
	lowestCost UINT
}
closedge[vexCounts] //Prim算法中的辅助信息
*/
func MiniSpanTreePrim(adjMat [vexCounts][vexCounts]UINT, s VertexData) {
	closeEdge := [vexCounts]CloseEdge{}
	for i := 0; i < vexCounts; i++ {
		closeEdge[i].lowestCost = INFINITE
	}
	closeEdge[s].data = s // 从顶点s开始
	closeEdge[s].lowestCost = 0
	for i := 0; i < vexCounts; i++ { // 初始化辅助数组
		if i != int(s) {
			closeEdge[i].data = s
			closeEdge[i].lowestCost = adjMat[s][i]
		}
	}
	for e := 1; e <= vexCounts-1; e++ { // n-1条边时退出
		k := Minimum(closeEdge) // 选择最小代价边
		fmt.Println(Minimum(closeEdge))
		fmt.Println(vextex[closeEdge[k].data], "--", vextex[k]) // 加入到最小生成树
		closeEdge[k].lowestCost = 0                             // 代价置为0
		for i := 0; i < vexCounts; i++ { // 更新v中顶点最小代价边信息
			if adjMat[k][i] < closeEdge[i].lowestCost {
				closeEdge[i].data = VertexData(k)
				closeEdge[i].lowestCost = adjMat[k][i]
			}
		}
	}
}

// Minimum 返回最小代价边
func Minimum(closeEdge [vexCounts]CloseEdge) int {
	var (
		min   UINT
		index int
	)
	min, index = INFINITE, -1
	for i := 0; i < vexCounts; i++ {
		if closeEdge[i].lowestCost < min && closeEdge[i].lowestCost != 0 {
			min = closeEdge[i].lowestCost
			index = i
		}
	}
	return index
}
