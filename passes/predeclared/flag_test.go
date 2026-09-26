package predeclared

import (
	"slices"
	"testing"
)

func TestParseMode(t *testing.T) {
	type testcase struct {
		input string
		modes []mode
		err   string
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
		switch {
		case err == nil && tt.err != "":
			t.Errorf("parseMode(%q): got nil error, want: %q", tt.input, tt.err)
		case err != nil && tt.err == "":
			t.Errorf("parseMode(%q): got error %q, want nil", tt.input, err)
		case err != nil && err.Error() != tt.err:
			t.Errorf("parseMode(%q): got error %q, want %q", tt.input, err, tt.err)
		case !slices.Equal(modes, tt.modes):
			t.Errorf("parseMode(%q): got: %v, want: %v", tt.input, modes, tt.modes)
		}
	}
}
