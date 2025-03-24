package cache

import (
	"context"
	"time"

	"github.com/langyan/web-gin-gorm-redis/config"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitRedis() {
 RedisClient = redis.NewClient(&redis.Options{
  Addr:     config.AppConfig.RedisAddr,
  Password: config.AppConfig.RedisPassword,
  DB:       config.AppConfig.RedisDB,
 })
}

func SetCache(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
 return RedisClient.Set(ctx, key, value, expiration).Err()
}

func GetCache(ctx context.Context, key string) (string, error) {
 return RedisClient.Get(ctx, key).Result()
}