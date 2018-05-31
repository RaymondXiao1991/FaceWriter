package common

import "testing"

func TestSliceEqual(t *testing.T) {
	s1 := []string{"1", "2"}
	s2 := []string{"1", "2"}

	if SliceEqual(s1, s2) {
		t.Log("test case passed")
	} else {
		t.Log("test case failed")
	}
}
