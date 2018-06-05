package queue

import "fmt"

// 队列是一种先进先出（First In First Out）的线性表，简称FIFO

type node struct {
	data interface{}
	next *node
}

//type Queue []*Node
type Queue struct {
	head  *node // 队头允许删除
	tail  *node // 队尾允许 插入
	count int
}

//	Creates a new pointer to a new queue.
func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(d interface{}) {
	n := &node{data: d}
	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	q.count++
}

func (q *Queue) Pop() (n interface{}) {
	if q.head == nil {
		return nil
	} else {
		n = q.head
		q.head = q.head.next
	}

	q.count--
	return n
}

func (q *Queue) Len() int {
	return q.count
}

func (q *Queue) Empty() bool {
	return q.count == 0
}

func (n *node) String() string {
	return fmt.Sprint(n.data)
}
