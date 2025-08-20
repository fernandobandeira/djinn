---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:34:02.513871'
profile: deep_research
source: https://blog.golang.org/error-handling-and-go
topic: Go Error Handling Official Blog
---

# Go Error Handling Official Blog

[ ![Go](https://go.dev/images/go-logo-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://go.dev/blog/error-handling-and-go#main-content)
  * [ Why Go _arrow_drop_down_ ](https://go.dev/blog/error-handling-and-go)
Press Enter to activate/deactivate dropdown 
    * [ Case Studies ](https://go.dev/solutions/case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions/use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [ Learn ](https://go.dev/learn/)
Press Enter to activate/deactivate dropdown 
  * [ Docs _arrow_drop_down_ ](https://go.dev/blog/error-handling-and-go)
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
  * [ Community _arrow_drop_down_ ](https://go.dev/blog/error-handling-and-go)
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
  * [Why Go _navigate_next_](https://go.dev/blog/error-handling-and-go)
[_navigate_before_ Why Go](https://go.dev/blog/error-handling-and-go)
    * [ Case Studies ](https://go.dev/solutions/case-studies)
    * [ Use Cases ](https://go.dev/solutions/use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [Docs _navigate_next_](https://go.dev/blog/error-handling-and-go)
[_navigate_before_ Docs](https://go.dev/blog/error-handling-and-go)
    * [ Go Spec ](https://go.dev/ref/spec)
    * [ Go User Manual ](https://go.dev/doc)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
    * [ Effective Go ](https://go.dev/doc/effective_go)
  * [Packages](https://pkg.go.dev)
  * [Community _navigate_next_](https://go.dev/blog/error-handling-and-go)
[_navigate_before_ Community](https://go.dev/blog/error-handling-and-go)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


# [The Go Blog](https://go.dev/blog/)
# Error handling and Go
Andrew Gerrand 12 July 2011 
## Introduction[¶](https://go.dev/blog/error-handling-and-go#introduction)
If you have written any Go code you have probably encountered the built-in `error` type. Go code uses `error` values to indicate an abnormal state. For example, the `os.Open` function returns a non-nil `error` value when it fails to open a file.
```
func Open(name string) (file *File, err error)

```

The following code uses `os.Open` to open a file. If an error occurs it calls `log.Fatal` to print the error message and stop.
```
f, err := os.Open("filename.ext")
if err != nil {
  log.Fatal(err)
}
// do something with the open *File f

```

You can get a lot done in Go knowing just this about the `error` type, but in this article we’ll take a closer look at `error` and discuss some good practices for error handling in Go.
## The error type[¶](https://go.dev/blog/error-handling-and-go#the-error-type)
The `error` type is an interface type. An `error` variable represents any value that can describe itself as a string. Here is the interface’s declaration:
```
type error interface {
  Error() string
}

```

The `error` type, as with all built in types, is [predeclared](https://go.dev/doc/go_spec.html#Predeclared_identifiers) in the [universe block](https://go.dev/doc/go_spec.html#Blocks).
The most commonly-used `error` implementation is the [errors](https://go.dev/pkg/errors/) package’s unexported `errorString` type.
```
// errorString is a trivial implementation of error.
type errorString struct {
  s string
}
func (e *errorString) Error() string {
  return e.s
}

```

You can construct one of these values with the `errors.New` function. It takes a string that it converts to an `errors.errorString` and returns as an `error` value.
```
// New returns an error that formats as the given text.
func New(text string) error {
  return &errorString{text}
}

```

Here’s how you might use `errors.New`:
```
func Sqrt(f float64) (float64, error) {
  if f < 0 {
    return 0, errors.New("math: square root of negative number")
  }
  // implementation
}

```

A caller passing a negative argument to `Sqrt` receives a non-nil `error` value (whose concrete representation is an `errors.errorString` value). The caller can access the error string (“math: square root of…”) by calling the `error`’s `Error` method, or by just printing it:
```
f, err := Sqrt(-1)
if err != nil {
  fmt.Println(err)
}

```

The [fmt](https://go.dev/pkg/fmt/) package formats an `error` value by calling its `Error() string` method.
It is the error implementation’s responsibility to summarize the context. The error returned by `os.Open` formats as “open /etc/passwd: permission denied,” not just “permission denied.” The error returned by our `Sqrt` is missing information about the invalid argument.
To add that information, a useful function is the `fmt` package’s `Errorf`. It formats a string according to `Printf`’s rules and returns it as an `error` created by `errors.New`.
```
if f < 0 {
  return 0, fmt.Errorf("math: square root of negative number %g", f)
}

```

In many cases `fmt.Errorf` is good enough, but since `error` is an interface, you can use arbitrary data structures as error values, to allow callers to inspect the details of the error.
For instance, our hypothetical callers might want to recover the invalid argument passed to `Sqrt`. We can enable that by defining a new error implementation instead of using `errors.errorString`:
```
type NegativeSqrtError float64
func (f NegativeSqrtError) Error() string {
  return fmt.Sprintf("math: square root of negative number %g", float64(f))
}

```

A sophisticated caller can then use a [type assertion](https://go.dev/doc/go_spec.html#Type_assertions) to check for a `NegativeSqrtError` and handle it specially, while callers that just pass the error to `fmt.Println` or `log.Fatal` will see no change in behavior.
As another example, the [json](https://go.dev/pkg/encoding/json/) package specifies a `SyntaxError` type that the `json.Decode` function returns when it encounters a syntax error parsing a JSON blob.
```
type SyntaxError struct {
  msg  string // description of error
  Offset int64 // error occurred after reading Offset bytes
}
func (e *SyntaxError) Error() string { return e.msg }

```

The `Offset` field isn’t even shown in the default formatting of the error, but callers can use it to add file and line information to their error messages:
```
if err := dec.Decode(&val); err != nil {
  if serr, ok := err.(*json.SyntaxError); ok {
    line, col := findLine(f, serr.Offset)
    return fmt.Errorf("%s:%d:%d: %v", f.Name(), line, col, err)
  }
  return err
}

```

(This is a slightly simplified version of some [actual code](https://github.com/camlistore/go4/blob/03efcb870d84809319ea509714dd6d19a1498483/jsonconfig/eval.go#L123-L135) from the [Camlistore](http://camlistore.org) project.)
The `error` interface requires only a `Error` method; specific error implementations might have additional methods. For instance, the [net](https://go.dev/pkg/net/) package returns errors of type `error`, following the usual convention, but some of the error implementations have additional methods defined by the `net.Error` interface:
```
package net
type Error interface {
  error
  Timeout() bool  // Is the error a timeout?
  Temporary() bool // Is the error temporary?
}

```

Client code can test for a `net.Error` with a type assertion and then distinguish transient network errors from permanent ones. For instance, a web crawler might sleep and retry when it encounters a temporary error and give up otherwise.
```
if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
  time.Sleep(1e9)
  continue
}
if err != nil {
  log.Fatal(err)
}

```

## Simplifying repetitive error handling[¶](https://go.dev/blog/error-handling-and-go#simplifying-repetitive-error-handling)
In Go, error handling is important. The language’s design and conventions encourage you to explicitly check for errors where they occur (as distinct from the convention in other languages of throwing exceptions and sometimes catching them). In some cases this makes Go code verbose, but fortunately there are some techniques you can use to minimize repetitive error handling.
Consider an [App Engine](https://cloud.google.com/appengine/docs/go/) application with an HTTP handler that retrieves a record from the datastore and formats it with a template.
```
func init() {
  http.HandleFunc("/view", viewRecord)
}
func viewRecord(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record := new(Record)
  if err := datastore.Get(c, key, record); err != nil {
    http.Error(w, err.Error(), 500)
    return
  }
  if err := viewTemplate.Execute(w, record); err != nil {
    http.Error(w, err.Error(), 500)
  }
}

```

This function handles errors returned by the `datastore.Get` function and `viewTemplate`’s `Execute` method. In both cases, it presents a simple error message to the user with the HTTP status code 500 (“Internal Server Error”). This looks like a manageable amount of code, but add some more HTTP handlers and you quickly end up with many copies of identical error handling code.
To reduce the repetition we can define our own HTTP `appHandler` type that includes an `error` return value:
```
type appHandler func(http.ResponseWriter, *http.Request) error

```

Then we can change our `viewRecord` function to return errors:
```
func viewRecord(w http.ResponseWriter, r *http.Request) error {
  c := appengine.NewContext(r)
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record := new(Record)
  if err := datastore.Get(c, key, record); err != nil {
    return err
  }
  return viewTemplate.Execute(w, record)
}

```

This is simpler than the original version, but the [http](https://go.dev/pkg/net/http/) package doesn’t understand functions that return `error`. To fix this we can implement the `http.Handler` interface’s `ServeHTTP` method on `appHandler`:
```
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if err := fn(w, r); err != nil {
    http.Error(w, err.Error(), 500)
  }
}

```

The `ServeHTTP` method calls the `appHandler` function and displays the returned error (if any) to the user. Notice that the method’s receiver, `fn`, is a function. (Go can do that!) The method invokes the function by calling the receiver in the expression `fn(w, r)`.
Now when registering `viewRecord` with the http package we use the `Handle` function (instead of `HandleFunc`) as `appHandler` is an `http.Handler` (not an `http.HandlerFunc`).
```
func init() {
  http.Handle("/view", appHandler(viewRecord))
}

```

With this basic error handling infrastructure in place, we can make it more user friendly. Rather than just displaying the error string, it would be better to give the user a simple error message with an appropriate HTTP status code, while logging the full error to the App Engine developer console for debugging purposes.
To do this we create an `appError` struct containing an `error` and some other fields:
```
type appError struct {
  Error  error
  Message string
  Code  int
}

```

Next we modify the appHandler type to return `*appError` values:
```
type appHandler func(http.ResponseWriter, *http.Request) *appError

```

(It’s usually a mistake to pass back the concrete type of an error rather than `error`, for reasons discussed in [the Go FAQ](https://go.dev/doc/go_faq.html#nil_error), but it’s the right thing to do here because `ServeHTTP` is the only place that sees the value and uses its contents.)
And make `appHandler`’s `ServeHTTP` method display the `appError`’s `Message` to the user with the correct HTTP status `Code` and log the full `Error` to the developer console:
```
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  if e := fn(w, r); e != nil { // e is *appError, not os.Error.
    c := appengine.NewContext(r)
    c.Errorf("%v", e.Error)
    http.Error(w, e.Message, e.Code)
  }
}

```

Finally, we update `viewRecord` to the new function signature and have it return more context when it encounters an error:
```
func viewRecord(w http.ResponseWriter, r *http.Request) *appError {
  c := appengine.NewContext(r)
  key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
  record := new(Record)
  if err := datastore.Get(c, key, record); err != nil {
    return &appError{err, "Record not found", 404}
  }
  if err := viewTemplate.Execute(w, record); err != nil {
    return &appError{err, "Can't display record", 500}
  }
  return nil
}

```

This version of `viewRecord` is the same length as the original, but now each of those lines has specific meaning and we are providing a friendlier user experience.
It doesn’t end there; we can further improve the error handling in our application. Some ideas:
  * give the error handler a pretty HTML template,
  * make debugging easier by writing the stack trace to the HTTP response when the user is an administrator,
  * write a constructor function for `appError` that stores the stack trace for easier debugging,
  * recover from panics inside the `appHandler`, logging the error to the console as “Critical,” while telling the user “a serious error has occurred.” This is a nice touch to avoid exposing the user to inscrutable error messages caused by programming errors. See the [Defer, Panic, and Recover](https://go.dev/doc/articles/defer_panic_recover.html) article for more details.


## Conclusion[¶](https://go.dev/blog/error-handling-and-go#conclusion)
Proper error handling is an essential requirement of good software. By employing the techniques described in this post you should be able to write more reliable and succinct Go code.
**Next article:**[Go for App Engine is now generally available](https://go.dev/blog/appengine-ga) **Previous article:**[First Class Functions in Go](https://go.dev/blog/functions-codewalk) **[Blog Index](https://go.dev/blog/all)**
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
- URL: https://blog.golang.org/error-handling-and-go
- Harvested: 2025-08-20T00:34:02.513871
- Profile: deep_research
- Agent: architect
