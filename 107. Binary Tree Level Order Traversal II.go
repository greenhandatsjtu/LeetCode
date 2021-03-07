package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	results := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		n := len(queue)
		result := make([]int, 0, n)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			result = append(result, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		results = append(results, result)
	}
	for i := 0; i < len(results)/2; i++ {
		results[i], results[len(results)-i-1] = results[len(results)-i-1], results[i]
	}
	return results
}
