package TREE

import "fmt"

type BinaryNode struct {
	Left,Right	*BinaryNode
	Data		interface{}
}

func NewBinaryNode(data interface{},left *BinaryNode,right *BinaryNode) *BinaryNode{
	return &BinaryNode{
		Data:data,
		Left:left,
		Right:right,
	}
}

func (bn *BinaryNode) IsLeaf() bool {
	return bn.Right == nil && bn.Left == nil
}

func (bn *BinaryNode) String() string {
	return fmt.Sprint(bn.Data)
}