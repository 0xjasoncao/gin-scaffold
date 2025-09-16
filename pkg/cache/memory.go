package cache

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Memory 内存缓存实现
type Memory struct {
	data  map[string]cacheItem
	mutex sync.RWMutex
}

// cacheItem 内存缓存项
type cacheItem struct {
	value      interface{}
	expiration time.Time
}

// NewMemory 创建新的内存缓存实例
func NewMemory() (*Memory, func()) {
	cache := &Memory{
		data: make(map[string]cacheItem),
	}

	// 启动定期清理过期项的 goroutine
	go cache.cleanupExpired()

	return cache, func() {
		cache.Close()
	}
}

// Set 存储键值对
func (m *Memory) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	var exp time.Time
	if expiration > 0 {
		exp = time.Now().Add(expiration)
	}

	m.data[key] = cacheItem{
		value:      value,
		expiration: exp,
	}

	return nil
}

// Get 获取键对应的值
func (m *Memory) Get(ctx context.Context, key string) (string, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return "", errors.New("key not found")
	}

	// 检查是否过期
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return "", errors.New("key expired")
	}

	// 如果value是字符串指针，直接赋值
	if val, ok := item.value.(string); ok {
		return val, nil
	}
	return "", errors.New("invalid value type")

}

// Delete 删除指定键
func (m *Memory) Delete(ctx context.Context, key string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.data, key)
	return nil
}

// Exists 判断键是否存在
func (m *Memory) Exists(ctx context.Context, key string) (bool, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	item, exists := m.data[key]
	if !exists {
		return false, nil
	}

	// 检查是否过期
	if !item.expiration.IsZero() && time.Now().After(item.expiration) {
		return false, nil
	}

	return true, nil
}

// Expire 设置键的过期时间
func (m *Memory) Expire(ctx context.Context, key string, expiration time.Duration) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item, exists := m.data[key]
	if !exists {
		return errors.New("key not found")
	}

	var exp time.Time
	if expiration > 0 {
		exp = time.Now().Add(expiration)
	}

	item.expiration = exp
	m.data[key] = item

	return nil
}

// Incr 对键的值进行自增
func (m *Memory) Incr(ctx context.Context, key string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item, exists := m.data[key]
	if !exists {
		m.data[key] = cacheItem{
			value:      int64(1),
			expiration: time.Time{},
		}
		return 1, nil
	}

	val, ok := item.value.(int64)
	if !ok {
		return 0, errors.New("value is not an integer")
	}

	val++
	item.value = val
	m.data[key] = item

	return val, nil
}

// Decr 对键的值进行自减
func (m *Memory) Decr(ctx context.Context, key string) (int64, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	item, exists := m.data[key]
	if !exists {
		m.data[key] = cacheItem{
			value:      int64(-1),
			expiration: time.Time{},
		}
		return -1, nil
	}

	val, ok := item.value.(int64)
	if !ok {
		return 0, errors.New("value is not an integer")
	}

	val--
	item.value = val
	m.data[key] = item

	return val, nil
}

// Close 关闭缓存（对于内存缓存，此方法什么也不做）
func (m *Memory) Close() error {
	return nil
}

// cleanupExpired 定期清理过期的缓存项
func (m *Memory) cleanupExpired() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		m.mutex.Lock()
		now := time.Now()
		for key, item := range m.data {
			if !item.expiration.IsZero() && now.After(item.expiration) {
				delete(m.data, key)
			}
		}
		m.mutex.Unlock()
	}
}
