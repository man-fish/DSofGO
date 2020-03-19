/*
	Package bst
	二叉查找树
*/
package bst

import (
	"errors"
	"fmt"
)

// Node 结点类
type Node struct {
	Data        int
	Left, Right *Node
}

// NewNode 结点类的构造方法
func NewNode(data int, left, right *Node) *Node {
	return &Node{
		data,
		left,
		right,
	}
}

// BinarySearchTree 二叉查找树
type BinarySearchTree struct {
	Root *Node
}

// New 二叉查找树构造方法
func New(datas []int) (bst *BinarySearchTree, err error) {
	if len(datas) < 1 {
		return nil, errors.New("params datas should has at least one element")
	}
	bst = new(BinarySearchTree)
	bst.Root = NewNode(datas[0], nil, nil)
	for i := 1; i < len(datas); i++ {
		bst.insert(bst.Root, datas[i])
	}
	return bst, nil
}

// Insert 向二叉查找树中插入节点
func (bst *BinarySearchTree) Insert(data int) {
	bst.insert(bst.Root, data)
}

func (bst *BinarySearchTree) insert(root *Node, data int) {
	if data > root.Data {
		if root.Right == nil {
			root.Right = NewNode(data, nil, nil)
		} else {
			bst.insert(root.Right, data)
		}
	} else {
		if root.Left == nil {
			root.Left = NewNode(data, nil, nil)
		} else {
			bst.insert(root.Left, data)
		}
	}
}

// Traverse 升序遍历二叉查找树
func (bst *BinarySearchTree) Traverse() {
	fmt.Println("遍历结果为：")
	bst.inOrder(bst.Root)
}

func (bst *BinarySearchTree) inOrder(root *Node) {
	if root != nil {
		bst.inOrder(root.Left)
		fmt.Println(root.Data)
		bst.inOrder(root.Right)
	}
}

// Search 在二叉查找树中搜索
func (bst *BinarySearchTree) Search(key int) bool {
	return bst.search(bst.Root, key)
}

func (bst *BinarySearchTree) search(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Data == key {
		return true
	}

	if root.Data > key {
		return bst.search(root.Left, key)
	} else {
		return bst.search(root.Right, key)
	}
}

// Delete 删除结点
func (bst *BinarySearchTree) Delete(key int) bool {
	return bst.delete(bst.Root, key)
}

func (bst *BinarySearchTree) delete(root *Node, key int) bool {
	if root == nil {
		return false
	}

	if root.Data == key {
		return bst.del(root)
	} else if root.Data < key {
		return bst.delete(root.Right, key)
	} else {
		return bst.delete(root.Left, key)
	}
}

func (bst *BinarySearchTree) del(root *Node) bool {
	var (
		cur, next *Node
	)

	if root.Left == nil {
		root = root.Left
	} else if root.Right == nil {
		root = root.Right
	} else {
		cur = root
		next = root.Left

		for next.Right != nil {
			cur = next
			next = next.Right
		}

		root.Data = next.Data

		if cur == root {
			cur.Left = next.Left
		} else {
			cur.Right = next.Left
		}
	}
	return true
}
