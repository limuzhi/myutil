/*
 * @PackageName: captcha
 * @Description: 验证存储redis重写
 * @Author: limuzhi
 * @Date: 2022/8/2 13:38
 */

package captcha

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strings"
	"time"
)

type RedisStore struct {
	redis *redis.Client
	exp   time.Duration
}

func NewRedisStore(rc *redis.Client, exp time.Duration) *RedisStore {
	if exp == 0 {
		exp = time.Minute * time.Duration(10)
	}
	return &RedisStore{
		redis: rc,
		exp:   exp,
	}
}

// Set sets the digits for the captcha id.
func (s *RedisStore) Set(id string, value string) error {
	s.redis.Set(context.Background(), id, value, s.exp)
	return nil
}

// Get returns stored digits for the captcha id. Clear indicates
// whether the captcha must be deleted from the store.
func (s *RedisStore) Get(id string, clear bool) string {
	if clear {
		defer func() {
			s.redis.Del(context.Background(), id)
		}()
	}
	return s.redis.Get(context.Background(), id).Val()
}

//Verify captcha's answer directly
func (s *RedisStore) Verify(id, answer string, clear bool) bool {
	val := s.Get(id, clear)

	return strings.ToLower(val) == strings.ToLower(answer)
}
