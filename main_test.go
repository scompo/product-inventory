package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestMain(t *testing.T) {
	g := Goblin(t)

	g.Describe("The ui should", func() {
		g.It("Do some stuff")
	})
}
