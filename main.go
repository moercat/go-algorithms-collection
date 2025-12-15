package main

import (
	"container/heap"
	"fmt"
	"go-algorithms-collection/algorithm"
	"go-algorithms-collection/problem"
	"go-algorithms-collection/utils"
)

func main() {
	fmt.Println("=== Go 算法、工具和刷题包演示 ===")

	demonstrateHeap()
	demonstrateBloomFilter()
	demonstrateCache()
	demonstrateLinkedList()
	demonstrateDynamicProgramming()
	demonstrateUtils()

	fmt.Println("演示完成！")
}

func demonstrateHeap() {
	fmt.Println("--- 堆 (Heap) 演示 ---")
	h := &algorithm.IntHeap{}
	heap.Init(h)
	values := []int{3, 1, 4, 1, 5, 9, 2, 6}

	for _, v := range values {
		heap.Push(h, v)
	}

	fmt.Printf("压入值 %v 后，按顺序弹出:\n", values)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}

func demonstrateBloomFilter() {
	fmt.Println("--- 布隆过滤器 (Bloom Filter) 演示 ---")
	bf := algorithm.NewBloomFilter(1000, 0.01)

	elements := []string{"apple", "banana", "cherry", "date"}

	for _, element := range elements {
		bf.Add([]byte(element))
	}

	testElements := []string{"apple", "banana", "grape", "cherry", "fig"}
	for _, element := range testElements {
		exists := bf.Contains([]byte(element))
		fmt.Printf("元素 '%s' 是否可能存在: %t\n", element, exists)
	}
	fmt.Println()
}

func demonstrateCache() {
	fmt.Println("--- LRU缓存 (Cache) 演示 ---")
	cache := algorithm.NewLRUCache(3)

	cache.Put(1, 100)
	cache.Put(2, 200)
	cache.Put(3, 300)
	fmt.Printf("添加 1=100, 2=200, 3=300\n")

	cache.Get(1)      // 访问1，提升优先级
	cache.Put(4, 400) // 添加新项，应该淘汰2
	fmt.Printf("访问1后添加 4=400，此时访问2得到: %d (应该为-1表示不存在)\n", cache.Get(2))

	cache.Put(1, 101) // 更新1
	result := cache.Get(1)
	fmt.Printf("更新1后获取其值: %d\n", result)
	fmt.Println()
}

func demonstrateLinkedList() {
	fmt.Println("--- 链表 (Linked List) 演示 ---")
	head := problem.NewListNode(1)
	head.Next = problem.NewListNode(2)
	head.Next.Next = problem.NewListNode(3)
	head.Next.Next.Next = problem.NewListNode(4)
	head.Next.Next.Next.Next = problem.NewListNode(5)

	fmt.Print("原始链表: ")
	printList(head)

	reversed := problem.ReverseList(head)
	fmt.Print("反转后链表: ")
	printList(reversed)
	fmt.Println()
}

func demonstrateDynamicProgramming() {
	fmt.Println("--- 动态规划 (Dynamic Programming) 演示 ---")

	// 最大子数组和 (Kadane算法)
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum := problem.MaxSubArray(nums)
	fmt.Printf("数组 %v 的最大子数组和: %d\n", nums, maxSum)

	// 斐波那契数列
	fibResult := problem.FibonacciDP(10)
	fmt.Printf("第10个斐波那契数: %d\n", fibResult)

	// 爬楼梯问题
	stairsResult := problem.ClimbingStairs(5)
	fmt.Printf("爬5阶楼梯的方法数: %d\n", stairsResult)
	fmt.Println()
}

func demonstrateUtils() {
	fmt.Println("--- 工具函数 (Utils) 演示 ---")

	// 数组工具
	arr := []int{1, 2, 2, 3, 4, 4, 5}
	fmt.Printf("原数组: %v\n", arr)
	uniqueArr := utils.RemoveDuplicates(arr)
	fmt.Printf("去重后: %v\n", uniqueArr)

	// 字符串工具
	str := "A man a plan a canal Panama"
	fmt.Printf("原字符串: %s\n", str)
	isPalin := utils.IsPalindrome(str)
	fmt.Printf("是否回文: %t\n", isPalin)

	// 数学工具
	a, b := 48, 18
	gcd := utils.GCD(a, b)
	lcm := utils.LCM(a, b)
	fmt.Printf("GCD(%d, %d) = %d\n", a, b, gcd)
	fmt.Printf("LCM(%d, %d) = %d\n", a, b, lcm)
	fmt.Println()
}

// 辅助函数打印链表
func printList(head *problem.ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ", current.Val)
		current = current.Next
	}
	fmt.Println()
}
