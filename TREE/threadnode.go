package TREE

import "fmt"

type ThreadNode struct {
	Data interface{}
	Left,Right *ThreadNode
	Ltag,Rtag	bool
}

func NewThreadNode(data interface{}, left *ThreadNode, right *ThreadNode, Ltag bool, Rtag bool) *ThreadNode {
	return &ThreadNode{
		data,
		left,
		right,
		Ltag,
		Rtag,
	}
}

func (tn *ThreadNode) IsLeaf() bool {
	return tn.Left == nil && tn.Right == nil
}

func (tn *ThreadNode) String() string {
	return fmt.Sprint(tn.Data)
}