package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const TTL = 72 * time.Hour

type URLRepository interface {
	Save(ctx context.Context, code string, url string) error
	Get(ctx context.Context, code string) (string, error)
}

type redisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) URLRepository {
	return &redisRepository{
		client: client,
	}
}


func (r *redisRepository) Save(ctx context.Context, code string, url string) error {
	return r.client.Set(ctx, code, url, TTL).Err()	
}

func (r *redisRepository) Get(ctx context.Context, code string) (string, error) {
	value, err := r.client.Get(ctx, code).Result()
	if err == redis.Nil {
		return "", nil
	}
	return value, err
}