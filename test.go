package main

import (
	"fmt"
	"go_dataStruct/util"
)

// func rangeStr(str string) {
// 	for i, v := range str {
// 		fmt.Println(v, str[i], i)
// 		fmt.Printf("%c,\n", v)
// 	}
// 	fmt.Println()
// }

func rangeStr(str string) {
	for i, v := range []rune(str) {
		fmt.Println(v, []rune(str)[i], i)
	}
	fmt.Println()
}

func forStr(str string) {
	fmt.Println(len([]rune(str)))
	for i := 0; i < len([]rune(str)); i++ {
		fmt.Printf("%c,", []rune(str)[i])
	}
	fmt.Println()
}

func reverseList(head *util.ListNode) *util.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := reverseList(head.Next)
	head.Next.Next = head.Next
	head.Next = nil
	return next
}

func main() {
	str1 := "csapp"
	str2 := "深入理解计算机网络"
	rangeStr(str1)
	rangeStr(str2)
	forStr(str1)
	forStr(str2)
}
