The predeclared static analysis finds declarations in Go source
code that shadow any of Go's predeclared identifiers.

The analysis and available flags are documented in the package
comment of package predeclared in the 'passes/predeclared'
directory.

<https://pkg.go.dev/github.com/nishanths/predeclared/passes/predeclared>

The list of predeclared identifiers can be found in the language
specification. The predeclared identifiers, as of go1.26, are
listed below for reference.

<https://golang.org/ref/spec#Predeclared_identifiers>

	Types:
	    any bool byte comparable
	    complex64 complex128 error float32 float64
	    int int8 int16 int32 int64 rune string
	    uint uint8 uint16 uint32 uint64 uintptr

	Constants:
	    true false iota

	Zero value:
	    nil

	Functions:
	    append cap clear close complex copy delete imag len
	    make max min new panic print println real recover

# Usage

The standalone predeclared command can be installed using 'go install'.

	go install github.com/nishanths/predeclared@latest

The Go package in the 'passes/predeclared' directory provides an
analysis.Analyzer value that can be used by analysis driver
programs. See <https://golang.org/x/tools/go/analysis> for more
details.

# Examples

For the source file:

	package example

	import "strconv"

	func clear[S ~[]E, E any](s S) S { ... }

	func baz(s string) {
		const max = 8192
		int, err := strconv.Atoi(s)
		...
	}

	func copy(dst, src string) { ... }

the predeclared analysis reports these diagnostics:

	/tmp/example.go:5:6: clear: shadows predeclared identifier
	/tmp/example.go:8:8: max: shadows predeclared identifier
	/tmp/example.go:9:2: int: shadows predeclared identifier
	/tmp/example.go:14:6: copy: shadows predeclared identifier
