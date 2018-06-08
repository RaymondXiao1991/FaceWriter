package sorts

// MergeSort 归并排序
/*
- 把长度为n的输入序列分成两个长度为n/2的子序列；
- 对这两个子序列分别采用归并排序；
- 将两个排序好的子序列合并成一个最终的排序序列。
 */
func MergeSort(l []int64) []int64 {
	length := len(l)
	if length < 2 {
		return l
	}
	mid := length / 2
	//fmt.Println(l)
	return merge(MergeSort(l[:mid]), MergeSort(l[mid:]))
}

func merge(left, right []int64) (result []int64) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] <= right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	//fmt.Println("left:", left, "right:", right)

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
