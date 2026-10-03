package a

import "strconv"

func clear[S ~[]E, E any](s S) S { /* ... */ return nil } // want "^clear: shadows predeclared identifier$"

func baz(s string) {
	const max = 8192            // want "^max: shadows predeclared identifier$"
	int, err := strconv.Atoi(s) // want "^int: shadows predeclared identifier$"
	/* ... */
	_, _ = int, err
}

func copy(dst, src string) { /* ... */ } // want "^copy: shadows predeclared identifier$"
