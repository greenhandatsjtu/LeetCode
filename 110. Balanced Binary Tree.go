package main

func isBalanced(root *TreeNode) bool {
	return maxDepth(root) != -1
}

// return -1 if child tree isn't balanced
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left == -1 || right == -1 || left-right > 1 || right-left > 1 {
		return -1
	}
	return max(left, right) + 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
