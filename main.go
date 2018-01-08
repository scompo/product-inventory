package main

import (
	"bufio"
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

func printStatus(p *Product) {
	fmt.Printf("Product id: %s\n", p.ID)
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

func readProductID(rc *bufio.Scanner) string {
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
			id := readProductID(rc)
			st, err := i.Status(id)
			if err == nil {
				printStatus(st)
			} else {
				fmt.Printf("%s\n", err.Error())
			}
		} else if option == insert {
			id := readProductID(rc)
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
			id := readProductID(rc)
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
