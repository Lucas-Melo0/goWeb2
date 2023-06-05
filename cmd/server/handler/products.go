package handler

import (
	"main/internal/products"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type request struct {
	Name         string `json:"nome"`
	Color        string `json:"cor"`
	Price        int    `json:"preco"`
	Stock        int    `json:"estoque"`
	Code         string `json:"codigo"`
	IsPublicated bool   `json:"publicacao"`
	CreationDate string `json:"data_de_criacao"`
}
type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{service: p}
}
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, p)
	}
}

func (p *Product) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := p.service.Insert(req.Name, req.Color, req.Price, req.Stock, req.Code, req.IsPublicated, req.CreationDate)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, p)
	}
}
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			c.JSON(400, gin.H{"error": "O nome do produto é obrigatório"})
			return
		}
		if req.Color == "" {
			c.JSON(400, gin.H{"error": "A cor do produto é obrigatório"})
			return
		}
		if req.Stock == 0 {
			c.JSON(400, gin.H{"error": "A quantidade em estoque é obrigatória"})
			return
		}
		if req.Price == 0 {
			c.JSON(400, gin.H{"error": "O preço é obrigatório"})
			return
		}
		if req.Code == "" {
			c.JSON(400, gin.H{"error": "O código é obrigatório"})
			return
		}
		if req.CreationDate == "" {
			c.JSON(400, gin.H{"error": "A data de criação é obrigatória"})
			return
		}
		p, err := p.service.Update(int(id), req.Name, req.Color, int(req.Price), int(req.Stock), req.Code, bool(req.IsPublicated), req.CreationDate)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}
func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" && req.Price == 0 {
			c.JSON(400, gin.H{"error": "O preco ou o nome tem que ser atualizados"})
			return
		}
		p, err := p.service.UpdateNameAndPrice(int(id), req.Name, int(req.Price))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, p)
	}
}
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123" {
			c.JSON(401, gin.H{"error": "token inválido"})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid ID"})
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, p)
	}
}
