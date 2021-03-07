package main

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dumb := &ListNode{Next: head, Val: -101}
	ptr := dumb
	for ptr.Next != nil && ptr.Next.Next != nil {
		if ptr.Next.Val == ptr.Next.Next.Val {
			val := ptr.Next.Val
			for ptr.Next != nil && ptr.Next.Val == val {
				ptr.Next = ptr.Next.Next
			}
		} else {
			ptr = ptr.Next
		}
	}
	return dumb.Next
}
