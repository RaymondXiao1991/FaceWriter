package queue

import (
	"testing"
	"fmt"
)

func TestPush(t *testing.T) {
	q := NewQueue()
	q.Push(77)
	q.Push(79)
	q.Push(11)
	q.Push(12)
	q.Push(321)
	q.Push(1)

	p := q.head
	for p != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
	fmt.Println()

	q.Pop()
	p = q.head
	for p != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
	fmt.Println()

	q.Pop()
	p = q.head
	for p != nil {
		fmt.Printf("%v ", p.data)
		p = p.next
	}
	fmt.Println()
}
