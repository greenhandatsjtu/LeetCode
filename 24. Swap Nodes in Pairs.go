package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next.Next
	next := head.Next
	next.Next = head
	head.Next = swapPairs(tmp)
	return next
}
