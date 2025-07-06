package routes

import (
	"product/cmd/product/handler"
	"product/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, productHandler handler.ProductHandler) {
	router.Use(middleware.RequestLogger(5))
	router.POST("/v1/product", productHandler.ProductManagement)
	router.POST("/v1/product-category", productHandler.ProductCategoryManagement)

	router.GET("/v1/product/:id", productHandler.GetProductByID)
	router.GET("/v1/product-category/:id", productHandler.GetProductCategoryByID)

	router.GET("v1/product/search", productHandler.SearchProduct)
}
