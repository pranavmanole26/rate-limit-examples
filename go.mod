module rate-limit-examples

go 1.18

require github.com/go-redis/redis v6.15.9+incompatible

require (
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/garyburd/redigo v1.6.3 // indirect
)

require (
	github.com/RussellLuo/slidingwindow v0.0.0-20200528002341-535bb99d338b
	github.com/go-redis/redis/v8 v8.11.5
	github.com/go-redis/redis_rate/v9 v9.1.2
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/oeegor/go-redis-ratelimiter v0.0.0-20160420134335-f8d22f28c1fc
)
