package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generate(1, n)
}

func generate(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		lefts := generate(start, i-1)
		rights := generate(i+1, end)
		for j := 0; j < len(lefts); j++ {
			for k := 0; k < len(rights); k++ {
				root := &TreeNode{Val: i}
				root.Left = lefts[j]
				root.Right = rights[k]
				res = append(res, root)
			}
		}
	}
	return res
}
