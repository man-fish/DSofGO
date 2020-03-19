/*
问题：
	给定一个二叉树，它的每个结点都存放一个 0-9 的数字，每条从根到叶子节点的路径都代表一个数字,计算从根到叶子节点生成的所有数字之和。
说明:
	叶子节点是指没有子节点的节点。
示例 1:
输入:
	[4,9,0,5,1]
    4
   / \
  9   0
 / \
5   1
输出:
	1026
解释:
	从根到叶子节点路径 4->9->5 代表数字 495.
	从根到叶子节点路径 4->9->1 代表数字 491.
	从根到叶子节点路径 4->0 代表数字 40.
	因此，数字总和 = 495 + 491 + 40 = 1026.
*/

package main

import "fmt"

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

func sumNumbers(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	summer(root, 0, &sum)
	return sum
}

func summer(root *TreeNode, cur int, sum *int) {
	cur = cur*10 + root.Val
	if root.Right == nil && root.Left == nil {
		*sum = *sum + cur
		return
	}

	if root.Left != nil {
		summer(root.Left, cur, sum)
	}

	if root.Right != nil {
		summer(root.Right, cur, sum)
	}
}

func sumNumbersV2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return helper(root, 0)
}

func helper(root *TreeNode, i int) int {
	if root == nil {
		return 0
	}
	tmp := i*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return tmp
	}
	return helper(root.Left, tmp) + helper(root.Right, tmp)
}

func main() {
	tree := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	res := sumNumbers(tree)
	fmt.Println(sumNumbersV2(tree))
	fmt.Println(res)
}
