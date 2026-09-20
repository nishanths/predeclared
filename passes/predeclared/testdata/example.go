package example

import "strconv"

func clear[S ~[]E, E any](s S) S { /* ... */ return nil }

func baz(s string) {
	const max = 8192
	int, err := strconv.Atoi(s)
	/* ... */
	_, _ = int, err
}

func copy(dst, src string) { /* ... */ }
