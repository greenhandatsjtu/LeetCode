package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	newSum := sum - root.Val
	if root.Left == nil && root.Right == nil {
		return newSum == 0
	} else {
		return hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum)
	}
}
