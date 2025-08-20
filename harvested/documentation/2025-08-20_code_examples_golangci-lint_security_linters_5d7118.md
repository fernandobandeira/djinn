---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:57:42.237066'
profile: code_examples
source: https://golangci-lint.run/usage/linters/
topic: GolangCI-Lint Security Linters
---

# GolangCI-Lint Security Linters

[![Logo](https://golangci-lint.run/images/logo-circle.svg) ![Dark Logo](https://golangci-lint.run/images/logo-nocircle.svg) Golangci-lint ](https://golangci-lint.run/)[Documentation ](https://golangci-lint.run/docs "Documentation")[Support us ](https://golangci-lint.run/docs/donate/ "Support us")[Linters ](https://golangci-lint.run/docs/linters/ "Linters")[Formatters](https://golangci-lint.run/docs/formatters/ "Formatters")
`CTRL K`
[ GitHub ](https://github.com/golangci/golangci-lint "GitHub")
`CTRL K`
  * [Golangci-lint Documentation ](https://golangci-lint.run/docs/)
    * [Welcome ](https://golangci-lint.run/docs/welcome/)
      * [Install](https://golangci-lint.run/docs/welcome/install/)
      * [Quick Start](https://golangci-lint.run/docs/welcome/quick-start/)
      * [Integrations](https://golangci-lint.run/docs/welcome/integrations/)
      * [FAQ](https://golangci-lint.run/docs/welcome/faq/)
    * [Configuration ](https://golangci-lint.run/docs/configuration/)
      * [Command-Line](https://golangci-lint.run/docs/configuration/cli/)
      * [Configuration File](https://golangci-lint.run/docs/configuration/file/)
    * [Linters ](https://golangci-lint.run/docs/linters/)
      * [All Linters](https://golangci-lint.run/usage/linters/#all-linters)
      * [Settings](https://golangci-lint.run/docs/linters/configuration/)
      * [False Positives](https://golangci-lint.run/docs/linters/false-positives/)
    * [Formatters ](https://golangci-lint.run/docs/formatters/)
      * [Settings](https://golangci-lint.run/docs/formatters/configuration/)
    * [Product ](https://golangci-lint.run/docs/product/)
      * [Thanks](https://golangci-lint.run/docs/product/thanks/)
      * [Changelog](https://golangci-lint.run/docs/product/changelog/)
      * [Migration guide](https://golangci-lint.run/docs/product/migration-guide/)
      * [Roadmap](https://golangci-lint.run/docs/product/roadmap/)
    * [Contributing ](https://golangci-lint.run/docs/contributing/)
      * [Workflow](https://golangci-lint.run/docs/contributing/workflow/)
      * [Architecture](https://golangci-lint.run/docs/contributing/architecture/)
      * [New linters](https://golangci-lint.run/docs/contributing/new-linters/)
      * [Debugging](https://golangci-lint.run/docs/contributing/debug/)
      * [FAQ](https://golangci-lint.run/docs/contributing/faq/)
      * [This Website](https://golangci-lint.run/docs/contributing/website/)
    * [Plugins ](https://golangci-lint.run/docs/plugins/)
      * [Module Plugin System](https://golangci-lint.run/docs/plugins/module-plugins/)
      * [Go Plugin System](https://golangci-lint.run/docs/plugins/go-plugins/)
    * [Support Us](https://golangci-lint.run/docs/donate/)
  * More
  * [Documentation v1 ↗](https://golangci.github.io/legacy-v1-doc/)


  * [Welcome ](https://golangci-lint.run/docs/welcome/)
    * [Install](https://golangci-lint.run/docs/welcome/install/)
    * [Quick Start](https://golangci-lint.run/docs/welcome/quick-start/)
    * [Integrations](https://golangci-lint.run/docs/welcome/integrations/)
    * [FAQ](https://golangci-lint.run/docs/welcome/faq/)
  * [Configuration ](https://golangci-lint.run/docs/configuration/)
    * [Command-Line](https://golangci-lint.run/docs/configuration/cli/)
    * [Configuration File](https://golangci-lint.run/docs/configuration/file/)
  * [Linters ](https://golangci-lint.run/docs/linters/)
    * [Settings](https://golangci-lint.run/docs/linters/configuration/)
    * [False Positives](https://golangci-lint.run/docs/linters/false-positives/)
  * [Formatters ](https://golangci-lint.run/docs/formatters/)
    * [Settings](https://golangci-lint.run/docs/formatters/configuration/)
  * [Product ](https://golangci-lint.run/docs/product/)
    * [Thanks](https://golangci-lint.run/docs/product/thanks/)
    * [Changelog](https://golangci-lint.run/docs/product/changelog/)
    * [Migration guide](https://golangci-lint.run/docs/product/migration-guide/)
    * [Roadmap](https://golangci-lint.run/docs/product/roadmap/)
  * [Contributing ](https://golangci-lint.run/docs/contributing/)
    * [Workflow](https://golangci-lint.run/docs/contributing/workflow/)
    * [Architecture](https://golangci-lint.run/docs/contributing/architecture/)
    * [New linters](https://golangci-lint.run/docs/contributing/new-linters/)
    * [Debugging](https://golangci-lint.run/docs/contributing/debug/)
    * [FAQ](https://golangci-lint.run/docs/contributing/faq/)
    * [This Website](https://golangci-lint.run/docs/contributing/website/)
  * [Plugins ](https://golangci-lint.run/docs/plugins/)
    * [Module Plugin System](https://golangci-lint.run/docs/plugins/module-plugins/)
    * [Go Plugin System](https://golangci-lint.run/docs/plugins/go-plugins/)
  * [Support Us](https://golangci-lint.run/docs/donate/)
  * More
  * [Documentation v1 ↗](https://golangci.github.io/legacy-v1-doc/)


LightDark
On this page
  * [All Linters](https://golangci-lint.run/usage/linters/#all-linters)


[Edit this page on GitHub →](https://github.com/golangci/golangci-lint/edit/main/docs/content/docs/linters/_index.md) Scroll to top
[Golangci-lint Documentation](https://golangci-lint.run/docs/)
Linters
# Linters
To see a list of supported linters and which linters are enabled/disabled:
```
golangci-lint help linters
```

To see a list of linters enabled by your configuration, use:
```
golangci-lint linters
```

[Quick Start](https://golangci-lint.run/docs/welcome/quick-start/#linting) [CLI](https://golangci-lint.run/docs/configuration/cli/#run) [Global Configuration](https://golangci-lint.run/docs/configuration/file/#linters-configuration) [Linter Settings](https://golangci-lint.run/docs/linters/configuration/)
## All Linters [](https://golangci-lint.run/usage/linters/#all-linters)
Default
New
Autofix
Fast
Slow
Deprecated
Reset
[arangolintOpinionated best practices for arangodb client.](https://golangci-lint.run/docs/linters/configuration/#arangolint)
[asasalintCheck for pass []any as any in variadic func(…any).](https://golangci-lint.run/docs/linters/configuration/#asasalint)
[asciicheckChecks that all code identifiers does not have non-ASCII symbols in the name.](https://golangci-lint.run/docs/linters/configuration/#asciicheck)
[bidichkChecks for dangerous unicode character sequences.](https://golangci-lint.run/docs/linters/configuration/#bidichk)
[bodycloseChecks whether HTTP response body is closed successfully.](https://golangci-lint.run/docs/linters/configuration/#bodyclose)
[canonicalheaderCanonicalheader checks whether net/http.Header uses canonical header.Autofix](https://golangci-lint.run/docs/linters/configuration/#canonicalheader)
[containedctxContainedctx is a linter that detects struct contained context.Context field.](https://golangci-lint.run/docs/linters/configuration/#containedctx)
[contextcheckCheck whether the function uses a non-inherited context.](https://golangci-lint.run/docs/linters/configuration/#contextcheck)
[copyloopvarA linter detects places where loop variables are copied.Autofix](https://golangci-lint.run/docs/linters/configuration/#copyloopvar)
[cyclopChecks function and package cyclomatic complexity.](https://golangci-lint.run/docs/linters/configuration/#cyclop)
[decorderCheck declaration order and count of types, constants, variables and functions.](https://golangci-lint.run/docs/linters/configuration/#decorder)
[depguardGo linter that checks if package imports are in a list of acceptable packages.](https://golangci-lint.run/docs/linters/configuration/#depguard)
[dogsledChecks assignments with too many blank identifiers (e.g. x, _, _, _, := f()).](https://golangci-lint.run/docs/linters/configuration/#dogsled)
[duplDetects duplicate fragments of code.](https://golangci-lint.run/docs/linters/configuration/#dupl)
[dupwordChecks for duplicate words in the source code.Autofix](https://golangci-lint.run/docs/linters/configuration/#dupword)
[durationcheckCheck for two durations multiplied together.](https://golangci-lint.run/docs/linters/configuration/#durationcheck)
[embeddedstructfieldcheckEmbedded types should be at the top of the field list of a struct, and there must be an empty line separating embedded fields from regular fields.](https://golangci-lint.run/docs/linters/configuration/#embeddedstructfieldcheck)
[err113Check errors handling expressions.Autofix](https://golangci-lint.run/docs/linters/configuration/#err113)
[errcheckErrcheck is a program for checking for unchecked errors in Go code. These unchecked errors can be critical bugs in some cases.](https://golangci-lint.run/docs/linters/configuration/#errcheck)
[errchkjsonChecks types passed to the json encoding functions. Reports unsupported types and reports occurrences where the check for the returned error can be omitted.](https://golangci-lint.run/docs/linters/configuration/#errchkjson)
[errnameChecks that sentinel errors are prefixed with the `Err` and error types are suffixed with the `Error`.](https://golangci-lint.run/docs/linters/configuration/#errname)
[errorlintFind code that can cause problems with the error wrapping scheme introduced in Go 1.13.Autofix](https://golangci-lint.run/docs/linters/configuration/#errorlint)
[exhaustiveCheck exhaustiveness of enum switch statements.](https://golangci-lint.run/docs/linters/configuration/#exhaustive)
[exhaustructChecks if all structure fields are initialized.](https://golangci-lint.run/docs/linters/configuration/#exhaustruct)
[exptostdDetects functions from golang.org/x/exp/ that can be replaced by std functions.Autofix](https://golangci-lint.run/docs/linters/configuration/#exptostd)
[fatcontextDetects nested contexts in loops and function literals.Autofix](https://golangci-lint.run/docs/linters/configuration/#fatcontext)
[forbidigoForbids identifiers.](https://golangci-lint.run/docs/linters/configuration/#forbidigo)
[forcetypeassertFind forced type assertions.](https://golangci-lint.run/docs/linters/configuration/#forcetypeassert)
[funcorderChecks the order of functions, methods, and constructors.](https://golangci-lint.run/docs/linters/configuration/#funcorder)
[funlenChecks for long functions.](https://golangci-lint.run/docs/linters/configuration/#funlen)
[ginkgolinterEnforces standards of using ginkgo and gomega.Autofix](https://golangci-lint.run/docs/linters/configuration/#ginkgolinter)
[gocheckcompilerdirectivesChecks that go compiler directive comments (//go:) are valid.](https://golangci-lint.run/docs/linters/configuration/#gocheckcompilerdirectives)
[gochecknoglobalsCheck that no global variables exist.](https://golangci-lint.run/docs/linters/configuration/#gochecknoglobals)
[gochecknoinitsChecks that no init functions are present in Go code.](https://golangci-lint.run/docs/linters/configuration/#gochecknoinits)
[gochecksumtypeRun exhaustiveness checks on Go “sum types”.](https://golangci-lint.run/docs/linters/configuration/#gochecksumtype)
[gocognitComputes and checks the cognitive complexity of functions.](https://golangci-lint.run/docs/linters/configuration/#gocognit)
[goconstFinds repeated strings that could be replaced by a constant.](https://golangci-lint.run/docs/linters/configuration/#goconst)
[gocriticProvides diagnostics that check for bugs, performance and style issues. Extensible without recompilation through dynamic rules. Dynamic rules are written declaratively with AST patterns, filters, report message and optional suggestion.Autofix](https://golangci-lint.run/docs/linters/configuration/#gocritic)
[gocycloComputes and checks the cyclomatic complexity of functions.](https://golangci-lint.run/docs/linters/configuration/#gocyclo)
[godotCheck if comments end in a period.Autofix](https://golangci-lint.run/docs/linters/configuration/#godot)
[godoxDetects usage of FIXME, TODO and other keywords inside comments.](https://golangci-lint.run/docs/linters/configuration/#godox)
[goheaderCheck if file header matches to pattern.Autofix](https://golangci-lint.run/docs/linters/configuration/#goheader)
[gomoddirectivesManage the use of ‘replace’, ‘retract’, and ’excludes’ directives in go.mod.](https://golangci-lint.run/docs/linters/configuration/#gomoddirectives)
[gomodguardAllow and blocklist linter for direct Go module dependencies. This is different from depguard where there are different block types for example version constraints and module recommendations.](https://golangci-lint.run/docs/linters/configuration/#gomodguard)
[goprintffuncnameChecks that printf-like functions are named with `f` at the end.](https://golangci-lint.run/docs/linters/configuration/#goprintffuncname)
[gosecInspects source code for security problems.](https://golangci-lint.run/docs/linters/configuration/#gosec)
[gosmopolitanReport certain i18n/l10n anti-patterns in your Go codebase.](https://golangci-lint.run/docs/linters/configuration/#gosmopolitan)
[govetVet examines Go source code and reports suspicious constructs. It is roughly the same as ‘go vet’ and uses its passes.Autofix](https://golangci-lint.run/docs/linters/configuration/#govet)
[grouperAnalyze expression groups.](https://golangci-lint.run/docs/linters/configuration/#grouper)
[ifaceDetect the incorrect use of interfaces, helping developers avoid interface pollution.Autofix](https://golangci-lint.run/docs/linters/configuration/#iface)
[importasEnforces consistent import aliases.Autofix](https://golangci-lint.run/docs/linters/configuration/#importas)
[inamedparamReports interfaces with unnamed method parameters.](https://golangci-lint.run/docs/linters/configuration/#inamedparam)
[ineffassignDetects when assignments to existing variables are not used.](https://golangci-lint.run/docs/linters/configuration/#ineffassign)
[interfacebloatA linter that checks the number of methods inside an interface.](https://golangci-lint.run/docs/linters/configuration/#interfacebloat)
[intrangeIntrange is a linter to find places where for loops could make use of an integer range.Autofix](https://golangci-lint.run/docs/linters/configuration/#intrange)
[ireturnAccept Interfaces, Return Concrete Types.](https://golangci-lint.run/docs/linters/configuration/#ireturn)
[lllReports long lines.](https://golangci-lint.run/docs/linters/configuration/#lll)
[loggercheckChecks key value pairs for common logger libraries (kitlog,klog,logr,slog,zap).](https://golangci-lint.run/docs/linters/configuration/#loggercheck)
[maintidxMaintidx measures the maintainability index of each function.](https://golangci-lint.run/docs/linters/configuration/#maintidx)
[makezeroFind slice declarations with non-zero initial length.](https://golangci-lint.run/docs/linters/configuration/#makezero)
[mirrorReports wrong mirror patterns of bytes/strings usage.Autofix](https://golangci-lint.run/docs/linters/configuration/#mirror)
[misspellFinds commonly misspelled English words.Autofix](https://golangci-lint.run/docs/linters/configuration/#misspell)
[mndAn analyzer to detect magic numbers.](https://golangci-lint.run/docs/linters/configuration/#mnd)
[musttagEnforce field tags in (un)marshaled structs.](https://golangci-lint.run/docs/linters/configuration/#musttag)
[nakedretChecks that functions with naked returns are not longer than a maximum size (can be zero).Autofix](https://golangci-lint.run/docs/linters/configuration/#nakedret)
[nestifReports deeply nested if statements.](https://golangci-lint.run/docs/linters/configuration/#nestif)
[nilerrFind the code that returns nil even if it checks that the error is not nil.](https://golangci-lint.run/docs/linters/configuration/#nilerr)
[nilnesserrReports constructs that checks for err != nil, but returns a different nil value error. Powered by nilness and nilerr.](https://golangci-lint.run/docs/linters/configuration/#nilnesserr)
[nilnilChecks that there is no simultaneous return of `nil` error and an invalid value.](https://golangci-lint.run/docs/linters/configuration/#nilnil)
[nlreturnChecks for a new line before return and branch statements to increase code clarity.Autofix](https://golangci-lint.run/docs/linters/configuration/#nlreturn)
[noctxDetects function and method with missing usage of context.Context.](https://golangci-lint.run/docs/linters/configuration/#noctx)
[noinlineerrDisallows inline error handling (`if err := ...; err != nil {`).](https://golangci-lint.run/docs/linters/configuration/#noinlineerr)
[nolintlintReports ill-formed or insufficient nolint directives.Autofix](https://golangci-lint.run/docs/linters/configuration/#nolintlint)
[nonamedreturnsReports all named returns.](https://golangci-lint.run/docs/linters/configuration/#nonamedreturns)
[nosprintfhostportChecks for misuse of Sprintf to construct a host with port in a URL.](https://golangci-lint.run/docs/linters/configuration/#nosprintfhostport)
[paralleltestDetects missing usage of t.Parallel() method in your Go test.](https://golangci-lint.run/docs/linters/configuration/#paralleltest)
[perfsprintChecks that fmt.Sprintf can be replaced with a faster alternative.Autofix](https://golangci-lint.run/docs/linters/configuration/#perfsprint)
[preallocFind slice declarations that could potentially be pre-allocated.](https://golangci-lint.run/docs/linters/configuration/#prealloc)
[predeclaredFind code that shadows one of Go’s predeclared identifiers.](https://golangci-lint.run/docs/linters/configuration/#predeclared)
[promlinterCheck Prometheus metrics naming via promlint.](https://golangci-lint.run/docs/linters/configuration/#promlinter)
[protogetterReports direct reads from proto message fields when getters should be used.Autofix](https://golangci-lint.run/docs/linters/configuration/#protogetter)
[reassignChecks that package variables are not reassigned.](https://golangci-lint.run/docs/linters/configuration/#reassign)
[recvcheckChecks for receiver type consistency.](https://golangci-lint.run/docs/linters/configuration/#recvcheck)
[reviveFast, configurable, extensible, flexible, and beautiful linter for Go. Drop-in replacement of golint.Autofix](https://golangci-lint.run/docs/linters/configuration/#revive)
[rowserrcheckChecks whether Rows.Err of rows is checked successfully.](https://golangci-lint.run/docs/linters/configuration/#rowserrcheck)
[sloglintEnsure consistent code style when using log/slog.Autofix](https://golangci-lint.run/docs/linters/configuration/#sloglint)
[spancheckChecks for mistakes with OpenTelemetry/Census spans.](https://golangci-lint.run/docs/linters/configuration/#spancheck)
[sqlclosecheckChecks that sql.Rows, sql.Stmt, sqlx.NamedStmt, pgx.Query are closed.](https://golangci-lint.run/docs/linters/configuration/#sqlclosecheck)
[staticcheckIt’s the set of rules from staticcheck.Autofix](https://golangci-lint.run/docs/linters/configuration/#staticcheck)
[tagalignCheck that struct tags are well aligned.Autofix](https://golangci-lint.run/docs/linters/configuration/#tagalign)
[tagliatelleChecks the struct tags.](https://golangci-lint.run/docs/linters/configuration/#tagliatelle)
[testableexamplesLinter checks if examples are testable (have an expected output).](https://golangci-lint.run/docs/linters/configuration/#testableexamples)
[testifylintChecks usage of github.com/stretchr/testify.Autofix](https://golangci-lint.run/docs/linters/configuration/#testifylint)
[testpackageLinter that makes you use a separate _test package.](https://golangci-lint.run/docs/linters/configuration/#testpackage)
[thelperThelper detects tests helpers which is not start with t.Helper() method.](https://golangci-lint.run/docs/linters/configuration/#thelper)
[tparallelTparallel detects inappropriate usage of t.Parallel() method in your Go test codes.](https://golangci-lint.run/docs/linters/configuration/#tparallel)
[unconvertRemove unnecessary type conversions.](https://golangci-lint.run/docs/linters/configuration/#unconvert)
[unparamReports unused function parameters.](https://golangci-lint.run/docs/linters/configuration/#unparam)
[unusedChecks Go code for unused constants, variables, functions and types.](https://golangci-lint.run/docs/linters/configuration/#unused)
[usestdlibvarsA linter that detect the possibility to use variables/constants from the Go standard library.Autofix](https://golangci-lint.run/docs/linters/configuration/#usestdlibvars)
[usetestingReports uses of functions with replacement inside the testing package.Autofix](https://golangci-lint.run/docs/linters/configuration/#usetesting)
[varnamelenChecks that the length of a variable’s name matches its scope.](https://golangci-lint.run/docs/linters/configuration/#varnamelen)
[wastedassignFinds wasted assignment statements.](https://golangci-lint.run/docs/linters/configuration/#wastedassign)
[whitespaceWhitespace is a linter that checks for unnecessary newlines at the start and end of functions, if, for, etc.Autofix](https://golangci-lint.run/docs/linters/configuration/#whitespace)
[wrapcheckChecks that errors returned from external packages are wrapped.](https://golangci-lint.run/docs/linters/configuration/#wrapcheck)
[wslNew major version. Use `wsl_v5` instead.Deprecated since v2.2.0](https://golangci-lint.run/docs/linters/configuration/#wsl)
[wsl_v5Add or remove empty lines.Autofix](https://golangci-lint.run/docs/linters/configuration/#wsl_v5)
[zerologlintDetects the wrong usage of `zerolog` that a user forgets to dispatch with `Send` or `Msg`.](https://golangci-lint.run/docs/linters/configuration/#zerologlint)
Last updated on 2025-08-18 06:41:21
[Powered by Hextra](https://github.com/imfing/hextra "Hextra GitHub Homepage")
[ ](https://bsky.app/profile/golangci-lint.run "Bluesky")[ ](https://fosstodon.org/@golangcilint "Mastodon")[ ](https://twitter.com/golangci "Twitter")[ ](https://gophers.slack.com/archives/CS0TBRKPC "Slack")[](https://github.com/golangci/golangci-lint "GitHub")


## Source Information
- URL: https://golangci-lint.run/usage/linters/
- Harvested: 2025-08-20T00:57:42.237066
- Profile: code_examples
- Agent: architect
