package services

import (
	"context"
	"github.com/redis/go-redis/v9"
	"fmt"
)

var ctx = context.Background()
var rdb

func init() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis", // docker-compose had setup the hostname for redis
		Password: "",      // No password. This is not recommended in production.
		DB:       0,       // Use default database
	})
}

func Get(val) (key) {
	fmt.Println("Get", key, val)
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func Set() (key, val) {
	fmt.Println("Set", key, val)
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		panic(err)
	}
}