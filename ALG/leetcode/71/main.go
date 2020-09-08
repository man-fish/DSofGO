package main

import (
	"fmt"
	"go_dataStruct/STACK"
	"strings"
)

func simplifyPath(path string) string {
	stk := STACK.NewSeqStack(len(path))
	parts := strings.Split(path, "/")
	for _, part := range parts {
		if part == "." || part == "" {
			continue
		}
		if part == ".." {
			if !stk.IsEmpty() {
				stk.Pop()
			}
			continue
		}
		stk.Push(part)
	}
	ans := ""
	for !stk.IsEmpty() {
		ans = "/" + stk.Pop().(string) + ans
	}
	if ans == "" {
		ans = "/"
	}
	return ans
}

func main() {
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
}
