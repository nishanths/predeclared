package predeclared

import (
	"slices"
	"testing"
)

func TestParseMode(t *testing.T) {
	type testcase struct {
		input string
		modes []mode
		error string
	}

	testcases := []testcase{
		{"", nil, "empty string"},
		{"d", []mode{'d'}, ""},
		{"s", []mode{'s'}, ""},
		{"ds", []mode{'d', 's'}, ""},
		{"x", nil, "invalid letter x"},
		{"dx", nil, "invalid letter x"},
		{"ssddsd", []mode{'d', 's'}, ""},
	}

	for _, tt := range testcases {
		modes, err := parseMode(tt.input)
		if err != nil && err.Error() != tt.error {
			t.Fatalf("parseMode(%q): got error: %q, want: %q", tt.input, err, tt.error)
		}
		if err == nil && tt.error != "" {
			t.Fatalf("parseMode(%q): got nil error, want: %q", tt.input, tt.error)
		}
		if !slices.Equal(modes, tt.modes) {
			t.Fatalf("parseMode(%q): got: %v, want: %v", tt.input, modes, tt.modes)
		}
	}
}
