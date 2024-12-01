package redis

import (
	"context"
	"fmt"
	"hello-go/configs"
	"hello-go/zlog"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

const (
	defaultTTL = 30 * time.Minute
)

var client *redis.Client

// 初始化redis客户端
func InitRedis() {
	redisConf := configs.Get().Redis
	client = redis.NewClient(&redis.Options{
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
	}
}

func init() {
	if client == nil {
		InitRedis()
	}
}

func Get(key string) (value string, err error) {
	return client.Get(context.Background(), key).Result()
}

func Set(key string, value any) (err error) {
	return client.Set(context.Background(), key, value, defaultTTL).Err()
}
func SetWithTTL(key string, value any, sec time.Duration) (err error) {
	return client.Set(context.Background(), key, value, sec*time.Second).Err()
}