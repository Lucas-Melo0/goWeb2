package products

import (
	"errors"
	"fmt"
	"main/pkg/store"
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
	UpdateNameAndPrice(id int, name string, price int) (Product, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func (r *repository) GetAll() ([]Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return []Product{}, err
	}
	return products, nil

}
func (r *repository) Insert(id int, name string, color string, price int, stock int, code string, isPublicated bool, creationDate string) (Product, error) {
	var products []Product
	err := r.db.Read(&products)
	if err != nil {
		return Product{}, err
	}
	for _, v := range products {
		if v.ID == id {
			return Product{}, errors.New("this id already exists")
		}
	}
	p := Product{id, name, color, price, stock, code, isPublicated, creationDate}
	products = append(products, p)
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
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

func (r *repository) UpdateNameAndPrice(id int, name string, price int) (Product, error) {
	updated := false
	var p Product
	for i := range ps {
		if ps[i].ID == id {
			ps[i].Name = name
			ps[i].Price = price
			p = ps[i]
			updated = true
		}
	}
	if !updated {
		return Product{}, fmt.Errorf("product %d not found", id)
	}
	return p, nil
}
func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].ID == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("product %d not found", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil
}

func (r *repository) LastId() (int, error) {
	return oldId, nil
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}
