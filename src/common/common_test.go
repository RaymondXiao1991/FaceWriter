package common

import (
	"testing"
)

func TestSliceEqual(t *testing.T) {
	s1 := []string{"1", "2"}
	s2 := []string{"1", "2"}

	if SliceEqual(s1, s2) {
		t.Log("test case1 passed")
	} else {
		t.Error("Test case1 failed, Got ", SliceEqual(s1, s2), "Expected true")
	}

	s3 := []int{1, 2}
	s4 := []int{1, 2}

	if SliceEqual(s3, s4) {
		t.Log("test case2 passed")
	} else {
		t.Error("Test case2 failed, Got ", SliceEqual(s3, s4), "Expected true")
	}
}

func TestGenerator(t *testing.T) {
	for i := range Generator() {
		if i == 5 {
			break
		}
		t.Log(i)
	}
}

func TestReverse(t *testing.T) {
	s := Reverse("cdfeaj")
	t.Log(s)
}
