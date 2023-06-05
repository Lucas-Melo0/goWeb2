package products

import (
	"errors"
	"fmt"
)

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"nome"`
	Color        string `json:"cor"`
	Price        int    `json:"preco"`
	Stock        int    `json:"estoque"`
	Code         string `json:"codigo"`
	IsPublicated bool   `json:"publicacao"`
	CreationDate string `json:"data_de_criacao"`
}

var ps []Product
var oldId int

type Repository interface {
	GetAll() ([]Product, error)
	Insert(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error)
	Update(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error)
	LastId() (int, error)
}

type repository struct{}

func (r *repository) GetAll() ([]Product, error) {
	return ps, nil

}
func (r *repository) Insert(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error) {
	for _, v := range ps {
		if v.ID == id {
			return Product{}, errors.New("this id already exists")
		}
	}
	p := Product{id, name, color, price, stock, code, isPublicated, creationDate}
	ps = append(ps, p)
	oldId = p.ID
	return p, nil
}
func (r *repository) Update(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error) {
	p := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, IsPublicated: isPublicated, CreationDate: creationDate}
	updated := false
	for i := range ps {
		if ps[i].ID == id {
			p.ID = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return p, nil
}

func (r *repository) LastId() (int, error) {
	return oldId, nil
}

func NewRepository() Repository {
	return &repository{}
}
