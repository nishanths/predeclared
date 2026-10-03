package rune

import (
	"test/miscpkgs/comparable"        // want "^comparable: shadows predeclared identifier$" "^comparable: same name as predeclared identifier$"
	_ "test/miscpkgs/comparable"      // OK
	cmprbl "test/miscpkgs/comparable" // OK
	clear "test/miscpkgs/f"           // want "^clear: shadows predeclared identifier$" "^clear: same name as predeclared identifier$"
)

const complex128 = 0 // want "^complex128: shadows predeclared identifier$" "^complex128: same name as predeclared identifier$"

type (
	complex struct{} // want "^complex: shadows predeclared identifier$" "^complex: same name as predeclared identifier$"
)

var nil = func() {}        // want "^nil: shadows predeclared identifier$" "^nil: same name as predeclared identifier$"
func copy(dst, src string) {} // want "^copy: shadows predeclared identifier$" "^copy: same name as predeclared identifier$"

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

type t1 byte

func (byte *t1) m1() {}
func (t *t1) byte()  {}

// suppress "imported and not used" compile errors
var _ = clear.X
var _ = comparable.X
var _ = cmprbl.X
