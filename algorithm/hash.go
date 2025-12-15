package algorithm

// HashTable 是一个简单的哈希表实现
type HashTable struct {
	buckets []*Bucket
	size    int
}

// Bucket 是哈希桶，用于处理冲突
type Bucket struct {
	pairs []*KVPair
}

// KVPair 是键值对
type KVPair struct {
	key   string
	value interface{}
}

// NewHashTable 创建一个新的哈希表
func NewHashTable(size int) *HashTable {
	ht := &HashTable{
		buckets: make([]*Bucket, size),
		size:    size,
	}
	
	for i := 0; i < size; i++ {
		ht.buckets[i] = &Bucket{
			pairs: make([]*KVPair, 0),
		}
	}
	
	return ht
}

// hash 简单的哈希函数
func (ht *HashTable) hash(key string) int {
	hash := 0
	for i := 0; i < len(key); i++ {
		hash = (hash + int(key[i])) % ht.size
	}
	return hash
}

// Put 向哈希表中插入键值对
func (ht *HashTable) Put(key string, value interface{}) {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	
	// 检查是否已存在该键
	for _, pair := range bucket.pairs {
		if pair.key == key {
			pair.value = value
			return
		}
	}
	
	// 添加新的键值对
	bucket.pairs = append(bucket.pairs, &KVPair{key: key, value: value})
}

// Get 从哈希表中获取值
func (ht *HashTable) Get(key string) (interface{}, bool) {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	
	for _, pair := range bucket.pairs {
		if pair.key == key {
			return pair.value, true
		}
	}
	
	return nil, false
}

// Delete 从哈希表中删除键值对
func (ht *HashTable) Delete(key string) bool {
	index := ht.hash(key)
	bucket := ht.buckets[index]
	
	for i, pair := range bucket.pairs {
		if pair.key == key {
			// 从切片中删除元素
			bucket.pairs = append(bucket.pairs[:i], bucket.pairs[i+1:]...)
			return true
		}
	}
	
	return false
}

// SwissMap 是一个类似 Google SwissTable 的哈希表实现
type SwissMap struct {
	buckets []swissBucket
	size    int
	count   int
}

const (
	swissBucketSize = 16 // 每个桶的槽位数
	maxLoadFactor   = 0.8
)

type swissBucket struct {
	control [swissBucketSize]byte
	pairs   [swissBucketSize]*KVPair
}

const (
	ctrlEmpty     = byte(0b11111111) // 空槽位
	ctrlDeleted   = byte(0b11111110) // 已删除槽位
	ctrlSentinel  = byte(0b00000000) // 哨兵槽位
)

// NewSwissMap 创建一个新的 SwissMap
func NewSwissMap() *SwissMap {
	return &SwissMap{
		buckets: make([]swissBucket, 1),
		size:    1,
	}
}

// putEmpty 初始化控制字节
func (s *SwissMap) putEmpty() {
	for i := 0; i < swissBucketSize; i++ {
		s.buckets[0].control[i] = ctrlEmpty
	}
}

// StringHash 简单的字符串哈希函数
func StringHash(s string) uint64 {
	var h uint64 = 5381
	for i := 0; i < len(s); i++ {
		h = ((h << 5) + h) + uint64(s[i])
	}
	return h
}

// Get 在 SwissMap 中查找值
func (s *SwissMap) Get(key string) (interface{}, bool) {
	hashValue := StringHash(key)
	bucketIndex := hashValue % uint64(s.size)
	
	// 简化的查找实现
	for i := 0; i < swissBucketSize; i++ {
		if s.buckets[bucketIndex].pairs[i] != nil && 
		   s.buckets[bucketIndex].pairs[i].key == key {
			return s.buckets[bucketIndex].pairs[i].value, true
		}
	}
	
	return nil, false
}

// Put 在 SwissMap 中插入键值对
func (s *SwissMap) Put(key string, value interface{}) {
	s.put(key, value)
}

func (s *SwissMap) put(key string, value interface{}) {
	if float64(s.count)/float64(s.size*swissBucketSize) > maxLoadFactor {
		s.resize()
	}
	
	hashValue := StringHash(key)
	bucketIndex := hashValue % uint64(s.size)
	
	// 简化的插入实现
	for i := 0; i < swissBucketSize; i++ {
		if s.buckets[bucketIndex].pairs[i] == nil {
			s.buckets[bucketIndex].pairs[i] = &KVPair{key: key, value: value}
			s.count++
			return
		} else if s.buckets[bucketIndex].pairs[i].key == key {
			s.buckets[bucketIndex].pairs[i].value = value
			return
		}
	}
}

// resize 重新分配空间
func (s *SwissMap) resize() {
	oldBuckets := s.buckets
	
	s.size *= 2
	s.buckets = make([]swissBucket, s.size)
	
	// 重新初始化控制字节
	for i := 0; i < s.size; i++ {
		for j := 0; j < swissBucketSize; j++ {
			s.buckets[i].control[j] = ctrlEmpty
		}
	}
	
	// 重新插入所有元素
	for _, bucket := range oldBuckets {
		for _, pair := range bucket.pairs {
			if pair != nil {
				s.Put(pair.key, pair.value)
			}
		}
	}
}