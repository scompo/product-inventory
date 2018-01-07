package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	QUIT   = "quit"
	VALUE  = "value"
	STATUS = "status"
	INSERT = "insert"
	UPDATE = "update"
)

// Product in the inventory.
type Product struct {
	Id       string
	Price    float64
	Quantity int64
}

// Returns the value of a Product.
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

// Returns the current value the Inventory.
func (i Inventory) Value() float64 {
	tot := 0.00
	for _, v := range i.data {
		tot = tot + v.Value()
	}
	return tot
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

func printStatus(p *Product) {
	fmt.Printf("Product id: %s\n", p.Id)
	fmt.Printf("Product price: %f\n", p.Price)
	fmt.Printf("Product quantity: %d\n", p.Quantity)
	fmt.Printf("Product value: %f\n", p.Value())
}

func printMenu() {
	fmt.Printf("Available operations:\n")
	fmt.Printf("%s\tExits\n", QUIT)
	fmt.Printf("%s\tPrints the inventory current value\n", VALUE)
	fmt.Printf("%s\tReturns the status of a product\n", STATUS)
	fmt.Printf("%s\tInserts a new product\n", INSERT)
	fmt.Printf("%s\tUpdates an existing product\n", UPDATE)
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

	fmt.Printf("Insert Product price: ")
	rc.Scan()
	temp = rc.Text()
	price, _ := strconv.ParseFloat(temp, 64)

	fmt.Printf("Insert Product quantity: ")
	rc.Scan()
	temp = rc.Text()
	qt, _ := strconv.ParseInt(temp, 10, 64)

	return NewProduct(id, price, qt)
}

func main() {
	i := EmptyInventory()
	rc := bufio.NewScanner(os.Stdin)

	var option string

	for option != QUIT {

		printHeader()
		printMenu()
		fmt.Printf("Select operation: ")

		rc.Scan()
		option = rc.Text()

		if option == VALUE {
			printValue(i)
		} else if option == STATUS {
			id := readProductId(rc)
			st, err := i.Status(id)
			if err == nil {
				printStatus(st)
			} else {
				fmt.Printf("%s\n", err.Error())
			}
		} else if option == INSERT {
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
		} else if option == UPDATE {
			id := readProductId(rc)
			if !i.Present(id) {
				fmt.Printf("Product id already present: %s\n", id)
			} else {
				p := readProduct(rc, id)
				p2, err := i.Update(p)
				if err != nil {
					fmt.Printf("%s\n", err.Error())
				} else {
					printStatus(p2)
				}
			}
		} else if option == QUIT {
			fmt.Println("Exiting.")
		} else {
			fmt.Println("Operation not recognized.")
		}
	}
}
