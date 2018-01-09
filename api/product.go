package api

// Product in the inventory.
type Product struct {
	ID       string
	Price    float64
	Quantity float64
}

// Value returns the value of a Product (price * quantity).
func (p Product) Value() float64 {
	return p.Price * p.Quantity
}

// NewProduct creates a new Product.
func NewProduct(id string, price float64, qt float64) *Product {
	p := new(Product)
	p.ID = id
	p.Price = price
	p.Quantity = qt
	return p
}
