package utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/google/uuid"
	"time"
)

// TryLock 尝试获取 Redis 分布式锁
func TryLock(ctx context.Context, redis *gredis.Redis, lockKey string, lockTTL time.Duration) (bool, error) {
	result, err := redis.Do(ctx, "SET", lockKey, "locked", "NX", "EX", int(lockTTL.Seconds()))

	if err != nil {
		return false, err
	}

	// 判断 SET NX EX 操作是否成功
	return result.String() == "OK", nil
}

// Unlock 释放 Redis 分布式锁
func Unlock(ctx context.Context, redis *gredis.Redis, lockKey string) error {
	_, err := redis.Do(ctx, "DEL", lockKey)
	return err
}

const (
	LockKey = "TASK_QUEUE_LOCK"
	LockTTL = 5
)

func PopWithLock(ctx context.Context, redis *gredis.Redis, listKey string) (result *gvar.Var, err error) {
	lockKey := "task_queue_lock" // 锁的键
	lockTTL := 5 * time.Second   // 锁的过期时间
	// 尝试获取锁
	locked, err := TryLock(ctx, redis, lockKey, lockTTL)
	if err != nil {
		return nil, err
	}
	if !locked {
		return nil, gerror.New("failed to acquire lock")
	}
	defer Unlock(ctx, redis, lockKey) // 确保释放锁

	// 执行 LPOP 操作
	result, err = redis.LPop(ctx, listKey)
	if err != nil {
		return nil, err
	}

	if result.IsNil() {
		return nil, gerror.New("Queue IS Empty") // 队列为空
	}

	return result, nil
}

// RedisQueue 是 Redis 队列的结构体，包含列表键和锁键
type RedisQueue struct {
	ListKey string // 任务队列的 Redis 键
	LockKey string // Redis 锁的键
	TTL     int    // 锁的过期时间（单位：秒）
}

// NewRedisQueue 创建一个新的 Redis 队列实例
func NewRedisQueue(listKey, lockKey string, ttl int) *RedisQueue {
	return &RedisQueue{
		ListKey: listKey,
		LockKey: lockKey,
		TTL:     ttl,
	}
}

// Pop 方法从任务队列中弹出任务（带锁，防止任务重复分发）
func (q *RedisQueue) Pop(ctx context.Context, redis *gredis.Redis) (result *gvar.Var, err error) {
	// Lua 脚本：加锁、弹出任务和解锁的原子操作
	luaScript := `
	if redis.call("SET", KEYS[1], ARGV[1], "NX", "EX", ARGV[2]) then
		local task = redis.call("LPOP", KEYS[2])
		if task then
			redis.call("DEL", KEYS[1])
			return task
		else
			redis.call("DEL", KEYS[1])
			return nil
		end
	else
		return nil
	end
	`

	// 生成唯一的锁标识（UUID）
	lockValue := uuid.NewString()

	// 执行 Lua 脚本
	result, err = redis.Do(ctx, "EVAL", luaScript, 2, q.LockKey, q.ListKey, lockValue, q.TTL)
	if err != nil {
		return nil, fmt.Errorf("failed to execute Lua script: %w", err)
	}

	// 如果返回值为空，表示任务队列为空或锁未获取到
	if result.IsNil() {
		return nil, nil
	}

	// 返回任务数据
	return result, nil
}

// KeyExists 判断 Redis 中某个 Key 是否存在
func KeyExists(ctx context.Context, redis *gredis.Redis, key string) (bool, error) {
	result, err := redis.Do(ctx, "EXISTS", key)
	if err != nil {
		return false, fmt.Errorf("failed to check if key exists: %v", err)
	}

	// Redis `EXISTS` 返回一个整数，1 表示存在，0 表示不存在
	count := result.Int()

	return count > 0, nil
}
