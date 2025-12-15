package algorithm

import (
	"container/heap"
	"math"
)

// Graph 表示图的邻接表表示
type Graph struct {
	vertices int
	adj      [][]Edge
}

// Edge 表示图的边
type Edge struct {
	To   int
	Weight int
}

// NewGraph 创建一个新的图
func NewGraph(vertices int) *Graph {
	return &Graph{
		vertices: vertices,
		adj:      make([][]Edge, vertices),
	}
}

// AddEdge 添加边
func (g *Graph) AddEdge(from, to, weight int) {
	g.adj[from] = append(g.adj[from], Edge{To: to, Weight: weight})
}

// Dijkstra 使用Dijkstra算法找到单源最短路径
func (g *Graph) Dijkstra(start int) []int {
	dist := make([]int, g.vertices)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[start] = 0

	pq := &MinHeap{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		u := item.value
		
		if item.priority > dist[u] {
			continue
		}
		
		for _, edge := range g.adj[u] {
			v := edge.To
			weight := edge.Weight
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, &Item{value: v, priority: dist[v]})
			}
		}
	}
	
	return dist
}

// Item 是优先队列中的元素
type Item struct {
	value    int
	priority int
	index    int
}

// MinHeap 是最小堆
type MinHeap []*Item

func (pq MinHeap) Len() int { return len(pq) }

func (pq MinHeap) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *MinHeap) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// UnionFind 并查集实现
type UnionFind struct {
	parent []int
	rank   []int
	count  int
}

// NewUnionFind 创建新的并查集
func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	rank := make([]int, n)
	
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	
	return &UnionFind{
		parent: parent,
		rank:   rank,
		count:  n,
	}
}

// Find 查找根节点（路径压缩）
func (uf *UnionFind) Find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x]) // 路径压缩
	}
	return uf.parent[x]
}

// Union 合并两个集合（按秩合并）
func (uf *UnionFind) Union(x, y int) bool {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	
	if rootX == rootY {
		return false // 已经在同一个集合中
	}
	
	// 按秩合并
	if uf.rank[rootX] < uf.rank[rootY] {
		uf.parent[rootX] = rootY
	} else if uf.rank[rootX] > uf.rank[rootY] {
		uf.parent[rootY] = rootX
	} else {
		uf.parent[rootY] = rootX
		uf.rank[rootX]++
	}
	
	uf.count--
	return true
}

// Connected 检查两个节点是否连通
func (uf *UnionFind) Connected(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

// Count 返回连通分量的数量
func (uf *UnionFind) Count() int {
	return uf.count
}

// TopologicalSort 拓扑排序
func TopologicalSort(vertices int, edges [][]int) []int {
	// 构建邻接表和入度数组
	adj := make([][]int, vertices)
	inDegree := make([]int, vertices)
	
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		adj[from] = append(adj[from], to)
		inDegree[to]++
	}
	
	// 找到所有入度为0的节点
	queue := []int{}
	for i := 0; i < vertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	
	result := []int{}
	
	// BFS处理
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)
		
		for _, neighbor := range adj[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	
	// 检查是否有环
	if len(result) != vertices {
		return []int{} // 存在环，无法拓扑排序
	}
	
	return result
}

// KMP 算法实现
func KMP(text, pattern string) []int {
	if len(pattern) == 0 {
		return []int{0}
	}

	// 构建部分匹配表（failure function）
	lps := make([]int, len(pattern))

	length := 0
	i := 1

	for i < len(pattern) {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	// 搜索模式
	result := []int{}
	ti := 0 // text索引
	pj := 0 // pattern索引

	for ti < len(text) {
		if pattern[pj] == text[ti] {
			pj++
			ti++
		}

		if pj == len(pattern) {
			result = append(result, ti-pj)
			pj = lps[pj-1]
		} else if ti < len(text) && pattern[pj] != text[ti] {
			if pj != 0 {
				pj = lps[pj-1]
			} else {
				ti++
			}
		}
	}

	return result
}

// QuickSelect 快速选择算法，找到第k小的元素
func QuickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}
	
	pivotIndex := partition(nums, left, right)
	
	if k == pivotIndex {
		return nums[k]
	} else if k < pivotIndex {
		return QuickSelect(nums, left, pivotIndex-1, k)
	} else {
		return QuickSelect(nums, pivotIndex+1, right, k)
	}
}

// partition 是快速选择的分区函数
func partition(nums []int, left, right int) int {
	pivot := nums[right]
	i := left
	
	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

// GetKthSmallest 获取第k小的元素
func GetKthSmallest(nums []int, k int) int {
	if k < 1 || k > len(nums) {
		return 0
	}
	
	// 复制数组以避免修改原数组
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	
	return QuickSelect(numsCopy, 0, len(numsCopy)-1, k-1)
}