package redis

import (
	"github.com/go-redis/redis/v7"
	"github.com/taufiqade/gowallet/models"
)

// AuthRedisRepository godoc
type AuthRedisRepository struct {
	client *redis.Client
}

// NewRedisAuthRepository godoc
func NewRedisAuthRepository(r *redis.Client) models.IRedisAuthRepository {
	return &AuthRedisRepository{client: r}
}

// Get godoc
func (a *AuthRedisRepository) Get(key string) (string, error) {
	res, err := a.client.Get(key).Result()
	return res, err
}

// Set godoc
func (a *AuthRedisRepository) Set(key, value string, exp int64) (err error) {
	_, err = a.client.Set(key, value, 0).Result()
	if exp > 0 {
		a.client.Do("EXPIREAT", key, exp)
	}
	return err
}
