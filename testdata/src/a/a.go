package a

import (
	_ "b" // import for b.Empty
	"context"
	"fmt"
)

type I interface{ M() }

type T1 struct {
	ctx context.Context // want "context.Context must not be in a field"
}

type T2 struct {
	fmt.Stringer
	ctx context.Context // OK
}

type T3 struct {
	ctx *context.Context // want "context.Context must not be in a field"
}

type T4 struct {
	context.Context // OK
}

type T5 struct {
	I
	ctx context.Context // want "context.Context must not be in a field"
}

type T6 struct {
	ctx context.Context // want "context.Context must not be in a field"
}

func (*T6) M() {}

type T7 struct {
	ctx context.Context // OK
}

var _ fmt.Stringer = (*T7)(nil)

func (*T7) String() string { return "" }

type T8 struct {
	N       *int
	context *context.Context // want "context.Context must not be in a field"
}

func (t *T8) M() {}
