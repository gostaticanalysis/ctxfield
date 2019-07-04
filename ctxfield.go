package ctxfield

import (
	"go/types"

	"github.com/gostaticanalysis/analysisutil"
	"github.com/gostaticanalysis/ident"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "ctxfield",
	Doc:  doc,
	Run:  new(runner).run,
	Requires: []*analysis.Analyzer{
		ident.Analyzer,
	},
}

const doc = `ctxfield reports context.Context which belongs to a struct as a field.
ctxfield ignores a struct which implements an interface.`

type runner struct {
	ctx types.Type
}

func (r *runner) run(pass *analysis.Pass) (interface{}, error) {
	r.ctx = analysisutil.TypeOf(pass, "context", "Context")
	m := pass.ResultOf[ident.Analyzer].(ident.Map)

	for o := range m {
		if r.ctxInField(o) {
			pass.Reportf(o.Pos(), "context.Context must not be in a field")
		}
	}

	return nil, nil
}

func (r *runner) ctxInField(o types.Object) bool {

	if o == nil || o.Parent() == types.Universe || o.Exported() {
		return false
	}

	v, isVar := o.(*types.Var)
	if !isVar || !v.IsField() || v.Anonymous() {
		return false
	}

	if !r.isCtx(v.Type()) {
		return false
	}

	var st types.Type
	for n, s := range analysisutil.Structs(v.Pkg()) {
		if analysisutil.HasField(s, v) {
			st = v.Pkg().Scope().Lookup(n).Type()
			break
		}
	}

	if st == nil {
		return false
	}

	stptr := types.NewPointer(st)

	for _, pkg := range v.Pkg().Imports() {
		for _, i := range analysisutil.Interfaces(pkg) {
			if types.Implements(st, i) ||
				types.Implements(stptr, i) {
				return false
			}
		}
	}

	return true
}

func (r *runner) isCtx(t types.Type) bool {
	if types.Identical(t, r.ctx) {
		return true
	}

	switch t := t.(type) {
	case *types.Pointer:
		return r.isCtx(t.Elem())
	case *types.Named:
		return r.isCtx(t.Underlying())
	}

	return false
}
