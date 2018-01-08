package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestProduct(t *testing.T) {
	g := Goblin(t)

	g.Describe("Product", func() {
		g.Describe("Value", func() {
			g.It("Returns the correct value of a product", func() {
				g.Assert(NewProduct("present", 10.00, 20).Value()).Equal(200.00)
			})
		})
	})
}
