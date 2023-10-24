package main

import (
	"log"
	"net/http"
	"os"

	"github.com/casikenegro/golang-gin-api/controller"
	"github.com/casikenegro/golang-gin-api/database"
	_ "github.com/casikenegro/golang-gin-api/docs"
	"github.com/casikenegro/golang-gin-api/repository"
	"github.com/casikenegro/golang-gin-api/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	productRepository repository.ProductRepository = repository.NewProductRepository()
	productService    services.ProductService      = services.New(productRepository)
	productController controller.ProductController = controller.New(productService)
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v\n", err)
	}
	database.DBConnection(&database.ConfigDB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
	})
}

// @title  Tag Service API
// @version 1.0
// @host 	localhost:3000
func main() {
	server := gin.Default()

	server.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, productController.FindAll())
	})
	server.POST("/products", func(ctx *gin.Context) {
		if err := productController.Save(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}

	})
	server.PUT("/products/:id", func(ctx *gin.Context) {
		if err := productController.Update(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}

	})
	server.DELETE("/products/:id", func(ctx *gin.Context) {
		if err := productController.Delete(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}

	})
	server.GET("/products/:sku", func(ctx *gin.Context) {
		if product := productController.FindOne(ctx); product == nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "product not exist",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"product": product,
			})
		}
	})
	server.POST("/products/:id/image", func(ctx *gin.Context) {
		if err := productController.AddImage(ctx); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		}

	})

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	server.Run(":3000")
}
