package repositories

import (
	"assignment3/config"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisRepository interface {
	Set(key string, value any, duration time.Duration) error
	Get(key string, model any) (any, error)
	Clear() error
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisClient(cfg *config.Config) (RedisRepository, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &redisRepository{client: client}, nil
}

func (r *redisRepository) Set(key string, value any, duration time.Duration) error {
	rValue, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(context.Background(), key, rValue, duration).Err()
}

func (r *redisRepository) Get(key string, model any) (any, error) {
	result, err := r.client.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(result), &model); err != nil {
		return nil, err
	}
	return model, nil
}
func (r *redisRepository) Clear() error {
	return r.client.FlushDB(context.Background()).Err()
}
