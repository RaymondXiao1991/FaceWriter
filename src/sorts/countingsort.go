package sorts

func CountingSort(l []int64, length int) {
	if l == nil || length <= 0 {
		return
	}

	var max int64
	for i := 0; i < length; i++ {
		if l[i] > max {
			max = l[i]
		}
	}

	tmpArr := make([]int64, max+1)
	for i := 0; i < length; i++ {
		tmpArr[l[i]]++
	}

	/*
	j := 0
	var t int64
	for t = 0; t < max+1; t++ {
		if tmpArr[t] > 0 {
			l[j] = t
			tmpArr[t]--
			j++
		}
		fmt.Println(l)
	}*/

	// 统计数组中，每个元素出现的次数
	for i := 0; i < length; i++ {
		tmpArr[l[i]]++
	}

	// 统计数组计数，每项存前N项和，这实质为排序过程
	for k := 1; k < length; k++ {
		tmpArr[k] += tmpArr[k-1]
	}

	//将计数排序结果转化为数组元素的真实排序结果
	for j := n - 1; j >= 0; j-- {
		elem := l[j]           //取待排序元素
		index := tmpArr[elem] - 1; //待排序元素在有序数组中的序号
		sorted_arr[index] = elem;    //将待排序元素存入结果数组中
		count_arr[elem]--;           //修正排序结果，其实是针对算得元素的修正
	}
}
