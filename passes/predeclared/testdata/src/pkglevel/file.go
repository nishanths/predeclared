package rune

import (
	"example.org/comparable"        // want "^comparable: shadows predeclared identifier$"
	_ "example.org/comparable"      // OK
	cmprbl "example.org/comparable" // OK
	clear "example.org/f"           // want "^clear: shadows predeclared identifier$"
)

func copy(dst, src string) {} // want "^copy: shadows predeclared identifier$"

type (
	complex struct{} // want "^complex: shadows predeclared identifier$"
	nil     = func() // want "^nil: shadows predeclared identifier$"
)

func f1() {
	const rune = 0
	const (
		int64 = 8
		max   = 4096
	)
	var (
		len, cap = 16, 32
		print    = func(any interface{}) {}
	)

	_, _ = len, cap
	_ = print
}

// suppress "imported and not used" compile errors
var _ = clear.X
var _ = comparable.X
var _ = cmprbl.X
