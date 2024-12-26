package utility

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"sms_backend/api/v1/sms"
	"sync"
)

const cacheLogger = "cache"

func AddValueToCacheAndRedisQueue(ctx context.Context, keyList string, message interface{}) (err error) {
	if v, ok := LocalCache.Get(keyList); ok {
		v.(*CacheQueue).Push(message)
	} else {
		cq := NewCacheQueue(0)
		cq.Push(message)
		LocalCache.Add(keyList, cq)
	}
	if _, err = g.Redis().LPush(ctx, keyList, message); err != nil {
		g.Log(cacheLogger).Error(ctx, err)
		return errors.New("将文件内容存储到Redis 失败")
	}
	return nil
}

func GetValueFromCacheOrRedisQueue(ctx context.Context, listKey string) (interface{}, error) {
	// Try to fetch the queue from local cache
	if cachedQueue, ok := LocalCache.Get(listKey); ok {
		queue, valid := cachedQueue.(*CacheQueue)
		if !valid {
			g.Log(cacheLogger).Errorf(ctx, "Invalid cache type for key: %s, expected *CacheQueue", listKey)
			//LocalCache.Remove(keyList) // Remove invalid cache
			return fetchFromRedis(ctx, listKey)
		}

		// Check if the queue is empty
		if queue.Len() == 0 {
			g.Log(cacheLogger).Warningf(ctx, "Local cache queue is empty for key: %s, removing cache and falling back to Redis", listKey)
			//LocalCache.Remove(keyList) // Remove empty cache
			return fetchFromRedis(ctx, listKey)
		}

		// Attempt to pop a value from the queue
		if value, ok := queue.Pop(); ok {
			return value, nil
		}

		// If pop fails, log and fall back to Redis
		g.Log(cacheLogger).Errorf(ctx, "Failed to pop value from cache queue for key: %s, falling back to Redis", listKey)
		//LocalCache.Remove(keyList) // Remove corrupted cache
	}

	// If cache is unavailable or invalid, fetch from Redis
	return fetchFromRedis(ctx, listKey)
}

// Helper function to fetch the value from Redis
func fetchFromRedis(ctx context.Context, listKey string) (interface{}, error) {
	// Check the length of the Redis list
	listLen, err := g.Redis().LLen(ctx, listKey)
	if err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to get length of Redis list for key '%s': %v", listKey, err)
		return nil, fmt.Errorf("failed to get length of Redis list for key '%s': %w", listKey, err)
	}

	// If the list is empty, clean up and return an error
	if listLen == 0 {
		g.Log(cacheLogger).Warningf(ctx, "Redis list '%s' is empty. Deleting the list and returning error.", listKey)
		if _, err := g.Redis().Del(ctx, listKey); err != nil {
			g.Log().Errorf(ctx, "Failed to delete empty Redis list for key '%s': %v", listKey, err)
			return nil, fmt.Errorf("failed to delete empty Redis list for key '%s': %w", listKey, err)
		}
		return nil, fmt.Errorf("Redis list '%s' is empty", listKey)
	}

	// If the list has items, try to pop one
	g.Log(cacheLogger).Infof(ctx, "Processing tasks from Redis list '%s'. List length: %d", listKey, listLen)
	// Local Queue is LPop -> Redis Queue use RPop
	messageData, err := g.Redis().RPop(ctx, listKey)
	if err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to LPop from Redis list for key '%s': %v", listKey, err)
		return nil, fmt.Errorf("failed to pop value from Redis list for key '%s': %w", listKey, err)
	}

	// Validate the popped message data
	if messageData == nil {
		g.Log(cacheLogger).Warningf(ctx, "No task found in Redis list '%s'. Queue is likely empty.", listKey)
		return nil, fmt.Errorf("no task found in Redis list '%s'", listKey)
	}
	Debug(ctx, "fetch_task_by_redis_cache", "cache_redis_fetch_task.txt")
	return messageData.Interface(), nil
}

func AddValueToSafeSliceAndRedis(ctx context.Context, listKey string, message interface{}) (err error) {
	if v, ok := LocalCache.Get(listKey); ok {
		v.(*SafeSlice).Append(message)
	} else {
		cq := NewSafeSlice()
		cq.Append(message)
		LocalCache.Add(listKey, cq)
	}
	if _, err = g.Redis().LPush(ctx, listKey, message); err != nil {
		g.Log(cacheLogger).Error(ctx, err)
		return errors.New("将文件内容存储到Redis 失败")
	}
	return nil
}

