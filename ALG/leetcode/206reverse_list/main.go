package main

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nxt := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return nxt
}

func reverseListWrong(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nxt := reverseList(head.Next)
	nxt.Next = head
	head.Next = nil

	return head
}

func main() {

}
