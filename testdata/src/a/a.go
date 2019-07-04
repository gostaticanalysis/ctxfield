package a

import (
	"context"
	"fmt"
)

type T1 struct {
	ctx context.Context // want "context.Context must not be in a field"
}

type T2 struct {
	fmt.Stringer
	ctx context.Context // OK
}