package lc

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1}

	fmt.Println(removeNthFromEnd(head, 1))
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
}

func Test_removeNthFromEnd2(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println(removeNthFromEnd(head, 2))
	for node := head; node != nil; node = node.Next {
		fmt.Printf("%d ", node.Val)
	}
}
