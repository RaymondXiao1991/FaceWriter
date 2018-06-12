package longest_increasing_subsequence

import (
	"testing"
	"common"
)

func TestLongestIncreasingSubSequence(t *testing.T) {
	s := []int{5, 6, 7, 1, 2, 8}
	length := LongestIncreasingSubSequence(s)

	if length == 4 {
		t.Log("test case passed")
	} else {
		t.Error("Test case failed, Got", length, ", Expected 4")
	}
}

func TestLongestIncreasingSubSequenceDetail(t *testing.T) {
	s := []int{5, 6, 7, 1, 2, 8}
	str := LongestIncreasingSubSequenceDetail(s)
	if common.SliceEqual(str, []int{5, 6, 7, 8}) {
		t.Log("test case passed")
	} else {
		t.Error("Test case failed, Got", str, ", Expected [5 6 7 8]")
	}

	s = []int{10, 22, 9, 33, 21, 50, 41, 60, 80}
	str = LongestIncreasingSubSequenceDetail(s)
	if common.SliceEqual(str, []int{10, 22, 33, 41, 60, 80}) {
		t.Log("test case passed")
	} else {
		t.Error("Test case failed, Got", str, ", Expected [10, 22, 33, 41, 60, 80]")
	}
}
