package queue

import (
	"testing"
	"fmt"
	"time"
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
}

func TestPop(t *testing.T) {
	q := NewQueue()
	q.Push(77)
	q.Push(79)
	q.Push(11)
	q.Push(12)
	q.Push(321)
	q.Push(1)

	q.Pop()
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
}

func TestTimeBefore(t *testing.T) {
	y, m, d := time.Now().Date()
	day := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	flag := day.Before(time.Unix(1527782400, 0))
	if !flag {
		t.Log("test case passed")
	} else {
		t.Log("test case failed")
	}
}
