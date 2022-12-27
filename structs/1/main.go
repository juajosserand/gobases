package main

import (
	"errors"
	"fmt"
)

func main() {
	products.GetAll()

	p, err := getById(3)
	if err != nil {
		panic(err)
	}
	fmt.Println("getById(3) =", p)

	p4 := Product{
		4,
		"product4",
		1000.0,
		"product4 desc",
		"product4 cat",
	}
	p4.Save()

	products.GetAll()
}

var products = Products{
	{
		ID:          1,
		Name:        "product1",
		Price:       250.0,
		Description: "product1 desc",
		Category:    "product1 cat",
	},
	{
		ID:          2,
		Name:        "product2",
		Price:       1250.0,
		Description: "product2 desc",
		Category:    "product2 cat",
	},
	{
		ID:          3,
		Name:        "product3",
		Price:       2250.0,
		Description: "product3 desc",
		Category:    "product3 cat",
	},
}

type Products []Product

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

func (p *Product) Save() {
	products = append(products, *p)
}

func (ps Products) GetAll() {
	for _, p := range ps {
		fmt.Println(p)
	}
}

func getById(id int) (product *Product, err error) {
	if id <= 0 {
		return product, errors.New("invalid negative id")
	}

	for _, p := range products {
		if p.ID == id {
			product = &p
			break
		}
	}

	return product, nil
}
