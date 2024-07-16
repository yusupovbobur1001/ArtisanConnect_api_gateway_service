package api

import (
	"api_service/api/handler"
	"api_service/api/middleware"
	"api_service/config"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/reservedesk.uz")
	api.Use(middleware.JWTMiddleware())

	h := handler.NewHandler(cfg)

	product := api.Group("/products")
	product.POST("/", h.CreateProduct)
	product.PUT("/:id", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)
	product.

}
