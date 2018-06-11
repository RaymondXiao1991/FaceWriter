package longest_increasing_subsequence

import (
	"testing"
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
	t.Log(str)
	/*for i := 0; i < len(str); i++ {
		fmt.Println(str[i].data)
	}*/
}
