package redisConn

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	mu                 sync.Mutex
	redisClient        *redis.Client
	redisClusterClient *redis.ClusterClient
)

func GetRedisConnection() *redis.Client {
	mu.Lock()
	defer mu.Unlock()
	if redisClient == nil {
		return redis.NewClient(&redis.Options{
			Addr: "0.0.0.0:6379",
		})
	}
	return redisClient
}

func GetRedisClusterConnection() *redis.ClusterClient {
	mu.Lock()
	defer mu.Unlock()
	if redisClusterClient == nil {
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{"0.0.0.0:6379"},
		})
	}
	return redisClusterClient
}
