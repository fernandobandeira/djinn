---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:56:57.758963'
profile: deep_research
source: https://blog.golang.org/module-mirror-launch
topic: Go Module Proxy and Security
---

# Go Module Proxy and Security

[ ![Go](https://go.dev/images/go-logo-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://go.dev/blog/module-mirror-launch#main-content)
  * [ Why Go _arrow_drop_down_ ](https://go.dev/blog/module-mirror-launch)
Press Enter to activate/deactivate dropdown 
    * [ Case Studies ](https://go.dev/solutions/case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions/use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [ Learn ](https://go.dev/learn/)
Press Enter to activate/deactivate dropdown 
  * [ Docs _arrow_drop_down_ ](https://go.dev/blog/module-mirror-launch)
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
  * [ Community _arrow_drop_down_ ](https://go.dev/blog/module-mirror-launch)
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
  * [Why Go _navigate_next_](https://go.dev/blog/module-mirror-launch)
[_navigate_before_ Why Go](https://go.dev/blog/module-mirror-launch)
    * [ Case Studies ](https://go.dev/solutions/case-studies)
    * [ Use Cases ](https://go.dev/solutions/use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [Docs _navigate_next_](https://go.dev/blog/module-mirror-launch)
[_navigate_before_ Docs](https://go.dev/blog/module-mirror-launch)
    * [ Go Spec ](https://go.dev/ref/spec)
    * [ Go User Manual ](https://go.dev/doc)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
    * [ Effective Go ](https://go.dev/doc/effective_go)
  * [Packages](https://pkg.go.dev)
  * [Community _navigate_next_](https://go.dev/blog/module-mirror-launch)
[_navigate_before_ Community](https://go.dev/blog/module-mirror-launch)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


# [The Go Blog](https://go.dev/blog/)
# Module Mirror and Checksum Database Launched
Katie Hockman 29 August 2019 
We are excited to share that our module [mirror](https://proxy.golang.org), [index](https://index.golang.org), and [checksum database](https://sum.golang.org) are now production ready! The `go` command will use the module mirror and checksum database by default for [Go 1.13 module users](https://go.dev/doc/go1.13#introduction). See [proxy.golang.org/privacy](https://proxy.golang.org/privacy) for privacy information about these services and the [go command documentation](https://go.dev/cmd/go/#hdr-Module_downloading_and_verification) for configuration details, including how to disable the use of these servers or use different ones. If you depend on non-public modules, see the [documentation for configuring your environment](https://go.dev/cmd/go/#hdr-Module_configuration_for_non_public_modules).
This post will describe these services and the benefits of using them, and summarizes some of the points from the [Go Module Proxy: Life of a Query](https://youtu.be/KqTySYYhPUE) talk at Gophercon 2019. See the [recording](https://youtu.be/KqTySYYhPUE) if you are interested in the full talk.
## Module Mirror[¶](https://go.dev/blog/module-mirror-launch#module-mirror)
[Modules](https://go.dev/blog/versioning-proposal) are sets of Go packages that are versioned together, and the contents of each version are immutable. That immutability provides new opportunities for caching and authentication. When `go get` runs in module mode, it must fetch the module containing the requested packages, as well as any new dependencies introduced by that module, updating your [go.mod](https://go.dev/cmd/go/#hdr-The_go_mod_file) and [go.sum](https://go.dev/cmd/go/#hdr-Module_downloading_and_verification) files as needed. Fetching modules from version control can be expensive in terms of latency and storage in your system: the `go` command may be forced to pull down the full commit history of a repository containing a transitive dependency, even one that isn’t being built, just to resolve its version.
The solution is to use a module proxy, which speaks an API that is better suited to the `go` command’s needs (see `go help goproxy`). When `go get` runs in module mode with a proxy, it will work faster by only asking for the specific module metadata or source code it needs, and not worrying about the rest. Below is an example of how the `go` command may use a proxy with `go get` by requesting the list of versions, then the info, mod, and zip file for the latest tagged version.
![](https://go.dev/blog/module-mirror-launch/proxy-protocol.png)
A module mirror is a special kind of module proxy that caches metadata and source code in its own storage system, allowing the mirror to continue to serve source code that is no longer available from the original locations. This can speed up downloads and protect you from disappearing dependencies. See [Go Modules in 2019](https://go.dev/blog/modules2019) for more information.
The Go team maintains a module mirror, served at [proxy.golang.org](https://proxy.golang.org), which the `go` command will use by default for module users as of Go 1.13. If you are running an earlier version of the `go` command, then you can use this service by setting `GOPROXY=https://proxy.golang.org` in your local environment.
## Checksum Database[¶](https://go.dev/blog/module-mirror-launch#checksum-database)
Modules introduced the `go.sum` file, which is a list of SHA-256 hashes of the source code and `go.mod` files of each dependency when it was first downloaded. The `go` command can use the hashes to detect misbehavior by an origin server or proxy that gives you different code for the same version.
The limitation of this `go.sum` file is that it works entirely by trust on _your_ first use. When you add a version of a dependency that you’ve never seen before to your module (possibly by upgrading an existing dependency), the `go` command fetches the code and adds lines to the `go.sum` file on the fly. The problem is that those `go.sum` lines aren’t being checked against anyone else’s: they might be different from the `go.sum` lines that the `go` command just generated for someone else, perhaps because a proxy intentionally served malicious code targeted to you.
Go’s solution is a global source of `go.sum` lines, called a [checksum database](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database), which ensures that the `go` command always adds the same lines to everyone’s `go.sum` file. Whenever the `go` command receives new source code, it can verify the hash of that code against this global database to make sure the hashes match, ensuring that everyone is using the same code for a given version.
The checksum database is served by [sum.golang.org](https://sum.golang.org), and is built on a [Transparent Log](https://research.swtch.com/tlog) (or “Merkle tree”) of hashes backed by [Trillian](https://github.com/google/trillian). The main advantage of a Merkle tree is that it is tamper proof and has properties that don’t allow for misbehavior to go undetected, which makes it more trustworthy than a simple database. The `go` command uses this tree to check “inclusion” proofs (that a specific record exists in the log) and “consistency” proofs (that the tree hasn’t been tampered with) before adding new `go.sum` lines to your module’s `go.sum` file. Below is an example of such a tree.
![](https://go.dev/blog/module-mirror-launch/tree.png)
The checksum database supports [a set of endpoints](https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md#checksum-database) used by the `go` command to request and verify `go.sum` lines. The `/lookup` endpoint provides a “signed tree head” (STH) and the requested `go.sum` lines. The `/tile` endpoint provides chunks of the tree called _tiles_ which the `go` command can use for proofs. Below is an example of how the `go` command may interact with the checksum database by doing a `/lookup` of a module version, then requesting the tiles required for the proofs.
![](https://go.dev/blog/module-mirror-launch/sumdb-protocol.png)
This checksum database allows the `go` command to safely use an otherwise untrusted proxy. Because there is an auditable security layer sitting on top of it, a proxy or origin server can’t intentionally, arbitrarily, or accidentally start giving you the wrong code without getting caught. Even the author of a module can’t move their tags around or otherwise change the bits associated with a specific version from one day to the next without the change being detected.
If you are using Go 1.12 or earlier, you can manually check a `go.sum` file against the checksum database with [gosumcheck](https://godoc.org/golang.org/x/mod/gosumcheck):
```
$ go get golang.org/x/mod/gosumcheck
$ gosumcheck /path/to/go.sum

```

In addition to verification done by the `go` command, third-party auditors can hold the checksum database accountable by iterating over the log looking for bad entries. They can work together and gossip about the state of the tree as it grows to ensure that it remains uncompromised, and we hope that the Go community will run them.
## Module Index[¶](https://go.dev/blog/module-mirror-launch#module-index)
The module index is served by [index.golang.org](https://index.golang.org), and is a public feed of new module versions that become available through [proxy.golang.org](https://proxy.golang.org). This is particularly useful for tool developers that want to keep their own cache of what’s available in [proxy.golang.org](https://proxy.golang.org), or keep up-to-date on some of the newest modules that people are using.
## Feedback or bugs[¶](https://go.dev/blog/module-mirror-launch#feedback-or-bugs)
We hope these services improve your experience with modules, and encourage you to [file issues](https://go.dev/issue/new?title=proxy.golang.org) if you run into problems or have feedback!
**Next article:**[Go 1.13 is released](https://go.dev/blog/go1.13) **Previous article:**[Migrating to Go Modules](https://go.dev/blog/migrating-to-go-modules) **[Blog Index](https://go.dev/blog/all)**
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
- URL: https://blog.golang.org/module-mirror-launch
- Harvested: 2025-08-20T00:56:57.758963
- Profile: deep_research
- Agent: architect
