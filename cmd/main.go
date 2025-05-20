package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	productRepository := repository.NewProductRepository(dbConnection)
	ProductUsecase := usecase.NewProductUseCase(productRepository)
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
