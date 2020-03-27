package QUEUE

import "go_DataStruct/LIST"

type DoubleLinkedQueue struct {
	front *LIST.DoubleNode
}

/**
双节点链式队列的结构可以说是非常巧妙了，DoubleNode的prev和next视作队列的出口和入口，front可以视作一个中间元素不储存值但是却尤为重要。
*/

func NewDoubleLinkedQueue() *DoubleLinkedQueue {
	return &DoubleLinkedQueue{new(LIST.DoubleNode)}
}
func (q *DoubleLinkedQueue) IsEmpty() bool {
	return q.front.Next == nil && q.front.Prev == nil
}

func (q *DoubleLinkedQueue) Add(x interface{}) bool {
	if x == nil {
		return false
	}
	if q.IsEmpty() {
		newnode := &LIST.DoubleNode{x, q.front, q.front}
		q.front.Next = newnode
		q.front.Prev = newnode
	} else {
		newnode := &LIST.DoubleNode{x, q.front.Prev, q.front}
		q.front.Prev.Next = newnode
		q.front.Prev = newnode
	}
	return true
}

func (q *DoubleLinkedQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	defer func() {
		q.front.Next.Next.Prev = q.front
		q.front.Next = q.front.Next.Next
	}()
	return q.front.Next.Data
}

func (q *DoubleLinkedQueue) Peek() interface{} {
	if !q.IsEmpty() {
		return q.front.Next.Data
	}
	return nil
}
