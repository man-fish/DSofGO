package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

type pq struct {
	queue []*ListNode
	size  int
}

func New(size int) *pq {
	return &pq{
		queue: make([]*Node, size+1),
	}
}

func (pq *pq) IsEmpty() bool {
	return pq.size == 0
}

func (pq *pq) DelMax() *ListNode {
	l := pq.queue[1]
	pq.swap(1, pq.size)

	pq.queue[pq.size] = nil
	pq.size--
	pq.sink(1)
	return l
}

func (pq *pq) Insert(l *ListNode) {
	pq.size++
	pq.queue[pq.size] = l
	pq.swim(pq.size)
}

func (pq *pq) swim(index int) {
	for index > 1 && pq.less(index/2, index) {
		pq.swap(index/2, index)
		index = index / 2
	}
}

func (pq *pq) sink(index int) {
	for index*2 <= pq.size {
		k := index * 2
		if k < pq.size && pq.less(k, k+1) {
			k = k + 1
		}
		if pq.less(k, index) {
			break
		}
		pq.swap(k, index)
		index = k
	}
}

func (pq *pq) swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
}

func (pq *pq) less(i, j int) bool {
	return pq.queue[i].Val > pq.queue[j].Val
}

func mergeKLists(lists []*ListNode) *ListNode {
	proQueue := New(len(lists))
	for _, node := range lists {
		if node == nil {
			continue
		}
		proQueue.Insert(node)
	}
	l := new(ListNode)
	head := l

	for {
		cur := proQueue.DelMax()
		l.Val = cur.Val

		if cur.Next != nil {
			proQueue.Insert(cur.Next)
		}

		if proQueue.IsEmpty() {
			break
		}
		l.Next = new(ListNode)
		l = l.Next
	}
	return head
}

func main() {
	// l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	// l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	// l3 := &ListNode{2, &ListNode{6, nil}}
	l4 := make([]*ListNode, 1)

	// l3 := &ListNode{1, &ListNode{3, &ListNode{8, &ListNode{9, nil}}}}
	l := mergeKLists(l4)
	for l != nil {
		fmt.Print("->", l.Val)
		l = l.Next
	}
}

// func (l *ListNode) IsEqual(a, b interface{}) bool {
// 	return a.(*ListNode).Val == b.(*ListNode).Val
// }

// func (l *ListNode) CompareTo(a interface{}, b interface{}) int {
// 	if a.(*ListNode).Val < b.(*ListNode).Val {
// 		return 1
// 	} else if a.(*ListNode).Val < b.(*ListNode).Val {
// 		return -1
// 	} else {
// 		return 0
// 	}
// }

// func mergeKLists(lists []*ListNode) *ListNode {
// 	var (
// 		proQueue heap.PriorityQueue

// 		l    *ListNode
// 		head *ListNode

// 		cur *ListNode
// 	)

// 	proQueue = heap.NewPqHeap(len(lists))
// 	for _, node := range lists {
// 		proQueue.Insert(node)
// 	}

// 	l = new(ListNode)
// 	head = l

// 	for !proQueue.IsEmpty() {
// 		cur = proQueue.DelMax().(*ListNode)
// 		l.Val = cur.Val
// 		l.Next = new(ListNode)
// 		l = l.Next
// 		if cur.Next != nil {
// 			proQueue.Insert(cur.Next)
// 		}
// 	}
// 	return head
// }
