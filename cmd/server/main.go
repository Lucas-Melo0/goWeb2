package main

import (
	"log"
	"main/cmd/server/handler"
	"main/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	products := r.Group("/products")
	{
		products.POST("/", p.Insert())
		products.GET("/", p.GetAll())
		products.PUT("/:id", p.Update())
		products.PATCH("/:id", p.UpdateNameAndPrice())
		products.DELETE("/:id", p.Delete())

	}
	r.Run()
}
