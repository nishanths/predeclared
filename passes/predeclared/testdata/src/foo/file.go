// Note: No diagnostics should be reported for the package clause.
// See the Go spec for reference: "The package clause is not a
// declaration; the package name does not appear in any scope."
package rune

import (
	"math/rand/v2"

	"example.org/comparable"        // want "^comparable: shadows predeclared identifier$" "^comparable: same name as predeclared identifier$"
	_ "example.org/comparable"      // OK
	cmprbl "example.org/comparable" // OK
	clear "example.org/f"           // want "^clear: shadows predeclared identifier$" "^clear: same name as predeclared identifier$"
)

func copy(dst, src string) {} // want "^copy: shadows predeclared identifier$" "^copy: same name as predeclared identifier$"

type (
	complex struct{} // want "^complex: shadows predeclared identifier$" "^complex: same name as predeclared identifier$"
	nil     = func() // want "^nil: shadows predeclared identifier$" "^nil: same name as predeclared identifier$"
)

func f1() {
	const rune = 0 // want "^rune: same name as predeclared identifier$" "^rune: shadows predeclared identifier$"
	const (
		int64 = 8    // want "^int64: shadows predeclared identifier$" "^int64: same name as predeclared identifier$"
		max   = 4096 // want "^max: shadows predeclared identifier$" "^max: same name as predeclared identifier$"
	)
	var (
		len, cap = 16, 32                   // want "^len: shadows predeclared identifier$" "^cap: shadows predeclared identifier$" "^len: same name as predeclared identifier$" "^cap: same name as predeclared identifier$"
		print    = func(any interface{}) {} // want "^print: shadows predeclared identifier$" "^any: shadows predeclared identifier$" "^print: same name as predeclared identifier$" "^any: same name as predeclared identifier$"
	)

	_, _ = len, cap
	_ = print
}

func f2() {
	// No 'shadow' diagnostics should be reported for labels.
	// Labels do not conflict with identifiers that are not
	// labels and thus cannot shadow the predeclared
	// identifiers.

complex64: // want "^complex64: same name as predeclared identifier$"
	for {
	real: // want "^real: same name as predeclared identifier$"
		select {
		default:
			break real
		}
		const complex64 = real(1) // want "^complex64: shadows predeclared identifier" "^complex64: same name as predeclared identifier$"
		break complex64
	}

normal:
	for {
		const normal = 42
		break normal
	}
}

// multi-level shadow.
func f3() {
	// Note: No 'shadow' diagnostics. The name 'nil' exists
	// in ancestor scopes and is not the predeclared identifier.
	// See the earlier 'type nil ...' declaration.
	nil := 0 // want "^nil: same name as predeclared identifier$"

	if randeven() {
		const iota = 9000 // want "^iota: shadows predeclared identifier$" "^iota: same name as predeclared identifier$"
	}
	const iota = 9999 // want "^iota: shadows predeclared identifier$" "^iota: same name as predeclared identifier$"
	if randeven() {
		// Note: no 'shadow' diagnostic here.
		// The following declaration does not shadow the
		// predeclared identifier. It shadows the earlier
		// declaration in this function.
		const iota = 10000 // want "^iota: same name as predeclared identifier$"
	}

	if randeven() {
		new := 2 // want "^new: shadows predeclared identifier$" "^new: same name as predeclared identifier$"
		_ = new
	}
	new := 1 // want "^new: shadows predeclared identifier$" "^new: same name as predeclared identifier$"
	if randeven() {
		new := 3 // want "^new: same name as predeclared identifier$"
		new = 4
		_ = new
	}

	// suppress "declared not and used" compile errors.
	_ = nil
	_ = new
}

func f4[T any](n int, make func() T) {}                    // want "^make: shadows predeclared identifier$" "^make: same name as predeclared identifier$"
func f5[error int32]()               {}                    // want "^error: shadows predeclared identifier$" "^error: same name as predeclared identifier$"
func f6() (true bool)                { return randeven() } // want "^true: shadows predeclared identifier$" "^true: same name as predeclared identifier$"

// variable in type switch header.
func f7() {
	someUsage := func(v any) {}

	var v any
	switch imag := v.(type) { // want "^imag: shadows predeclared identifier$" "^imag: same name as predeclared identifier$"
	case int:
		someUsage(imag)
	default:
		someUsage(imag)
	}

	// Sanity check: No diagnostics should be reported.
	// The code is same as the above type switch except for
	// the variable name.
	switch vv := v.(type) {
	case int:
		someUsage(vv)
	default:
		someUsage(vv)
	}
}

// type alias.
func f8() {
	type iota = int64   // want "^iota: shadows predeclared identifier$" "^iota: same name as predeclared identifier$"
	type int32 = uint16 // want "^int32: shadows predeclared identifier$" "^int32: same name as predeclared identifier$"
	type byte = uint8   // want "^byte: shadows predeclared identifier$" "^byte: same name as predeclared identifier$"
}

type t1 byte

func (byte *t1) m1() {} // want "^byte: shadows predeclared identifier$" "^byte: same name as predeclared identifier$"
func (t *t1) byte()  {}

type _ struct {
	int8
	int16   uint16
	normal  int32
	recover interface{ uint() uint32 }
}

type i1 interface {
	uintptr
	~float64
	normal() uint32
	new() any
}

func randeven() bool { return rand.Int()%2 == 0 }

// suppress "imported and not used" compile errors
var _ = clear.X
var _ = comparable.X
var _ = cmprbl.X
