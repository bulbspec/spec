package bulb

import (
	"errors"
	"fmt"
	"reflect"
)

var errNilResolver = errors.New("cannot resolve instances from nil Resolver")

// An invalidResolution is an [error] indicating that a [Resolver] returned a a value that could
// not be assigned to the requested type.
//
// NOTE: This error ALWAYS points to a broken implementation of [Resolver]. The contract of
// [Resolver] requires returned values to be assignable to the requested type, but the absence of
// generic methods in Go's type system prevents this from being enforced statically.
type invalidResolution struct {
	requested reflect.Type
	returned  reflect.Type
}

// Error implements [error].
func (err invalidResolution) Error() string {
	return fmt.Sprintf(
		"value from Resolver has type %v when %v was requested",
		err.returned,
		err.requested)
}

// Resolve obtains an instance of the requested type from a [Resolver]. An [error] is returned when
// the [Resolver] returns an [error] or a value that is not assignable to T.
func Resolve[T any](resolver Resolver) (T, error) {
	var zero T

	if resolver == nil {
		return zero, errNilResolver
	}

	typ := reflect.TypeFor[T]()

	resolved, err := resolver.Get(typ)
	if err != nil {
		return zero, err
	}

	typed, ok := resolved.(T)
	if !ok {
		return zero, invalidResolution{
			requested: typ,
			returned:  reflect.TypeOf(resolved),
		}
	}

	return typed, nil
}
