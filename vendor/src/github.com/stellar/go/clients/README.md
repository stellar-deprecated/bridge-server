# Clients package

Packages contained by this package provide client libraries for accessing the ecosystem of stellar services.  At present, it only contains a simple horizon client library, but in the future it will contain clients to interact with stellar-core, federation, the bridge server and more.

See [godoc](https://godoc.org/github.com/stellar/go/clients) for details about each package.

## Adding new client packages

Ideally, each one of our client packages will have commonalities in their API to ease the cost of learning each.  It's recommended that we follow a pattern similar to the `net/http` package's client shape:

A type, `Client`, is the central type of any client package, and its methods should provide the bulk of the functionality for the package.  A `DefaultClient` var is provided for consumers that don't need client-level customization of behavior.  Each method on the `Client` type should have a corresponding func at the package level that proxies a call through to the default client.  For example, `http.Get()` is the equivalent of `http.DefaultClient.Get()`.