package handler

import (
	"main/internal/products"
	"main/pkg/store/web"
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		p, err := p.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, err, ""))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

// InsertProducts godoc
// @Summary Insert products
// @Tags Products
// @Description insert products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to insert"
// @Success 200 {object} web.Response
// @Router /products [post]
func (p *Product) Insert() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, ""))
			return
		}
		p, err := p.service.Insert(req.Name, req.Color, req.Price, req.Stock, req.Code, req.IsPublicated, req.CreationDate)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(404, nil, ""))
			return
		}
		c.JSON(http.StatusCreated, web.NewResponse(200, p, ""))
	}
}

// UpdateProduct godoc
// @Summary Update a product
// @Tags Products
// @Description update product details
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/{id} [put]
func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(400, nil, ""))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(400, nil, ""))
			return
		}
		if req.Name == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.Color == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.Stock == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.Price == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.Code == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.CreationDate == "" {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		p, err := p.service.Update(int(id), req.Name, req.Color, int(req.Price), int(req.Stock), req.Code, bool(req.IsPublicated), req.CreationDate)
		if err != nil {
			c.JSON(http.StatusNotFound, web.NewResponse(400, nil, ""))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

// UpdateProductNameAndPrice godoc
// @Summary Update a product's name and price
// @Tags Products
// @Description update product's name and price
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Param product body request true "Product to update"
// @Success 200 {object} web.Response
// @Router /products/{id} [patch]
func (p *Product) UpdateNameAndPrice() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "bad request"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		if req.Name == "" && req.Price == 0 {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}
		p, err := p.service.UpdateNameAndPrice(int(id), req.Name, int(req.Price))
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, ""))
			return
		}
		c.JSON(http.StatusOK, web.NewResponse(200, p, ""))
	}
}

// DeleteProduct godoc
// @Summary Delete a product
// @Tags Products
// @Description delete a product
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Success 204 {object} web.Response
// @Router /products/{id} [delete]
func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, "bad request"))
			return
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, web.NewResponse(400, nil, ""))
			return
		}

		err = p.service.Delete(int(id))
		if err != nil {
			c.JSON(http.StatusInternalServerError, web.NewResponse(500, nil, ""))
			return
		}
		c.JSON(http.StatusNoContent, web.NewResponse(204, p, ""))
	}
}
