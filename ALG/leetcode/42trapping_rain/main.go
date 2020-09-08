package main

import (
	"fmt"
	"go_dataStruct/STACK"
)

// trapping-rain-water

func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	barrel := STACK.NewSeqStack(n)
	curStart := height[0]
	count := 0
	for i := 1; i < n; i++ {
		if barrel.IsEmpty() {
			if curStart < height[i] {
				curStart = height[i]
				continue
			}
			barrel.Push(curStart)
		}
		if height[i] < curStart {
			barrel.Push(height[i])
		} else {
			for !barrel.IsEmpty() {
				count += curStart - barrel.Pop().(int)
			}
			curStart = height[i]
		}
	}
	curStart = 0
	for !barrel.IsEmpty() {
		cur := barrel.Pop()

	}
	return count
}

func main() {
	res := trap([]int{3, 1, 2})
	fmt.Println(res)
}
