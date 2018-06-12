package longest_increasing_subsequence

/*
  实例一：最长递增子序列（Longest Increasing Subsequence）。
  问题描述。给定长度为N的数组A，计算A的最长单调递增的子序列（不一定连续）。如给定数组A{5，6，7，1，2，8}，则A的LIS为{5，6，7，8}，长度为4.
  思路：因为子序列要求是递增的，所以重点是子序列的起始字符和结尾字符，因此我们可以利用结尾字符。
  想到：以A[0]结尾的最长递增子序列有多长？以A[1]结尾的最长递增子序列有多长？……以A[n-1]结尾的最长递增子序列有多长？
  （动态规划solution）所以我们可以使用一个额外的空间来保存前面已经算得的最长递增子序列，然后每次更新当前的即可。也就是问题演化成：已经计算得到了b[0,1,2,……,i-1]，如何计算得到b[i]呢？
  显然，如果ai>=aj，则可以将ai放到b[j]的后面，得到比b[j]更长的子序列。从而：b[i] = max{b[j]}+1. s.t. A[i] > A[j] && 0 <= j < i.
  所以计算b[i]的过程是，遍历b[i]之前的所有位置j，找出满足关系式的最大的b[j].
  得到b[0...n-1]之后，遍历所有的b[i]找到最大值，即为最大递增子序列。 总的时间复杂度为O(N2).
 */
func LongestIncreasingSubSequence(s []int) int {
	if s == nil {
		return 0
	}
	l := len(s)
	if l == 0 {
		return 0
	}

	t := make([]int, l)
	t[0] = 1
	result := 1
	for i := 1; i < l; i++ {
		max := -1
		for j := 0; j < i; j++ {
			if s[j] < s[i] && t[j] > max {
				max = t[j]
			}
		}
		t[i] = max + 1
		result = themax(result, t[i])
	}
	return result
}

func LongestIncreasingSubSequence2(arr []int) int {
	l := len(arr)
	longest := make([]int, l)

	for i := 0; i < l; i++ {
		longest[i] = 1
	}

	for j := 1; j < l; j++ {
		for i := 0; i < j; i++ {
			if arr[j] > arr[i] && longest[j] < longest[i]+1 {
				longest[j] = longest[i] + 1
			}
		}
	}

	max := 0
	for j := 0; j < l; j++ {
		if longest[j] > max {
			max = longest[j]
		}
	}

	return max
}

type Channel struct {
	data int
}

// 如果不仅是求LIS的长度，而要求LIS本身呢？我们可以通过记录前驱的方式，从该位置找到其前驱，进而找到前驱的前驱……
func generateLIS(s, dp []int) []int {
	index, length := 0, 0
	for i := 0; i < len(s); i++ {
		if dp[i] > length {
			length = dp[i]
			index = i
		}
	}

	subArr := make([]int, length)
	subArr[length-1] = s[index]

	for j := index - 1; j >= 0; j-- {
		if (dp[index] == dp[j]+1) && (s[index] > s[j]) {
			subArr[length-2] = s[j]
			index = j
			length--
		}
	}

	return subArr
}

// LongestIncreasingSubSequenceDetail O(N2)时间复杂度
func LongestIncreasingSubSequenceDetail(s []int) []int {
	dp := make([]int, len(s))
	dp[0] = 1
	for i := 1; i < len(s); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if s[i] > s[j] {
				//求dp[i]时遍历，dp[0...i-1],找出arr[j]<arr[i]小且dp[j]是最大的
				//dp[i]=dp[j]+1;
				dp[i] = themax(dp[i], dp[j]+1)
			}
		}
	}

	return generateLIS(s, dp)
}

func themax(x, y int) int {
	if x > y {
		return x
	}
	return y
}
