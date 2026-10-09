// Command predeclared applies the predeclared static analysis to
// the specified Go packages.
package main

import (
	"flag"
	"log"
	"os"

	"github.com/nishanths/predeclared/internal/driver"
	"github.com/nishanths/predeclared/passes/predeclared"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("predeclared: ")

	a := predeclared.Analyzer
	driver.ParseFlags("predeclared", a)
	os.Exit(driver.Run(flag.Args(), a))
}
