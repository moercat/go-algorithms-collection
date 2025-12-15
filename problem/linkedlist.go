package problem

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewListNode 创建新的链表节点
func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// ReverseList 反转链表
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	
	return prev
}

// HasCycle 检查链表是否有环
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	
	slow := head
	fast := head
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		
		if slow == fast {
			return true
		}
	}
	
	return false
}

// MergeTwoLists 合并两个有序链表
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	
	return dummy.Next
}

// RemoveNthFromEnd 删除链表的倒数第n个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast := dummy
	slow := dummy
	
	// 让fast指针先走n+1步
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	
	// 同时移动fast和slow指针，直到fast到达链表末尾
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	
	// 删除第n个节点
	slow.Next = slow.Next.Next
	
	return dummy.Next
}

// MiddleNode 找到链表的中间节点
func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return slow
}