package algorithm

import "container/list"

// LRUCache LRU缓存实现
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element // key -> list element
	list     *list.List           // 双向链表，存储(key, value)
}

// Pair 键值对
type Pair struct {
	key, value int
}

// NewLRUCache 创建新的LRU缓存
func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

// Get 获取值
func (lru *LRUCache) Get(key int) int {
	if elem, ok := lru.cache[key]; ok {
		// 将访问的节点移到链表头部
		lru.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return -1 // 未找到
}

// Put 插入或更新值
func (lru *LRUCache) Put(key, value int) {
	if elem, ok := lru.cache[key]; ok {
		// 更新现有节点
		elem.Value.(*Pair).value = value
		lru.list.MoveToFront(elem)
		return
	}
	
	// 检查容量
	if lru.list.Len() >= lru.capacity {
		// 删除链表尾部节点
		back := lru.list.Back()
		if back != nil {
			pair := back.Value.(*Pair)
			delete(lru.cache, pair.key)
			lru.list.Remove(back)
		}
	}
	
	// 添加新节点到链表头部
	front := lru.list.PushFront(&Pair{key: key, value: value})
	lru.cache[key] = front
}

// LFUCache LFU缓存实现
type LFUCache struct {
	capacity int
	minFreq  int
	keyToVal map[int]int
	keyToFreq map[int]int
	freqToKeys map[int]*list.List // 频次 -> 对应频次的键的双向链表
	keyToListElement map[int]*list.Element // 键 -> 在对应频次链表中的元素
}

// NewLFUCache 创建新的LFU缓存
func NewLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		keyToVal: make(map[int]int),
		keyToFreq: make(map[int]int),
		freqToKeys: make(map[int]*list.List),
		keyToListElement: make(map[int]*list.Element),
	}
}

// Get 获取值
func (lfu *LFUCache) Get(key int) int {
	if _, ok := lfu.keyToVal[key]; !ok {
		return -1
	}
	
	lfu.increaseFreq(key)
	return lfu.keyToVal[key]
}

// Put 插入或更新值
func (lfu *LFUCache) Put(key, value int) {
	if lfu.capacity <= 0 {
		return
	}
	
	if _, ok := lfu.keyToVal[key]; !ok {
		if len(lfu.keyToVal) >= lfu.capacity {
			lfu.removeMinFreqKey()
		}
		lfu.keyToVal[key] = value
		lfu.keyToFreq[key] = 1
		if _, ok := lfu.freqToKeys[1]; !ok {
			lfu.freqToKeys[1] = list.New()
		}
		elem := lfu.freqToKeys[1].PushBack(key)
		lfu.keyToListElement[key] = elem
		lfu.minFreq = 1
		return
	}
	
	lfu.keyToVal[key] = value
	lfu.increaseFreq(key)
}

// increaseFreq 增加键的使用频次
func (lfu *LFUCache) increaseFreq(key int) {
	freq := lfu.keyToFreq[key]
	lfu.keyToFreq[key]++

	// 从旧频次链表中删除
	oldList := lfu.freqToKeys[freq]
	elem := lfu.keyToListElement[key]
	oldList.Remove(elem)
	delete(lfu.keyToListElement, key)

	// 如果旧频次链表为空，删除该频次
	if oldList.Len() == 0 {
		delete(lfu.freqToKeys, freq)
		if freq == lfu.minFreq {
			lfu.minFreq++
		}
	}

	// 添加到新频次链表
	newFreq := freq + 1
	if _, ok := lfu.freqToKeys[newFreq]; !ok {
		lfu.freqToKeys[newFreq] = list.New()
	}
	newElem := lfu.freqToKeys[newFreq].PushBack(key)
	lfu.keyToListElement[key] = newElem
}

// removeMinFreqKey 删除最小频次的键
func (lfu *LFUCache) removeMinFreqKey() {
	keyList := lfu.freqToKeys[lfu.minFreq]
	back := keyList.Back()
	if back == nil {
		return
	}
	
	key := back.Value.(int)
	
	// 从各映射中删除
	delete(lfu.keyToVal, key)
	delete(lfu.keyToFreq, key)
	delete(lfu.keyToListElement, key)
	keyList.Remove(back)
	
	if keyList.Len() == 0 {
		delete(lfu.freqToKeys, lfu.minFreq)
	}
}

// BitManipulation 位操作算法集合

// CountOneBits 计算二进制中1的个数
func CountOneBits(n uint) int {
	count := 0
	for n != 0 {
		n &= n - 1 // 清除最右边的1
		count++
	}
	return count
}

// ReverseBits 反转二进制位
func ReverseBits(n uint32) uint32 {
	result := uint32(0)
	
	for i := 0; i < 32; i++ {
		result <<= 1
		result |= n & 1
		n >>= 1
	}
	
	return result
}

// IsPowerOfTwo 检查是否是2的幂
func IsPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// NextPowerOfTwo 获取大于等于n的最小2的幂
func NextPowerOfTwo(n int) int {
	if n <= 1 {
		return 1
	}
	
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++
	
	return n
}

// GrayCode 生成n位格雷码序列
func GrayCode(n int) []int {
	size := 1 << n // 2^n
	result := make([]int, size)
	
	for i := 0; i < size; i++ {
		// 格雷码公式: i ^ (i >> 1)
		result[i] = i ^ (i >> 1)
	}
	
	return result
}

// SwapWithoutTemp 不使用临时变量交换两个数
func SwapWithoutTemp(a, b *int) {
	*a ^= *b
	*b ^= *a
	*a ^= *b
}

// SingleNumber 找出数组中只出现一次的数字（其他数字都出现两次）
func SingleNumber(nums []int) int {
	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}