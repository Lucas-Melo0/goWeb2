package main

import (
	"main/cmd/server/handler"
	"main/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	products := r.Group("/products")
	{
		products.POST("/", p.Insert())
		products.GET("/", p.GetAll())
		products.PUT("/:id", p.Update())

	}
	r.Run()
}
