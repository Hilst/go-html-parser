package templates

import (
	opt "github.com/moznion/go-optional"
)

// GENERATORS
func props(els ...any) []any {
	return els
}

// TYPE ENFORCER
func array(parr opt.Option[any]) opt.Option[[]any] {
	if parr, err := parr.Take(); err == nil {
		return opt.Some[[]any](parr.([]any))
	}
	return opt.None[[]any]()
}

// VALIDATOR
func resolveArray(op opt.Option[[]any]) []any {
	return op.TakeOr([]any{})
}

func emptyArray(a []any) bool {
	return len(a) == 0
}
