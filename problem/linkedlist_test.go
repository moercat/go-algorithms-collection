package problem

import (
	"testing"
)

func TestReverseList(t *testing.T) {
	// 创建测试链表: 1 -> 2 -> 3 -> nil
	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)

	// 反转链表
	reversed := ReverseList(head)

	// 验证反转后的链表: 3 -> 2 -> 1 -> nil
	current := reversed
	expectedValues := []int{3, 2, 1}
	
	for i, expected := range expectedValues {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			return
		}
		if current.Val != expected {
			t.Errorf("At index %d: expected %d, got %d", i, expected, current.Val)
		}
		current = current.Next
	}
	
	if current != nil {
		t.Error("List is longer than expected")
	}
}

func TestHasCycle(t *testing.T) {
	// 测试无环链表
	head1 := NewListNode(1)
	head1.Next = NewListNode(2)
	head1.Next.Next = NewListNode(3)
	
	if HasCycle(head1) {
		t.Error("Expected no cycle in linear list")
	}
	
	// 测试有环链表
	head2 := NewListNode(1)
	head2.Next = NewListNode(2)
	head2.Next.Next = NewListNode(3)
	head2.Next.Next.Next = head2.Next // 创建一个环: 3 -> 2
	
	if !HasCycle(head2) {
		t.Error("Expected cycle in circular list")
	}
}

func TestMergeTwoLists(t *testing.T) {
	// 创建第一个有序链表: 1 -> 2 -> 4
	l1 := NewListNode(1)
	l1.Next = NewListNode(2)
	l1.Next.Next = NewListNode(4)

	// 创建第二个有序链表: 1 -> 3 -> 4
	l2 := NewListNode(1)
	l2.Next = NewListNode(3)
	l2.Next.Next = NewListNode(4)

	// 合并两个链表
	merged := MergeTwoLists(l1, l2)

	// 验证合并后链表: 1 -> 1 -> 2 -> 3 -> 4 -> 4
	current := merged
	expectedValues := []int{1, 1, 2, 3, 4, 4}
	
	for i, expected := range expectedValues {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			return
		}
		if current.Val != expected {
			t.Errorf("At index %d: expected %d, got %d", i, expected, current.Val)
		}
		current = current.Next
	}
	
	if current != nil {
		t.Error("List is longer than expected")
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	// 创建链表: 1 -> 2 -> 3 -> 4 -> 5
	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)
	head.Next.Next.Next.Next = NewListNode(5)

	// 删除倒数第2个节点 (4)
	newHead := RemoveNthFromEnd(head, 2)

	// 验证结果链表: 1 -> 2 -> 3 -> 5
	current := newHead
	expectedValues := []int{1, 2, 3, 5}
	
	for i, expected := range expectedValues {
		if current == nil {
			t.Errorf("List is shorter than expected at index %d", i)
			return
		}
		if current.Val != expected {
			t.Errorf("At index %d: expected %d, got %d", i, expected, current.Val)
		}
		current = current.Next
	}
	
	if current != nil {
		t.Error("List is longer than expected")
	}
}

func TestMiddleNode(t *testing.T) {
	// 创建链表: 1 -> 2 -> 3 -> 4 -> 5
	head := NewListNode(1)
	head.Next = NewListNode(2)
	head.Next.Next = NewListNode(3)
	head.Next.Next.Next = NewListNode(4)
	head.Next.Next.Next.Next = NewListNode(5)

	// 获取中间节点
	middle := MiddleNode(head)

	// 验证中间节点是3
	if middle.Val != 3 {
		t.Errorf("Expected middle node value to be 3, got %d", middle.Val)
	}
}