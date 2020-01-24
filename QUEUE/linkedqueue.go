package QUEUE

import "go_DataStruct/LIST"

type LinkedQueue struct {
	front,rear *LIST.Node

}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{nil,nil}
}
func (q *LinkedQueue) IsEmpty() bool {
	return q.rear == nil && q.front == nil
}

func (q *LinkedQueue) Add(x interface{}) bool {
	if x == nil {
		return false
	}
	newnode := &LIST.Node{x,nil}
	if q.front == nil {
		q.front = newnode
		q.rear = newnode
	}else{
		q.rear.Next = newnode
	}
	q.rear = newnode
	return true
}

func (q *LinkedQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	defer func() {
		q.front = q.front.Next
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