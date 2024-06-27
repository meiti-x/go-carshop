package main

import (
	"github.com/meiti-x/golang-web-api/api"
	"github.com/meiti-x/golang-web-api/config"
	"github.com/meiti-x/golang-web-api/data/cache"
	"github.com/meiti-x/golang-web-api/data/db"
	"github.com/meiti-x/golang-web-api/pkg/logging"
)

// @securityDefinitions.apikey AuthBearer
// @in header
// @name Authorization
func main() {
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)

	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDbClient()
	api.InitServer()

}
