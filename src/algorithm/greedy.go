package algorithm

import (
	"math"
	"fmt"
)

// 贪心算法是使所做的选择看起来都是当前最佳的，期望通过所做的局部最优选择来产生一个全局最优解。
/*1.将优化问题转换成这样一个问题，即先做出选择，再解决剩下的一个子问题。
　　2.证明原问题总是有一个最优解是贪心选择的得到的，从而说明贪心选择的安全。
　　3.说明在做出贪心选择后，剩下的子问题具有这样一个性质。即如果将子问题的最优解和我们所做的贪心选择联合起来，可以得到一个更加负责的动态规划解。
 */

/*
  案例一:
  给你一个长度为n的绳子，请把绳子剪成m段（m，n都是整数，且都大于1）每段绳子的长度即为K[0],K[1],K[2]...K[m]。
  请问K[0]*k[1]..*k[m]可能的最大乘积是多少？
*/
// 当n>=5,我们尽可能地剪长度为3的绳子;当剩下的绳子长度为4时,把绳子剪为长度为2的绳子.
func GreedyCutRope(length int) int { // 迭代法
	if length < 2 {
		return 0
	}
	switch length {
	case 2:
		return 1
	case 3:
		return 2
	}

	// 尽可能多地去减长度为3的绳子段
	timesOfThree := length / 3
	// 当绳子最后剩下的长度为4的时候，不能再去剪去长度为3的绳子段
	if length-timesOfThree*3 == 1 {
		timesOfThree--
	}
	timesOfTwo := (length - timesOfThree*3) / 2

	return int(math.Pow(3, float64(timesOfThree)) * math.Pow(2, float64(timesOfTwo)))
}

func GreedyCutRopeByRecursive(length int) int { // 递归法
	switch length {
	case 0, 1:
		return 1
	case 2, 3, 4:
		return length
	}
	return (GreedyCutRopeByRecursive(length - 3)) * 3
}

/*
  案例二:
  给定N个物品和一个容量为C的背包,物品i的重量为Wi，其价值为Vi，
  背包问题是如何选择装入背包的物品，使得装入背包中物品的总价值最大。
  注意在背包问题中，可以将某种物品的一部分装入背包中，但是不可以重复装入。
 */
// 选择单位重量价值最大的物品
func GreedyKnapSack(weight, value []int, c int) float64 {
	// 假设物品已按单位重量降序排列
	var (
		i        int
		maxValue float64
		x        [10]float64
	)
	for i = 0; weight[i] < c; i++ {
		x[i] = 1                      // 将物品i装入书包
		maxValue += float64(value[i]) // 总价值累加
		c -= weight[i]                // 背包剩余数量
	}
	fmt.Println(c, weight[i])
	x[i] = float64(c / weight[i]) // 物品i装入一部分
	fmt.Println("x[i]", x[i])
	fmt.Println("value[i]", value[i])
	maxValue += x[i] * float64(value[i])
	return maxValue

	/*for (i=1;i<=n;i++)//物品整件被装下,x[i]=1
	{
	if (w[i]>c)
	{
	break;
	}
	x[i]=1;
	c-=w[i];
	}

	//物品i只有部分被装下
	if (i<=n)
	{
		x[i]=c/w[i];
	}*/
}
