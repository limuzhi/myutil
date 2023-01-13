/*
 * @PackageName: myredis
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 17:52
 */

package myredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func TestNewRedisClient(t *testing.T) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()
	rdb, err := NewRedisClient(ctx, &redis.Options{
		Addr:     "127.0.0.1:6379",
		Username: "",
		Password: "",
		DB:       1,
	})
	if err != nil {
		t.Errorf("NewRedisClient() error = %v", err)
		return
	}
	rdb.Set(ctx, "aa", 11, time.Second*50)
	t.Log("success")
}
