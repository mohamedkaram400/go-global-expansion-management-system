package conn

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func ConnectRedis(redisHost string) (*redis.Client, error) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisHost,
	})

	// Test connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis successfully")
	return RedisClient, nil
}
