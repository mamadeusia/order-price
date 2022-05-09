package db

import (
	"context"
	"fmt"
	"preh/config"

	"github.com/go-redis/redis/v8"
)

var (
	InitDb        = initDb
	SetToRedis    = setToRedis
	GetFromRedis  = getFromRedis
	SetHToRedis   = setHToRedis
	GetHFromRedis = getHFromRedis
	UpdateHRedis  = updateHRedis
)

var redisClient *redis.Client

func initDb(ctx context.Context) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.GetRedisUrl(),
		Password: config.GetRedisPass(),
		DB:       0,
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}
func setToRedis(ctx context.Context, key, val string) {
	err := redisClient.Set(ctx, key, val, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
}
func getFromRedis(ctx context.Context, key string) string {
	val, err := redisClient.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}
func setHToRedis(ctx context.Context, key, field string, value interface{}) error {
	fmt.Println("KEY :: ", key, "\n FIELD :: ", field, "\n VALUE :: ", value)
	if _, err := redisClient.HSetNX(ctx, key, field, value).Result(); err != nil {
		fmt.Println("passed setHToRedisError")
		return fmt.Errorf("set: redis error: %w", err)
	}
	return nil
}
func getHFromRedis(ctx context.Context, key, field string) ([]byte, error) {
	result, err := redisClient.HGet(ctx, key, field).Result()
	if err != nil && err != redis.Nil {
		return nil, fmt.Errorf("find: redis error: %w", err)
	}
	if result == "" {
		return nil, fmt.Errorf("find: not found")
	}
	return []byte(result), nil

}
func updateHRedis(ctx context.Context, key, field string, value interface{}) error {
	if _, err := redisClient.HSet(ctx, key, field, value).Result(); err != nil {
		return fmt.Errorf("update: redis error: %w", err)
	}
	return nil
}
