# Go 算法、工具和刷题包

这是一个全面的 Go 语言库，包含了常见的数据结构、算法、实用工具和刷题解决方案。

## 目录结构

```
go-standLib/
├── algorithm/     # 算法实现
│   ├── heap.go      # 堆数据结构及堆排序
│   ├── ring.go      # 环形缓冲区和一致性哈希环
│   ├── hash.go      # 哈希表和 SwissMap 实现
│   └── bloom.go     # 布隆过滤器
├── utils/         # 实用工具函数
│   ├── array.go     # 数组操作工具
│   ├── string.go    # 字符串操作工具
│   └── math.go      # 数学计算工具
├── problem/       # 刷题算法
│   ├── linkedlist.go # 链表相关算法
│   ├── dp.go        # 动态规划算法
│   └── sort_search.go # 排序和搜索算法
├── main.go        # 演示程序
└── go.mod         # Go 模块定义
```

## 功能特性

### 算法包 (algorithm)

#### 1. 堆 (Heap)
- `IntHeap`: 最小堆实现
- `MaxHeap`: 最大堆实现
- `HeapSort`: 使用堆排序算法对数组进行排序

#### 2. 环 (Ring)
- `RingBuffer`: 环形缓冲区实现
  - `Write()`: 向缓冲区写入数据
  - `Read()`: 从缓冲区读取数据
  - `IsEmpty()/IsFull()`: 检查缓冲区状态
- `ConsistentHashRing`: 一致性哈希环实现
  - `GetNode()`: 获取键对应的节点
  - `AddNode()/RemoveNode()`: 添加/移除节点

#### 3. 哈希 (Hash)
- `HashTable`: 简单哈希表实现
  - `Put()/Get()/Delete()`: 基本操作
- `SwissMap`: 类似 Google SwissTable 的哈希表实现
  - 高性能哈希存储

#### 4. 布隆过滤器 (Bloom Filter)
- `NewBloomFilter()`: 创建布隆过滤器
- `Add()`: 添加元素
- `Contains()`: 检查元素是否存在
- 高效的空间利用，允许误报但不误删

#### 5. 图算法 (Graph Algorithms)
- `Graph`: 图的邻接表表示
  - `Dijkstra()`: Dijkstra最短路径算法
- `UnionFind`: 并查集数据结构
  - `Find()/Union()`: 查找和合并操作
  - `Connected()`: 检查连通性
- `TopologicalSort()`: 拓扑排序算法
- `KMP()`: KMP字符串匹配算法

#### 6. 缓存算法 (Cache Algorithms)
- `LRUCache`: LRU缓存实现
  - `Get()/Put()`: 获取和插入操作
- `LFUCache`: LFU缓存实现
  - `Get()/Put()`: 获取和插入操作
- `CountOneBits()`: 计算二进制中1的个数
- `IsPowerOfTwo()`: 检查是否是2的幂
- `NextPowerOfTwo()`: 获取大于等于n的最小2的幂
- `GrayCode()`: 生成格雷码序列
- `SingleNumber()`: 找出数组中只出现一次的数字

### 工具包 (utils)

#### 1. 数组工具
- `ReverseArray()`: 反转数组
- `RemoveDuplicates()`: 移除重复元素
- `Contains()`: 检查元素是否存在
- `Max()/Min()/Sum()`: 基本统计
- `Filter()`: 条件过滤

#### 2. 字符串工具
- `ReverseString()`: 反转字符串
- `IsPalindrome()`: 检查回文
- `CountWords()`: 计算单词数
- `RemoveWhitespace()`: 移除空白字符
- `CapitalizeWords()`: 首字母大写
- `CommonPrefix()`: 最长公共前缀
- `IsAnagram()`: 检查字母异位词

#### 3. 数学工具
- `GCD()/LCM()`: 最大公约数/最小公倍数
- `IsPrime()`: 质数判断
- `Fibonacci()`: 斐波那契数列
- `Factorial()`: 阶乘
- `Power()`: 幂运算
- `Sqrt()`: 平方根（牛顿法）
- `Combination()/Permutation()`: 组合与排列

### 刷题包 (problem)

#### 1. 链表算法
- `ReverseList()`: 反转链表
- `HasCycle()`: 检查链表是否有环
- `MergeTwoLists()`: 合并两个有序链表
- `RemoveNthFromEnd()`: 删除倒数第n个节点
- `MiddleNode()`: 找到中间节点

#### 2. 动态规划
- `FibonacciDP()`: 动态规划斐波那契
- `MaxSubArray()`: 最大子数组和（Kadane算法）
- `ClimbingStairs()`: 爬楼梯问题
- `UniquePaths()`: 不同路径问题
- `CoinChange()`: 零钱兑换问题
- `LongestCommonSubsequence()`: 最长公共子序列
- `Rob()`: 打家劫舍问题

#### 3. 排序和搜索
- `QuickSort()`: 快速排序
- `MergeSort()`: 归并排序
- `BinarySearch()`: 二分搜索
- `BinarySearchFirst()/BinarySearchLast()`: 查找首次/末次出现位置
- `SearchInRotatedSortedArray()`: 在旋转排序数组中搜索
- `FindKthLargest()`: 找到第K大的元素

## 使用方法

### 1. 构建项目
```bash
go mod init go-standLib
go build .
```

### 2. 运行演示
```bash
./go-standLib
```

### 3. 在你的项目中使用
```go
import (
    "go-standLib/algorithm"
    "go-standLib/utils"
    "go-standLib/problem"
)

// 使用示例
bf := algorithm.NewBloomFilter(1000, 0.01)
bf.Add([]byte("example"))
exists := bf.Contains([]byte("example"))
```

## 特点

1. **完整的数据结构实现**: 包含堆、环、哈希表、布隆过滤器等常见数据结构
2. **丰富的工具函数**: 提供数组、字符串、数学等常用操作
3. **刷题算法集合**: 包含链表、动态规划、排序搜索等经典算法
4. **高效的设计**: 包含高性能的 SwissMap 等优化数据结构
5. **清晰的文档**: 每个函数都有明确的注释和说明
6. **易于使用**: API 设计简洁，易于集成到现有项目

## 适用场景

- Go 语言学习和练习
- 算法竞赛准备
- 面试准备
- 需要高效数据结构的应用
- 工具函数库