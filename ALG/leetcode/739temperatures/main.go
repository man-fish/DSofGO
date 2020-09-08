package main

import "go_dataStruct/STACK"

func dailyTemperatures(T []int) []int {
	n := len(T)
	stk := STACK.NewSeqStack(n)
	ans := make([]int, n)
	for i := range T {
		for !stk.IsEmpty() && T[stk.Peek().(int)] < T[i] {
			currentIdx := stk.Pop().(int)
			ans[currentIdx] = i - currentIdx
		}
		stk.Push(i)
	}
	return ans
}
