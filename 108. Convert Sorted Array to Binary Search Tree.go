package main

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	var construct func(low, high int) *TreeNode
	construct = func(low, high int) *TreeNode {
		if low > high {
			return nil
		}
		mid := (low + high) / 2
		root := &TreeNode{Val: nums[mid], Left: construct(low, mid-1), Right: construct(mid+1, high)}
		return root
	}
	return construct(0, len(nums)-1)
}
