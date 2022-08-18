package russell_luo

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	redisConn "rate-limit-examples/common"

	sw "github.com/RussellLuo/slidingwindow"
	"github.com/go-redis/redis/v8"
)

var (
	store *RedisDatastore
	once  sync.Once
	lim   *sw.Limiter
)

type RedisDatastore struct {
	client redis.Cmdable
	ttl    time.Duration
}

func NewRedisDatastore(client redis.Cmdable, ttl time.Duration) *RedisDatastore {
	if store == nil {
		store = &RedisDatastore{client: client, ttl: ttl}
	}
	return store
}

func (d *RedisDatastore) fullKey(key string, start int64) string {
	return fmt.Sprintf("%s@%d", key, start)
}

func (d *RedisDatastore) Add(key string, start, value int64) (int64, error) {
	k := d.fullKey(key, start)
	c, err := d.client.IncrBy(context.TODO(), k, value).Result()
	if err != nil {
		return 0, err
	}
	// Ignore the possible error from EXPIRE command.
	d.client.Expire(context.TODO(), k, d.ttl).Result() // nolint:errcheck
	return c, err
}

func (d *RedisDatastore) Get(key string, start int64) (int64, error) {
	k := d.fullKey(key, start)
	value, err := d.client.Get(context.TODO(), k).Result()
	if err != nil {
		if err == redis.Nil {
			// redis.Nil is not an error, it only indicates the key does not exist.
			err = nil
		}
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

func AllowRequest(key string) bool {

	once.Do(func() {
		size := time.Minute
		store := NewRedisDatastore(
			redisConn.GetRedisConnection(),
			2*size,
		)
		lim, _ = sw.NewLimiter(size, 1000, func() (sw.Window, sw.StopFunc) {
			return sw.NewSyncWindow(key, sw.NewNonblockingSynchronizer(store, 1*time.Second))
		})

	})

	return lim.Allow()
}
