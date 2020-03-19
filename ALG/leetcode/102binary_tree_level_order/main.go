/*
给定一个二叉树，返回其按层次遍历的节点值。 （即逐层地，从左到右访问所有节点）。
例如:
给定二叉树: [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：
[
	[3],
	[9,20],
	[15,7]
]
*/

package main

import (
	"fmt"
	"go_dataStruct/QUEUE"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	queue := QUEUE.NewLinkedQueue()
	queue.Add(root)

	for !queue.IsEmpty() {
		store := make([]int, 0)
		preNodes := make([]*TreeNode, 0)
		for !queue.IsEmpty() { // 如果说你不能访问queue内部的元素，那么可以采用以下方法，但是实际上是可以的。
			cur := queue.Poll().(*TreeNode)
			store = append(store, cur.Val)

			if cur.Left != nil {
				preNodes = append(preNodes, cur.Left)
			}

			if cur.Right != nil {
				preNodes = append(preNodes, cur.Right)
			}
		}
		for _, v := range preNodes {
			queue.Add(v)
		}
		res = append(res, store)
	}

	return res
}

func main() {
	tree := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}
	tree = &TreeNode{}
	res := levelOrder(tree)
	fmt.Println(res)
}
