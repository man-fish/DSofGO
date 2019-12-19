package QUEUE

import (
	"fmt"
)

type SeqQueue struct {
	list		[]interface{}
	front,rear	int
}

func NewSequeue(length int) *SeqQueue {
	return &SeqQueue{
		list:make([]interface{},length),
		front:0,
		rear:0,
	}
}

func (q *SeqQueue) IsEmpty() bool {
	return q.rear == 0 && q.front == 0
}

func (q *SeqQueue) Add(x interface{}) bool {
	if x == nil {
		return false
	}

	if (q.rear+1)%len(q.list) == q.front {
		temp := q.list
		q.list = make([]interface{},len(temp)*2)
		j := 0
		for i := q.front; i != q.rear ; i = (i+1)%len(temp)  {
			q.list[j] = temp[i]
			j++
		}
		q.rear = 0
		q.rear = j
	}
	fmt.Println("当前数组长度:",len(q.list))
	q.list[q.rear] = x
	q.rear = (q.rear+1)%len(q.list)
	return true
}

func (q *SeqQueue) Poll() interface{} {
	if q.IsEmpty() {
		return nil
	}
	defer func() {
		q.front = (q.front+1)%len(q.list)
	}()
	return q.list[q.front]
}

func (q *SeqQueue) Peek() interface{} {
	if q.IsEmpty() {
		return nil
	}
	return q.list[q.front]
}