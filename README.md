# bulb spec

[![License](https://img.shields.io/badge/License-BSD_3--Clause-blue.svg)](https://opensource.org/licenses/BSD-3-Clause)

Bulb enables projects to leverage dependency injection and dynamic composition without being coupled to a specific service provider or IoC container implementation.

The [`bulb`][bulb-package] package provides a set of interfaces for resolving values of arbitrary types at runtime. Implementations of [`bulb`][bulb-package] interfaces provide mechanisms for specifying how the values should be constructed and how their dependencies are satisfied. The documentation is a guide for both consumers and implementors of the interfaces.

## Conventions

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED", "NOT RECOMMENDED", "MAY", and "OPTIONAL" in this document are to be interpreted as described in [BCP 14] [[RFC2119]] [[RFC8174]] when, and only when, they appear in all capitals, as shown here.

## Documentation

This [README](./README.md) is high-level overview of the [`bulb`][bulb-package] package. All the important [concepts](#concepts) are introduced here.

The [`bulb`][bulb-package] package docs are a reference for working with and implementing the API contract.

## Concepts

### Closing Values

When [`scope.Close`][bulb.Scope.Close] is called, the lifetime of the [`Transient`][bulb.Transient] and [`Scoped`][bulb.Scoped] values from its [`Resolver`][bulb.Resolver] ends. Implementations MUST call [`Close`][bulb.Closer.Close] on any value that implements [`Closer`][bulb.Closer].

### Dependencies

[`Resolver`][bulb.Resolver] implementations SHOULD support initializing values recursively such that the fields of resolved values have the correct [lifetime](#lifetimes). For example, an implementation may initialize struct values by recursively initializing its exported fields. In this case, the values for those fields are the dependencies. Another implementation may support definition initialization logic using a function that takes a [`Resolver`][bulb.Resolver] and returns the target value. In this case the dependencies would be any values obtained from the [`Resolver`][bulb.Resolver] in the given function.

The semantics of [`Scoped`][bulb.Scoped] include some RECOMMENDED behaviour for how implementations handle dependencies when it can impact lifetime management.

### Lifetimes

A [`Lifetime`][bulb.Lifetime] describes the lifetime and sharing semantics of values created by a [`Resolver`][bulb.Resolver].

See the docs for [`Transient`][bulb.Transient], [`Scoped`][bulb.Scoped], and [`Singleton`][bulb.Singleton] for details.

### Resolvers

The [`Resolver`][bulb.Resolver] interface is the primary contract between consumers and implementors. Here's an overview of a typical integration:

- Some setup code configures a specific [`Resolver`][bulb.Resolver] implementation to satisfy the types the application or framework code needs.
- The setup code passes the configured [`Resolver`][bulb.Resolver] to the application or framework code.
- The application or framework code uses the [`Resolver`][bulb.Resolver] to resolve the values it needs without any coupling to how the were initialized.

### Root And Scoped Resolvers

A root resolver is any instance of [`Resolver`][bulb.Resolver] that did not come from a call to [`Scope.Resolver`][bulb.Scope.Resolver]. The mechanism for obtaining a root resolver is implementation-defined. A single instance of each [`Singleton`][bulb.Singleton] type is [shared](#shareable-types) by a root resolver and any scoped resolver created from it, but not across distinct root resolvers. Root resolvers SHOULD NOT resolve [`Scoped`][bulb.Scoped] types.

A scoped resolver is any [`Resolver`][bulb.Resolver] obtained from a call to [`Scope.Resolver`][bulb.Scope.Resolver]. A single instance of each [`Scoped`][bulb.Scoped] type is [shared](#shareable-types) by scoped resolver, but not by any other [`Resolver`][bulb.Resolver].

### Shareable Types

Sometimes [`Resolver.Get`][bulb.Resolver.Get] needs to return "the same" instance when the same type is resolved more than once. This is not technically possible because Go doesn't have references, but for some types and usecases a copy is _effectively_ the same instance. Whether a type is _shareable_ depends on its copy semantics, method implementations, and how it's used. For example, copying the value of a struct type could be considered instance sharing so long as the struct's own state is never mutated. Whether the struct is just holding primative values, or holding pointers to shared mutable data, one copy of the struct would be indistinguishable from another. On the other hand, even using pointers directly can't guarantee shared instance semantics because one of the pointers could be reassigned to point a different value.

Since sharing semantics can't be determined at compile time, it's up to implementations and users to decide which types are sharable. Bulb is only concerned with the concept because the semantics of [lifetimes](#lifetimes) and [scopes](#root-and-scoped-resolvers) depend on it.

[bulb-package]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb
[bulb.Lifetime]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Lifetime
[bulb.Closer]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Closer
[bulb.Closer.Close]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Closer.Close
[bulb.Resolver]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Resolver
[bulb.Resolver.Get]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Resolver.Get
[bulb.Scope.Close]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Scope.Close
[bulb.Scoped]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Scoped
[bulb.Singleton]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Singleton
[bulb.Transient]: https://pkg.go.dev/github.com/bulbspec/spec/pkg/bulb#Transient
