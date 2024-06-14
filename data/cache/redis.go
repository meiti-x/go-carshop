package cache

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/meiti-x/golang-web-api/config"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:         cfg.Redis.Host + ":" + cfg.Redis.Port,
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  cfg.Redis.DialTimeout,
		ReadTimeout:  cfg.Redis.ReadTimeout,
		WriteTimeout: cfg.Redis.WriteTimeout,
		PoolSize:     cfg.Redis.PoolSize,
		PoolTimeout:  cfg.Redis.PoolTimeout,
		IdleTimeout:  cfg.Redis.IdleCheckFrequency * time.Millisecond,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedisClient() {
	redisClient.Close()
}
