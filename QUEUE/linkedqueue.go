package QUEUE

import "go_DataStruct/LIST"

type LinkedQueue struct {
	front,  rear *LIST.Node
	size int64
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil, nil, 0}
} 
func (q *LinkedQueue) IsEmpty() bool {
	return q.rear == nil && q.front == nil
}

func (q *LinkedQueue) Size() int64 {
	return q.size
}

func (q *LinkedQueue) Add(x interface{}) bool {
	if x == nil {
		return false
	}
	newnode := &LIST.Node{x, nil}
	if q.front == nil { 
		q.front = newnode
		q.rear = newnode
	} else {
		 q.rear.Next = newnode
	}
	q.rear = newnode
	q.size++
	return true
}

func (q *LinkedQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	defer func() {
		q.front = q.front.Next
		q.size--
		if q.front == nil {
			q.rear = nil
		}
	}()
	return q.front.Data
}

func (q *LinkedQueue) Peek() interface{} {
	if !q.IsEmpty() {
		return q.front.Data
	}
	return nil
}

