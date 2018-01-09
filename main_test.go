package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestMain(t *testing.T) {
	g := Goblin(t)

	g.Describe("FormatCommandDescription", func() {
		g.It("Works as expected", func() {
			g.Assert(FormatCommandDescription("cmd", "desc")).Equal("cmd\tdesc")
		})
	})
}
