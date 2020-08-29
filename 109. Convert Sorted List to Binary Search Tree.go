package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getMedium(low, high *ListNode) *ListNode {
	fast, slow := low, low
	for fast != high && fast.Next != high {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var construct func(low, high *ListNode) *TreeNode
	construct = func(low, high *ListNode) *TreeNode {
		if low == high {
			return nil
		}
		mid := getMedium(low, high)
		root := &TreeNode{Val: mid.Val, Left: construct(low, mid), Right: construct(mid.Next, high)}
		return root
	}
	return construct(head, nil)
}
