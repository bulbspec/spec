package bulb

// A ServiceProvider is a [Resolver] that can also create new [Scope]s.
type ServiceProvider interface {
	Resolver

	// NewScope creates a new [Scope].
	NewScope() Scope
}
