package main

import (
	"fmt"
	"go_DataStruct/STACK"
)

type Stack interface {
	Push(x interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
}

func testStack(stk Stack) {
	stk.Push(1)
	fmt.Println(stk.Peek())
	fmt.Println(stk.Pop())
	stk.Push(1)
	stk.Push(1)
	stk.Push(1)
	stk.Push(1)
	stk.Push(1)
	stk.Push(1)
	fmt.Println(stk.Pop())
}



func main () {
	sgk1 := STACK.NewSingleStack()
	testStack(sgk1)
	sgk2 := STACK.NewSeqStack(5)
	testStack(sgk2)
}
