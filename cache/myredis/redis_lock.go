/*
 * @PackageName: myredis
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 10:15
 */

package myredis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var mutex sync.Mutex

//枷锁
func Lock(ctx context.Context, rdb *redis.Client, lockKey string, lockVal interface{}, d time.Duration) (bool, error) {
	mutex.Lock()
	defer mutex.Unlock()
	lockSuccess, err := rdb.SetNX(ctx, lockKey, lockVal, d).Result()
	if err != nil || !lockSuccess {
		return false, err
	}
	return lockSuccess, nil
}

//解锁
func UnLock(ctx context.Context, rdb *redis.Client, lockKey string) int64 {
	resNum, err := rdb.Del(ctx, lockKey).Result()
	if err != nil {
		return 0
	}
	return resNum
}

func Lock2(ctx context.Context, rdb *redis.Client, ) {

}