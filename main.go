package main

import (
	"errors"
	"fmt"
)

// Product in the inventory.
type Product struct {
	Id       string
	Price    float64
	Quantity int64
}

func (p Product) Value() float64 {
	return p.Price * float64(p.Quantity)
}

// An inventory of products.
type Inventory struct {
	data map[string]*Product
}

// Creates an empty Inventory
func EmptyInventory() Inventory {
	return Inventory{data: map[string]*Product{}}
}

// Creates a new Product
func NewProduct(id string, price float64, qt int64) *Product {
	p := new(Product)
	p.Id = id
	p.Price = price
	p.Quantity = qt
	return p
}

// Returns true if the id of the product is present, false otherwise.
func (i Inventory) Present(id string) bool {
	_, present := i.data[id]
	return present
}

// Returns the current status of a Product if present, otherwise an error.
func (i Inventory) Status(id string) (*Product, error) {
	if !i.Present(id) {
		return nil, errors.New("missing product: " + id)
	}
	return i.data[id], nil
}

// Adds a new Product.
func (i *Inventory) Add(p *Product) (*Product, error) {
	if i.Present(p.Id) {
		return nil, errors.New("already present: " + p.Id)
	} else {
		i.data[p.Id] = p
	}
	return i.data[p.Id], nil
}

// Updates the status of an existing Product.
func (i *Inventory) Update(p *Product) (*Product, error) {
	if !i.Present(p.Id) {
		return nil, errors.New("missing product: " + p.Id)
	} else {
		i.data[p.Id].Price = p.Price
		i.data[p.Id].Quantity = i.data[p.Id].Quantity + p.Quantity
	}
	return i.data[p.Id], nil
}

func main() {
	fmt.Println("hello!")
}
