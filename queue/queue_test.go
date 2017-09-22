package queue

import "testing"

func TestAddNode(t *testing.T) {
	q := NewQueue()

	q.Enqueue("Queue_1")
	q.Enqueue("Queue_2")
	q.Enqueue("Queue_3")
}
