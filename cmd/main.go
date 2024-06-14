package main

import (
	"github.com/meiti-x/golang-web-api/api"
	"github.com/meiti-x/golang-web-api/config"
	"github.com/meiti-x/golang-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer()
	defer cache.CloseRedisClient()
	cache.InitRedis(cfg)
}
