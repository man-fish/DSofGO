/*
给定一个二叉树，判断其是否是一个有效的二叉搜索树。
假设一个二叉搜索树具有如下特征：
节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:
    输入:
     5
    / \
    1  4
      / \
     3   6
    输出: false
    解释: 输入为: [5,1,4,null,null,3,6]。
         根节点的值为 5 ，但是其右子节点值为 4
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

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sortedArr := make([]int, 0)
	traverse(root, &sortedArr)
	for i, v := range sortedArr {
		if i != 0 && sortedArr[i-1] >= v {
			return false
		}
	}
	return true
}

func traverse(root *TreeNode, sortedArr *[]int) {
	if root != nil {
		traverse(root.Left, sortedArr)
		*sortedArr = append(*sortedArr, root.Val)
		traverse(root.Right, sortedArr)
	}
}

func isValidBSTV2(root *TreeNode) bool {
	return helper(root, -1, -1)
}

func helper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if min != -1 && root.Val <= min {
		return false
	}
	if max != -1 && root.Val >= max {
		return false
	}
	if !helper(root.Left, min, root.Val) {
		return false
	}

	if !helper(root.Right, root.Val, max) {
		return false
	}
	return true
}

func append_test(arr *[]int) {
	*arr = append(*arr, 1)
}

func main() {
	arr := make([]int, 0)
	append_test(&arr)
	fmt.Print(arr)
	ar := &arr
	fmt.Println((*ar)[0] == (*ar)[0])
}
