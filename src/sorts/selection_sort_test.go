package sorts

import (
	"testing"
	"common"
)

func TestSelectionSort(t *testing.T) {
	l := []int64{49, 38, 65, 97, 76, 13, 27, 49, 78, 34, 12, 64, 1}
	t.Log(l)

	SelectionSort(l)
	t.Log(l)

	tl := []int64{1, 12, 13, 27, 34, 38, 49, 49, 64, 65, 76, 78, 97}
	if common.SliceEqual(l, tl) {
		t.Log("test case passed")
	} else {
		t.Error("test case failed")
	}
}
