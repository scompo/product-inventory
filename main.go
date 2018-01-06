package main

import "fmt"

// Product in the inventory.
type Product struct {
	Id       string
	Price    float64
	Quantity int64
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

// Adds or updates the status of a Product.
func (i *Inventory) Add(p *Product) *Product {
	if i.Present(p.Id) {
		i.data[p.Id].Price = p.Price
		i.data[p.Id].Quantity = i.data[p.Id].Quantity + p.Quantity
	} else {
		i.data[p.Id] = p
	}
	return i.data[p.Id]
}

func main() {
	fmt.Println("hello!")
}
