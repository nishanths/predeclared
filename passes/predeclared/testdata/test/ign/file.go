package rune

func copy(dst, src string) {} // want "^copy: shadows predeclared identifier$" "^copy: same name as predeclared identifier$"

type (
	complex struct{}
	nil     = func() // want "^nil: shadows predeclared identifier$" "^nil: same name as predeclared identifier$"
)

func f1() {
	const rune = 0 // want "^rune: same name as predeclared identifier$" "^rune: shadows predeclared identifier$"
	const (
		int64 = 8 // want "^int64: shadows predeclared identifier$" "^int64: same name as predeclared identifier$"
		max   = 4096
	)
	var (
		len, cap = 16, 32                   // want "^cap: shadows predeclared identifier$" "^cap: same name as predeclared identifier$"
		print    = func(any interface{}) {} // want "^any: shadows predeclared identifier$" "^any: same name as predeclared identifier$"
	)

	_, _ = len, cap
	_ = print
}
