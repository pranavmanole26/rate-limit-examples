package redis_ratelimiter

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	ratelimiter "github.com/oeegor/go-redis-ratelimiter"
)

var limitCtx *ratelimiter.LimitCtx

func GetRedisPool() *redis.Pool {
	return redis.NewPool(func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", "0.0.0.0:6379", redis.DialPassword(""))
		if err != nil {
			return nil, err
		}
		return c, err
	}, 100)
}

func AllowRequest(key string) bool {
	if limitCtx != nil && limitCtx.Current >= limitCtx.Limit {
		if limitCtx.ExpireAt < int64(time.Now().Unix()) {
			limitCtx = nil
		}
		return false
	}
	cur := time.Now()
	limitCtx, _ = ratelimiter.Incr(GetRedisPool(), key, 5, time.Minute, 1)
	fmt.Println(time.Since(cur))
	return true
}
