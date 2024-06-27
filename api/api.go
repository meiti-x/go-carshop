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
	"github.com/meiti-x/golang-web-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterSwagger(r *gin.Engine, cfg *config.Config) {
	docs.SwaggerInfo.Title = "golang web api"
	docs.SwaggerInfo.Description = "golang web api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", cfg.Server.ExternalPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
func InitServer() {
	cfg := config.GetConfig()
	fmt.Printf("cfg: %v\n", cfg)
	r := gin.New()

	val, ok := binding.Validator.Engine().(*validator.Validate)

	if ok {
		val.RegisterValidation("isAdult", validations.IsAdult)
	}
	r.Use(gin.Logger(), gin.Recovery(), middlewares.TestMiddleware(), middlewares.RateLimitMiddleware())
	api := r.Group("api/")

	v1 := api.Group("/v1/")
	{
		health := v1.Group("/health")
		test_router := v1.Group("/test")
		routers.Health(health)
		routers.TestRouter(test_router)
	}
	RegisterSwagger(r, cfg)

	r.Run(":5005")

}
