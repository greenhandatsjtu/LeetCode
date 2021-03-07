package main

func deleteDuplicates(head *ListNode) *ListNode {
	ptr := head
	for ptr != nil && ptr.Next != nil {
		for ptr.Next != nil && ptr.Next.Val == ptr.Val {
			ptr.Next = ptr.Next.Next
		}
		ptr = ptr.Next
	}
	return head
}
