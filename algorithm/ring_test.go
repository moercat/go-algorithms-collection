package algorithm

import (
	"testing"
)

// 测试RingBuffer
func TestRingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	// 测试写入
	data1 := "abc"
	success := rb.Write(data1)
	if !success {
		t.Error("Write failed")
	}

	// 测试读取
	readData, ok := rb.Read()
	if !ok {
		t.Error("Read failed")
	}

	expected := "abc"
	if readData != expected {
		t.Errorf("Read data mismatch: expected %v, got %v", expected, readData)
	}

	// 测试缓冲区状态
	if !rb.IsEmpty() {
		t.Error("Buffer should be empty after read")
	}

	// 再次写入测试
	rb.Write("xyz")
	if rb.IsEmpty() {
		t.Error("Buffer should not be empty after write")
	}

	if rb.IsFull() {
		t.Error("Buffer should not be full with one element")
	}
}

func TestHashTable(t *testing.T) {
	ht := NewHashTable(10)

	// 测试Put和Get
	ht.Put("key1", "value1")
	ht.Put("key2", "value2")

	if value, exists := ht.Get("key1"); !exists || value != "value1" {
		t.Errorf("Get(key1) = %v, %t; expected value1, true", value, exists)
	}

	if value, exists := ht.Get("key2"); !exists || value != "value2" {
		t.Errorf("Get(key2) = %v, %t; expected value2, true", value, exists)
	}

	// 测试更新
	ht.Put("key1", "updated_value")
	if value, exists := ht.Get("key1"); !exists || value != "updated_value" {
		t.Errorf("Get(key1) after update = %v, %t; expected updated_value, true", value, exists)
	}

	// 测试删除
	ht.Delete("key1")
	if value, exists := ht.Get("key1"); exists {
		t.Errorf("Get(key1) after Delete should return false, got %v, %t", value, exists)
	}
}

func TestSwissMap(t *testing.T) {
	sm := NewSwissMap()
	sm.putEmpty() // 初始化空控制字节

	// 测试Put和Get
	sm.Put("key1", "value1")
	sm.Put("key2", "value2")

	if value, exists := sm.Get("key1"); !exists || value != "value1" {
		t.Errorf("Get(key1) = %v, %t; expected value1, true", value, exists)
	}

	if value, exists := sm.Get("key2"); !exists || value != "value2" {
		t.Errorf("Get(key2) = %v, %t; expected value2, true", value, exists)
	}

	// 测试更新
	sm.Put("key1", "updated_value")
	if value, exists := sm.Get("key1"); !exists || value != "updated_value" {
		t.Errorf("Get(key1) after update = %v, %t; expected updated_value, true", value, exists)
	}
}