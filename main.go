package main

import (
	"product/cmd/product/handler"
	"product/cmd/product/repository"
	"product/cmd/product/resource"
	"product/cmd/product/service"
	"product/cmd/product/usecase"
	"product/config"
	"product/infrastructure/log"
	"product/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	cfg := config.LoadConfig()
	redis := resource.InitRedis(&cfg)
	db := resource.InitDb(&cfg)

	// logger
	log.SetupLogger()

	// init
	productRepository := repository.NewProductRepository(redis, db)
	productService := service.NewProductService(*productRepository)
	productUsecase := usecase.NewProductUsecase(*productService)
	productHandler := handler.NewProductHandler(*productUsecase)

	// gin
	port := cfg.App.Port
	router := gin.Default()
	routes.SetupRoutes(router, *productHandler)
	router.Run(":" + port)

	log.Logger.Printf("Server running on port: %s", port)

}
