/*
 * @PackageName: myredis
 * @Description: 延迟队列
 * @Author: limuzhi
 * @Date: 2022/12/2 18:06
 */

package myredis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// PushDelayQueue
// 1. add the timePiece(sample: "1645614542") to sorted set
// 2. rpush the real data to timePiece
//
// @delaySecond, the expected delay seconds, 600 means delay 600 second
// @maxTTL, the max time data will live if there is no consumer
func PushDelayQueue(ctx context.Context, rdb *redis.Client, key, value string, delaySecond, maxTTL int64) error {
	expireSecond := time.Now().Unix() + delaySecond
	timePiece := fmt.Sprintf("dq:%s:%d", key, expireSecond)
	z := redis.Z{Score: float64(expireSecond), Member: timePiece}
	v, err := rdb.ZAddNX(ctx, key, &z).Result()
	if err != nil {
		return err
	}
	_, err = rdb.RPush(ctx, timePiece, value).Result()
	if err != nil {
		return err
	}

	// new timePiece will set expire time
	if v > 0 {
		// consumer will also deleted the item
		rdb.Expire(ctx, timePiece, time.Second*time.Duration(maxTTL+delaySecond))
		// sorted set max live time
		rdb.Expire(ctx, key, time.Hour*24*3)
	}
	return err
}

// PopDelayQueue
// 1. get a timePiece from sorted set which is before time.Now()
// 2. lpop the real data from timePiece
//
// Usage: Use it in a script or goroutine
func PopDelayQueue(ctx context.Context, rdb *redis.Client, key string) (chan string, chan error) {
	resCh := make(chan string, 0)
	errCh := make(chan error, 1)
	go func() {
		defer close(resCh)
		defer close(errCh)
		for {
			now := time.Now().Unix()
			opt := redis.ZRangeBy{Min: "0", Max: strconv.FormatInt(now, 10), Count: 1}
			val, err := rdb.ZRangeByScore(ctx, key, &opt).Result()
			if err != nil {
				errCh <- err
				return
			}
			// sleep 1s if the queue is empty
			if len(val) == 0 {
				select {
				case <-ctx.Done():
					errCh <- ctx.Err()
					return
				case <-time.After(time.Second):
					continue
				}
			}
			for _, listK := range val {
				for {
					// read from the timePiece
					s, lErr := rdb.LPop(ctx, listK).Result()
					if lErr == nil {
						select {
						case resCh <- s:
						case <-ctx.Done():
							errCh <- ctx.Err()
							return
						}
					} else if lErr == redis.Nil {
						rdb.ZRem(ctx, key, listK)
						rdb.Del(ctx, listK)
						break
					} else {
						errCh <- lErr
						return
					}
				}
			}
		}
	}()
	return resCh, errCh
}
