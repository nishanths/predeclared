// Command predeclared applies the predeclared static analysis to
// the specified Go packages.
package main

import (
	"flag"
	"log"

	"github.com/nishanths/predeclared/internal/driver"
	"github.com/nishanths/predeclared/passes/predeclared"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/packages"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("predeclared: ")
	driver.ParseFlags("predeclared", []*analysis.Analyzer{predeclared.Analyzer})
	driver.Run(flag.Args(), packages.LoadSyntax, []*analysis.Analyzer{predeclared.Analyzer})
}
