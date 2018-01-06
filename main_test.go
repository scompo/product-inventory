package main

import (
	"testing"
)

func TestValue(t *testing.T) {
	p := NewProduct("present", 10.00, 20)
	if p.Value() != 200.00 {
		t.Errorf("wrong value")
	}
}

func TestPresentWithNewElement(t *testing.T) {
	i := EmptyInventory()
	present := i.Present("not present")
	if present {
		t.Errorf("present and should not!")
	}
}

func TestStatusWithExistentElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	i.Add(p)

	p, err := i.Status(p.Id)

	if err != nil {
		t.Errorf("Why an error?")
	}

	if p.Id != "present" {
		t.Errorf("id not correct")
	}
	if p.Price != 10.00 {
		t.Errorf("price not correct")
	}
	if p.Quantity != 20 {
		t.Errorf("quantity not correct")
	}
}

func TestStatusWithNonExistentElement(t *testing.T) {
	i := EmptyInventory()
	p, err := i.Status("not present")

	if err != nil {
		if err.Error() != "missing product: not present" {
			t.Errorf("expected a descriptive error message")
		}
		if p != nil {
			t.Errorf("expected nil value")
		}
	} else {
		t.Errorf("Expected an error")
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

func TestAddExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 10.00, 20)

	i.Add(p)

	added, err := i.Add(p2)
	if err != nil {
		if err.Error() != "already present: present" {
			t.Errorf("expected a descriptive error message")
		}
		if added != nil {
			t.Errorf("expected nil value")
		}
	} else {
		t.Errorf("Expected an error")
	}
}

func TestAddNotExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)

	added, err := i.Add(p)

	if err != nil {
		t.Errorf("Why an error?")
	}
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

func TestUpdateElementNotPresent(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("not present", 10.00, 20)

	added, err := i.Update(p)

	if err != nil {
		if err.Error() != "missing product: not present" {
			t.Errorf("expected a descriptive error message")
		}
		if added != nil {
			t.Errorf("expected nil value")
		}
	} else {
		t.Errorf("Expected an error")
	}
}

func TestUpdateWithExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 20.20, 10)

	i.Add(p)

	updated, err := i.Update(p2)

	if err != nil {
		t.Errorf("Why an error?")
	}

	if updated.Id != "present" {
		t.Errorf("id not correct")
	}
	if updated.Price != 20.20 {
		t.Errorf("price not correct")
	}
	if updated.Quantity != 30 {
		t.Errorf("quantity not correct")
	}
}

func TestUpdateWithExistingElementSubtracting(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 10.00, -20)

	i.Add(p)

	updated, err := i.Update(p2)

	if err != nil {
		t.Errorf("Why an error?")
	}

	if updated.Id != "present" {
		t.Errorf("id not correct")
	}
	if updated.Price != 10.00 {
		t.Errorf("price not correct")
	}
	if updated.Quantity != 0 {
		t.Errorf("quantity not correct")
	}
}

func TestInventoryValue(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("p1", 10.00, 20)
	p2 := NewProduct("p2", 10.00, 30)

	i.Add(p)
	i.Add(p2)

	if i.Value() != 500.00 {
		t.Errorf("wrong value")
	}
}
