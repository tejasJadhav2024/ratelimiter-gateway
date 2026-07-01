package ratelimiter

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisFixedWindowLimiter struct {
	client     *redis.Client
	limit      int
	windowSize time.Duration
}

func NewRedisFixedWindowLimiter(client *redis.Client, limit int, windowSize time.Duration) *RedisFixedWindowLimiter {
	return &RedisFixedWindowLimiter{
		client:     client,
		limit:      limit,
		windowSize: windowSize,
	}
}

func (l *RedisFixedWindowLimiter) Allow(clientID string) bool {
	ctx := context.Background()

	windowKey := fmt.Sprintf("ratelimit:%s:%d",
		clientID,
		time.Now().Unix()/int64(l.windowSize.Seconds()),
	)

	pipe := l.client.Pipeline()
	incrCmd := pipe.Incr(ctx, windowKey)
	pipe.Expire(ctx, windowKey, l.windowSize)
	_, err := pipe.Exec(ctx)

	if err != nil {
		// fail open — if Redis is down, allow the request
		return true
	}

	count := incrCmd.Val()
	return count <= int64(l.limit)
}
