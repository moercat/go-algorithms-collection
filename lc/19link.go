package lc

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := 0, 0
	var leftNode = &ListNode{
		Next: head,
	}
	var newHead = leftNode
	for node := head; node != nil; node = node.Next {
		right++

		if right > n {
			left++
			leftNode = leftNode.Next
		}
	}
	if leftNode.Next == nil {
		return head
	}
	next := leftNode.Next.Next
	leftNode.Next = next

	return newHead.Next
}
