---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:32:55.055960'
profile: api_documentation
source: https://pkg.go.dev/errors
topic: Go Standard Library Errors Package
---

# Go Standard Library Errors Package

[ ![Go](https://pkg.go.dev/static/shared/logo/go-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://pkg.go.dev/errors#main-content)
![](https://pkg.go.dev/static/shared/icon/search_gm_grey_24dp.svg) ![](https://pkg.go.dev/static/shared/icon/search_gm_grey_24dp.svg)
  * [ Why Go ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/errors)
    * [ Case Studies ](https://go.dev/solutions#case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions#use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [Learn](https://go.dev/learn/)
  * [ Docs ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/errors)
    * [ Effective Go ](https://go.dev/doc/effective_go)
Tips for writing clear, performant, and idiomatic Go code
    * [ Go User Manual ](https://go.dev/doc/)
A complete introduction to building software with Go
    * [ Standard library ](https://pkg.go.dev/std)
Reference documentation for Go's standard library
    * [ Release Notes ](https://go.dev/doc/devel/release)
Learn what's new in each Go release
  * [Packages](https://pkg.go.dev/)
  * [ Community ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/errors)
    * [ Recorded Talks ](https://go.dev/talks/)
Videos from prior events
    * [ Meetups _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://www.meetup.com/pro/go)
Meet other local Go developers
    * [ Conferences _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://github.com/golang/go/wiki/Conferences)
Learn and network with Go developers from around the world
    * [ Go blog ](https://go.dev/blog)
The Go project's official blog.
    * [ Go project ](https://go.dev/help)
Get help and stay informed from Go
    * Get connected 
[ ![](https://pkg.go.dev/static/shared/logo/social/google-groups.svg) ](https://groups.google.com/g/golang-nuts "Get connected with google-groups \(Opens in new window\)") [ ![](https://pkg.go.dev/static/shared/logo/social/github.svg) ](https://github.com/golang "Get connected with github \(Opens in new window\)") [ ![](https://pkg.go.dev/static/shared/logo/social/twitter.svg) ](https://twitter.com/golang "Get connected with twitter \(Opens in new window\)") [ ![](https://pkg.go.dev/static/shared/logo/social/reddit.svg) ](https://www.reddit.com/r/golang/ "Get connected with reddit \(Opens in new window\)") [ ![](https://pkg.go.dev/static/shared/logo/social/slack.svg) ](https://invite.slack.golangbridge.org/ "Get connected with slack \(Opens in new window\)") [ ![](https://pkg.go.dev/static/shared/logo/social/stack-overflow.svg) ](https://stackoverflow.com/collectives/go)


[ ![Go.](https://pkg.go.dev/static/shared/logo/go-blue.svg) ](https://go.dev/)
  * [ Why Go _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/errors)
[ _![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Why Go ](https://pkg.go.dev/errors)
    * [ Case Studies ](https://go.dev/solutions#case-studies)
    * [ Use Cases ](https://go.dev/solutions#use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [ Docs _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/errors)
[_![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Docs ](https://pkg.go.dev/errors)
    * [ Effective Go ](https://go.dev/doc/effective_go)
    * [ Go User Manual ](https://go.dev/doc/)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
  * [Packages](https://pkg.go.dev/)
  * [ Community _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/errors)
[ _![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Community ](https://pkg.go.dev/errors)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://www.meetup.com/pro/go)
    * [ Conferences _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://github.com/golang/go/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://pkg.go.dev/static/shared/logo/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://pkg.go.dev/static/shared/logo/social/github.svg)](https://github.com/golang) [![](https://pkg.go.dev/static/shared/logo/social/twitter.svg)](https://twitter.com/golang) [![](https://pkg.go.dev/static/shared/logo/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://pkg.go.dev/static/shared/logo/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://pkg.go.dev/static/shared/logo/social/stack-overflow.svg)](https://stackoverflow.com/collectives/go)


  1. [Discover Packages](https://pkg.go.dev/)
  2. [Standard library](https://pkg.go.dev/std)
  3. [ errors ](https://pkg.go.dev/errors@go1.25.0) ![](https://pkg.go.dev/static/shared/icon/content_copy_gm_grey_24dp.svg)


[ ![Go](https://pkg.go.dev/static/shared/logo/go-blue.svg) ](https://go.dev/)
# errors
package standard library ![](https://pkg.go.dev/static/shared/icon/content_copy_gm_grey_24dp.svg)
[ Version:  go1.25.0 ](https://pkg.go.dev/errors?tab=versions)
Opens a new window with list of versions in this module. 
Latest Latest  ![Warning](https://pkg.go.dev/static/shared/icon/alert_gm_grey_24dp.svg)
This package is not in the latest version of its module.
[ Go to latest ](https://pkg.go.dev/errors) Published: Aug 12, 2025  License: [BSD-3-Clause](https://pkg.go.dev/errors?tab=licenses)
Opens a new window with license information. 
[ Imports: 2 ](https://pkg.go.dev/errors?tab=imports)
Opens a new window with list of imports. 
[ Imported by: 1,524,852 ](https://pkg.go.dev/errors?tab=importedby)
Opens a new window with list of known importers. 
Main Versions  Licenses  Imports  Imported By 
## Details
  * ![checked](https://pkg.go.dev/static/shared/icon/check_circle_gm_grey_24dp.svg) Valid [go.mod](https://cs.opensource.google/go/go/+/go1.25.0:src/go.mod) file ![](https://pkg.go.dev/static/shared/icon/help_gm_grey_24dp.svg)
The Go module system was introduced in Go 1.11 and is the official dependency management solution for Go. 
  * ![checked](https://pkg.go.dev/static/shared/icon/check_circle_gm_grey_24dp.svg) Redistributable license ![](https://pkg.go.dev/static/shared/icon/help_gm_grey_24dp.svg)
Redistributable licenses place minimal restrictions on how software can be used, modified, and redistributed. 
  * ![checked](https://pkg.go.dev/static/shared/icon/check_circle_gm_grey_24dp.svg) Tagged version ![](https://pkg.go.dev/static/shared/icon/help_gm_grey_24dp.svg)
Modules with tagged versions give importers more predictable builds.
  * ![checked](https://pkg.go.dev/static/shared/icon/check_circle_gm_grey_24dp.svg) Stable version ![](https://pkg.go.dev/static/shared/icon/help_gm_grey_24dp.svg)
When a project reaches major version v1 it is considered stable.
  * [Learn more about best practices](https://pkg.go.dev/about#best-practices)


## Repository
[ cs.opensource.google/go/go ](https://cs.opensource.google/go/go "https://cs.opensource.google/go/go")
## Links
  * [ ![](https://pkg.go.dev/static/shared/icon/security_grey_24dp.svg) Report a Vulnerability ](https://go.dev/security/policy "Report security issues in the Go standard library and sub-repositories")


Jump to ... 
  * [ Documentation ](https://pkg.go.dev/errors#section-documentation)
    * [Overview](https://pkg.go.dev/errors#pkg-overview)
    * [ Index ](https://pkg.go.dev/errors#pkg-index)
      * [ Examples ](https://pkg.go.dev/errors#pkg-examples)
    * [ Constants ](https://pkg.go.dev/errors#pkg-constants)
    * [ Variables ](https://pkg.go.dev/errors#pkg-variables)
    * [ Functions ](https://pkg.go.dev/errors#pkg-functions)
      * [ As(err, target) ](https://pkg.go.dev/errors#As "As\(err, target\)")
      * [ Is(err, target) ](https://pkg.go.dev/errors#Is "Is\(err, target\)")
      * [ Join(errs) ](https://pkg.go.dev/errors#Join "Join\(errs\)")
      * [ New(text) ](https://pkg.go.dev/errors#New "New\(text\)")
      * [ Unwrap(err) ](https://pkg.go.dev/errors#Unwrap "Unwrap\(err\)")
    * [ Types ](https://pkg.go.dev/errors#pkg-types)
  * [ Source Files ](https://pkg.go.dev/errors#section-sourcefiles)


DocumentationSource FilesOverviewIndexConstantsVariablesFunctionsTypesExamplesAs(err, target)Is(err, target)Join(errs)New(text)Unwrap(err)
##  ![](https://pkg.go.dev/static/shared/icon/code_gm_grey_24dp.svg) Documentation [¶](https://pkg.go.dev/errors#section-documentation "Go to Documentation")
### Overview [¶](https://pkg.go.dev/errors#pkg-overview "Go to Overview")
Package errors implements functions to manipulate errors. 
The [New](https://pkg.go.dev/errors#New) function creates errors whose only content is a text message. 
An error e wraps another error if e's type has one of the methods 
```
Unwrap() error
Unwrap() []error

```

If e.Unwrap() returns a non-nil error w or a slice containing w, then we say that e wraps w. A nil error returned from e.Unwrap() indicates that e does not wrap any error. It is invalid for an Unwrap method to return an []error containing a nil error value. 
An easy way to create wrapped errors is to call [fmt.Errorf](https://pkg.go.dev/fmt#Errorf) and apply the %w verb to the error argument: 
```
wrapsErr := fmt.Errorf("... %w ...", ..., err, ...)

```

Successive unwrapping of an error creates a tree. The [Is](https://pkg.go.dev/errors#Is) and [As](https://pkg.go.dev/errors#As) functions inspect an error's tree by examining first the error itself followed by the tree of each of its children in turn (pre-order, depth-first traversal). 
See <https://go.dev/blog/go1.13-errors> for a deeper discussion of the philosophy of wrapping and when to wrap. 
[Is](https://pkg.go.dev/errors#Is) examines the tree of its first argument looking for an error that matches the second. It reports whether it finds a match. It should be used in preference to simple equality checks: 
```
if errors.Is(err, fs.ErrExist)

```

is preferable to 
```
if err == fs.ErrExist

```

because the former will succeed if err wraps [io/fs.ErrExist](https://pkg.go.dev/io/fs#ErrExist). 
[As](https://pkg.go.dev/errors#As) examines the tree of its first argument looking for an error that can be assigned to its second argument, which must be a pointer. If it succeeds, it performs the assignment and returns true. Otherwise, it returns false. The form 
```
var perr *fs.PathError
if errors.As(err, &perr) {
	fmt.Println(perr.Path)
}

```

is preferable to 
```
if perr, ok := err.(*fs.PathError); ok {
	fmt.Println(perr.Path)
}

```

because the former will succeed if err wraps an [*io/fs.PathError](https://pkg.go.dev/io/fs#PathError). 
Example [¶](https://pkg.go.dev/errors#example-package "Go to Example")
```
Output:
1989-03-15 22:30:00 +0000 UTC: the file system has gone away

```

Share Format Run
### Index [¶](https://pkg.go.dev/errors#pkg-index "Go to Index")
  * [Variables](https://pkg.go.dev/errors#pkg-variables)
  * [func As(err error, target any) bool](https://pkg.go.dev/errors#As)
  * [func Is(err, target error) bool](https://pkg.go.dev/errors#Is)
  * [func Join(errs ...error) error](https://pkg.go.dev/errors#Join)
  * [func New(text string) error](https://pkg.go.dev/errors#New)
  * [func Unwrap(err error) error](https://pkg.go.dev/errors#Unwrap)


#### Examples [¶](https://pkg.go.dev/errors#pkg-examples "Go to Examples")
  * [Package](https://pkg.go.dev/errors#example-package)
  * [As](https://pkg.go.dev/errors#example-As)
  * [Is](https://pkg.go.dev/errors#example-Is)
  * [Join](https://pkg.go.dev/errors#example-Join)
  * [New](https://pkg.go.dev/errors#example-New)
  * [New (Errorf)](https://pkg.go.dev/errors#example-New-Errorf)
  * [Unwrap](https://pkg.go.dev/errors#example-Unwrap)


### Constants [¶](https://pkg.go.dev/errors#pkg-constants "Go to Constants")
This section is empty.
### Variables [¶](https://pkg.go.dev/errors#pkg-variables "Go to Variables")
[View Source](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/errors.go;l=90) ```
var ErrUnsupported = New[](https://pkg.go.dev/errors#New)("unsupported operation")
```

ErrUnsupported indicates that a requested operation cannot be performed, because it is unsupported. For example, a call to [os.Link](https://pkg.go.dev/os#Link) when using a file system that does not support hard links. 
Functions and methods should not return this error but should instead return an error including appropriate context that satisfies 
```
errors.Is(err, errors.ErrUnsupported)

```

either by directly wrapping ErrUnsupported or by implementing an [Is](https://pkg.go.dev/errors#Is) method. 
Functions and methods should document the cases in which an error wrapping this will be returned. 
### Functions [¶](https://pkg.go.dev/errors#pkg-functions "Go to Functions")
####  func [As](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/wrap.go;l=97) [¶](https://pkg.go.dev/errors#As "Go to As") added in go1.13
```
func As(err error[](https://pkg.go.dev/builtin#error), target any[](https://pkg.go.dev/builtin#any)) bool[](https://pkg.go.dev/builtin#bool)
```

As finds the first error in err's tree that matches target, and if one is found, sets target to that error value and returns true. Otherwise, it returns false. 
The tree consists of err itself, followed by the errors obtained by repeatedly calling its Unwrap() error or Unwrap() []error method. When err wraps multiple errors, As examines err followed by a depth-first traversal of its children. 
An error matches target if the error's concrete value is assignable to the value pointed to by target, or if the error has a method As(any) bool such that As(target) returns true. In the latter case, the As method is responsible for setting target. 
An error type might provide an As method so it can be treated as if it were a different error type. 
As panics if target is not a non-nil pointer to either a type that implements error, or to any interface type. 
Example [¶](https://pkg.go.dev/errors#example-As "Go to Example")
```
Output:
Failed at path: non-existing

```

Share Format Run
####  func [Is](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/wrap.go;l=44) [¶](https://pkg.go.dev/errors#Is "Go to Is") added in go1.13
```
func Is(err, target error[](https://pkg.go.dev/builtin#error)) bool[](https://pkg.go.dev/builtin#bool)
```

Is reports whether any error in err's tree matches target. 
The tree consists of err itself, followed by the errors obtained by repeatedly calling its Unwrap() error or Unwrap() []error method. When err wraps multiple errors, Is examines err followed by a depth-first traversal of its children. 
An error is considered to match a target if it is equal to that target or if it implements a method Is(error) bool such that Is(target) returns true. 
An error type might provide an Is method so it can be treated as equivalent to an existing error. For example, if MyError defines 
```
func (m MyError) Is(target error) bool { return target == fs.ErrExist }

```

then Is(MyError{}, fs.ErrExist) returns true. See [syscall.Errno.Is](https://pkg.go.dev/syscall#Errno.Is) for an example in the standard library. An Is method should only shallowly compare err and the target and not call [Unwrap](https://pkg.go.dev/errors#Unwrap) on either. 
Example [¶](https://pkg.go.dev/errors#example-Is "Go to Example")
```
Output:
file does not exist

```

Share Format Run
####  func [Join](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/join.go;l=19) [¶](https://pkg.go.dev/errors#Join "Go to Join") added in go1.20
```
func Join(errs ...error[](https://pkg.go.dev/builtin#error)) error[](https://pkg.go.dev/builtin#error)
```

Join returns an error that wraps the given errors. Any nil error values are discarded. Join returns nil if every value in errs is nil. The error formats as the concatenation of the strings obtained by calling the Error method of each element of errs, with a newline between each string. 
A non-nil error returned by Join implements the Unwrap() []error method. 
Example [¶](https://pkg.go.dev/errors#example-Join "Go to Example")
```
Output:
err1
err2
err is err1
err is err2
[err1 err2]

```

Share Format Run
####  func [New](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/errors.go;l=64) [¶](https://pkg.go.dev/errors#New "Go to New")
```
func New(text string[](https://pkg.go.dev/builtin#string)) error[](https://pkg.go.dev/builtin#error)
```

New returns an error that formats as the given text. Each call to New returns a distinct error value even if the text is identical. 
Example [¶](https://pkg.go.dev/errors#example-New "Go to Example")
```
Output:
emit macho dwarf: elf header corrupted

```

Share Format Run
Example (Errorf) [¶](https://pkg.go.dev/errors#example-New-Errorf "Go to Example \(Errorf\)")
The fmt package's Errorf function lets us use the package's formatting features to create descriptive error messages. 
```
Output:
user "bimmler" (id 17) not found

```

Share Format Run
####  func [Unwrap](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/wrap.go;l=17) [¶](https://pkg.go.dev/errors#Unwrap "Go to Unwrap") added in go1.13
```
func Unwrap(err error[](https://pkg.go.dev/builtin#error)) error[](https://pkg.go.dev/builtin#error)
```

Unwrap returns the result of calling the Unwrap method on err, if err's type contains an Unwrap method returning error. Otherwise, Unwrap returns nil. 
Unwrap only calls a method of the form "Unwrap() error". In particular Unwrap does not unwrap errors returned by [Join](https://pkg.go.dev/errors#Join). 
Example [¶](https://pkg.go.dev/errors#example-Unwrap "Go to Example")
```
Output:
error2: [error1]
error1

```

Share Format Run
### Types [¶](https://pkg.go.dev/errors#pkg-types "Go to Types")
This section is empty.
##  ![](https://pkg.go.dev/static/shared/icon/insert_drive_file_gm_grey_24dp.svg) Source Files [¶](https://pkg.go.dev/errors#section-sourcefiles "Go to Source Files")
[View all Source files](https://cs.opensource.google/go/go/+/go1.25.0:src/errors)
  * [errors.go](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/errors.go "errors.go")
  * [join.go](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/join.go "join.go")
  * [wrap.go](https://cs.opensource.google/go/go/+/go1.25.0:src/errors/wrap.go "wrap.go")


Click to show internal directories. 
Click to hide internal directories. 
[ Why Go ](https://go.dev/solutions) [ Use Cases ](https://go.dev/solutions#use-cases) [ Case Studies ](https://go.dev/solutions#case-studies)
[ Get Started ](https://learn.go.dev/) [ Playground ](https://play.golang.org) [ Tour ](https://tour.golang.org) [ Stack Overflow ](https://stackoverflow.com/questions/tagged/go?tab=Newest) [ Help ](https://go.dev/help)
[ Packages ](https://pkg.go.dev) [ Standard Library ](https://pkg.go.dev/std) [ Sub-repositories ](https://pkg.go.dev/golang.org/x) [ About Go Packages ](https://pkg.go.dev/about)
[ About ](https://go.dev/project) [Download](https://go.dev/dl/) [Blog](https://go.dev/blog) [ Issue Tracker ](https://github.com/golang/go/issues) [ Release Notes ](https://go.dev/doc/devel/release.html) [ Brand Guidelines ](https://go.dev/brand) [ Code of Conduct ](https://go.dev/conduct)
[ Connect ](https://www.twitter.com/golang) [ Twitter ](https://www.twitter.com/golang) [GitHub](https://github.com/golang) [ Slack ](https://invite.slack.golangbridge.org/) [ r/golang ](https://reddit.com/r/golang) [ Meetup ](https://www.meetup.com/pro/go) [ Golang Weekly ](https://golangweekly.com/)
![Gopher in flight goggles](https://pkg.go.dev/static/shared/gopher/pilot-bust-1431x901.svg)
  * [Copyright](https://go.dev/copyright)
  * [Terms of Service](https://go.dev/tos)
  * [ Privacy Policy ](http://www.google.com/intl/en/policies/privacy/)
  * [ Report an Issue ](https://go.dev/s/pkgsite-feedback)
  * ![System theme](https://pkg.go.dev/static/shared/icon/brightness_6_gm_grey_24dp.svg) ![Dark theme](https://pkg.go.dev/static/shared/icon/brightness_2_gm_grey_24dp.svg) ![Light theme](https://pkg.go.dev/static/shared/icon/light_mode_gm_grey_24dp.svg)
Theme Toggle 
  * ![](https://pkg.go.dev/static/shared/icon/keyboard_grey_24dp.svg)
Shortcuts Modal 

[ ![Google logo](https://pkg.go.dev/static/shared/logo/google-white.svg) ](https://google.com)
## Jump to
![](https://pkg.go.dev/static/shared/icon/close_gm_grey_24dp.svg)
Close
## Keyboard shortcuts
![](https://pkg.go.dev/static/shared/icon/close_gm_grey_24dp.svg)
**?**|  : This menu  
---|---  
**/**|  : Search site  
**f** or **F**|  : Jump to  
**y** or **Y** |  : Canonical URL  
Close
go.dev uses cookies from Google to deliver and enhance the quality of its services and to analyze traffic. [Learn more.](https://policies.google.com/technologies/cookies)
Okay


## Source Information
- URL: https://pkg.go.dev/errors
- Harvested: 2025-08-20T00:32:55.055960
- Profile: api_documentation
- Agent: architect
