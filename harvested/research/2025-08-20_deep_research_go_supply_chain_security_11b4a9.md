---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:56:50.291878'
profile: deep_research
source: https://go.dev/blog/supply-chain
topic: Go Supply Chain Security
---

# Go Supply Chain Security

[ ![Go](https://go.dev/images/go-logo-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://go.dev/blog/supply-chain#main-content)
  * [ Why Go _arrow_drop_down_ ](https://go.dev/blog/supply-chain)
Press Enter to activate/deactivate dropdown 
    * [ Case Studies ](https://go.dev/solutions/case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions/use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [ Learn ](https://go.dev/learn/)
Press Enter to activate/deactivate dropdown 
  * [ Docs _arrow_drop_down_ ](https://go.dev/blog/supply-chain)
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
  * [ Community _arrow_drop_down_ ](https://go.dev/blog/supply-chain)
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
  * [Why Go _navigate_next_](https://go.dev/blog/supply-chain)
[_navigate_before_ Why Go](https://go.dev/blog/supply-chain)
    * [ Case Studies ](https://go.dev/solutions/case-studies)
    * [ Use Cases ](https://go.dev/solutions/use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [Docs _navigate_next_](https://go.dev/blog/supply-chain)
[_navigate_before_ Docs](https://go.dev/blog/supply-chain)
    * [ Go Spec ](https://go.dev/ref/spec)
    * [ Go User Manual ](https://go.dev/doc)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
    * [ Effective Go ](https://go.dev/doc/effective_go)
  * [Packages](https://pkg.go.dev)
  * [Community _navigate_next_](https://go.dev/blog/supply-chain)
[_navigate_before_ Community](https://go.dev/blog/supply-chain)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


# [The Go Blog](https://go.dev/blog/)
# How Go Mitigates Supply Chain Attacks
Filippo Valsorda 31 March 2022 
Modern software engineering is collaborative, and based on reusing Open Source software. That exposes targets to supply chain attacks, where software projects are attacked by compromising their dependencies.
Despite any process or technical measure, every dependency is unavoidably a trust relationship. However, the Go tooling and design help mitigate risk at various stages.
## All builds are “locked”[¶](https://go.dev/blog/supply-chain#all-builds-are-locked)
There is no way for changes in the outside world—such as a new version of a dependency being published—to automatically affect a Go build.
Unlike most other package managers files, Go modules don’t have a separate list of constraints and a lock file pinning specific versions. The version of every dependency contributing to any Go build is fully determined by the [`go.mod` file](https://go.dev/ref/mod#go-mod-file) of the main module.
Since Go 1.16, this determinism is enforced by default, and build commands (`go build`, `go test`, `go install`, `go run`, …) [will fail if the go.mod is incomplete](https://go.dev/ref/mod#go-mod-file-updates). The only commands that will change the `go.mod` (and therefore the build) are `go get` and `go mod tidy`. These commands are not expected to be run automatically or in CI, so changes to dependency trees must be made deliberately and have the opportunity to go through code review.
This is very important for security, because when a CI system or new machine runs `go build`, the checked-in source is the ultimate and complete source of truth for what will get built. There is no way for third parties to affect that.
Moreover, when a dependency is added with `go get`, its transitive dependencies are added at the version specified in the dependency’s `go.mod` file, not at their latest versions, thanks to [Minimal version selection](https://go.dev/ref/mod#minimal-version-selection). The same happens for invocations of `go install example.com/cmd/devtoolx@latest`, [the equivalents of which in some ecosystems bypass pinning](https://research.swtch.com/npm-colors). In Go, the latest version of `example.com/cmd/devtoolx` will be fetched, but then all the dependencies will be set by its `go.mod` file.
If a module gets compromised and a new malicious version is published, no one will be affected until they explicitly update that dependency, providing the opportunity to review the changes and time for the ecosystem to detect the event.
## Version contents never change[¶](https://go.dev/blog/supply-chain#version-contents-never-change)
Another key property necessary to ensure third parties can’t affect builds is that the contents of a module version are immutable. If an attacker that compromises a dependency could re-upload an existing version, they could automatically compromise all projects that depend on it.
That’s what the [`go.sum` file](https://go.dev/ref/mod#go-sum-files) is for. It contains a list of cryptographic hashes of each dependency that contributes to the build. Again, an incomplete `go.sum` causes an error, and only `go get` and `go mod tidy` will modify it, so any changes to it will accompany a deliberate dependency change. Other builds are guaranteed to have a full set of checksums.
This is a common feature of most lock files. Go goes beyond it with the [Checksum Database](https://go.dev/ref/mod#checksum-database) (sumdb for short), a global append-only cryptographically-verifiable list of go.sum entries. When `go get` needs to add an entry to the `go.sum` file, it fetches it from the sumdb along with cryptographic proof of the sumdb integrity. This ensures that not only every build of a certain module uses the same dependency contents, but that every module out there uses the same dependency contents!
The sumdb makes it impossible for compromised dependencies or even Google-operated Go infrastructure to target specific dependents with modified (e.g. backdoored) source. You’re guaranteed to be using the exact same code that everyone else who’s using e.g. v1.9.2 of `example.com/modulex` is using and has reviewed.
Finally, my favorite features of the sumdb: it doesn’t require any key management on the part of module authors, and it works seamlessly with the decentralized nature of Go modules.
## The VCS is the source of truth[¶](https://go.dev/blog/supply-chain#the-vcs-is-the-source-of-truth)
Most projects are developed through some version control system (VCS) and then, in other ecosystems, uploaded to the package repository. This means there are two accounts that could be compromised, the VCS host and the package repository, the latter of which is used more rarely and more likely to be overlooked. It also means it’s easier to hide malicious code in the version uploaded to the repository, especially if the source is routinely modified as part of the upload, for example to minimize it.
In Go, there is no such thing as a package repository account. The import path of a package embeds the information that `go mod download` [needs in order to fetch its module](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths) directly from the VCS, where tags define versions.
We do have the [Go Module Mirror](https://go.dev/blog/module-mirror-launch), but that’s only a proxy. Module authors don’t register an account and don’t upload versions to the proxy. The proxy uses the same logic that the `go` tool uses (in fact, the proxy runs `go mod download`) to fetch and cache a version. Since the Checksum Database guarantees that there can be only one source tree for a given module version, everyone using the proxy will see the same result as everyone bypassing it and fetching directly from the VCS. (If the version is not available anymore in the VCS or if its contents changed, fetching directly will lead to an error, while fetching from the proxy might still work, improving availability and protecting the ecosystem from [“left-pad” issues](https://blog.npmjs.org/post/141577284765/kik-left-pad-and-npm).)
Running VCS tools on the client exposes a pretty large attack surface. That’s another place the Go Module Mirror helps: the `go` tool on the proxy runs inside a robust sandbox and is configured to support every VCS tool, while [the default is to only support the two major VCS systems](https://go.dev/ref/mod#vcs-govcs) (git and Mercurial). Anyone using the proxy can still fetch code published using off-by-default VCS systems, but attackers can’t reach that code in most installations.
## Building code doesn’t execute it[¶](https://go.dev/blog/supply-chain#building-code-doesnt-execute-it)
It is an explicit security design goal of the Go toolchain that neither fetching nor building code will let that code execute, even if it is untrusted and malicious. This is different from most other ecosystems, many of which have first-class support for running code at package fetch time. These “post-install” hooks have been used in the past as the most convenient way to turn a compromised dependency into compromised developer machines, and to [worm](https://en.wikipedia.org/wiki/Computer_worm) through module authors.
To be fair, if you’re fetching some code it’s often to execute it shortly afterwards, either as part of tests on a developer machine or as part of a binary in production, so lacking post-install hooks is only going to slow down attackers. (There is no security boundary within a build: any package that contributes to a build can define an `init` function.) However, it can be a meaningful risk mitigation, since you might be executing a binary or testing a package that only uses a subset of the module’s dependencies. For example, if you build and execute `example.com/cmd/devtoolx` on macOS there is no way for a Windows-only dependency or a dependency of `example.com/cmd/othertool` to compromise your machine.
In Go, modules that don’t contribute code to a specific build have no security impact on it.
## “A little copying is better than a little dependency”[¶](https://go.dev/blog/supply-chain#a-little-copying-is-better-than-a-little-dependency)
The final and maybe most important software supply chain risk mitigation in the Go ecosystem is the least technical one: Go has a culture of rejecting large dependency trees, and of preferring a bit of copying to adding a new dependency. It goes all the way back to one of the Go proverbs: [“a little copying is better than a little dependency”](https://youtube.com/clip/UgkxWCEmMJFW0-TvSMzcMEAHZcpt2FsVXP65). The label “zero dependencies” is proudly worn by high-quality reusable Go modules. If you find yourself in need of a library, you’re likely to find it will not cause you to take on a dependency on dozens of other modules by other authors and owners.
That’s enabled also by the rich standard library and additional modules (the `golang.org/x/...` ones), which provide commonly used high-level building blocks such as an HTTP stack, a TLS library, JSON encoding, etc.
All together this means it’s possible to build rich, complex applications with just a handful of dependencies. No matter how good the tooling is, it can’t eliminate the risk involved in reusing code, so the strongest mitigation will always be a small dependency tree.
**Next article:**[Get familiar with workspaces](https://go.dev/blog/get-familiar-with-workspaces) **Previous article:**[An Introduction To Generics](https://go.dev/blog/intro-generics) **[Blog Index](https://go.dev/blog/all)**
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
- URL: https://go.dev/blog/supply-chain
- Harvested: 2025-08-20T00:56:50.291878
- Profile: deep_research
- Agent: architect
