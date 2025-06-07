package bulb

import "context"

// The Closer interface defines a hook for cleanup when a value’s lifetime ends.
type Closer interface {

	// Close is called on [Scoped] and [Transient] values when the [Scope] they were resolved from
	// is closed. It is not called on values resolved from [root] [Resolver]s, nor on [Singleton]
	// values resolved from a [Scope].
	//
	// [root]: https://github.com/bulbspec/spec?tab=readme-ov-file#root-and-scoped-resolvers
	Close(ctx context.Context) error
}
