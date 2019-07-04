package main

import (
	"github.com/gostaticanalysis/ctxfield"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(ctxfield.Analyzer) }
