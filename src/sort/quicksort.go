package sort

import "fmt"

// QuickSort 快速排序
/*
- 从数列中挑出一个元素，称为 “基准”（pivot）；
- 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
- 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
 */
func QuickSort(l []int64, left, right int) {
	if left < right {
		p := partition(l, left, right)
		QuickSort(l, left, p-1)
		QuickSort(l, p+1, right)
	}
}

// 数组分区，左小右大
func partition(l []int64, left, right int) int {
	pivot := left // 设定基准值（pivot）
	storeIndex := pivot + 1
	for i := storeIndex; i <= right; i++ {
		if l[i] < l[pivot] {
			l[storeIndex], l[i] = l[i], l[storeIndex] // 交换位置
			storeIndex++
		}
	}
	l[pivot], l[ storeIndex-1] = l[ storeIndex-1], l[pivot] // 将基准元素放置到最后的正确位置上
	return storeIndex - 1
}

func QuickSort2(l []int64) {
	length := len(l)
	if length <= 1 {
		return
	}
	mid, i := l[0], 1         // 取第一个元素作为分水岭，i下标初始为1，即分水岭右侧的第一个元素的下标
	head, tail := 0, length-1 // 头尾的下标

	// 如果头和尾没有相遇，就会一直触发交换
	for head < tail {
		fmt.Println(l, head, tail, i)
		if l[i] > mid {
			// 如果分水岭右侧的元素大于分水岭，就将右侧的尾部元素和分水岭右侧元素交换
			l[i], l[tail] = l[tail], l[i]
			tail-- // 尾下标左移一位
		} else {
			// 如果分水岭右侧的元素小于等于分水岭，就将分水岭右移一位
			l[i], l[head] = l[head], l[i]
			head++ // 头下标右移一位
			i++    // i下标右移一位
		}
		fmt.Printf("----{head:%d,tail:%d,i:%d,mid:%d} values:%v\n", head, tail, i, mid, l)
	}

	// 分水岭左右的元素递归做同样处理
	QuickSort2(l[:head])
	QuickSort2(l[head+1:])
}
