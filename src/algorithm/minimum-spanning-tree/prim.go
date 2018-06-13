package minimum_spanning_tree

type CloseEdge struct {
	data       VertexData // 表示u中顶点信息
	lowestCost UINT       // 最小代价
}
struct node
{
VertexData data;
unsigned int lowestcost;
}closedge[vexCounts] //Prim算法中的辅助信息

func MiniSpanTree_Prim(adjMat [vexCounts][vexCounts]UINT, s VertexData) {
	closeEdge := [...]CloseEdge{}
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
		k := Minmum(closeEdge) // 选择最小代价边
	}
}

// Minmum 返回最小代价边
func Minmum(closeEdge *node) int {

}
