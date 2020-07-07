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
		if newSum == 0 {
			return true
		} else {
			return false
		}
	} else {
		return hasPathSum(root.Left, newSum) || hasPathSum(root.Right, newSum)
	}
}
