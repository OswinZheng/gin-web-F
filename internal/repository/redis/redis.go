package redis

import (
	"context"
	"fmt"

	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

var ctx = context.Background()

func Client() *redis.Client {
	if redisClient == nil {
		ctx := context.Background()
		redisClient = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", configs.Get().Redis.Host, configs.Get().Redis.Port),
			Password: configs.Get().Redis.Password, // no password set
			DB:       configs.Get().Redis.Db,       // use default DB
			PoolSize: configs.Get().Redis.PoolSize,
		})
		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
	}
	return redisClient
}

func Close() {
	redisClient.Close()
}

func Set(key, value string) {
	err := redisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	return redisClient.Get(ctx, key).Result()
}
