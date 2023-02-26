package services

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis", // docker-compose had setup the hostname for redis
		Password: "",      // No password. This is not recommended in production.
		DB:       0,       // Use default database
	})
}

func Get(key string) string {
	fmt.Println("Get", key)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return val
}

func Set(key string, val string) {
	fmt.Println("Set", key, val)
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}

	// Expire the key after 1 day.
	rdb.Expire(ctx, key, 86400)
}
