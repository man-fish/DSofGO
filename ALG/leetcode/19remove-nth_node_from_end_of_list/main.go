package main

import "fmt"
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
type ListNode struct {
	Val int
	Next *ListNode
}     

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	for n != 0 {
		cur = cur.Next
		n--
	}
	
	lastN := head

	if cur == nil {
		return lastN.Next
	}

	for cur.Next != nil {
		cur = cur.Next
		lastN = lastN.Next
	}
	lastN.Next = lastN.Next.Next
	return head

}

func main() {
	l1 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l1 = removeNthFromEnd(l1, 3)
	for l1 != nil {
		fmt.Println(l1)
		l1 = l1.Next
	}
}
