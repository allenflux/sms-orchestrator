package utility

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
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

func PopWithLock(ctx context.Context, redis *gredis.Redis, listKey string) (result *gvar.Var, err error) {
	lockKey := "task_queue_lock" + listKey // 锁的键
	lockTTL := 5 * time.Second             // 锁的过期时间
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
		return nil, nil // 队列为空
	}

	return result, nil
}
