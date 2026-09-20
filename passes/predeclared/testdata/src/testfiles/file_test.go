package rune

import "testing"

func TestF1(t *testing.T) {
	type println func() // want "^println: shadows predeclared identifier$" "^println: same name as predeclared identifier$"
}
