package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/nishanths/predeclared/api"
)

var config *api.Config

func outPath(p string) string { return strings.TrimSuffix(p, ".go") + ".out" }

func equalBytes(t *testing.T, a, b []byte, normalize func([]byte) []byte) {
	t.Helper()
	if normalize != nil {
		a = normalize(a)
		b = normalize(b)
	}
	if !bytes.Equal(a, b) {
		t.Errorf(`bytes not equal
want: %s
got:  %s
`, a, b)
	}
}

func setupConfig(p string) {
	// Get the first line.
	b, err := ioutil.ReadFile(p)
	if err != nil {
		panic(fmt.Sprintf("failed to read file: %s", p))
	}
	idx := bytes.IndexByte(b, '\n')
	if idx == -1 {
		panic(fmt.Sprintf("no lines in file: %s", p))
	}
	// Does it have the prefix?
	const prefix = "//predeclared"
	line := string(b[:idx])
	if !strings.HasPrefix(line, prefix) {
		config = initConfig()
		return
	} else {
		line = strings.TrimPrefix(line, prefix)
	}
	// Parse.
	args := strings.Fields(line)
	for i := 0; i < len(args); {
		arg := args[i]
		switch arg {
		case "-ignore":
			i++
			*ignore = args[i]
		case "-q":
			*qualified = true
		default:
			panic("unhandled flag")
		}
		i++
	}

	config = initConfig()
}

func resetConfig() {
	config = nil
	*ignore = ""
	*qualified = false
}

func TestAll(t *testing.T) {
	filenames := []string{
		"testdata/example1.go",
		"testdata/example2.go",
		"testdata/example3.go",
		"testdata/ignore.go",
		"testdata/all.go",
		"testdata/all-q.go",
		"testdata/no-issues.go",
		"testdata/no-issues2.go",
	}

	for i, path := range filenames {
		if testing.Verbose() {
			t.Logf("test [%d]: %s", i, path)
		}
		resetConfig()
		setupConfig(path)
		runOneFile(t, path)
	}
}

func runOneFile(t *testing.T, path string) {
	src, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("failed to read file: %s", err)
		return
	}

	outContent, err := ioutil.ReadFile(outPath(path))
	if err != nil {
		t.Errorf("failed to read out file: %s", err)
		return
	}

	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, src, parserMode())
	if err != nil {
		t.Errorf("failed to parse file")
		return
	}

	issues := api.ProcessFile(config, fset, file)
	var buf bytes.Buffer
	for _, issue := range issues {
		fmt.Fprintf(&buf, "%s\n", issue)
	}

	equalBytes(t, outContent, buf.Bytes(), bytes.TrimSpace)
}
