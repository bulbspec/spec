package bulb

import (
	"reflect"
)

// A Resolver provides values of arbitrary types.
type Resolver interface {

	// Get returns an instance of the requested type.
	//
	// If no error is returned then the returned value MUST be assigable to the requested type.
	Get(typ reflect.Type) (any, error)
}
