package ctxfield_test

import (
	"testing"

	"github.com/gostaticanalysis/ctxfield"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, ctxfield.Analyzer, "a")
}