package algorithm

import (
	"hash/fnv"
	"math"
)

// BloomFilter 是布隆过滤器的实现
type BloomFilter struct {
	bitSet    []uint64  // 位数组
	bitCount  uint      // 位总数
	hashCount int       // 哈希函数数量
}

// NewBloomFilter 创建一个新的布隆过滤器
// capacity: 预期元素数量
// errorRate: 误报率
func NewBloomFilter(capacity uint, errorRate float64) *BloomFilter {
	// 计算最优的位数组大小
	bitCount := uint(math.Ceil(float64(capacity) * math.Log(errorRate) / (math.Log(2) * math.Log(2))))
	// 计算最优的哈希函数数量
	hashCount := int(math.Ceil(math.Ln2 * float64(bitCount) / float64(capacity)))
	
	// 确保位数组大小是64的倍数以方便使用uint64
	wordCount := int((bitCount + 63) / 64)
	
	return &BloomFilter{
		bitSet:    make([]uint64, wordCount),
		bitCount:  uint(wordCount * 64),
		hashCount: hashCount,
	}
}

// hash1 第一个哈希函数
func (bf *BloomFilter) hash1(data []byte) uint {
	h := fnv.New64a()
	h.Write(data)
	return uint(h.Sum64()) % bf.bitCount
}

// hash2 第二个哈希函数
func (bf *BloomFilter) hash2(data []byte) uint {
	h := fnv.New64a()
	h.Write(data)
	sum := h.Sum64()
	return uint(sum>>32) ^ uint(sum) % uint(bf.bitCount)
}

// Add 向布隆过滤器中添加元素
func (bf *BloomFilter) Add(data []byte) {
	for i := 0; i < bf.hashCount; i++ {
		// 使用双重哈希来生成多个不同的哈希值
		h1 := bf.hash1(data)
		h2 := bf.hash2(data)
		index := (h1 + uint(i)*h2) % bf.bitCount
		
		// 设置对应位为1
		wordIndex := index / 64
		bitIndex := index % 64
		bf.bitSet[wordIndex] |= 1 << bitIndex
	}
}

// Contains 检查元素是否可能存在于集合中
func (bf *BloomFilter) Contains(data []byte) bool {
	for i := 0; i < bf.hashCount; i++ {
		h1 := bf.hash1(data)
		h2 := bf.hash2(data)
		index := (h1 + uint(i)*h2) % bf.bitCount
		
		// 检查对应位是否为1
		wordIndex := index / 64
		bitIndex := index % 64
		if (bf.bitSet[wordIndex] & (1 << bitIndex)) == 0 {
			return false
		}
	}
	
	return true
}

// Reset 重置布隆过滤器
func (bf *BloomFilter) Reset() {
	for i := range bf.bitSet {
		bf.bitSet[i] = 0
	}
}

// BitSetSize 返回位数组大小
func (bf *BloomFilter) BitSetSize() uint {
	return bf.bitCount
}

// HashCount 返回哈希函数数量
func (bf *BloomFilter) HashCount() int {
	return bf.hashCount
}