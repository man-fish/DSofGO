package TREE

import "fmt"

type TriNode struct {
	Data	int
	Parent,Left,Right	int
}

func NewTriNode(data,parent,left,right int) *TriNode {
	return &TriNode{
		Data:data,
		Parent:parent,
		Left:left,
		Right:right,
	}
}

func NewEmptyTriNode(data int) *TriNode {
	return NewTriNode(data,-1,-1,-1)
}

func (tn *TriNode) IsLeaf() bool {
	return tn.Left == -1 && tn.Right == -1
}

func (tn *TriNode) String() string {
	return fmt.Sprint(tn.Data)
}