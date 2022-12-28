package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
	"zusammen/internal/domain/entity"
)

type RedisCache struct {
	Host    string
	Db      int
	Expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) RedisCache {
	return RedisCache{
		Host:    host,
		Db:      db,
		Expires: exp,
	}
}

func (cache *RedisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: "root",
		DB:       cache.Db,
	})
}

func (cache *RedisCache) Set(key string, value *entity.User) {
	client := cache.getClient()

	data, err := json.Marshal(value)
	if err != nil {
		log.Println(err)
	}
	var ctx context.Context
	client.Set(ctx, key, data, cache.Expires*time.Second)
}

func (cache *RedisCache) Get(key string) *entity.User {
	client := cache.getClient()

	var ctx context.Context
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}

	user := entity.User{}
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		log.Println(err)
	}
	return &user
}
