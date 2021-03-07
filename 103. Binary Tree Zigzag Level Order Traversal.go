package main

func zigzagLevelOrder(root *TreeNode) [][]int {
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
	for i := 0; i < len(results); i++ {
		if i&1 == 1 {
			for j := 0; j < len(results[i])/2; j++ {
				results[i][j], results[i][len(results[i])-j-1] = results[i][len(results[i])-j-1], results[i][j]
			}
		}
	}
	return results
}
