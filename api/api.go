package api

import (
	"api_service/api/handler"
	"api_service/api/middleware"
	"api_service/config"

	_ "api_service/api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ReserveDesk API
// @version 1.0
// @description ReserveDesk is program to book seats in restaurants order food before arrival.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /reservedesk.uz

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @schemes http

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @security [ApiKeyAuth]
func NewRouter(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/reservedesk.uz")
	api.Use(middleware.JWTMiddleware())

	h := handler.NewHandler(cfg)

	//product api
	product := api.Group("/products")
	product.POST("/", h.CreateProduct)
	product.PUT("/:id", h.UpdateProduct)
	product.DELETE("/:id", h.DeleteProduct)
	product.GET("/aa", h.ListProducts)
	product.GET("/:id", h.GetProduct)
	product.GET("/search", h.SearchProducts)
	product.POST("/rating/:product_id", h.AddProductRating)
	product.GET("ratings/:product_id", h.GetProductRatings)

	// auth api
	auth := api.Group("/profiles")
	auth.PUT("/:id", h.UpdateProfile)
	auth.DELETE("/:id", h.DeleteOrder)
	auth.GET("/:id", h.GetByIdProfile)
	auth.GET("/", h.GetAllProfil)

	// order api
	order := api.Group("/orders")
	order.PUT("/:id", h.UpdateOrder)
	order.DELETE("/:id", h.DeleteOrder)
	order.POST("/", h.CreateOrder)
	order.GET("/{page}/{limit}", h.ListOrders)
	order.GET("/:id", h.GetOrder)
	order.GET("payment/status/:order_id", h.GetPaymentStatus)

	return router
}
