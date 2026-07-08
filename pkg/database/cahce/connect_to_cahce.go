package cahce

import (
	"Bank-repository-service/internal/configs"
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

func NewRedisDB(cfg *configs.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Addr,
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     cfg.Redis.PoolSize,
		MinIdleConns: cfg.Redis.MinIdleConns,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	return client
}

