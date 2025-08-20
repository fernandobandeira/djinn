---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:09:11.827090'
profile: deep_research
source: https://golang.org/doc/tutorial/add-a-test
topic: Go Testing Frameworks and Practices 2025
---

# Go Testing Frameworks and Practices 2025

[ ![Go](https://go.dev/images/go-logo-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://go.dev/doc/tutorial/add-a-test#main-content)
  * [ Why Go _arrow_drop_down_ ](https://go.dev/doc/tutorial/add-a-test)
Press Enter to activate/deactivate dropdown 
    * [ Case Studies ](https://go.dev/solutions/case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions/use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [ Learn ](https://go.dev/learn/)
Press Enter to activate/deactivate dropdown 
  * [ Docs _arrow_drop_down_ ](https://go.dev/doc/tutorial/add-a-test)
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
  * [ Community _arrow_drop_down_ ](https://go.dev/doc/tutorial/add-a-test)
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
  * [Why Go _navigate_next_](https://go.dev/doc/tutorial/add-a-test)
[_navigate_before_ Why Go](https://go.dev/doc/tutorial/add-a-test)
    * [ Case Studies ](https://go.dev/solutions/case-studies)
    * [ Use Cases ](https://go.dev/solutions/use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [Docs _navigate_next_](https://go.dev/doc/tutorial/add-a-test)
[_navigate_before_ Docs](https://go.dev/doc/tutorial/add-a-test)
    * [ Go Spec ](https://go.dev/ref/spec)
    * [ Go User Manual ](https://go.dev/doc)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
    * [ Effective Go ](https://go.dev/doc/effective_go)
  * [Packages](https://pkg.go.dev)
  * [Community _navigate_next_](https://go.dev/doc/tutorial/add-a-test)
[_navigate_before_ Community](https://go.dev/doc/tutorial/add-a-test)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _open_in_new_ ](https://www.meetup.com/pro/go)
    * [ Conferences _open_in_new_ ](https://go.dev/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://go.dev/images/logos/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://go.dev/images/logos/social/github.svg)](https://github.com/golang) [![](https://go.dev/images/logos/social/twitter.svg)](https://twitter.com/golang) [![](https://go.dev/images/logos/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://go.dev/images/logos/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://go.dev/images/logos/social/stack-overflow.svg)](https://stackoverflow.com/tags/go)


  1. [ Documentation ](https://go.dev/doc/)
  2. [ Tutorials ](https://go.dev/doc/tutorial/)
  3. [ Add a test ](https://go.dev/doc/tutorial/add-a-test)


# Add a test
Now that you've gotten your code to a stable place (nicely done, by the way), add a test. Testing your code during development can expose bugs that find their way in as you make changes. In this topic, you add a test for the `Hello` function. 
**Note:** This topic is part of a multi-part tutorial that begins with [Create a Go module](https://go.dev/doc/tutorial/create-module.html). 
Go's built-in support for unit testing makes it easier to test as you go. Specifically, using naming conventions, Go's `testing` package, and the `go test` command, you can quickly write and execute tests. 
  1. In the greetings directory, create a file called greetings_test.go. 
Ending a file's name with _test.go tells the `go test` command that this file contains test functions. 
  2. In greetings_test.go, paste the following code and save the file. ```
package greetings
import (
  "testing"
  "regexp"
)
// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
  name := "Gladys"
  want := regexp.MustCompile(`\b`+name+`\b`)
  msg, err := Hello("Gladys")
  if !want.MatchString(msg) || err != nil {
    t.Errorf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
  }
}
// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
  msg, err := Hello("")
  if msg != "" || err == nil {
    t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
  }
}

```

In this code, you: 
     * Implement test functions in the same package as the code you're testing. 
     * Create two test functions to test the `greetings.Hello` function. Test function names have the form `Test_Name_`, where _Name_ says something about the specific test. Also, test functions take a pointer to the `testing` package's [`testing.T` type](https://go.dev/pkg/testing/#T) as a parameter. You use this parameter's methods for reporting and logging from your test. 
     * Implement two tests: 
       * `TestHelloName` calls the `Hello` function, passing a `name` value with which the function should be able to return a valid response message. If the call returns an error or an unexpected response message (one that doesn't include the name you passed in), you use the `t` parameter's [ `Errorf` method](https://go.dev/pkg/testing/#T.Errorf) to print a message to the console. 
       * `TestHelloEmpty` calls the `Hello` function with an empty string. This test is designed to confirm that your error handling works. If the call returns a non-empty string or no error, you use the `t` parameter's [`Errorf       ` method](https://go.dev/pkg/testing/#T.Errorf) to print a message to the console. 
  3. At the command line in the greetings directory, run the [`go test` command](https://go.dev/cmd/go/#hdr-Test_packages) to execute the test. 
The `go test` command executes test functions (whose names begin with `Test`) in test files (whose names end with _test.go). You can add the `-v` flag to get verbose output that lists all of the tests and their results. 
The tests should pass. 
```
$ go test
PASS
ok   example.com/greetings  0.364s
$ go test -v
=== RUN  TestHelloName
--- PASS: TestHelloName (0.00s)
=== RUN  TestHelloEmpty
--- PASS: TestHelloEmpty (0.00s)
PASS
ok   example.com/greetings  0.372s

```

  4. Break the `greetings.Hello` function to view a failing test. 
The `TestHelloName` test function checks the return value for the name you specified as a `Hello` function parameter. To view a failing test result, change the `greetings.Hello` function so that it no longer includes the name. 
In greetings/greetings.go, paste the following code in place of the `Hello` function. Note that the highlighted lines change the value that the function returns, as if the `name` argument had been accidentally removed. 
```
// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
  // If no name was given, return an error with a message.
  if name == "" {
    return name, errors.New("empty name")
  }
  // Create a message using a random format.
  // message := fmt.Sprintf(randomFormat(), name)
  message := fmt.Sprint(randomFormat())
  return message, nil
}

```

  5. At the command line in the greetings directory, run `go test` to execute the test. 
This time, run `go test` without the `-v` flag. The output will include results for only the tests that failed, which can be useful when you have a lot of tests. The `TestHelloName` test should fail -- `TestHelloEmpty` still passes. 
```
$ go test
--- FAIL: TestHelloName (0.00s)
  greetings_test.go:15: Hello("Gladys") = "Hail, %v! Well met!", <nil>, want match for `\bGladys\b`, nil
FAIL
exit status 1
FAIL  example.com/greetings  0.182s

```



In the next (and last) topic, you'll see how to compile and install your code to run it locally. 
[< Return greetings for multiple people](https://go.dev/doc/tutorial/greetings-multiple-people.html) [Compile and install the application >](https://go.dev/doc/tutorial/compile-install.html)
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
- URL: https://golang.org/doc/tutorial/add-a-test
- Harvested: 2025-08-19T23:09:11.827090
- Profile: deep_research
- Agent: architect
