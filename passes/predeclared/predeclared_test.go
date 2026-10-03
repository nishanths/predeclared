package predeclared

import (
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	run := func(t *testing.T, setup func(), patterns ...string) {
		t.Helper()
		resetFlags()
		setup()
		analysistest.Run(t, filepath.Join(testdata, "test"), Analyzer, patterns...)
	}

	run(t, func() { fMode = "ds" }, "test/foo")
	run(t, func() { fMode = "ds" }, "test/testfiles")
	run(t, func() { fMode = "ds"; fField = true }, "test/field")
	run(t, func() { fMode = "ds"; fPkgLevel = true }, "test/pkglevel")
	run(t, func() { fMode = "ds"; fIgnore = "complex,print,max,len" }, "test/ign")
	run(t, func() { fMode = "ds"; fField = true }, "test/syncmap")
	run(t, func() {}, "test/readmedoc/...")
}
