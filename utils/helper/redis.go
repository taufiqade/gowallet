package helper

import (
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
)

var redisConn *redis.Client

// GetInstance godoc
func GetInstance() *redis.Client {
	dsn := fmt.Sprintf("%s:%d", os.Getenv("REDIS_HOST"), 6379)
	redisConn = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := redisConn.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return redisConn
}

// Get godoc
func Get(key string) (string, error) {
	return GetInstance().Get(key).Result()
}

// Set godoc
func Set(key, value string, exp int64) (err error) {
	_, err = GetInstance().Set(key, value, 0).Result()
	if exp > 0 {
		GetInstance().Do("EXPIREAT", key, exp)
	}
	return
}
