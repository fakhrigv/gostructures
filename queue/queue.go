package queue

import (
	"github.com/altay13/godatastructures/doublylinkedlist"
)

// Queue implementation
type Queue struct {
	list *doublylinkedlist.List
	size int
}

// NewQueue method creates the stack
func NewQueue() *Queue {
	l := doublylinkedlist.NewList()
	q := &Queue{
		list: l,
	}
	return q
}

// Enqueue ...
func (q *Queue) Enqueue(item interface{}) {
	q.list.AddNode(item)
	q.size++
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	n := q.list.GetLastNode()
	q.list.RemoveLastNode()
	if n != nil {
		q.size--
	}
	return n.Item
}
