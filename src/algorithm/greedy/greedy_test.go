package greedy

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
	store := GreedyKnapSack([]int{9, 6, 4, 1}, []int{3, 5, 1, 7}, 4, 3)
	t.Log(store)

	if store == 36 {
		t.Log("test case1 passed")
	} else {
		t.Error("Test case1 failed, Got ", store, "Expected 36")
	}
}

func TestGreedyActivitySelector(t *testing.T) {
	b := GreedyActivitySelector([]int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}, []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})
	t.Log(b)

}
