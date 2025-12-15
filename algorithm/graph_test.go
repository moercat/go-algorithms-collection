package algorithm

import (
	"testing"
)

// 图算法测试
func TestGraph(t *testing.T) {
	// 创建一个简单的图进行测试
	graph := NewGraph(5)

	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 1)
	graph.AddEdge(1, 3, 5)
	graph.AddEdge(2, 3, 8)
	graph.AddEdge(2, 4, 10)
	graph.AddEdge(3, 4, 2)

	// 测试Dijkstra算法
	distances := graph.Dijkstra(0)
	expectedDistances := []int{0, 4, 2, 9, 11}

	for i, expectedDist := range expectedDistances {
		if i < len(distances) && distances[i] != expectedDist {
			t.Errorf("Dijkstra from 0: distance to %d = %d; expected %d", i, distances[i], expectedDist)
		}
	}
}

func TestUnionFind(t *testing.T) {
	uf := NewUnionFind(5)
	
	// 初始状态下，每个元素都在自己的集合中
	if uf.Connected(0, 1) {
		t.Error("Elements should not be connected initially")
	}
	
	// 连接一些元素
	uf.Union(0, 1)
	uf.Union(2, 3)
	
	if !uf.Connected(0, 1) {
		t.Error("Elements 0 and 1 should be connected after Union(0, 1)")
	}
	
	if uf.Connected(0, 2) {
		t.Error("Elements 0 and 2 should not be connected")
	}
	
	// 连接更多的元素
	uf.Union(1, 3)
	if !uf.Connected(0, 2) {
		t.Error("Elements 0 and 2 should be connected after Union(1, 3)")
	}
}