package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/meiti-x/golang-web-api/api/middlewares"
	"github.com/meiti-x/golang-web-api/api/routers"
	"github.com/meiti-x/golang-web-api/api/validations"
	"github.com/meiti-x/golang-web-api/config"
)

func InitServer() {
	cfg := config.GetConfig()
	fmt.Printf("cfg: %v\n", cfg)
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("isAdult", validations.IsAdult)
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleware())
	api := r.Group("api/")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}

	r.Run(":5005")

}
