package sort

import (
	"testing"
	"common"
)

/*func TestQuickSort(t *testing.T) {
	l := []int64{49, 38, 65, 97, 76, 13, 27, 49, 78, 34, 12, 64, 1}
	t.Log(l)

	QuickSort(l, 0, len(l))
	t.Log(l)

	tl := []int64{1, 12, 13, 27, 34, 38, 49, 49, 64, 65, 76, 78, 97}
	if common.SliceEqual(l, tl) {
		t.Log("test case passed")
	} else {
		t.Log("test case failed")
	}
}*/

func TestQuickSort2(t *testing.T) {
	l := []int64{49, 38, 65, 97, 76, 13, 27, 49, 78, 34, 12, 64, 1}
	t.Log(l)

	QuickSort2(l)
	t.Log(l)

	tl := []int64{1, 12, 13, 27, 34, 38, 49, 49, 64, 65, 76, 78, 97}
	if common.SliceEqual(l, tl) {
		t.Log("test case passed")
	} else {
		t.Log("test case failed")
	}
}
