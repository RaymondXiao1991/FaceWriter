package minimum_spanning_tree

const INFINITE = 0xFFFFFFFF
const vexCounts = 6   //顶点数量
type VertexData uint8 //顶点数据
type UINT uint

var vextex = []string{"A", "B", "C", "D", "E", "F"}

func MinSpanTree() {
	var adjMat [vexCounts][vexCounts]UINT
	AdjMatrix(adjMat) //邻接矩阵
}

func AdjMatrix(adjMat [vexCounts][vexCounts]UINT) [vexCounts][vexCounts]UINT { //邻接矩阵表示法
	for i := 0; i < vexCounts; i++ { //初始化邻接矩阵
		for j := 0; j < vexCounts; j++ {
			adjMat[i][j] = INFINITE
		}
	}
	adjMat[0][1] = 6
	adjMat[0][2] = 1
	adjMat[0][3] = 5
	adjMat[1][0] = 6
	adjMat[1][2] = 5
	adjMat[1][4] = 3
	adjMat[2][0] = 1
	adjMat[2][1] = 5
	adjMat[2][3] = 5
	adjMat[2][4] = 6
	adjMat[2][5] = 4
	adjMat[3][0] = 5
	adjMat[3][2] = 5
	adjMat[3][5] = 2
	adjMat[4][1] = 3
	adjMat[4][2] = 6
	adjMat[4][5] = 6
	adjMat[5][2] = 4
	adjMat[5][3] = 2
	adjMat[5][4] = 6
	return adjMat
}
