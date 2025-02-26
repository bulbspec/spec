# bulb spec

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

Bulb enables projects to leverage dependency injection and inversion of control (IoC) without being coupled to a specific service provider or IoC container implementation.

The [`bulb`][bulb-package] package provides a set of interfaces for resolving values of arbitrary types at runtime. The documentation is a guide for both consumers and implementors of the interfaces.

## Conventions

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [BCP 14] [[RFC2119]] [[RFC8174]] when, and only when, they appear in all capitals, as shown here.

## Documentation

This [README](./README.md) is high-level overview of the [`bulb`][bulb-package] package. All the important [concepts](#concepts) are introduced here.

The [`bulb`][bulb-package] package docs are a reference for working with and implementing the API contract.

## Concepts

### Closing Values

When [`Close`][bulb.Scope.Close] is called on a [`Scope`][bulb.Scope], the lifetime of the [`Transient`][bulb.Transient] and [`Scoped`][bulb.Scoped] values from its [`Resolver`][bulb.Resolver] is over. Implementations MUST call [`Close`][bulb.Closer.Close] on any value that implements [`Closer`][bulb.Closer].

### Dependencies

[`Resolver`][bulb.Resolver] implementations SHOULD support initializing values recursively such that the fields of resolved values have the correct [lifetime](#lifetimes). For example, an implementation may initialize struct values by recursively initializing its exported fields. In this case, the values for those fields are the dependencies. Another implementation may support definition initialization logic using a function that takes a [`Resolver`][bulb.Resolver] and returns the target value. In this case the dependencies would be any values obtained from the [`Resolver`][bulb.Resolver] in the given function.

The semantics of [`Scoped`][bulb.Scoped] include some RECOMMENDED behaviour for how implementations handle dependencies when it can impact lifetime management.

### Lifetimes

A [`Lifetime`][bulb.Lifetime] describes the lifetime and sharing semantics of values created by a [`Resolver`][bulb.Resolver].

See the docs for [`Transient`][bulb.Transient], [`Scoped`][bulb.Scoped], and [`Singleton`][bulb.Singleton] for details.

### Resolvers

The [`Resolver`][bulb.Resolver] interface is the primary contract between consumers and implementors. Here's an overview of a typical integration:

- Some application or framework code uses a [`Resolver`][bulb.Resolver] to resolve the values it needs.
- Some setup code configures a specific [`Resolver`][bulb.Resolver] implementation to satisfy the types the application or framework code needs.
- The setup code passes the configured [`Resolver`][bulb.Resolver] to the application or framework code.
- They all live happily ever after.

### Root And Scoped Resolvers

A root resolver is any instance of [`Resolver`][bulb.Resolver] that did not come from a call to [`NewScope`][bulb.Resolver.NewScope]. The mechanism for obtaining a root resolver is implementation-defined. A single instance of each [`Singleton`][bulb.Singleton] type is [shared](#shareable-types) by a root resolver and any scoped resolver created from it, but not across distinct root resolvers. Root resolvers SHOULD NOT resolve [`Scoped`][bulb.Scoped] types.

A scoped resolver is any instance of [`Resolver`][bulb.Resolver] obtained from a call to [`NewScope`][bulb.Resolver.NewScope]. A single instance of each [`Scoped`][bulb.Scoped] type is [shared](#shareable-types) by scoped resolver, but not by any other [`Resolver`][bulb.Resolver].

There is no such thing as a nested [`Scope`][bulb.Scope]. A scoped resolver obtained from another scoped resolver will not share instances of [`Scoped`][bulb.Scoped] types with the original [`Resolver`][bulb.Resolver].

### ServiceProviders

The [`ServiceProvider`][bulb.ServiceProvider] interface is an extension of the [`Resolver`][bulb.Resolver] interface that provides the ability to create new [resolution scopes](#root-and-scoped-resolvers). The primary usecase for a [`ServiceProvider`][bulb.ServiceProvider] is supporting logic that defines isolation boundaries within an application or framework. E.g. an HTTP API may use a [`ServiceProvider`][bulb.ServiceProvider] to create a new scope for each request it handles.

### Shareable Types

Sometimes [`Resolver.Get`][bulb.Resolver.Get] needs to return "the same" instance when the same type is resolved more than once. This is not technically possible because Go doesn't have references, but for some types and usecases a copy is _effectively_ the same instance. Unfortunately, _shareability_ can't be determined statically because it's not a property of the type alone. It also depends on how the values are used.

For example, copying the value of a struct type could be considered instance sharing so long as the struct's own state is never mutated. Whether the struct is just holding primative values, or holding pointers to shared mutable data, one copy of the struct would be indistinguishable from another. On the other hand, even using pointers directly can't guarantee shared instance semantics because one of the pointers could be reassigned to point a different value.

Since Go doesn't provide mechanisms to validate sharing semantics at compile time it's up to implementations and users to decide which types are sharable. Bulb is only concerned with the concept because the semantics of [lifetimes](#lifetimes) and [scopes](#root-and-scoped-resolvers) depend on it.

[bulb-package]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb
[bulb.Lifetime]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Lifetime
[bulb.Closer]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Closer
[bulb.Closer.Close]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Closer.Close
[bulb.Resolver]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Resolver
[bulb.Resolver.Get]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Resolver.Get
[bulb.Resolver.NewScope]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Resolver.NewScope
[bulb.Scope]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Scope
[bulb.Scope.Close]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Scope.Close
[bulb.Scoped]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Scoped
[bulb.ServiceProvider]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#ServiceProvider
[bulb.Singleton]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Singleton
[bulb.Transient]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Transient
