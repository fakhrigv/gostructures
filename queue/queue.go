package queue

import "github.com/altay13/godatastructures/linkedlist"

// Queue implementation
type Queue struct {
	list *linkedlist.List
}

// NewQueue method creates the stack
func NewQueue() *Queue {
	l := linkedlist.NewList()
	q := &Queue{
		list: l,
	}
	return q
}

// Enqueue ...
func (q *Queue) Enqueue(item interface{}) {
	q.list.AddNode(item)
}

// Dequeue ...
func (q *Queue) Dequeue() interface{} {
	return q.list.GetLastNode()
}
