package main

import "fmt"

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代
func swapPairs(head *ListNode) *ListNode {
	res := head
	if head != nil && head.Next != nil { // 特殊处理前两个节点,因为前两个顶点没有下面迭代中所谓的head。
		next := head.Next
		head.Next = next.Next
		next.Next = head
		res = next
	}
	// 2-> 1 -> 3 -> 4 -> nil
	//     head cur next
	for head != nil {
		cur := head.Next
		if cur == nil {
			break
		}
		next := cur.Next
		if next == nil {
			break
		}
		cur.Next = next.Next
		next.Next = cur
		head.Next = next
		head = head.Next.Next
	}
	return res
}

// 递归
func swapPairsV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fir := head
	sec := head.Next

	fir.Next = swapPairsV2(sec.Next)
	sec.Next = fir
	return sec
}

// 翻转

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	node = swapPairsV2(node)
	index := 0

	for node != nil {
		if index > 4 {
			break
		}
		index++
		fmt.Println(node.Val)
		node = node.Next
	}
}
