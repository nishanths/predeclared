package predeclared

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	run := func(t *testing.T, setup func(), patterns ...string) {
		t.Helper()
		resetFlags()
		setup()
		analysistest.Run(t, analysistest.TestData(), Analyzer, patterns...)
	}

	run(t, func() { fMode = "ds" }, "foo")
	run(t, func() { fMode = "ds" }, "testfiles")
	run(t, func() { fMode = "ds"; fField = true }, "field")
	run(t, func() { fMode = "ds"; fPkgLevel = true }, "pkglevel")
	run(t, func() { fMode = "ds"; fIgnore = "complex,print,max,len" }, "ign")
	run(t, func() { fMode = "ds"; fField = true }, "syncmap")
}
