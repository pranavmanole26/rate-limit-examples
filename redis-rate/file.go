package redisrate

import (
	"context"
	"sync"
	"time"

	redisConn "rate-limit-examples/common"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redis_rate/v9"
)

var (
	once    sync.Once
	rdb     *redis.Client
	rcdb    *redis.ClusterClient
	limiter *redis_rate.Limiter
)

func AllowRequest(key string) bool {
	ctx := context.Background()
	once.Do(func() {
		rdb = redisConn.GetRedisConnection()
		limiter = redis_rate.NewLimiter(rdb)
	})

	res, err := limiter.Allow(ctx, key, redis_rate.Limit{Rate: 1000, Burst: 1000, Period: 60 * time.Second})
	if err != nil {
		panic(err)
	}

	return res.Allowed > 0
}
