package permutation_tree

/*
排列树
  （为了便于描述算法，下列方法使用了较多的全局变量）
    I.输出集合S中所有的排列，即limit为all；
    II.输出集合S中元素奇偶性相间的排列，即limit为sp。
 */
s := []int{1, 2, 3, 4, 5, 6, 7, 8}
n := len(s)
x :=[n]int{}

/**
 * 输出集合的排列
 * @param limit  决定选出特定条件的子集
 * 注：all为所有排列,sp为限定元素奇偶性相间。
 */
func all_permutation(limit string) {
	switch limit {
	case "all":
		backtrack(0)
		break
	case "sp":
		backtrack1(0)
		break
	}
}

/**
 * 回溯法求集合的所有排列，依次递归
 * 注：是否回溯的条件为精髓
 * @param t
 */
func backtrack(t int) {
	if t >= n {
		output(s)
	} else {
		for i := t; i < n; i++ {
			swap(i, t, s)
			backtrack(t + 1)
			swap(i, t, s)
		}
	}
}

/**
 * 回溯法求集合中元素奇偶性相间的排列,依次递归
 * @param t
 */
func backtrack1(t int) {
	if t >= n {
		output(s)
	} else {
		for i := t; i < n; i++ {
			swap(i, t, s)
			if legal(x, t)
			backtrack1(t + 1)
			swap(i, t, s)
		}
	}
}

/**
 * 对子集中元素奇偶性进行判断
 * @param x
 * @param t
 * @return
 */
func legal(x []int, t int) bool {
	bRet := true //判断是否需要剪枝

	//奇偶相间,即每隔一个数判断奇偶相同
	for i := 0; i < t-2; i++ {
		bRet = bRet & ((s[i+2]-s[i])%2 == 0)
	}

	return bRet
}

/**
 * 元素交换
 * @param i
 * @param j
 */
private
static
void
swap(int
i, int
j, int[] s) {
int tmp = s[i];
s[i] = s[j];
s[j] = tmp;
}

/**
 * 子集输出函数
 * @param x
 */
private
static
void
output(int[] s) {
for (int i = 0; i < s.length; i++) {
System.out.print(s[i]);
}
System.out.println();
}
}
