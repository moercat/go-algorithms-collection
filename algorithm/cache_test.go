package algorithm

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := NewLRUCache(2)

	// 测试Put和Get
	cache.Put(1, 100)
	cache.Put(2, 200)

	if cache.Get(1) != 100 {
		t.Error("Expected value 100 for key 1")
	}

	if cache.Get(2) != 200 {
		t.Error("Expected value 200 for key 2")
	}

	// 访问1，使其变为最近使用的
	cache.Get(1)

	// 添加第三个元素，因为1是最近使用的，所以2应该是最久未使用的，会被淘汰
	cache.Put(3, 300)

	if cache.Get(2) != -1 {
		t.Error("Expected key 2 to be evicted after putting key 3")
	}

	if cache.Get(1) != 100 {
		t.Error("Expected value 100 for key 1 after eviction of key 2")
	}

	if cache.Get(3) != 300 {
		t.Error("Expected value 300 for key 3")
	}
}

func TestLRUCacheUpdate(t *testing.T) {
	cache := NewLRUCache(2)
	
	cache.Put(1, 100)
	cache.Put(1, 101) // 更新同一键
	
	if cache.Get(1) != 101 {
		t.Error("Expected updated value 101 for key 1")
	}
}

func TestLFUCache(t *testing.T) {
	lfu := NewLFUCache(2)

	lfu.Put(1, 100)
	lfu.Put(2, 200)

	// 访问1使其频率增加
	lfu.Get(1)
	lfu.Get(1)

	// 添加新值，应该淘汰2（频率1）而不是1（频率2）
	lfu.Put(3, 300)

	if lfu.Get(2) != -1 {
		t.Error("Expected key 2 to be evicted")
	}

	if lfu.Get(1) != 100 {
		t.Error("Expected value 100 for key 1")
	}

	if lfu.Get(3) != 300 {
		t.Error("Expected value 300 for key 3")
	}
}