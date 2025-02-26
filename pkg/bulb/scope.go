package bulb

import (
	"context"
)

// A Scope is an isolation boundary in which instances of [Scoped] types are shared.
type Scope interface {

	// ServiceProvider returns a [ServiceProvider] for the target [Scope].
	ServiceProvider() ServiceProvider

	// Close closes the [Scope] and calls the `Close(context.Context) error` or `Close() error`
	// methods on any [Scoped] and [Transient] values resolved within it.
	Close(ctx context.Context) error
}
