package redis

import (
	"errors"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

func NewTestRedis() *redismock.ClientMock {
	mr, err := miniredis.Run()
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return redismock.NewNiceMock(client)
}
func Test_redis_Get(t *testing.T) {

}

func RedisIsAvailable(client redis.Cmdable) bool {
	return client.Ping().Err() == nil
}

// Test Redis is down.
func TestRedisCannotBePinged(t *testing.T) {
	r := NewTestRedis()
	r.On("Ping").
		Return(redis.NewStatusResult("", errors.New("server not available")))

	assert.False(t, RedisIsAvailable(r))
}
