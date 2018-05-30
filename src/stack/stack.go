package stack

type Node struct {
	Value interface{}
}

type Stack []*Node

func NewStack() *Stack {
	return &Stack{}
}

func (q *Stack) Push(n *Node) {
	*q = append(*q, n)
}

func (q *Stack) Pop() (n *Node) {
	x := q.Len() - 1
	n = (*q)[x]
	*q = (*q)[:x]
	return
}

func (q *Stack) Len() int {
	return len(*q)
}
