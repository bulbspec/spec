package bulb

// Lifetimes describe the lifetime and sharing semantics of values created by [Resolver]s.
//
// The zero value for [Lifetime] is [Undefined] because none of the defined values are suitable.
type Lifetime int

const (

	// Undefined is the zero value because none of the defined values are suitable. The appropriate
	// lifetime and sharing semantics for a given type depends on the contract it provides, the way
	// it's implemented, and how it's used by the application.
	Undefined Lifetime = iota

	// Transient means every call to [Resolver.Get] returns a distinct value. Consumers can safely
	// [close] resolved instances because they will never be [shared]. Implementations MUST
	// [close] [Transient] values when [scoped] [Resolver]s are closed.
	//
	// [close]: https://github.com/bulbspec/spec?tab=readme-ov-file#closing-values
	// [scoped]: https://github.com/bulbspec/spec?tab=readme-ov-file#root-vs-scoped-resolvers
	// [shared]: https://github.com/bulbspec/spec?tab=readme-ov-file#shareable-types
	Transient

	// Scoped means every call to [Resolver.Get] for the same [reflect.Type] on any [Resolver] from
	// from the same [Scope] MUST return the [same instance], and calls to [Resolver.Get] for the
	// same type on a [Resolver] from a different [Scope] MUST return a different instance.
	// Consumers SHOULD NOT [close] resolved instances of [Scoped] types because they may be in use
	// by, or subsequently returned to, other consumers. Implementations MUST [close] [Scoped]
	// values when [scoped] [Resolver]s are closed.
	//
	// There is no such thing as a nested [Scope]. A [scoped] [Resolver] obtained from another
	// [scoped] [Resolver] will not share instances of [Scoped] types with the original [Resolver].
	//
	// # Implemenation Guidelines
	//
	// Implementations SHOULD NOT resolve [Scoped] types from a [root] [Resolver]. Associating a
	// type with the [Scoped] [Lifetime] conveys the intention to constrain instances of that type
	// to the lifetime of the [Scope] of the [Resolver] that produced the instance. A [root]
	// [Resolver] has no [Scope]. Attempting to resolve a [Scoped] type from a [root] [Resolver]
	// conveys the intention to constrain the instance to the lifetime of a [Scope] that doesn't
	// exist. There are only two ways to interpret this:
	//
	//		1. The target type is intended to be used both with and without lifetimes, and the
	// 		[Resolver] an instance is obtained from corresponds to the desired usecase.
	//
	//		2. The attempt to resolve the type is a bug. Maybe the type was configured with the
	// 		wrong [Lifetime]. Maybe the [Resolver] was supposed to be [scoped]. Maybe the type is
	//		being resolved to satisfy a dependency and it's the dependency relationship that's in
	//		a bad state.
	//
	// Given there are several more explicit ways to express the first interpretation, and the
	// potential consequences of leaking an instance of a type whose lifetime is supposed to be
	// managed, the RECOMMENDED approach is to assume it's a bug and not resolve the value.
	//
	// Implementations SHOULD NOT resolve [Scoped] types to satisfy [dependencies] for [Singleton]
	// types. Even if they're resolved from a [scoped] [Resolver], instances of [Singleton] types
	// are [shared] with the originating [root] [Resolver] and don't inherit the lifetime of the
	// [Scope]. Allowing a [Singleton] type to depend on a [Scoped] type carries the same risks as
	// resolving [Scoped] values directly from a [root] [Resolver] plus the additional risk that a
	// [Scoped] value held by a [Singleton] will be [closed] while the [Singleton] is still in use
	// and produce undefined behaviour. Therefore, the RECOMMENDED approach is to assume it's a bug
	// and not resolve the value.
	//
	// [close]: https://github.com/bulbspec/spec?tab=readme-ov-file#closing-values
	// [closed]: https://github.com/bulbspec/spec?tab=readme-ov-file#closing-values
	// [dependencies]: https://github.com/bulbspec/spec?tab=readme-ov-file#dependencies
	// [root]: https://github.com/bulbspec/spec?tab=readme-ov-file#root-vs-scoped-resolvers
	// [same instance]: https://github.com/bulbspec/spec?tab=readme-ov-file#shareable-types
	// [scoped]: https://github.com/bulbspec/spec?tab=readme-ov-file#root-vs-scoped-resolvers
	Scoped

	// Singleton means that calling [Resolver.Get] for the same type on a [root] [Resolver] or any
	// [scoped] [Resolver] derived from it recursively MUST return the [same instance] every time
	// regardless of which [Resolver] returned it first. Consumers SHOULD NOT [close] resolved
	// instances of [Singleton] types because they may be in use by, or subsequently returned to,
	// other consumers. Implementations MUST NOT [close] instances of [Singleton] types that were
	// resolved by a [scoped] [Resolver] when the [Resolver] is closed.
	Singleton
)

// Defined indicates whether a [Lifetime] is one of the defined values: [Transient], [Scoped], or
// [Singleton].
func (lifetime Lifetime) Defined() bool {
	return lifetime > Undefined && lifetime <= Singleton
}

// String returns the name of a [Lifetime] as a string representation.
func (lifetime Lifetime) String() string {
	switch lifetime {
	case Transient:
		return "Transient"
	case Scoped:
		return "Scoped"
	case Singleton:
		return "Singleton"
	}
	return "Undefined"
}
