package middlewares

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-gin-boilerplate/config"
	"time"
)

var ctx = context.Background()
var rdb *redis.Client

func RedisInit(){
	c := config.GetConfig()
	rdb = redis.NewClient(&redis.Options{
		Addr:     c.GetString("db.redis"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func GetRedis() *redis.Client{
	return rdb
}

func RedisHSet(key, field, val string) {
	rdb.HSet(ctx, key, field, val)
	rdb.Expire(ctx, key, 24*time.Hour)
}