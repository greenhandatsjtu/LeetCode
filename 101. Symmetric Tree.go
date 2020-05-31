package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func addNodesToStack(root *TreeNode, stack *[]int) {
	if root == nil {
		*stack = append(*stack, math.MinInt32)
		return
	}
	*stack = append(*stack, root.Val)
	addNodesToStack(root.Left, stack)
	addNodesToStack(root.Right, stack)
}

func checkSymmetry(root *TreeNode, stack *[]int, index *int) bool {
	if *index >= len(*stack) {
		return root == nil
	}
	if root == nil {
		*index++
		return (*stack)[*index-1] == math.MinInt32
	}
	if root.Val != (*stack)[*index] {
		return false
	}
	*index++
	return checkSymmetry(root.Right, stack, index) && checkSymmetry(root.Left, stack, index)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var stack []int
	var index int
	addNodesToStack(root.Left, &stack)
	return checkSymmetry(root.Right, &stack, &index)
}
