/*
 * @PackageName: myredis
 * @Description: 普通队列
 * @Author: limuzhi
 * @Date: 2022/12/2 17:57
 */

package myredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

//入队
func PushQueue(rdb *redis.Client, ctx context.Context, queueKey string, value interface{}) error {
	return rdb.RPush(ctx, queueKey, value).Err()
}

//出队
func PopQueue(rdb *redis.Client, queueKey string, timeout time.Duration) string {
	for {
		value, err := rdb.BLPop(context.Background(), timeout, queueKey).Result()
		if err == redis.Nil {
			time.Sleep(timeout)
			continue
		}
		if err != nil {
			time.Sleep(timeout)
			continue
		}
		if len(value) > 1 {
			if value[0] != queueKey {
				continue
			}
			return value[1]
		}
		time.Sleep(timeout)
	}
}
