package main

import (
	"log"
	"main/cmd/server/handler"
	docs "main/docs"
	"main/internal/products"
	"main/middleware"
	"main/pkg/store"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	store := store.Factory("file", "products.json")
	if store == nil {
		log.Fatal("Não foi possivel criar a store")
	}
	repo := products.NewRepository(store)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	products := r.Group("/products")
	products.Use(middleware.TokenAuthMiddleware())
	{
		products.POST("/", p.Insert())
		products.GET("/", p.GetAll())
		products.PUT("/:id", p.Update())
		products.PATCH("/:id", p.UpdateNameAndPrice())
		products.DELETE("/:id", p.Delete())

	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8081")

}
