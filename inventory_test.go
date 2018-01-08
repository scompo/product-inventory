package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestInventory(t *testing.T) {
	g := Goblin(t)

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
				r, err := i.Status(existent.ID)

				g.It("Returns the correct Product", func() {
					g.Assert(r).Equal(existent)
				})
				g.It("Returns no errors", func() {
					g.Assert(err).Equal(nil)
				})
			})

			g.Describe("Requesting a non existent Product", func() {
				r, err := i.Status("not existent")

				g.It("Returns nil as Product", func() {
					g.Assert(r == nil).IsTrue()
				})
				g.It("Returns the correct error", func() {
					g.Assert(err.Error()).Equal("missing product: not existent")
				})
			})
		})

		g.Describe("Add", func() {
			i := EmptyInventory()
			existent := NewProduct("present", 10.00, 20)
			i.Add(existent)

			g.Describe("Adding an existent Product", func() {
				p2 := NewProduct("present", 10.00, 20)
				r, err := i.Add(p2)

				g.It("Returns an error", func() {
					g.Assert(err.Error()).Equal("Already present: present")
				})
				g.It("Returns nil as status", func() {
					g.Assert(r == nil).IsTrue()
				})
			})

			g.Describe("Adding a non existent Product", func() {
				p2 := NewProduct("another", 10.00, 20)
				r, err := i.Add(p2)

				g.It("Does not return an error", func() {
					g.Assert(err == nil).IsTrue()
				})
				g.It("Returns the updated status of the added element", func() {
					g.Assert(r).Equal(p2)
				})
				g.It("Does not return an error", func() {
					g.Assert(err == nil).IsTrue()
				})
			})
		})

		g.Describe("Update", func() {
			i := EmptyInventory()
			existent := NewProduct("present", 10.00, 20)
			i.Add(existent)

			g.Describe("With an existent Product", func() {
				p2 := NewProduct("present", 20.00, 20)
				r, err := i.Update(p2)

				g.It("Returns the correct Product ID", func() {
					g.Assert(r.ID).Equal(existent.ID)
				})
				g.It("Returns the correct Product price", func() {
					g.Assert(r.Price).Equal(p2.Price)
				})
				g.It("Returns the correct Product quantity", func() {
					g.Assert(r.Quantity).Equal(40.00)
				})
				g.It("Returns no errors", func() {
					g.Assert(err).Equal(nil)
				})
				g.It("Subtract a quantity if passed a negative number", func() {
					r2, _ := i.Update(NewProduct("present", 20.00, -20))
					g.Assert(r2.Quantity).Equal(20.00)
				})
				g.It("Works as well with floats", func() {
					r2, _ := i.Update(NewProduct("present", 20.00, -0.5))
					g.Assert(r2.Quantity).Equal(19.5)
				})
			})

			g.Describe("With a non existent Product", func() {
				p2 := NewProduct("not existent", 20.00, 20)
				r, err := i.Update(p2)

				g.It("Returns nil as Product", func() {
					g.Assert(r == nil).IsTrue()
				})
				g.It("Returns the correct error", func() {
					g.Assert(err.Error()).Equal("Missing product: not existent")
				})
			})
		})

		g.Describe("Value", func() {
			g.It("Works as expected for an existing Product", func() {
				i := EmptyInventory()
				i.Add(NewProduct("p1", 1.00, 20))
				i.Add(NewProduct("p2", 50.00, 0.5))
				i.Add(NewProduct("p3", 3.00, 3))

				g.Assert(i.Value()).Equal(54.00)
			})
			g.It("Returns 0 for an empty Inventory", func() {
				i := EmptyInventory()
				g.Assert(i.Value()).Equal(0.00)
			})
		})
	})
}
