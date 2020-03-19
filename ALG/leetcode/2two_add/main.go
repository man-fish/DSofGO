/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
示例：
	输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
	输出：7 -> 0 -> 8
	原因：342 + 465 = 807
*/
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	list := l1

	for l1 != nil && l2 != nil {
		curResult := l1.Val + l2.Val
		if curResult >= 10 {
			curResult = curResult % 10
			if l1.Next == nil {
				l1.Next = &ListNode{Val: 1, Next: nil}
			} else {
				l1.Next.Val++
			}
		}

		l1.Val = curResult
		if l1.Next == nil && l2.Next != nil {
			l1.Next = l2.Next
			return list
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if l1.Val >= 10 {
			if l1.Next != nil {
				l1.Next.Val++
			} else {
				l1.Next = &ListNode{Val: 1, Next: nil}
			}
		}

		l1.Val = l1.Val % 10
		l1 = l1.Next
	}

	return list
}

// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9, Next: nil}}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
	l3 := addTwoNumbers(l1, l2)
	for l3 != nil {
		fmt.Print(l3.Val)
		l3 = l3.Next
	}
}
