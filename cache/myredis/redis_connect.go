/*
 * @PackageName: cache
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 15:42
 */

package myredis

import (
	"context"
	"github.com/go-redis/redis/extra/redisotel"
	"github.com/go-redis/redis/v8"
)

func NewRedisClient(ctx context.Context, opt *redis.Options) (*redis.Client, error) {
	rdb := redis.NewClient(opt)
	rdb.AddHook(redisotel.TracingHook{})
	_, err := rdb.Ping(ctx).Result()
	return rdb, err
}

//func FlushDB(ctx context.Context, rdb *redis.Client) error {
//	return rdb.FlushDB(ctx).Err()
//}
