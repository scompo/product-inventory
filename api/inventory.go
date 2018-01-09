package api

import "errors"

// Inventory represents an inventory of products.
type Inventory struct {
	data map[string]*Product
}

// EmptyInventory creates an empty Inventory.
func EmptyInventory() Inventory {
	return Inventory{data: map[string]*Product{}}
}

// Present returns true if the id of the product is present, false otherwise.
func (i Inventory) Present(id string) bool {
	_, present := i.data[id]
	return present
}

// Status returns the current status of a Product if present.
// If the the id of the product does not exist it returns an error.
func (i Inventory) Status(id string) (*Product, error) {
	if !i.Present(id) {
		return nil, errors.New("missing product: " + id)
	}
	return i.data[id], nil
}

// Value returns the current value the Inventory.
func (i Inventory) Value() float64 {
	tot := 0.00
	for _, v := range i.data {
		tot = tot + v.Value()
	}
	return tot
}

// Add adds a new Product to the Inventory.
// If the product is already present it returns an error.
func (i *Inventory) Add(p *Product) (*Product, error) {
	if i.Present(p.ID) {
		return nil, errors.New("Already present: " + p.ID)
	}
	i.data[p.ID] = p
	return i.data[p.ID], nil
}

// Update updates the status of an existing Product.
// If the product is already present it returns an error.
func (i *Inventory) Update(p *Product) (*Product, error) {
	if !i.Present(p.ID) {
		return nil, errors.New("Missing product: " + p.ID)
	}
	i.data[p.ID].Price = p.Price
	i.data[p.ID].Quantity = i.data[p.ID].Quantity + p.Quantity
	return i.data[p.ID], nil
}
