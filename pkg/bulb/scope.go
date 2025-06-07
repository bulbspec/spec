package bulb

import (
	"context"
)

// A Scoper creates [Scope]s to provide [scoped] [Resolver]s.
//
// [scoped]: https://github.com/bulbspec/spec?tab=readme-ov-file#root-vs-scoped-resolvers
type Scoper interface {

	// NewScope creates a new [Scope].
	NewScope() Scope
}

// A Scope is an isolation boundary in which instances of [Scoped] types are shared.
type Scope interface {

	// Resolver returns the [Resolver] associated with the [Scope].
	Resolver() Resolver

	// Close closes the [Scope] and calls [Closer.Close] on any [Scoped] and [Transient] values
	// resolved from it's [Resolver] that implement [Closer].
	Close(ctx context.Context) error
}
