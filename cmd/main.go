package main

import (
	"github.com/meiti-x/golang-web-api/api"
	"github.com/meiti-x/golang-web-api/config"
	"github.com/meiti-x/golang-web-api/data/cache"
	"github.com/meiti-x/golang-web-api/data/db"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	err := cache.InitRedis(cfg)
	if err != nil {
		panic(err)
	}
	defer cache.CloseRedis()
	print(cfg)
	err = db.InitDb(cfg)
	if err != nil {
		panic(err)
	}
	defer db.CloseDbClient()
	api.InitServer()

}
