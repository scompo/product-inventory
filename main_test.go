package main

import (
	"fmt"
	"testing"
)

func TestPresentWithNewElement(t *testing.T) {
	i := EmptyInventory()
	present := i.Present("not present")
	if present {
		t.Errorf("present and should not!")
	}
}

func TestPresentWithExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	i.Add(p)
	present := i.Present("present")
	if !present {
		t.Errorf("not present and should be!")
	}
}

func TestAddNewElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	added := i.Add(p)
	if added.Id != "present" {
		t.Errorf("id not correct")
	}
	if added.Price != 10.00 {
		t.Errorf("price not correct")
	}
	if added.Quantity != 20 {
		t.Errorf("quantity not correct")
	}
}

func TestAddWithExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 20.20, 10)
	i.Add(p)
	added := i.Add(p2)

	if added.Id != "present" {
		t.Errorf("id not correct")
	}
	if added.Price != 20.20 {
		t.Errorf("price not correct")
	}
	if added.Quantity != 30 {
		t.Errorf("quantity not correct")
	}
}

func TestAddWithExistingElementSubtracting(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 10.00, -20)
	i.Add(p)
	added := i.Add(p2)

	if added.Id != "present" {
		t.Errorf("id not correct")
	}
	if added.Price != 10.00 {
		t.Errorf("price not correct")
	}
	if added.Quantity != 0 {
		t.Errorf("quantity not correct")
	}
}
