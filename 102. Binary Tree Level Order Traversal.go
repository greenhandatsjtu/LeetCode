package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	var level []int
	var queue []*TreeNode
	if root == nil {
		return res
	}
	queue = append(queue, root, nil)
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		if top == nil {
			res = append(res, level)
			level = []int{}
			if len(queue) != 0 {
				queue = append(queue, nil) // use nil to imply next level
			}
		} else {
			level = append(level, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
		}
	}
	return res
}
