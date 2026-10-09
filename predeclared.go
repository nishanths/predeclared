// Command predeclared applies the predeclared static analysis to
// the specified Go packages.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/nishanths/predeclared/internal/driver"
	"github.com/nishanths/predeclared/passes/predeclared"

	"golang.org/x/tools/go/packages"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("predeclared: ")
	driver.ParseFlags("predeclared", predeclared.Analyzer)
	// As of x/tools@v0.50.0, neither the predeclared
	// analyzer nor the analyzers (recursively) required by
	// it uses facts. Therefore loading packages with
	// packages.LoadSyntax is sufficient.
	// See usage site of func needFacts in file
	// x/tools/go/analysis/internal/checker/checker.go for more details.
	os.Exit(driver.Run(flag.Args(), packages.LoadSyntax, predeclared.Analyzer))
}
