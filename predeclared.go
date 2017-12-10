// Command predeclared prints declarations and fields in the given files
// that have the same name as one of Go's predeclared identifiers.
//
// Exit code
//
// The command exits with exit code 1 if an error occured parsing the given
// files. Otherwise the exit code is 0, even if issues were found. Use the
// '-exit' flag to use a exit code 1 when issues are found.
//
// Usage
//
// Run 'predeclared -h' for help.
//
// If the '-q' flag isn't specified, the command does not report the names of
// fields in struct types, methods in interface types, and method declarations
// even if they have the same name as predeclared identifier. (These kinds aren't
// included by default since fields and method are always accessed via a
// qualifier, à la obj.A).
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/scanner"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const help = `Find declarations and fields that override predeclared identifiers.

Usage: 
  predeclared [flags] [path...]

Flags:
  -e	 Report all parse errors, not just the first 10 on different lines
  -exit  Set exit status to 1 if issues are found
  -q     Include methods and fields that have the same name as predeclared identifiers
`

func usage() {
	fmt.Fprintf(os.Stderr, help)
	os.Exit(2)
}

var (
	allErrors = flag.Bool("e", false, "")
	qualified = flag.Bool("q", false, "")
	setExit   = flag.Bool("exit", false, "")
)

var exitCode = 0

func main() {
	log.SetFlags(0)
	log.SetPrefix("declfmt: ")

	flag.Usage = usage
	flag.Parse()

	var fset = token.NewFileSet()
	if flag.NArg() == 0 {
		handleFile(fset, true, "<standard input>", os.Stdout) // use the same filename that gofmt uses
	} else {
		for i := 0; i < flag.NArg(); i++ {
			path := flag.Arg(i)
			info, err := os.Stat(path)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				exitCode = 1
			} else if info.IsDir() {
				handleDir(fset, path)
			} else {
				handleFile(fset, false, path, os.Stdout)
			}
		}
	}

	os.Exit(exitCode)
}

func parserMode() parser.Mode {
	if *allErrors {
		return parser.ParseComments | parser.AllErrors
	}
	return parser.ParseComments
}

func handleFile(fset *token.FileSet, stdin bool, filename string, out io.Writer) {
	var src []byte
	var err error
	if stdin {
		src, err = ioutil.ReadAll(os.Stdin)
	} else {
		src, err = ioutil.ReadFile(filename)
	}
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		exitCode = 1
		return
	}

	file, err := parser.ParseFile(fset, filename, src, parserMode())
	if err != nil {
		scanner.PrintError(os.Stderr, err)
		exitCode = 1
		return
	}

	issues := processFile(fset, file)
	if len(issues) == 0 {
		return
	}

	if *setExit {
		exitCode = 1
	}

	for _, issue := range issues {
		fmt.Fprintf(out, "%s\n", issue)
	}
}

func handleDir(fset *token.FileSet, p string) {
	if err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !isGoFile(info) {
			return nil
		}
		handleFile(fset, false, path, os.Stdout)
		return nil
	}); err != nil {
		fmt.Fprint(os.Stderr, err)
		exitCode = 1
	}
}

func isGoFile(f os.FileInfo) bool {
	// ignore non-Go files
	name := f.Name()
	return !f.IsDir() && !strings.HasPrefix(name, ".") && !strings.HasPrefix(name, "_") && strings.HasSuffix(name, ".go")
}

// https://golang.org/ref/spec#Predeclared_identifiers
var predeclaredIdents = map[string]bool{
	"bool":       true,
	"byte":       true,
	"complex64":  true,
	"complex128": true,
	"error":      true,
	"float32":    true,
	"float64":    true,
	"int":        true,
	"int8":       true,
	"int16":      true,
	"int32":      true,
	"int64":      true,
	"rune":       true,
	"string":     true,
	"uint":       true,
	"uint8":      true,
	"uint16":     true,
	"uint32":     true,
	"uint64":     true,
	"uintptr":    true,

	"true":  true,
	"false": true,
	"iota":  true,

	"nil": true,

	"append":  true,
	"cap":     true,
	"close":   true,
	"complex": true,
	"copy":    true,
	"delete":  true,
	"imag":    true,
	"len":     true,
	"make":    true,
	"new":     true,
	"panic":   true,
	"print":   true,
	"println": true,
	"real":    true,
	"recover": true,
}

type Issue struct {
	ident *ast.Ident
	kind  string
	fset  *token.FileSet
}

func (i Issue) String() string {
	pos := i.fset.Position(i.ident.Pos())
	return fmt.Sprintf("%s: %s %q has same name as predeclared identifier", pos, i.kind, i.ident.Name)
}

func processFile(fset *token.FileSet, file *ast.File) []Issue {
	var issues []Issue

	maybeAdd := func(x *ast.Ident, kind string) {
		if predeclaredIdents[x.Name] {
			issues = append(issues, Issue{x, kind, fset})
		}
	}

	// TODO: consider deduping package name issues for files in the
	// same directory.
	maybeAdd(file.Name, "package name")

	// Handle declarations and fields.
	// https://golang.org/ref/spec#Declarations_and_scope
	ast.Inspect(file, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.GenDecl:
			var kind string
			switch x.Tok {
			case token.CONST:
				kind = "const"
			case token.VAR:
				kind = "variable"
			default:
				return true
			}
			for _, spec := range x.Specs {
				if vspec, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vspec.Names {
						maybeAdd(name, kind)
					}
				}
			}
			// Shouldn't look at the specs again.
			// Also, specs can't nest other specs, so it's okay to not look deeper.
			return false
		case *ast.TypeSpec:
			maybeAdd(x.Name, "type")
			return true
		case *ast.StructType:
			if *qualified && x.Fields != nil {
				for _, field := range x.Fields.List {
					for _, name := range field.Names {
						maybeAdd(name, "field")
					}
				}
			}
			return true
		case *ast.InterfaceType:
			if *qualified && x.Methods != nil {
				for _, meth := range x.Methods.List {
					for _, name := range meth.Names {
						maybeAdd(name, "method")
					}
				}
			}
			return true
		case *ast.FuncDecl:
			if x.Recv == nil {
				// it's a function
				maybeAdd(x.Name, "function")
			} else {
				// it's a method
				if *qualified {
					maybeAdd(x.Name, "method")
				}
			}
			// add receivers idents
			if x.Recv != nil {
				for _, field := range x.Recv.List {
					for _, name := range field.Names {
						maybeAdd(name, "variable")
					}
				}
			}
			// add params idents
			for _, field := range x.Type.Params.List {
				for _, name := range field.Names {
					maybeAdd(name, "variable")
				}
			}
			// add returns idents
			if x.Type.Results != nil {
				for _, field := range x.Type.Results.List {
					for _, name := range field.Names {
						maybeAdd(name, "variable")
					}
				}
			}
			return true
		case *ast.LabeledStmt:
			maybeAdd(x.Label, "label")
			return true
		case *ast.AssignStmt:
			// We only care about short variable declarations, which use token.DEFINE.
			if x.Tok == token.DEFINE {
				for _, expr := range x.Lhs {
					if ident, ok := expr.(*ast.Ident); ok {
						maybeAdd(ident, "variable")
					}
				}
			}
			return true
		default:
			return true
		}
	})

	return issues
}
