package sorts

// DirectInsertSort 直接插入排序
// 从后向前找到合适位置后插入
// 基本思想：每步将一个待排序的记录，按其顺序码大小插入到前面已经排序的字序列的合适位置（从后向前找到合适位置后），直到全部插入排序完为止。
func DirectInsertSort(l []int64) {
	for i := 1; i < len(l); i++ {
		temp := l[i] // 待插入元素
		j := 0
		for j = i - 1; j >= 0; j-- {
			// 将大于temp的往后移动一位
			if l[j] > temp {
				l[j+1] = l[j]
			} else {
				break
			}
		}
		l[j+1] = temp
	}
}

// BinaryInsertSort 二分法插入排序
// 基本思想：二分法插入排序的思想和直接插入一样，只是找合适的插入位置的方式不同，这里是按二分法找到合适的位置，可以减少比较的次数。
func BinaryInsertSort(l []int64) {
	for i := 1; i < len(l); i++ {
		temp := l[i] // 待插入元素
		left, mid, right := 0, 0, i-1

		for left <= right {
			mid = (left + right) / 2
			if temp < l[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}

		// 循环完后，start=end+1,此时start为当前插入数字所待坑位
		// 把坑位给当前插入的数据挪出来
		for j := i - 1; j >= left; j-- {
			l[j+1] = l[j]
		}
		// 将当前插入数字挪入它该待的坑位
		l[left] = temp
	}
}

// ShellSort 希尔排序
// 基本思想：先取一个小于n的整数d1作为第一个增量，把文件的全部记录分成d1个组。所有距离为d1的倍数的记录放在同一个组中。先在各组内进行直接插入排序；然后，取第二个增量d2
func ShellSort(l []int64) {
	/*length := len(l)
	tmp, gap := 0, 1
	for gap < length/3 {
		gap = gap*3 + 1
	}

	for gap; gap > 0; gap = math.Floor(float64(gap / 3)) {
		for i := gap; i < length; i++ {
			tmp = l[i]
			for j := i - gap; j > 0 && l[j] > tmp; j -= gap {
				l[j+gap] = l[j]
			}
			l[j+gap] = tmp
		}
	}*/
}
