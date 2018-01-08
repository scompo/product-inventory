package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	quit   = "quit"
	value  = "value"
	status = "status"
	insert = "insert"
	update = "update"
)

// Product in the inventory.
type Product struct {
	Id       string
	Price    float64
	Quantity float64
}

// Value returns the value of a Product (price * quantity).
func (p Product) Value() float64 {
	return p.Price * p.Quantity
}

// Inventory represents an inventory of products.
type Inventory struct {
	data map[string]*Product
}

// EmptyInventory creates an empty Inventory.
func EmptyInventory() Inventory {
	return Inventory{data: map[string]*Product{}}
}

// NewProduct creates a new Product.
func NewProduct(id string, price float64, qt float64) *Product {
	p := new(Product)
	p.Id = id
	p.Price = price
	p.Quantity = qt
	return p
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
	if i.Present(p.Id) {
		return nil, errors.New("Already present: " + p.Id)
	}
	i.data[p.Id] = p
	return i.data[p.Id], nil
}

// Update updates the status of an existing Product.
// If the product is already present it returns an error.
func (i *Inventory) Update(p *Product) (*Product, error) {
	if !i.Present(p.Id) {
		return nil, errors.New("Missing product: " + p.Id)
	}
	i.data[p.Id].Price = p.Price
	i.data[p.Id].Quantity = i.data[p.Id].Quantity + p.Quantity
	return i.data[p.Id], nil
}

func printStatus(p *Product) {
	fmt.Printf("Product id: %s\n", p.Id)
	fmt.Printf("Product price: %f\n", p.Price)
	fmt.Printf("Product quantity: %f\n", p.Quantity)
	fmt.Printf("Product value: %f\n", p.Value())
}

func printMenu() {
	fmt.Printf("Available operations:\n")
	fmt.Printf("%s\tExits\n", quit)
	fmt.Printf("%s\tPrints the inventory current value\n", value)
	fmt.Printf("%s\tReturns the status of a product\n", status)
	fmt.Printf("%s\tInserts a new product\n", insert)
	fmt.Printf("%s\tUpdates an existing product\n", update)
}

func printHeader() {
	fmt.Println("=================")
	fmt.Println("Product Inventory")
	fmt.Println("=================")
}

func printValue(i Inventory) {
	inventoryValue := i.Value()
	fmt.Printf("Inventory value: %.2f€\n", inventoryValue)
}

func readProductId(rc *bufio.Scanner) string {
	fmt.Printf("Insert product id: ")
	rc.Scan()
	return rc.Text()
}

func readProduct(rc *bufio.Scanner, id string) *Product {
	var temp string

	fmt.Printf("Insert price: ")
	rc.Scan()
	temp = rc.Text()
	price, _ := strconv.ParseFloat(temp, 64)

	fmt.Printf("Insert quantity: ")
	rc.Scan()
	temp = rc.Text()
	qt, _ := strconv.ParseFloat(temp, 64)

	return NewProduct(id, price, qt)
}

func printFunction(f string) {
	fmt.Printf("Operation %s selected, executing...\n", f)
}

func main() {
	i := EmptyInventory()
	rc := bufio.NewScanner(os.Stdin)

	var option string

	for option != quit {

		printHeader()
		printMenu()
		fmt.Printf("Select operation: ")

		rc.Scan()
		option = rc.Text()

		printFunction(option)

		if option == value {
			printValue(i)
		} else if option == status {
			id := readProductId(rc)
			st, err := i.Status(id)
			if err == nil {
				printStatus(st)
			} else {
				fmt.Printf("%s\n", err.Error())
			}
		} else if option == insert {
			id := readProductId(rc)
			if i.Present(id) {
				fmt.Printf("Product id already present: %s\n", id)
			} else {
				p := readProduct(rc, id)
				p2, err := i.Add(p)
				if err != nil {
					fmt.Printf("%s\n", err.Error())
				} else {
					printStatus(p2)
				}
			}
		} else if option == update {
			id := readProductId(rc)
			if !i.Present(id) {
				fmt.Printf("Missing product: %s\n", id)
			} else {
				p := readProduct(rc, id)
				p2, err := i.Update(p)
				if err != nil {
					fmt.Printf("%s\n", err.Error())
				} else {
					printStatus(p2)
				}
			}
		} else if option == quit {
			fmt.Println("Exiting.")
		} else {
			fmt.Println("Operation not recognized.")
		}
	}
}
