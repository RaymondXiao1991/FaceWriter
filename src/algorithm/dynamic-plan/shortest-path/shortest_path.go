package shortest_path

/*
  实例二：格子取数/走棋盘问题
  问题描述。给定一个m*n的矩阵，每个位置是一个非负整数，
  从左上角开始放一个机器人，它每次只能朝右和下走，走到右下角，
  求机器人的所有路径中，总和最小的那条路径。如下图所示，其中图中所示的彩色方块是已知的某些非负整数值。
 */

// ShortestPath 最短路径: 状态转移法
// 考虑一般情况下位于机器人位于某点(x, y)处，那么它是怎么来的呢？只可能来自于左边或者上边。即：
//  dp[x, y] = min(dp[x-1, y], dp[x, y-1]) + a[x, y],其中a[x, y]是棋盘中(x, y)点的权重取值。
/* 状态转移方程:
    dp(i,0) = ∑(i,k=0)chess[k][0]
    dp(0,j) = ∑(j,k=0)chess[0][k]
    dp(i,j) = min(dp(i-1,j),dp(i,j-1)) + chess[i][j]
 */
func ShortestPath(chess [][]int) int {
	row, col := len(chess), len(chess[0])
	dp := make([][]int, len(chess))

	dp[0][0] = chess[0][0]
	for i := 1; i < row; i++ {
		dp[i][0] = dp[i-1][0] + chess[i][0]
	}
	for j := 1; j < col; j++ {
		dp[0][j] = dp[0][j-1] + chess[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if dp[i-1][j] < dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
			dp[i][j] += chess[i][j]
		}
	}

	return dp[row-1][col-1]
}

// ShortestPath2 最短路径: 滚动数组法
// 观察状态转移方程发现，每次更新(x, y)，只需要最多知道上一行即可，没必要知道更早的数据。凡是满足这样条件的动态规划问题，都可以用“滚动数组”的方式做空间上的优化。
/* 滚动数组方程:
    dp(j) = ∑(j,k=0)chess[0][k]
    dp(j) = min(dp(j),dp(j-1)) + chess[i][j]
 */
func ShortestPath2(chess [][]int) int {
	row, col := len(chess), len(chess[0])
	dp := make([]int, col)

	dp[0] = chess[0][0]
	for j := 1; j < col; j++ {
		dp[j] = dp[j-1] + chess[0][j]
	}
	for i := 1; i < row; i++ {
		for j := 0; j < col; j++ {
			if j == 0 {
				dp[j] += chess[i][j]
			} else {
				if dp[j] < dp[j-1] {
					dp[j] = dp[j] + chess[i][j]
				} else {
					dp[j] = dp[j-1] + chess[i][j]
				}
			}
		}
	}

	return dp[col-1]
}
