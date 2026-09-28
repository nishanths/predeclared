package driver

import (
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/checker"
	"golang.org/x/tools/go/packages"
)

const (
	includeTests = true
	contextLines = -1
)

func ParseFlags(progname string, as []*analysis.Analyzer) {
	for _, a := range as {
		a.Flags.VisitAll(func(f *flag.Flag) {
			if flag.Lookup(f.Name) != nil {
				log.Fatalf("internal error: conflicting flag -%s", f.Name)
			}
			flag.Var(f.Value, f.Name, f.Usage)
		})
	}
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] [packages]\n", progname)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func Run(args []string, loadMode packages.LoadMode, as []*analysis.Analyzer) {
	cfg := packages.Config{
		Mode:  loadMode | packages.NeedModule,
		Tests: includeTests,
	}
	initial, err := packages.Load(&cfg, args...)
	if err != nil {
		log.Fatalf("error loading packages: %s", err)
	}

	if len(initial) == 0 {
		log.Fatal("matched no packages")
	}

	if packages.PrintErrors(initial) > 0 {
		os.Exit(1)
	}

	graph, err := checker.Analyze(as, initial, nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := graph.PrintText(os.Stderr, contextLines); err != nil {
		log.Fatal(err)
	}

	var errors, diags int
	for act := range graph.All() {
		if act.Err != nil {
			errors++
			continue
		}
		if act.IsRoot {
			diags += len(act.Diagnostics)
		}
	}
	// Note: The exit status values match those used by
	// analysis/singlechecker.
	if errors > 0 {
		os.Exit(1)
	}
	if diags > 0 {
		os.Exit(3)
	}
	os.Exit(0)
}
