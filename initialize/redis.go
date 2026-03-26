package initialize

import (
	"os"
	"server/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func ConnectRedis() redis.Client {
	redisCfg := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		global.Log.Error("Failed to connect to Redis", zap.Error(err))
		os.Exit(1)
	}
	return *rdb
}
