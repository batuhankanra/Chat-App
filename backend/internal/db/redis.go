package db

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
	ctx    context.Context
}

func RedisConnect(addr, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &RedisClient{
		Client: rdb,
		ctx:    context.Background(),
	}
}

func (rc *RedisClient) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if _, err := rc.Client.Ping(ctx).Result(); err != nil {
		return fmt.Errorf("Redis bağlantı hatası: %w", err)
	}
	fmt.Println("Ping successful")
	return nil
}

func (rc *RedisClient) Close() error {
	return rc.Client.Close()
}
