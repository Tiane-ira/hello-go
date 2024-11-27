package redis

import (
	"context"
	"fmt"
	"hello-go/configs"
	"hello-go/zlog"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// 初始化redis客户端
func NewClinet() *redis.Client {
	redisConf := configs.Get().Redis
	client := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", redisConf.Host, redisConf.Port),
		Password:     redisConf.Password,
		DB:           redisConf.Db,
		MaxRetries:   redisConf.MaxRetries,
		MaxIdleConns: redisConf.MaxIdleConns,
		PoolSize:     redisConf.PoolSize,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		zlog.Logger.Error("redis connect ping failed", zap.Error(err))
		panic(err)
	} else {
		zlog.Logger.Info("redis connected", zap.String("pong", pong))
		return client
	}
}
