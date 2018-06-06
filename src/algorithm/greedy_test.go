package algorithm

import "testing"

func TestGreedyCutRope(t *testing.T) {
	store := GreedyCutRope(10)

	if store == 36 {
		t.Log("test case1 passed")
	} else {
		t.Error("Test case1 failed, Got ", store, "Expected 36")
	}
}

func TestGreedyCutRopeByRecursive(t *testing.T) {
	store := GreedyCutRopeByRecursive(10)

	if store == 36 {
		t.Log("test case1 passed")
	} else {
		t.Error("Test case1 failed, Got ", store, "Expected 36")
	}
}

func TestGreedyKnapSack(t *testing.T) {
	store := GreedyKnapSack([]int{9, 6, 4, 1}, []int{3, 5, 1, 7}, 3)
	t.Log(store)

	if store == 36 {
		t.Log("test case1 passed")
	} else {
		t.Error("Test case1 failed, Got ", store, "Expected 36")
	}
}
