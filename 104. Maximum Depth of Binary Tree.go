package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func depth(root *TreeNode, n int) int {
	if root == nil {
		return n
	}
	n += 1
	return max(depth(root.Left, n), depth(root.Right, n))
}

func maxDepth(root *TreeNode) int {
	return depth(root, 0)
}
