package a

import (
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
