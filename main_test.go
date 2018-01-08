package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)

	g.Describe("Product", func() {
		g.Describe("Value", func() {
			g.It("Returns the correct value of a product", func() {
				g.Assert(NewProduct("present", 10.00, 20).Value()).Equal(200.00)
			})
		})
	})

	g.Describe("Inventory", func() {

		g.Describe("Present", func() {

			i := EmptyInventory()
			i.Add(NewProduct("present", 10.00, 20))

			g.It("Returns true for an existing Product", func() {
				g.Assert(i.Present("present")).Equal(true)
			})
			g.It("Returns false for a non existing Product", func() {
				g.Assert(i.Present("not present")).Equal(false)
			})
		})

		g.Describe("Status", func() {

			i := EmptyInventory()
			existent := NewProduct("present", 10.00, 20)
			i.Add(existent)

			g.Describe("Requesting an existent Product", func() {
				g.It("Returns the correct Product", func() {
					r, _ := i.Status(existent.ID)
					g.Assert(r).Equal(existent)
				})
				g.It("Returns no errors", func() {
					_, err := i.Status(existent.ID)
					g.Assert(err).Equal(nil)
				})
			})

			g.Describe("Requesting a non existent Product", func() {
				g.It("Returns nil as Product", func() {
					r, _ := i.Status("not existent")
					g.Assert(r == nil).IsTrue()
				})
				g.It("Returns the correct error", func() {
					_, err := i.Status("not existent")
					g.Assert(err.Error()).Equal("missing product: not existent")
				})
			})
		})
	})
}

func TestAddExistingElement(t *testing.T) {
	i := EmptyInventory()
	p := NewProduct("present", 10.00, 20)
	p2 := NewProduct("present", 10.00, 20)

	i.Add(p)

	added, err := i.Add(p2)
	if err != nil {
		if err.Error() != "Already present: present" {
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
	if added.ID != "present" {
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
		if err.Error() != "Missing product: not present" {
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

	if updated.ID != "present" {
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

	if updated.ID != "present" {
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
