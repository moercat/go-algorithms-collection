package lc

import "container/list"

type LRUCache struct {
	Link     *list.List
	Search   map[int]*list.Element
	Capacity int
}

type Node struct {
	Key   int
	Value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Link:     list.New(),
		Search:   make(map[int]*list.Element, 2),
		Capacity: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.Search[key]
	if ok {
		this.Link.MoveToBack(elem)
		return elem.Value.(Node).Value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	oldItem, exist := this.Search[key]
	if !exist {
		if len(this.Search) >= this.Capacity {
			elem := this.Link.Front()
			if elem != nil {
				delete(this.Search, elem.Value.(Node).Key)
				this.Link.Remove(elem)
			}
		}
		oldItem = this.Link.PushBack(Node{
			Key:   key,
			Value: value,
		})
	}

	oldItem.Value = Node{
		Key:   key,
		Value: value,
	}
	this.Link.MoveToBack(oldItem)
	this.Search[key] = oldItem
}