func RetrieveValueFromCacheAndRedis(ctx context.Context, listKey string, index int) (interface{}, error) {
	// 尝试从本地缓存中获取
	if v, ok := LocalCache.Get(listKey); ok {
		slice, ok := v.(*SafeSlice)
		if !ok {
			g.Log(cacheLogger).Errorf(ctx, "LocalCache for key [%s] is not of type SafeSlice", listKey)
			return nil, fmt.Errorf("invalid type in local cache for key [%s]", listKey)
		}

		// 从本地缓存的 SafeSlice 获取值
		message, err := slice.Get(index)
		if err == nil {
			return message, nil
		}

		g.Log(cacheLogger).Warningf(ctx, "Failed to get index [%d] from SafeSlice for key [%s]: %v", index, listKey, err)
	}

	// 如果本地缓存获取失败，从 Redis 中获取
	v, err := g.Redis().LIndex(ctx, listKey, int64(index))
	if err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to get index [%d] from Redis for key [%s]: %v", index, listKey, err)
		return nil, fmt.Errorf("failed to get value from Redis for key [%s], index [%d]: %w", listKey, index, err)
	}

	// 返回 Redis 获取的值
	return v.Interface(), nil
}

type GenericMap struct {
	data  map[string]interface{}
	mutex sync.RWMutex // 读写锁
}

// NewGenericMap 构造函数，用于初始化 GenericMap
func NewGenericMap() *GenericMap {
	return &GenericMap{
		data: make(map[string]interface{}),
	}
}

// Set 设置键值对，加锁以保护 map
func (m *GenericMap) Set(key string, value interface{}) {
	m.mutex.Lock() // 在写操作前加锁
	m.data[key] = value
	m.mutex.Unlock() // 写操作后解锁
}

// Get 获取值，使用读锁
func (m *GenericMap) Get(key string) (interface{}, bool) {
	m.mutex.RLock() // 在读操作前加读锁
	value, exists := m.data[key]
	m.mutex.RUnlock() // 读操作后解读锁
	return value, exists
}

// Delete 删除键值对，加锁以保护 map
func (m *GenericMap) Delete(key string) {
	m.mutex.Lock() // 在写操作前加锁
	delete(m.data, key)
	m.mutex.Unlock() // 写操作后解锁
}

// GetMap 获取整个 map 的副本，使用读锁
func (m *GenericMap) GetMap() map[string]interface{} {
	m.mutex.RLock() // 在读操作前加读锁
	copiedMap := make(map[string]interface{})
	for key, value := range m.data {
		copiedMap[key] = value
	}
	m.mutex.RUnlock() // 读操作后解读锁
	return copiedMap
}
func CheckMapExistsByCacheAndRedis(ctx context.Context, key, hashKey string) (bool, error) {
	// 从本地缓存获取数据
	cacheValue, cacheExists := LocalCache.Get(key)
	if cacheExists {
		// 类型断言为 GenericMap
		if localMap, ok := cacheValue.(*GenericMap); ok {
			// 检查是否存在指定的 hashKey
			if _, exists := localMap.Get(hashKey); exists {
				g.Log(cacheLogger).Infof(ctx, "Key '%s' found in local cache, hashKey '%s' exists", key, hashKey)
				return true, nil
			}
		} else {
			g.Log(cacheLogger).Warningf(ctx, "Key '%s' exists in cache but cannot be converted to GenericMap", key)
		}
	}

	// 如果本地缓存不存在或未命中，再从 Redis 检查
	g.Log(cacheLogger).Infof(ctx, "Checking Redis for key '%s' and hashKey '%s'", key, hashKey)
	exists, err := g.Redis().HExists(ctx, key, hashKey)
	if err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to check hashKey existence in Redis for key '%s': %v", key, err)
		return false, err
	}

	// 返回 Redis 的检查结果
	if exists > 0 {
		g.Log(cacheLogger).Infof(ctx, "Key '%s' and hashKey '%s' found in Redis", key, hashKey)
		return true, nil
	}

	g.Log(cacheLogger).Infof(ctx, "Key '%s' and hashKey '%s' not found in cache or Redis", key, hashKey)
	return false, nil
}

func SetMapValueToCacheAndRedis(ctx context.Context, key, hashKey string, message interface{}) error {
	var myMap *GenericMap
	if v, ok := LocalCache.Get(key); ok {
		myMap = v.(*GenericMap)

	} else {
		myMap = NewGenericMap()
		LocalCache.Add(key, myMap)
	}
	myMap.Set(hashKey, message)
	_, err := g.Redis().HSet(ctx, key, map[string]interface{}{hashKey: message})
	if err != nil {
		g.Log(cacheLogger).Error(ctx, "Failed to set hash fields:", err)
		return err
	}
	g.Log(cacheLogger).Info(ctx, "Hash fields set successfully")
	return nil
}

