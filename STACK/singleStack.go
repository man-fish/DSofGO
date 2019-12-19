package STACK

import (
	"go_DataStruct/LIST"
)

type SingleStack struct {
	list *LIST.SingleList
}

func NewSingleStack() *SingleStack {
	return &SingleStack{list:LIST.NewEmptySingleList()}
}

func (sgk *SingleStack) IsEmpty() bool {
	return sgk.list.IsEmpty()
}

func (sgk *SingleStack) Push(x interface{}) {
	sgk.list.Insert(0,x)
}

func (sgk *SingleStack) Pop() interface{} {
	return sgk.list.Delete(0).Data
}

func (sgk *SingleStack) Peek() interface{} {
	return sgk.list.Get(0)
}