/*
	package main
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
	有效字符串需满足：
	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。
	示例:
		输入: "()[]{}"
		输出: true
	思路：
		棧
*/

package main

import (
	"fmt"
	"go_DataStruct/STACK"
)

func isValid(s string) bool {
	stk := STACK.NewSeqStack(len(s))
	for i, v := range s {
		if i == 0 {
			stk.Push(v)
			continue
		}
		if stk.Peek() != nil {
			if stk.Peek().(int32)+2 == v || stk.Peek().(int32)+1 == v {
				stk.Pop()
			} else {
				stk.Push(v)
			}
		} else {
			stk.Push(v)
		}
	}

	return stk.IsEmpty()
}

func main() {
	fmt.Println(isValid("([)"))
}
