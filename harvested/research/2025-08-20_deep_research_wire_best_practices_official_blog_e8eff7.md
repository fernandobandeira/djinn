---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:36:59.648826'
profile: deep_research
source: https://blog.golang.org/wire
topic: Wire Best Practices Official Blog
---

# Wire Best Practices Official Blog

[ ![Go](https://go.dev/images/go-logo-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://go.dev/blog/wire#main-content)
  * [ Why Go _arrow_drop_down_ ](https://go.dev/blog/wire)
Press Enter to activate/deactivate dropdown 
    * [ Case Studies ](https://go.dev/solutions/case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions/use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [ Learn ](https://go.dev/learn/)
Press Enter to activate/deactivate dropdown 
  * [ Docs _arrow_drop_down_ ](https://go.dev/blog/wire)
Press Enter to activate/deactivate dropdown 
    * [ Go Spec ](https://go.dev/ref/spec)
The official Go language specification
    * [ Go User Manual ](https://go.dev/doc)
A complete introduction to building software with Go
    * [ Standard library ](https://pkg.go.dev/std)
Reference documentation for Go's standard library
    * [ Release Notes ](https://go.dev/doc/devel/release)
Learn what's new in each Go release
    * [ Effective Go ](https://go.dev/doc/effective_go)
Tips for writing clear, performant, and idiomatic Go code
  * [ Packages ](https://pkg.go.dev)
Press Enter to activate/deactivate dropdown 
  * [ Community _arrow_drop_down_ ](https://go.dev/blog/wire)
Press Enter to activate/deactivate dropdown 
    * [ Recorded Talks ](https://go.dev/talks/)
Videos from prior events
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
Meet other local Go developers
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
Learn and network with Go developers from around the world
    * [ Go blog ](https://go.dev/blog)
The Go project's official blog.
    * [ Go project ](https://go.dev/help)
Get help and stay informed from Go
    * Get connected 
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


[ ![Go.](https://go.dev/images/go-logo-blue.svg) ](https://go.dev/)
  * [Why Go _navigate_next_](https://go.dev/blog/wire)
[_navigate_before_ Why Go](https://go.dev/blog/wire)
    * [ Case Studies ](https://go.dev/solutions/case-studies)
    * [ Use Cases ](https://go.dev/solutions/use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [Docs _navigate_next_](https://go.dev/blog/wire)
[_navigate_before_ Docs](https://go.dev/blog/wire)
    * [ Go Spec ](https://go.dev/ref/spec)
    * [ Go User Manual ](https://go.dev/doc)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
    * [ Effective Go ](https://go.dev/doc/effective_go)
  * [Packages](https://pkg.go.dev)
  * [Community _navigate_next_](https://go.dev/blog/wire)
[_navigate_before_ Community](https://go.dev/blog/wire)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


# [The Go Blog](https://go.dev/blog/)
# Compile-time Dependency Injection With Go Cloud's Wire
Robert van Gent 9 October 2018 
## Overview[¶](https://go.dev/blog/wire#overview)
The Go team recently [announced](https://go.dev/blog/go-cloud) the open source project [Go Cloud](https://github.com/google/go-cloud), with portable Cloud APIs and tools for [open cloud](https://cloud.google.com/open-cloud/) development. This post goes into more detail about Wire, a dependency injection tool used in Go Cloud.
## What problem does Wire solve?[¶](https://go.dev/blog/wire#what-problem-does-wire-solve)
[Dependency injection](https://en.wikipedia.org/wiki/Dependency_injection) is a standard technique for producing flexible and loosely coupled code, by explicitly providing components with all of the dependencies they need to work. In Go, this often takes the form of passing dependencies to constructors:
```
// NewUserStore returns a UserStore that uses cfg and db as dependencies.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}

```

This technique works great at small scale, but larger applications can have a complex graph of dependencies, resulting in a big block of initialization code that’s order-dependent but otherwise not very interesting. It’s often hard to break up this code cleanly, especially because some dependencies are used multiple times. Replacing one implementation of a service with another can be painful because it involves modifying the dependency graph by adding a whole new set of dependencies (and their dependencies…), and removing unused old ones. In practice, making changes to initialization code in applications with large dependency graphs is tedious and slow.
Dependency injection tools like Wire aim to simplify the management of initialization code. You describe your services and their dependencies, either as code or as configuration, then Wire processes the resulting graph to figure out ordering and how to pass each service what it needs. Make changes to an application’s dependencies by changing a function signature or adding or removing an initializer, and then let Wire do the tedious work of generating initialization code for the entire dependency graph.
## Why is this part of Go Cloud?[¶](https://go.dev/blog/wire#why-is-this-part-of-go-cloud)
Go Cloud’s goal is to make it easier to write portable Cloud applications by providing idiomatic Go APIs for useful Cloud services. For example, [blob.Bucket](https://godoc.org/github.com/google/go-cloud/blob) provides a storage API with implementations for Amazon’s S3 and Google Cloud Storage (GCS); applications written using `blob.Bucket` can swap implementations without changing their application logic. However, the initialization code is inherently provider-specific, and each provider has a different set of dependencies.
For example, [constructing a GCS `blob.Bucket`](https://godoc.org/github.com/google/go-cloud/blob/gcsblob#OpenBucket) requires a `gcp.HTTPClient`, which eventually requires `google.Credentials`, while [constructing one for S3](https://godoc.org/github.com/google/go-cloud/blob/s3blob) requires an `aws.Config`, which eventually requires AWS credentials. Thus, updating an application to use a different `blob.Bucket` implementation involves exactly the kind of tedious update to the dependency graph that we described above. The driving use case for Wire is to make it easy to swap implementations of Go Cloud portable APIs, but it’s also a general-purpose tool for dependency injection.
## Hasn’t this been done already?[¶](https://go.dev/blog/wire#hasnt-this-been-done-already)
There are a number of dependency injection frameworks out there. For Go, [Uber’s dig](https://github.com/uber-go/dig) and [Facebook’s inject](https://github.com/facebookgo/inject) both use reflection to do runtime dependency injection. Wire was primarily inspired by Java’s [Dagger 2](https://google.github.io/dagger/), and uses code generation rather than reflection or [service locators](https://en.wikipedia.org/wiki/Service_locator_pattern).
We think this approach has several advantages:
  * Runtime dependency injection can be hard to follow and debug when the dependency graph gets complex. Using code generation means that the initialization code that’s executed at runtime is regular, idiomatic Go code that’s easy to understand and debug. Nothing is obfuscated by an intervening framework doing “magic”. In particular, problems like forgetting a dependency become compile-time errors, not run-time errors.
  * Unlike [service locators](https://en.wikipedia.org/wiki/Service_locator_pattern), there’s no need to make up arbitrary names or keys to register services. Wire uses Go types to connect components with their dependencies.
  * It’s easier to avoid dependency bloat. Wire’s generated code will only import the dependencies you need, so your binary won’t have unused imports. Runtime dependency injectors can’t identify unused dependencies until runtime.
  * Wire’s dependency graph is knowable statically, which provides opportunities for tooling and visualization.


## How does it work?[¶](https://go.dev/blog/wire#how-does-it-work)
Wire has two basic concepts: providers and injectors.
_Providers_ are ordinary Go functions that “provide” values given their dependencies, which are described simply as parameters to the function. Here’s some sample code that defines three providers:
```
// NewUserStore is the same function we saw above; it is a provider for UserStore,
// with dependencies on *Config and *mysql.DB.
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {...}
// NewDefaultConfig is a provider for *Config, with no dependencies.
func NewDefaultConfig() *Config {...}
// NewDB is a provider for *mysql.DB based on some connection info.
func NewDB(info *ConnectionInfo) (*mysql.DB, error) {...}

```

Providers that are commonly used together can be grouped into `ProviderSets`. For example, it’s common to use a default `*Config` when creating a `*UserStore`, so we can group `NewUserStore` and `NewDefaultConfig` in a `ProviderSet`:
```
var UserStoreSet = wire.ProviderSet(NewUserStore, NewDefaultConfig)

```

_Injectors_ are generated functions that call providers in dependency order. You write the injector’s signature, including any needed inputs as arguments, and insert a call to `wire.Build` with the list of providers or provider sets that are needed to construct the end result:
```
func initUserStore() (*UserStore, error) {
  // We're going to get an error, because NewDB requires a *ConnectionInfo
  // and we didn't provide one.
  wire.Build(UserStoreSet, NewDB)
  return nil, nil // These return values are ignored.
}

```

Now we run go generate to execute wire:
```
$ go generate
wire.go:2:10: inject initUserStore: no provider found for ConnectionInfo (required by provider of *mysql.DB)
wire: generate failed

```

Oops! We didn’t include a `ConnectionInfo` or tell Wire how to build one. Wire helpfully tells us the line number and types involved. We can either add a provider for it to `wire.Build`, or add it as an argument:
```
func initUserStore(info ConnectionInfo) (*UserStore, error) {
  wire.Build(UserStoreSet, NewDB)
  return nil, nil // These return values are ignored.
}

```

Now `go generate` will create a new file with the generated code:
```
// File: wire_gen.go
// Code generated by Wire. DO NOT EDIT.
//go:generate wire
//+build !wireinject
func initUserStore(info ConnectionInfo) (*UserStore, error) {
  defaultConfig := NewDefaultConfig()
  db, err := NewDB(info)
  if err != nil {
    return nil, err
  }
  userStore, err := NewUserStore(defaultConfig, db)
  if err != nil {
    return nil, err
  }
  return userStore, nil
}

```

Any non-injector declarations are copied into the generated file. There is no dependency on Wire at runtime: all of the written code is just normal Go code.
As you can see, the output is very close to what a developer would write themselves. This was a trivial example with just three components, so writing the initializer by hand wouldn’t be too painful, but Wire saves a lot of manual toil for components and applications with more complex dependency graphs.
## How can I get involved and learn more?[¶](https://go.dev/blog/wire#how-can-i-get-involved-and-learn-more)
The [Wire README](https://github.com/google/wire/blob/master/README.md) goes into more detail about how to use Wire and its more advanced features. There’s also a [tutorial](https://github.com/google/wire/tree/master/_tutorial) that walks through using Wire in a simple application.
We appreciate any input you have about your experience with Wire! [Wire’s](https://github.com/google/wire) development is conducted on GitHub, so you can [file an issue](https://github.com/google/wire/issues/new/choose) to tell us what could be better. For updates and discussion about the project, join [the Go Cloud mailing list](https://groups.google.com/forum/#!forum/go-cloud).
Thank you for taking the time to learn about Go Cloud’s Wire. We’re excited to work with you to make Go the language of choice for developers building portable cloud applications.
**Next article:**[Announcing App Engine’s New Go 1.11 Runtime](https://go.dev/blog/appengine-go111) **Previous article:**[Participate in the 2018 Go Company Questionnaire](https://go.dev/blog/survey2018-company) **[Blog Index](https://go.dev/blog/all)**
[ Why Go ](https://go.dev/solutions/) [ Use Cases ](https://go.dev/solutions/use-cases) [ Case Studies ](https://go.dev/solutions/case-studies)
[ Get Started ](https://go.dev/learn/) [ Playground ](https://go.dev/play) [ Tour ](https://go.dev/tour/) [ Stack Overflow ](https://stackoverflow.com/questions/tagged/go?tab=Newest) [ Help ](https://go.dev/help/)
[ Packages ](https://pkg.go.dev) [ Standard Library ](https://go.dev/pkg/) [ About Go Packages ](https://pkg.go.dev/about)
[ About ](https://go.dev/project) [ Download ](https://go.dev/dl/) [ Blog ](https://go.dev/blog/) [ Issue Tracker ](https://github.com/golang/go/issues) [ Release Notes ](https://go.dev/doc/devel/release) [ Brand Guidelines ](https://go.dev/brand) [ Code of Conduct ](https://go.dev/conduct)
[ Connect ](https://www.twitter.com/golang) [ Twitter ](https://www.twitter.com/golang) [ GitHub ](https://github.com/golang) [ Slack ](https://invite.slack.golangbridge.org/) [ r/golang ](https://reddit.com/r/golang) [ Meetup ](https://www.meetup.com/pro/go) [ Golang Weekly ](https://golangweekly.com/)
Opens in new window. 
![The Go Gopher](https://go.dev/images/gophers/pilot-bust.svg)
  * [Copyright](https://go.dev/copyright)
  * [Terms of Service](https://go.dev/tos)
  * [ Privacy Policy ](http://www.google.com/intl/en/policies/privacy/)
  * [ Report an Issue ](https://go.dev/s/website-issue)
  * ![System theme](https://go.dev/images/icons/brightness_6_gm_grey_24dp.svg) ![Dark theme](https://go.dev/images/icons/brightness_2_gm_grey_24dp.svg) ![Light theme](https://go.dev/images/icons/light_mode_gm_grey_24dp.svg)

[ ![Google logo](https://go.dev/images/google-white.png) ](https://google.com)
go.dev uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic. [Learn more.](https://policies.google.com/technologies/cookies)
Okay


## Source Information
- URL: https://blog.golang.org/wire
- Harvested: 2025-08-20T00:36:59.648826
- Profile: deep_research
- Agent: architect
