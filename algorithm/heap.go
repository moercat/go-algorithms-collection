package algorithm

import "container/heap"

// IntHeap 是一个整数堆的实现
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MaxHeap 是一个最大堆的实现
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// HeapSort 使用堆排序算法对数组进行排序
func HeapSort(arr []int) []int {
	h := &IntHeap{}
	heap.Init(h)
	
	for _, v := range arr {
		heap.Push(h, v)
	}
	
	result := make([]int, 0, len(arr))
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(int))
	}
	
	return result
}