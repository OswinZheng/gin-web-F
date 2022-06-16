package middleware

import (
	"time"

	"github.com/OswinZheng/gin-web-F/configs"
	"github.com/OswinZheng/gin-web-F/internal/repository/redis"
	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	mGin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

var RateLimitMiddleWare gin.HandlerFunc

func RateLimit() gin.HandlerFunc {
	if RateLimitMiddleWare == nil {
		rate := limiter.Rate{
			Period: time.Duration(configs.Get().Server.RateLimit) * time.Second,
			Limit:  int64(configs.Get().Server.RateLimit),
		}
		redisClient := redis.Client()
		store, err := sRedis.NewStoreWithOptions(redisClient, limiter.StoreOptions{
			Prefix: "seed_project_limiter",
		})
		if err != nil {
			panic(err)
		}
		RateLimitMiddleWare = mGin.NewMiddleware(limiter.New(store, rate))
	}
	return RateLimitMiddleWare
}