func ParseCacheValueToRecord(v interface{}) (*sms.SubPostConversationRecordData, error) {
	// 检查输入是否为 nil
	if v == nil {
		return nil, fmt.Errorf("input value is nil")
	}

	// 初始化返回结果
	res := &sms.SubPostConversationRecordData{}

	switch value := v.(type) {
	case string:
		fmt.Printf("The type is string and the value is: %s\n", value)
		// 解析 JSON 字符串到结构体
		err := json.Unmarshal([]byte(value), res)
		if err != nil {
			fmt.Printf("Error decoding JSON: %v\n", err)
			return nil, err
		}
		return res, nil

	case *sms.SubPostConversationRecordData:
		// 如果输入本身已经是目标类型，直接返回
		return value, nil

	default:
		// 未知类型
		fmt.Printf("Unknown type: %T\n", value)
		return nil, fmt.Errorf("unknown type: %T", value)
	}
}

// 释放垃圾缓存

func RemoveCacheAndRedisKey(ctx context.Context, key string) (err error) {
	if _, ok := LocalCache.Get(key); ok {
		LocalCache.Remove(key)
	}
	if _, err = g.Redis().Del(ctx, key); err != nil {
		g.Log(cacheLogger).Error(ctx, err)
		return errors.New("将文件内容从Redis删除 失败")
	}
	return nil
}

// RemoveCacheAndRedisHashField removes a specific field from both the local cache and the Redis hash.
// If the field exists in the local cache, it will be deleted from the cached map.
// It also attempts to delete the same field from the Redis hash.
//
// Parameters:
// - ctx: The context for handling the operation.
// - key: The key of the hash in both local cache and Redis.
// - hashKey: The field to delete from the hash.
//
// Returns:
// - error: If any error occurs during the operation, it will be returned.
func RemoveCacheAndRedisHashField(ctx context.Context, key, hashKey string) error {
	// Attempt to remove the field from the local cache
	if cachedValue, exists := LocalCache.Get(key); exists {
		if myMap, valid := cachedValue.(*GenericMap); valid {
			myMap.Delete(hashKey)
			g.Log(cacheLogger).Infof(ctx, "Successfully deleted field '%s' from local cache for key '%s'", hashKey, key)
		} else {
			g.Log(cacheLogger).Warningf(ctx, "Invalid type for cached value under key '%s'. Skipping local cache deletion.", key)
		}
	}

	// Attempt to delete the field from the Redis hash
	if _, err := g.Redis().HDel(ctx, key, hashKey); err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to delete field '%s' from Redis hash '%s': %v", hashKey, key, err)
		return fmt.Errorf("failed to delete field '%s' from Redis hash '%s': %w", hashKey, key, err)
	}

	g.Log(cacheLogger).Infof(ctx, "Successfully deleted field '%s' from Redis hash '%s'", hashKey, key)
	return nil
}

// CleanUpDeviceCache cleans up a specific device cache (e.g., unregistered or unassigned) from both local cache and Redis.
// It checks if the cache exists and removes it if found.
//
// Parameters:
// - ctx: The context for handling the operation.
// - cacheKeyPrefix: The cache prefix to distinguish the type of cache (e.g., "unregistered" or "unassigned").
// - deviceNumber: The device number to clean up.
//
// Returns:
// - error: An error if the cache cleanup fails.
func CleanUpDeviceCache(ctx context.Context, cacheKeyPrefix, deviceNumber string) error {
	cacheKey := cacheKeyPrefix
	hashKey := cacheKey + deviceNumber

	// Check if the cache exists
	exists, err := CheckMapExistsByCacheAndRedis(ctx, cacheKey, hashKey)
	if err != nil {
		g.Log(cacheLogger).Errorf(ctx, "Failed to check device cache for device_number '%s' with prefix '%s': %v", deviceNumber, cacheKeyPrefix, err)
		return fmt.Errorf("failed to check device cache: %w", err)
	}

	// Remove the cache if it exists
	if exists {
		if err := RemoveCacheAndRedisHashField(ctx, cacheKey, hashKey); err != nil {
			g.Log(cacheLogger).Errorf(ctx, "Failed to remove device cache for device_number '%s' with prefix '%s': %v", deviceNumber, cacheKeyPrefix, err)
			return fmt.Errorf("failed to remove device cache: %w", err)
		}
		g.Log(cacheLogger).Infof(ctx, "Removed device cache for device_number '%s' with prefix '%s'", deviceNumber, cacheKeyPrefix)
	}

	return nil
}
