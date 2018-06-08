package sorts

import (
	"common"
)

// HeapSort 堆排序
/*
堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。
堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。
- 将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区；
- 将堆顶元素R[1]与最后一个元素R[n]交换，此时得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]；
- 由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)。不断重复此过程直到有序区的元素个数为n-1，则整个排序过程完成。
*/
func HeapSort(l []int64) []int64 {
	heapSort(common.BySortIndex(l), 0, len(l))
	return l
}

func heapSort(data common.Interface, f, e int) {
	first, lo, hi := f, 0, e-f

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		heapAdjust(data, i, hi, first)
	}

	// Pop elements, largest first, into end of data.
	// 二叉树结构当中最后一个有子结点的结点
	for i := hi - 1; i >= 0; i-- {
		data.Swap(first, first+i)
		heapAdjust(data, lo, i, first)
	}
}

// heapAdjust 堆调整
func heapAdjust(data common.Interface, lo, hi, first int) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi { // child超出了数组长度，也就是该结点无孩子结点，返回
			break
		}
		if child+1 < hi && data.Less(first+child, first+child+1) { // 右孩子结点存在
			child++
		}
		// 以上是为了在孩子结点当中找到较大的结点，用child表示
		if !data.Less(first+root, first+child) {
			return
		}
		// 如果根结点小于较大的孩子结点，则进行交换；该孩子结点的堆结构可能会被影响，继续去处理孩子结点
		data.Swap(first+root, first+child)
		root = child
	}
}
