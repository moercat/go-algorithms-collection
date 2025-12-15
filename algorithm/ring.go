package algorithm

// RingBuffer 是一个环形缓冲区的实现
type RingBuffer struct {
	data     []interface{}
	size     int
	readPos  int
	writePos int
	count    int
}

// NewRingBuffer 创建一个新的环形缓冲区
func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		data: make([]interface{}, size),
		size: size,
	}
}

// Write 向环形缓冲区写入数据
func (rb *RingBuffer) Write(item interface{}) bool {
	if rb.IsFull() {
		return false
	}
	
	rb.data[rb.writePos] = item
	rb.writePos = (rb.writePos + 1) % rb.size
	rb.count++
	return true
}

// Read 从环形缓冲区读取数据
func (rb *RingBuffer) Read() (interface{}, bool) {
	if rb.IsEmpty() {
		return nil, false
	}
	
	item := rb.data[rb.readPos]
	rb.data[rb.readPos] = nil // 清空引用
	rb.readPos = (rb.readPos + 1) % rb.size
	rb.count--
	return item, true
}

// IsEmpty 检查环形缓冲区是否为空
func (rb *RingBuffer) IsEmpty() bool {
	return rb.count == 0
}

// IsFull 检查环形缓冲区是否已满
func (rb *RingBuffer) IsFull() bool {
	return rb.count == rb.size
}

// Size 返回环形缓冲区的大小
func (rb *RingBuffer) Size() int {
	return rb.size
}

// Count 返回当前元素数量
func (rb *RingBuffer) Count() int {
	return rb.count
}

// ConsistentHashRing 是一致性哈希环的实现
type ConsistentHashRing struct {
	ring map[string]string // 虚拟节点 -> 真实节点
	nodes []string         // 真实节点列表
	hash func(string) uint32
}

// NewConsistentHashRing 创建一个新的一致性哈希环
func NewConsistentHashRing(nodes []string) *ConsistentHashRing {
	chr := &ConsistentHashRing{
		ring: make(map[string]string),
		hash: hashFunc,
	}
	
	chr.nodes = make([]string, len(nodes))
	copy(chr.nodes, nodes)
	
	// 为每个节点添加虚拟节点
	for _, node := range nodes {
		for i := 0; i < 10; i++ {
			virtualNode := node + "#" + itoa(i)
			key := virtualNode
			chr.ring[key] = node
		}
	}
	
	return chr
}

// AddNode 添加节点到一致性哈希环
func (chr *ConsistentHashRing) AddNode(node string) {
	for i := 0; i < 10; i++ {
		virtualNode := node + "#" + itoa(i)
		key := virtualNode
		chr.ring[key] = node
	}
	chr.nodes = append(chr.nodes, node)
}

// RemoveNode 从一致性哈希环移除节点
func (chr *ConsistentHashRing) RemoveNode(node string) {
	for i := 0; i < len(chr.nodes); i++ {
		if chr.nodes[i] == node {
			chr.nodes = append(chr.nodes[:i], chr.nodes[i+1:]...)
			i-- // 调整索引
		}
	}
	
	// 移除虚拟节点
	for key, val := range chr.ring {
		if val == node {
			delete(chr.ring, key)
		}
	}
}

// GetNode 获取给定键对应的节点
func (chr *ConsistentHashRing) GetNode(key string) string {
	if len(chr.ring) == 0 {
		return ""
	}
	
	// 简化实现：直接返回哈希值对应的节点
	hashValue := chr.hash(key)
	nodes := make([]string, 0, len(chr.ring))
	for node := range chr.ring {
		nodes = append(nodes, node)
	}
	
	if len(nodes) == 0 {
		return ""
	}
	
	// 简化：使用模运算选择节点
	selectedIndex := int(hashValue) % len(nodes)
	return chr.ring[nodes[selectedIndex]]
}

// 简单的哈希函数模拟
func hashFunc(s string) uint32 {
	var h uint32
	for i := 0; i < len(s); i++ {
		h = h*31 + uint32(s[i])
	}
	return h
}

func itoa(i int) string {
	if i < 10 {
		return string(rune('0' + i))
	}
	return string(rune('0' + i/10)) + string(rune('0' + i%10))
}