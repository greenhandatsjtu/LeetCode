package main

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var record []*TreeNode
	tmp := math.MinInt64
	for len(record) > 0 || root != nil {
		for root != nil {
			record = append(record, root)
			root = root.Left
		}
		root = record[len(record)-1]
		record = record[:len(record)-1]
		if (*root).Val <= tmp {
			return false
		}
		tmp = (*root).Val
		root = root.Right
	}
	return true
}
