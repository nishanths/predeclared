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

func ParseFlags(progname string, as ...*analysis.Analyzer) {
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

func Run(args []string, as ...*analysis.Analyzer) (exitStatus int) {
	exitAtLeast := func(s int) {
		if s > exitStatus {
			exitStatus = s
		}
	}

	cfg := packages.Config{
		Mode:  loadMode(as) | packages.NeedModule,
		Tests: includeTests,
	}
	pkgs, err := packages.Load(&cfg, args...)
	if err != nil {
		log.Printf("error loading packages: %s", err)
		exitAtLeast(1)
		return
	}

	if len(pkgs) == 0 {
		log.Println("no packages matched")
		exitAtLeast(1)
		return
	}

	if packages.PrintErrors(pkgs) > 0 {
		exitAtLeast(1)
		// Do not return, the analysis can still proceed.
	}

	graph, err := checker.Analyze(as, pkgs, nil)
	if err != nil {
		log.Println(err)
		exitAtLeast(1)
		return
	}

	if err := graph.PrintText(os.Stderr, contextLines); err != nil {
		log.Println(err)
		exitAtLeast(1)
		return
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
	// Note: These exit status values match those used by
	// analysis/singlechecker.
	if errors > 0 {
		exitAtLeast(1)
	}
	if diags > 0 {
		exitAtLeast(3)
	}
	return
}

func loadMode(as []*analysis.Analyzer) packages.LoadMode {
	if needFacts(as) {
		return packages.LoadAllSyntax
	}
	return packages.LoadSyntax
}

// The needFacts function was copied from file
// go/analysis/internal/checker/checker.go in the
// golang.org/x/tools module.
// See LICENSE file.

func needFacts(analyzers []*analysis.Analyzer) bool {
	seen := make(map[*analysis.Analyzer]bool)
	var q []*analysis.Analyzer // for BFS
	q = append(q, analyzers...)
	for len(q) > 0 {
		a := q[0]
		q = q[1:]
		if !seen[a] {
			seen[a] = true
			if len(a.FactTypes) > 0 {
				return true
			}
			q = append(q, a.Requires...)
		}
	}
	return false
}
