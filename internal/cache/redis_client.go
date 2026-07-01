package cache

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedis(addr string) error {
	Client = redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("could not connect to Redis: %w", err)
	}

	fmt.Println("Connected to Redis successfully")
	return nil
}
