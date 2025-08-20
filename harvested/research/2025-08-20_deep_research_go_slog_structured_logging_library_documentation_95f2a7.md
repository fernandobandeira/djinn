---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-20T00:35:38.846164'
profile: deep_research
source: https://pkg.go.dev/log/slog
topic: Go slog structured logging library documentation
---

# Go slog structured logging library documentation

[ ![Go](https://pkg.go.dev/static/shared/logo/go-white.svg) ](https://go.dev/)
[ Skip to Main Content ](https://pkg.go.dev/log/slog#main-content)
![](https://pkg.go.dev/static/shared/icon/search_gm_grey_24dp.svg) ![](https://pkg.go.dev/static/shared/icon/search_gm_grey_24dp.svg)
  * [ Why Go ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/log/slog)
    * [ Case Studies ](https://go.dev/solutions#case-studies)
Common problems companies solve with Go
    * [ Use Cases ](https://go.dev/solutions#use-cases)
Stories about how and why companies use Go
    * [ Security ](https://go.dev/security/)
How Go can help keep you secure by default
  * [Learn](https://go.dev/learn/)
  * [ Docs ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/log/slog)
    * [ Effective Go ](https://go.dev/doc/effective_go)
Tips for writing clear, performant, and idiomatic Go code
    * [ Go User Manual ](https://go.dev/doc/)
A complete introduction to building software with Go
    * [ Standard library ](https://pkg.go.dev/std)
Reference documentation for Go's standard library
    * [ Release Notes ](https://go.dev/doc/devel/release)
Learn what's new in each Go release
  * [Packages](https://pkg.go.dev/)
  * [ Community ![submenu dropdown icon](https://pkg.go.dev/static/shared/icon/arrow_drop_down_gm_grey_24dp.svg) ](https://pkg.go.dev/log/slog)
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
  * [ Why Go _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/log/slog)
[ _![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Why Go ](https://pkg.go.dev/log/slog)
    * [ Case Studies ](https://go.dev/solutions#case-studies)
    * [ Use Cases ](https://go.dev/solutions#use-cases)
    * [ Security ](https://go.dev/security/)
  * [Learn](https://go.dev/learn/)
  * [ Docs _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/log/slog)
[_![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Docs ](https://pkg.go.dev/log/slog)
    * [ Effective Go ](https://go.dev/doc/effective_go)
    * [ Go User Manual ](https://go.dev/doc/)
    * [ Standard library ](https://pkg.go.dev/std)
    * [ Release Notes ](https://go.dev/doc/devel/release)
  * [Packages](https://pkg.go.dev/)
  * [ Community _![](https://pkg.go.dev/static/shared/icon/navigate_next_gm_grey_24dp.svg) _ ](https://pkg.go.dev/log/slog)
[ _![](https://pkg.go.dev/static/shared/icon/navigate_before_gm_grey_24dp.svg) _ Community ](https://pkg.go.dev/log/slog)
    * [ Recorded Talks ](https://go.dev/talks/)
    * [ Meetups _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://www.meetup.com/pro/go)
    * [ Conferences _![](https://pkg.go.dev/static/shared/icon/launch_gm_grey_24dp.svg) _ ](https://github.com/golang/go/wiki/Conferences)
    * [ Go blog ](https://go.dev/blog)
    * [ Go project ](https://go.dev/help)
    * Get connected
[![](https://pkg.go.dev/static/shared/logo/social/google-groups.svg)](https://groups.google.com/g/golang-nuts) [![](https://pkg.go.dev/static/shared/logo/social/github.svg)](https://github.com/golang) [![](https://pkg.go.dev/static/shared/logo/social/twitter.svg)](https://twitter.com/golang) [![](https://pkg.go.dev/static/shared/logo/social/reddit.svg)](https://www.reddit.com/r/golang/) [![](https://pkg.go.dev/static/shared/logo/social/slack.svg)](https://invite.slack.golangbridge.org/) [![](https://pkg.go.dev/static/shared/logo/social/stack-overflow.svg)](https://stackoverflow.com/collectives/go)


  1. [Discover Packages](https://pkg.go.dev/)
  2. [Standard library](https://pkg.go.dev/std)
  3. [log](https://pkg.go.dev/log)
  4. [ slog ](https://pkg.go.dev/log/slog@go1.25.0) ![](https://pkg.go.dev/static/shared/icon/content_copy_gm_grey_24dp.svg)


[ ![Go](https://pkg.go.dev/static/shared/logo/go-blue.svg) ](https://go.dev/)
# slog
package standard library ![](https://pkg.go.dev/static/shared/icon/content_copy_gm_grey_24dp.svg)
[ Version:  go1.25.0 ](https://pkg.go.dev/log/slog?tab=versions)
Opens a new window with list of versions in this module. 
Latest Latest  ![Warning](https://pkg.go.dev/static/shared/icon/alert_gm_grey_24dp.svg)
This package is not in the latest version of its module.
[ Go to latest ](https://pkg.go.dev/log/slog) Published: Aug 12, 2025  License: [BSD-3-Clause](https://pkg.go.dev/log/slog?tab=licenses)
Opens a new window with license information. 
[ Imports: 23 ](https://pkg.go.dev/log/slog?tab=imports)
Opens a new window with list of imports. 
[ Imported by: 53,844 ](https://pkg.go.dev/log/slog?tab=importedby)
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
  * [ Documentation ](https://pkg.go.dev/log/slog#section-documentation)
    * [Overview](https://pkg.go.dev/log/slog#pkg-overview)
    * [ Index ](https://pkg.go.dev/log/slog#pkg-index)
      * [ Examples ](https://pkg.go.dev/log/slog#pkg-examples)
    * [ Constants ](https://pkg.go.dev/log/slog#pkg-constants)
    * [ Variables ](https://pkg.go.dev/log/slog#pkg-variables)
    * [ Functions ](https://pkg.go.dev/log/slog#pkg-functions)
      * [ Debug(msg, args) ](https://pkg.go.dev/log/slog#Debug "Debug\(msg, args\)")
      * [ DebugContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#DebugContext "DebugContext\(ctx, msg, args\)")
      * [ Error(msg, args) ](https://pkg.go.dev/log/slog#Error "Error\(msg, args\)")
      * [ ErrorContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#ErrorContext "ErrorContext\(ctx, msg, args\)")
      * [ Info(msg, args) ](https://pkg.go.dev/log/slog#Info "Info\(msg, args\)")
      * [ InfoContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#InfoContext "InfoContext\(ctx, msg, args\)")
      * [ Log(ctx, level, msg, args) ](https://pkg.go.dev/log/slog#Log "Log\(ctx, level, msg, args\)")
      * [ LogAttrs(ctx, level, msg, attrs) ](https://pkg.go.dev/log/slog#LogAttrs "LogAttrs\(ctx, level, msg, attrs\)")
      * [ NewLogLogger(h, level) ](https://pkg.go.dev/log/slog#NewLogLogger "NewLogLogger\(h, level\)")
      * [ SetDefault(l) ](https://pkg.go.dev/log/slog#SetDefault "SetDefault\(l\)")
      * [ Warn(msg, args) ](https://pkg.go.dev/log/slog#Warn "Warn\(msg, args\)")
      * [ WarnContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#WarnContext "WarnContext\(ctx, msg, args\)")
    * [ Types ](https://pkg.go.dev/log/slog#pkg-types)
      * [ type Attr ](https://pkg.go.dev/log/slog#Attr "type Attr")
        * [ Any(key, value) ](https://pkg.go.dev/log/slog#Any "Any\(key, value\)")
        * [ Bool(key, v) ](https://pkg.go.dev/log/slog#Bool "Bool\(key, v\)")
        * [ Duration(key, v) ](https://pkg.go.dev/log/slog#Duration "Duration\(key, v\)")
        * [ Float64(key, v) ](https://pkg.go.dev/log/slog#Float64 "Float64\(key, v\)")
        * [ Group(key, args) ](https://pkg.go.dev/log/slog#Group "Group\(key, args\)")
        * [ GroupAttrs(key, attrs) ](https://pkg.go.dev/log/slog#GroupAttrs "GroupAttrs\(key, attrs\)")
        * [ Int(key, value) ](https://pkg.go.dev/log/slog#Int "Int\(key, value\)")
        * [ Int64(key, value) ](https://pkg.go.dev/log/slog#Int64 "Int64\(key, value\)")
        * [ String(key, value) ](https://pkg.go.dev/log/slog#String "String\(key, value\)")
        * [ Time(key, v) ](https://pkg.go.dev/log/slog#Time "Time\(key, v\)")
        * [ Uint64(key, v) ](https://pkg.go.dev/log/slog#Uint64 "Uint64\(key, v\)")
        * [ (a) Equal(b) ](https://pkg.go.dev/log/slog#Attr.Equal "\(a\) Equal\(b\)")
        * [ (a) String() ](https://pkg.go.dev/log/slog#Attr.String "\(a\) String\(\)")
      * [ type Handler ](https://pkg.go.dev/log/slog#Handler "type Handler")
      * [ type HandlerOptions ](https://pkg.go.dev/log/slog#HandlerOptions "type HandlerOptions")
      * [ type JSONHandler ](https://pkg.go.dev/log/slog#JSONHandler "type JSONHandler")
        * [ NewJSONHandler(w, opts) ](https://pkg.go.dev/log/slog#NewJSONHandler "NewJSONHandler\(w, opts\)")
        * [ (h) Enabled(_, level) ](https://pkg.go.dev/log/slog#JSONHandler.Enabled "\(h\) Enabled\(_, level\)")
        * [ (h) Handle(_, r) ](https://pkg.go.dev/log/slog#JSONHandler.Handle "\(h\) Handle\(_, r\)")
        * [ (h) WithAttrs(attrs) ](https://pkg.go.dev/log/slog#JSONHandler.WithAttrs "\(h\) WithAttrs\(attrs\)")
        * [ (h) WithGroup(name) ](https://pkg.go.dev/log/slog#JSONHandler.WithGroup "\(h\) WithGroup\(name\)")
      * [ type Kind ](https://pkg.go.dev/log/slog#Kind "type Kind")
        * [ (k) String() ](https://pkg.go.dev/log/slog#Kind.String "\(k\) String\(\)")
      * [ type Level ](https://pkg.go.dev/log/slog#Level "type Level")
        * [ SetLogLoggerLevel(level) ](https://pkg.go.dev/log/slog#SetLogLoggerLevel "SetLogLoggerLevel\(level\)")
        * [ (l) AppendText(b) ](https://pkg.go.dev/log/slog#Level.AppendText "\(l\) AppendText\(b\)")
        * [ (l) Level() ](https://pkg.go.dev/log/slog#Level.Level "\(l\) Level\(\)")
        * [ (l) MarshalJSON() ](https://pkg.go.dev/log/slog#Level.MarshalJSON "\(l\) MarshalJSON\(\)")
        * [ (l) MarshalText() ](https://pkg.go.dev/log/slog#Level.MarshalText "\(l\) MarshalText\(\)")
        * [ (l) String() ](https://pkg.go.dev/log/slog#Level.String "\(l\) String\(\)")
        * [ (l) UnmarshalJSON(data) ](https://pkg.go.dev/log/slog#Level.UnmarshalJSON "\(l\) UnmarshalJSON\(data\)")
        * [ (l) UnmarshalText(data) ](https://pkg.go.dev/log/slog#Level.UnmarshalText "\(l\) UnmarshalText\(data\)")
      * [ type LevelVar ](https://pkg.go.dev/log/slog#LevelVar "type LevelVar")
        * [ (v) AppendText(b) ](https://pkg.go.dev/log/slog#LevelVar.AppendText "\(v\) AppendText\(b\)")
        * [ (v) Level() ](https://pkg.go.dev/log/slog#LevelVar.Level "\(v\) Level\(\)")
        * [ (v) MarshalText() ](https://pkg.go.dev/log/slog#LevelVar.MarshalText "\(v\) MarshalText\(\)")
        * [ (v) Set(l) ](https://pkg.go.dev/log/slog#LevelVar.Set "\(v\) Set\(l\)")
        * [ (v) String() ](https://pkg.go.dev/log/slog#LevelVar.String "\(v\) String\(\)")
        * [ (v) UnmarshalText(data) ](https://pkg.go.dev/log/slog#LevelVar.UnmarshalText "\(v\) UnmarshalText\(data\)")
      * [ type Leveler ](https://pkg.go.dev/log/slog#Leveler "type Leveler")
      * [ type LogValuer ](https://pkg.go.dev/log/slog#LogValuer "type LogValuer")
      * [ type Logger ](https://pkg.go.dev/log/slog#Logger "type Logger")
        * [ Default() ](https://pkg.go.dev/log/slog#Default "Default\(\)")
        * [ New(h) ](https://pkg.go.dev/log/slog#New "New\(h\)")
        * [ With(args) ](https://pkg.go.dev/log/slog#With "With\(args\)")
        * [ (l) Debug(msg, args) ](https://pkg.go.dev/log/slog#Logger.Debug "\(l\) Debug\(msg, args\)")
        * [ (l) DebugContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#Logger.DebugContext "\(l\) DebugContext\(ctx, msg, args\)")
        * [ (l) Enabled(ctx, level) ](https://pkg.go.dev/log/slog#Logger.Enabled "\(l\) Enabled\(ctx, level\)")
        * [ (l) Error(msg, args) ](https://pkg.go.dev/log/slog#Logger.Error "\(l\) Error\(msg, args\)")
        * [ (l) ErrorContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#Logger.ErrorContext "\(l\) ErrorContext\(ctx, msg, args\)")
        * [ (l) Handler() ](https://pkg.go.dev/log/slog#Logger.Handler "\(l\) Handler\(\)")
        * [ (l) Info(msg, args) ](https://pkg.go.dev/log/slog#Logger.Info "\(l\) Info\(msg, args\)")
        * [ (l) InfoContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#Logger.InfoContext "\(l\) InfoContext\(ctx, msg, args\)")
        * [ (l) Log(ctx, level, msg, args) ](https://pkg.go.dev/log/slog#Logger.Log "\(l\) Log\(ctx, level, msg, args\)")
        * [ (l) LogAttrs(ctx, level, msg, attrs) ](https://pkg.go.dev/log/slog#Logger.LogAttrs "\(l\) LogAttrs\(ctx, level, msg, attrs\)")
        * [ (l) Warn(msg, args) ](https://pkg.go.dev/log/slog#Logger.Warn "\(l\) Warn\(msg, args\)")
        * [ (l) WarnContext(ctx, msg, args) ](https://pkg.go.dev/log/slog#Logger.WarnContext "\(l\) WarnContext\(ctx, msg, args\)")
        * [ (l) With(args) ](https://pkg.go.dev/log/slog#Logger.With "\(l\) With\(args\)")
        * [ (l) WithGroup(name) ](https://pkg.go.dev/log/slog#Logger.WithGroup "\(l\) WithGroup\(name\)")
      * [ type Record ](https://pkg.go.dev/log/slog#Record "type Record")
        * [ NewRecord(t, level, msg, pc) ](https://pkg.go.dev/log/slog#NewRecord "NewRecord\(t, level, msg, pc\)")
        * [ (r) Add(args) ](https://pkg.go.dev/log/slog#Record.Add "\(r\) Add\(args\)")
        * [ (r) AddAttrs(attrs) ](https://pkg.go.dev/log/slog#Record.AddAttrs "\(r\) AddAttrs\(attrs\)")
        * [ (r) Attrs(f) ](https://pkg.go.dev/log/slog#Record.Attrs "\(r\) Attrs\(f\)")
        * [ (r) Clone() ](https://pkg.go.dev/log/slog#Record.Clone "\(r\) Clone\(\)")
        * [ (r) NumAttrs() ](https://pkg.go.dev/log/slog#Record.NumAttrs "\(r\) NumAttrs\(\)")
        * [ (r) Source() ](https://pkg.go.dev/log/slog#Record.Source "\(r\) Source\(\)")
      * [ type Source ](https://pkg.go.dev/log/slog#Source "type Source")
      * [ type TextHandler ](https://pkg.go.dev/log/slog#TextHandler "type TextHandler")
        * [ NewTextHandler(w, opts) ](https://pkg.go.dev/log/slog#NewTextHandler "NewTextHandler\(w, opts\)")
        * [ (h) Enabled(_, level) ](https://pkg.go.dev/log/slog#TextHandler.Enabled "\(h\) Enabled\(_, level\)")
        * [ (h) Handle(_, r) ](https://pkg.go.dev/log/slog#TextHandler.Handle "\(h\) Handle\(_, r\)")
        * [ (h) WithAttrs(attrs) ](https://pkg.go.dev/log/slog#TextHandler.WithAttrs "\(h\) WithAttrs\(attrs\)")
        * [ (h) WithGroup(name) ](https://pkg.go.dev/log/slog#TextHandler.WithGroup "\(h\) WithGroup\(name\)")
      * [ type Value ](https://pkg.go.dev/log/slog#Value "type Value")
        * [ AnyValue(v) ](https://pkg.go.dev/log/slog#AnyValue "AnyValue\(v\)")
        * [ BoolValue(v) ](https://pkg.go.dev/log/slog#BoolValue "BoolValue\(v\)")
        * [ DurationValue(v) ](https://pkg.go.dev/log/slog#DurationValue "DurationValue\(v\)")
        * [ Float64Value(v) ](https://pkg.go.dev/log/slog#Float64Value "Float64Value\(v\)")
        * [ GroupValue(as) ](https://pkg.go.dev/log/slog#GroupValue "GroupValue\(as\)")
        * [ Int64Value(v) ](https://pkg.go.dev/log/slog#Int64Value "Int64Value\(v\)")
        * [ IntValue(v) ](https://pkg.go.dev/log/slog#IntValue "IntValue\(v\)")
        * [ StringValue(value) ](https://pkg.go.dev/log/slog#StringValue "StringValue\(value\)")
        * [ TimeValue(v) ](https://pkg.go.dev/log/slog#TimeValue "TimeValue\(v\)")
        * [ Uint64Value(v) ](https://pkg.go.dev/log/slog#Uint64Value "Uint64Value\(v\)")
        * [ (v) Any() ](https://pkg.go.dev/log/slog#Value.Any "\(v\) Any\(\)")
        * [ (v) Bool() ](https://pkg.go.dev/log/slog#Value.Bool "\(v\) Bool\(\)")
        * [ (v) Duration() ](https://pkg.go.dev/log/slog#Value.Duration "\(v\) Duration\(\)")
        * [ (v) Equal(w) ](https://pkg.go.dev/log/slog#Value.Equal "\(v\) Equal\(w\)")
        * [ (v) Float64() ](https://pkg.go.dev/log/slog#Value.Float64 "\(v\) Float64\(\)")
        * [ (v) Group() ](https://pkg.go.dev/log/slog#Value.Group "\(v\) Group\(\)")
        * [ (v) Int64() ](https://pkg.go.dev/log/slog#Value.Int64 "\(v\) Int64\(\)")
        * [ (v) Kind() ](https://pkg.go.dev/log/slog#Value.Kind "\(v\) Kind\(\)")
        * [ (v) LogValuer() ](https://pkg.go.dev/log/slog#Value.LogValuer "\(v\) LogValuer\(\)")
        * [ (v) Resolve() ](https://pkg.go.dev/log/slog#Value.Resolve "\(v\) Resolve\(\)")
        * [ (v) String() ](https://pkg.go.dev/log/slog#Value.String "\(v\) String\(\)")
        * [ (v) Time() ](https://pkg.go.dev/log/slog#Value.Time "\(v\) Time\(\)")
        * [ (v) Uint64() ](https://pkg.go.dev/log/slog#Value.Uint64 "\(v\) Uint64\(\)")
  * [ Source Files ](https://pkg.go.dev/log/slog#section-sourcefiles)
  * [ Directories ](https://pkg.go.dev/log/slog#section-directories)


DocumentationSource FilesDirectoriesOverviewIndexConstantsVariablesFunctionsTypesExamplesDebug(msg, args)DebugContext(ctx, msg, args)Error(msg, args)ErrorContext(ctx, msg, args)Info(msg, args)InfoContext(ctx, msg, args)Log(ctx, level, msg, args)LogAttrs(ctx, level, msg, attrs)NewLogLogger(h, level)SetDefault(l)Warn(msg, args)WarnContext(ctx, msg, args)type Attrtype Handlertype HandlerOptionstype JSONHandlertype Kindtype Leveltype LevelVartype Levelertype LogValuertype Loggertype Recordtype Sourcetype TextHandlertype ValueAny(key, value)Bool(key, v)Duration(key, v)Float64(key, v)Group(key, args)GroupAttrs(key, attrs)Int(key, value)Int64(key, value)String(key, value)Time(key, v)Uint64(key, v)(a) Equal(b)(a) String()NewJSONHandler(w, opts)(h) Enabled(_, level)(h) Handle(_, r)(h) WithAttrs(attrs)(h) WithGroup(name)(k) String()SetLogLoggerLevel(level)(l) AppendText(b)(l) Level()(l) MarshalJSON()(l) MarshalText()(l) String()(l) UnmarshalJSON(data)(l) UnmarshalText(data)(v) AppendText(b)(v) Level()(v) MarshalText()(v) Set(l)(v) String()(v) UnmarshalText(data)Default()New(h)With(args)(l) Debug(msg, args)(l) DebugContext(ctx, msg, args)(l) Enabled(ctx, level)(l) Error(msg, args)(l) ErrorContext(ctx, msg, args)(l) Handler()(l) Info(msg, args)(l) InfoContext(ctx, msg, args)(l) Log(ctx, level, msg, args)(l) LogAttrs(ctx, level, msg, attrs)(l) Warn(msg, args)(l) WarnContext(ctx, msg, args)(l) With(args)(l) WithGroup(name)NewRecord(t, level, msg, pc)(r) Add(args)(r) AddAttrs(attrs)(r) Attrs(f)(r) Clone()(r) NumAttrs()(r) Source()NewTextHandler(w, opts)(h) Enabled(_, level)(h) Handle(_, r)(h) WithAttrs(attrs)(h) WithGroup(name)AnyValue(v)BoolValue(v)DurationValue(v)Float64Value(v)GroupValue(as)Int64Value(v)IntValue(v)StringValue(value)TimeValue(v)Uint64Value(v)(v) Any()(v) Bool()(v) Duration()(v) Equal(w)(v) Float64()(v) Group()(v) Int64()(v) Kind()(v) LogValuer()(v) Resolve()(v) String()(v) Time()(v) Uint64()
##  ![](https://pkg.go.dev/static/shared/icon/code_gm_grey_24dp.svg) Documentation [¶](https://pkg.go.dev/log/slog#section-documentation "Go to Documentation")
### Overview [¶](https://pkg.go.dev/log/slog#pkg-overview "Go to Overview")
  * [Levels](https://pkg.go.dev/log/slog#hdr-Levels)
  * [Groups](https://pkg.go.dev/log/slog#hdr-Groups)
  * [Contexts](https://pkg.go.dev/log/slog#hdr-Contexts)
  * [Attrs and Values](https://pkg.go.dev/log/slog#hdr-Attrs_and_Values)
  * [Customizing a type's logging behavior](https://pkg.go.dev/log/slog#hdr-Customizing_a_type_s_logging_behavior)
  * [Wrapping output methods](https://pkg.go.dev/log/slog#hdr-Wrapping_output_methods)
  * [Working with Records](https://pkg.go.dev/log/slog#hdr-Working_with_Records)
  * [Performance considerations](https://pkg.go.dev/log/slog#hdr-Performance_considerations)
  * [Writing a handler](https://pkg.go.dev/log/slog#hdr-Writing_a_handler)


Package slog provides structured logging, in which log records include a message, a severity level, and various other attributes expressed as key-value pairs. 
It defines a type, [Logger](https://pkg.go.dev/log/slog#Logger), which provides several methods (such as [Logger.Info](https://pkg.go.dev/log/slog#Logger.Info) and [Logger.Error](https://pkg.go.dev/log/slog#Logger.Error)) for reporting events of interest. 
Each Logger is associated with a [Handler](https://pkg.go.dev/log/slog#Handler). A Logger output method creates a [Record](https://pkg.go.dev/log/slog#Record) from the method arguments and passes it to the Handler, which decides how to handle it. There is a default Logger accessible through top-level functions (such as [Info](https://pkg.go.dev/log/slog#Info) and [Error](https://pkg.go.dev/log/slog#Error)) that call the corresponding Logger methods. 
A log record consists of a time, a level, a message, and a set of key-value pairs, where the keys are strings and the values may be of any type. As an example, 
```
slog.Info("hello", "count", 3)

```

creates a record containing the time of the call, a level of Info, the message "hello", and a single pair with key "count" and value 3. 
The [Info](https://pkg.go.dev/log/slog#Info) top-level function calls the [Logger.Info](https://pkg.go.dev/log/slog#Logger.Info) method on the default Logger. In addition to [Logger.Info](https://pkg.go.dev/log/slog#Logger.Info), there are methods for Debug, Warn and Error levels. Besides these convenience methods for common levels, there is also a [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log) method which takes the level as an argument. Each of these methods has a corresponding top-level function that uses the default logger. 
The default handler formats the log record's message, time, level, and attributes as a string and passes it to the [log](https://pkg.go.dev/log) package. 
```
2022/11/08 15:28:26 INFO hello count=3

```

For more control over the output format, create a logger with a different handler. This statement uses [New](https://pkg.go.dev/log/slog#New) to create a new logger with a [TextHandler](https://pkg.go.dev/log/slog#TextHandler) that writes structured records in text form to standard error: 
```
logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

```

[TextHandler](https://pkg.go.dev/log/slog#TextHandler) output is a sequence of key=value pairs, easily and unambiguously parsed by machine. This statement: 
```
logger.Info("hello", "count", 3)

```

produces this output: 
```
time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3

```

The package also provides [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler), whose output is line-delimited JSON: 
```
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("hello", "count", 3)

```

produces this output: 
```
{"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}

```

Both [TextHandler](https://pkg.go.dev/log/slog#TextHandler) and [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler) can be configured with [HandlerOptions](https://pkg.go.dev/log/slog#HandlerOptions). There are options for setting the minimum level (see Levels, below), displaying the source file and line of the log call, and modifying attributes before they are logged. 
Setting a logger as the default with 
```
slog.SetDefault(logger)

```

will cause the top-level functions like [Info](https://pkg.go.dev/log/slog#Info) to use it. [SetDefault](https://pkg.go.dev/log/slog#SetDefault) also updates the default logger used by the [log](https://pkg.go.dev/log) package, so that existing applications that use [log.Printf](https://pkg.go.dev/log#Printf) and related functions will send log records to the logger's handler without needing to be rewritten. 
Some attributes are common to many log calls. For example, you may wish to include the URL or trace identifier of a server request with all log events arising from the request. Rather than repeat the attribute with every log call, you can use [Logger.With](https://pkg.go.dev/log/slog#Logger.With) to construct a new Logger containing the attributes: 
```
logger2 := logger.With("url", r.URL)

```

The arguments to With are the same key-value pairs used in [Logger.Info](https://pkg.go.dev/log/slog#Logger.Info). The result is a new Logger with the same handler as the original, but additional attributes that will appear in the output of every call. 
#### Levels [¶](https://pkg.go.dev/log/slog#hdr-Levels "Go to Levels")
A [Level](https://pkg.go.dev/log/slog#Level) is an integer representing the importance or severity of a log event. The higher the level, the more severe the event. This package defines constants for the most common levels, but any int can be used as a level. 
In an application, you may wish to log messages only at a certain level or greater. One common configuration is to log messages at Info or higher levels, suppressing debug logging until it is needed. The built-in handlers can be configured with the minimum level to output by setting [HandlerOptions.Level]. The program's `main` function typically does this. The default value is LevelInfo. 
Setting the [HandlerOptions.Level] field to a [Level](https://pkg.go.dev/log/slog#Level) value fixes the handler's minimum level throughout its lifetime. Setting it to a [LevelVar](https://pkg.go.dev/log/slog#LevelVar) allows the level to be varied dynamically. A LevelVar holds a Level and is safe to read or write from multiple goroutines. To vary the level dynamically for an entire program, first initialize a global LevelVar: 
```
var programLevel = new(slog.LevelVar) // Info by default

```

Then use the LevelVar to construct a handler, and make it the default: 
```
h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
slog.SetDefault(slog.New(h))

```

Now the program can change its logging level with a single statement: 
```
programLevel.Set(slog.LevelDebug)

```

#### Groups [¶](https://pkg.go.dev/log/slog#hdr-Groups "Go to Groups")
Attributes can be collected into groups. A group has a name that is used to qualify the names of its attributes. How this qualification is displayed depends on the handler. [TextHandler](https://pkg.go.dev/log/slog#TextHandler) separates the group and attribute names with a dot. [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler) treats each group as a separate JSON object, with the group name as the key. 
Use [Group](https://pkg.go.dev/log/slog#Group) to create a Group attribute from a name and a list of key-value pairs: 
```
slog.Group("request",
  "method", r.Method,
  "url", r.URL)

```

TextHandler would display this group as 
```
request.method=GET request.url=http://example.com

```

JSONHandler would display it as 
```
"request":{"method":"GET","url":"http://example.com"}

```

Use [Logger.WithGroup](https://pkg.go.dev/log/slog#Logger.WithGroup) to qualify all of a Logger's output with a group name. Calling WithGroup on a Logger results in a new Logger with the same Handler as the original, but with all its attributes qualified by the group name. 
This can help prevent duplicate attribute keys in large systems, where subsystems might use the same keys. Pass each subsystem a different Logger with its own group name so that potential duplicates are qualified: 
```
logger := slog.Default().With("id", systemID)
parserLogger := logger.WithGroup("parser")
parseInput(input, parserLogger)

```

When parseInput logs with parserLogger, its keys will be qualified with "parser", so even if it uses the common key "id", the log line will have distinct keys. 
#### Contexts [¶](https://pkg.go.dev/log/slog#hdr-Contexts "Go to Contexts")
Some handlers may wish to include information from the [context.Context](https://pkg.go.dev/context#Context) that is available at the call site. One example of such information is the identifier for the current span when tracing is enabled. 
The [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log) and [Logger.LogAttrs](https://pkg.go.dev/log/slog#Logger.LogAttrs) methods take a context as a first argument, as do their corresponding top-level functions. 
Although the convenience methods on Logger (Info and so on) and the corresponding top-level functions do not take a context, the alternatives ending in "Context" do. For example, 
```
slog.InfoContext(ctx, "message")

```

It is recommended to pass a context to an output method if one is available. 
#### Attrs and Values [¶](https://pkg.go.dev/log/slog#hdr-Attrs_and_Values "Go to Attrs and Values")
An [Attr](https://pkg.go.dev/log/slog#Attr) is a key-value pair. The Logger output methods accept Attrs as well as alternating keys and values. The statement 
```
slog.Info("hello", slog.Int("count", 3))

```

behaves the same as 
```
slog.Info("hello", "count", 3)

```

There are convenience constructors for [Attr](https://pkg.go.dev/log/slog#Attr) such as [Int](https://pkg.go.dev/log/slog#Int), [String](https://pkg.go.dev/log/slog#String), and [Bool](https://pkg.go.dev/log/slog#Bool) for common types, as well as the function [Any](https://pkg.go.dev/log/slog#Any) for constructing Attrs of any type. 
The value part of an Attr is a type called [Value](https://pkg.go.dev/log/slog#Value). Like an [any], a Value can hold any Go value, but it can represent typical values, including all numbers and strings, without an allocation. 
For the most efficient log output, use [Logger.LogAttrs](https://pkg.go.dev/log/slog#Logger.LogAttrs). It is similar to [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log) but accepts only Attrs, not alternating keys and values; this allows it, too, to avoid allocation. 
The call 
```
logger.LogAttrs(ctx, slog.LevelInfo, "hello", slog.Int("count", 3))

```

is the most efficient way to achieve the same output as 
```
slog.InfoContext(ctx, "hello", "count", 3)

```

#### Customizing a type's logging behavior [¶](https://pkg.go.dev/log/slog#hdr-Customizing_a_type_s_logging_behavior "Go to Customizing a type's logging behavior")
If a type implements the [LogValuer](https://pkg.go.dev/log/slog#LogValuer) interface, the [Value](https://pkg.go.dev/log/slog#Value) returned from its LogValue method is used for logging. You can use this to control how values of the type appear in logs. For example, you can redact secret information like passwords, or gather a struct's fields in a Group. See the examples under [LogValuer](https://pkg.go.dev/log/slog#LogValuer) for details. 
A LogValue method may return a Value that itself implements [LogValuer](https://pkg.go.dev/log/slog#LogValuer). The [Value.Resolve](https://pkg.go.dev/log/slog#Value.Resolve) method handles these cases carefully, avoiding infinite loops and unbounded recursion. Handler authors and others may wish to use [Value.Resolve](https://pkg.go.dev/log/slog#Value.Resolve) instead of calling LogValue directly. 
#### Wrapping output methods [¶](https://pkg.go.dev/log/slog#hdr-Wrapping_output_methods "Go to Wrapping output methods")
The logger functions use reflection over the call stack to find the file name and line number of the logging call within the application. This can produce incorrect source information for functions that wrap slog. For instance, if you define this function in file mylog.go: 
```
func Infof(logger *slog.Logger, format string, args ...any) {
  logger.Info(fmt.Sprintf(format, args...))
}

```

and you call it like this in main.go: 
```
Infof(slog.Default(), "hello, %s", "world")

```

then slog will report the source file as mylog.go, not main.go. 
A correct implementation of Infof will obtain the source location (pc) and pass it to NewRecord. The Infof function in the package-level example called "wrapping" demonstrates how to do this. 
#### Working with Records [¶](https://pkg.go.dev/log/slog#hdr-Working_with_Records "Go to Working with Records")
Sometimes a Handler will need to modify a Record before passing it on to another Handler or backend. A Record contains a mixture of simple public fields (e.g. Time, Level, Message) and hidden fields that refer to state (such as attributes) indirectly. This means that modifying a simple copy of a Record (e.g. by calling [Record.Add](https://pkg.go.dev/log/slog#Record.Add) or [Record.AddAttrs](https://pkg.go.dev/log/slog#Record.AddAttrs) to add attributes) may have unexpected effects on the original. Before modifying a Record, use [Record.Clone](https://pkg.go.dev/log/slog#Record.Clone) to create a copy that shares no state with the original, or create a new Record with [NewRecord](https://pkg.go.dev/log/slog#NewRecord) and build up its Attrs by traversing the old ones with [Record.Attrs](https://pkg.go.dev/log/slog#Record.Attrs). 
#### Performance considerations [¶](https://pkg.go.dev/log/slog#hdr-Performance_considerations "Go to Performance considerations")
If profiling your application demonstrates that logging is taking significant time, the following suggestions may help. 
If many log lines have a common attribute, use [Logger.With](https://pkg.go.dev/log/slog#Logger.With) to create a Logger with that attribute. The built-in handlers will format that attribute only once, at the call to [Logger.With](https://pkg.go.dev/log/slog#Logger.With). The [Handler](https://pkg.go.dev/log/slog#Handler) interface is designed to allow that optimization, and a well-written Handler should take advantage of it. 
The arguments to a log call are always evaluated, even if the log event is discarded. If possible, defer computation so that it happens only if the value is actually logged. For example, consider the call 
```
slog.Info("starting request", "url", r.URL.String()) // may compute String unnecessarily

```

The URL.String method will be called even if the logger discards Info-level events. Instead, pass the URL directly: 
```
slog.Info("starting request", "url", &r.URL) // calls URL.String only if needed

```

The built-in [TextHandler](https://pkg.go.dev/log/slog#TextHandler) will call its String method, but only if the log event is enabled. Avoiding the call to String also preserves the structure of the underlying value. For example [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler) emits the components of the parsed URL as a JSON object. If you want to avoid eagerly paying the cost of the String call without causing the handler to potentially inspect the structure of the value, wrap the value in a fmt.Stringer implementation that hides its Marshal methods. 
You can also use the [LogValuer](https://pkg.go.dev/log/slog#LogValuer) interface to avoid unnecessary work in disabled log calls. Say you need to log some expensive value: 
```
slog.Debug("frobbing", "value", computeExpensiveValue(arg))

```

Even if this line is disabled, computeExpensiveValue will be called. To avoid that, define a type implementing LogValuer: 
```
type expensive struct { arg int }
func (e expensive) LogValue() slog.Value {
  return slog.AnyValue(computeExpensiveValue(e.arg))
}

```

Then use a value of that type in log calls: 
```
slog.Debug("frobbing", "value", expensive{arg})

```

Now computeExpensiveValue will only be called when the line is enabled. 
The built-in handlers acquire a lock before calling [io.Writer.Write](https://pkg.go.dev/io#Writer.Write) to ensure that exactly one [Record](https://pkg.go.dev/log/slog#Record) is written at a time in its entirety. Although each log record has a timestamp, the built-in handlers do not use that time to sort the written records. User-defined handlers are responsible for their own locking and sorting. 
#### Writing a handler [¶](https://pkg.go.dev/log/slog#hdr-Writing_a_handler "Go to Writing a handler")
For a guide to writing a custom handler, see <https://golang.org/s/slog-handler-guide>. 
Example (DiscardHandler) [¶](https://pkg.go.dev/log/slog#example-package-DiscardHandler "Go to Example \(DiscardHandler\)")
```
Output:
level=INFO msg="message 1"

```

Share Format Run
Example (Wrapping) [¶](https://pkg.go.dev/log/slog#example-package-Wrapping "Go to Example \(Wrapping\)")
```
Output:
level=INFO source=example_wrap_test.go:43 msg="message, formatted"

```

Share Format Run
### Index [¶](https://pkg.go.dev/log/slog#pkg-index "Go to Index")
  * [Constants](https://pkg.go.dev/log/slog#pkg-constants)
  * [func Debug(msg string, args ...any)](https://pkg.go.dev/log/slog#Debug)
  * [func DebugContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#DebugContext)
  * [func Error(msg string, args ...any)](https://pkg.go.dev/log/slog#Error)
  * [func ErrorContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#ErrorContext)
  * [func Info(msg string, args ...any)](https://pkg.go.dev/log/slog#Info)
  * [func InfoContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#InfoContext)
  * [func Log(ctx context.Context, level Level, msg string, args ...any)](https://pkg.go.dev/log/slog#Log)
  * [func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)](https://pkg.go.dev/log/slog#LogAttrs)
  * [func NewLogLogger(h Handler, level Level) *log.Logger](https://pkg.go.dev/log/slog#NewLogLogger)
  * [func SetDefault(l *Logger)](https://pkg.go.dev/log/slog#SetDefault)
  * [func Warn(msg string, args ...any)](https://pkg.go.dev/log/slog#Warn)
  * [func WarnContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#WarnContext)
  * [type Attr](https://pkg.go.dev/log/slog#Attr)
  *     * [func Any(key string, value any) Attr](https://pkg.go.dev/log/slog#Any)
    * [func Bool(key string, v bool) Attr](https://pkg.go.dev/log/slog#Bool)
    * [func Duration(key string, v time.Duration) Attr](https://pkg.go.dev/log/slog#Duration)
    * [func Float64(key string, v float64) Attr](https://pkg.go.dev/log/slog#Float64)
    * [func Group(key string, args ...any) Attr](https://pkg.go.dev/log/slog#Group)
    * [func GroupAttrs(key string, attrs ...Attr) Attr](https://pkg.go.dev/log/slog#GroupAttrs)
    * [func Int(key string, value int) Attr](https://pkg.go.dev/log/slog#Int)
    * [func Int64(key string, value int64) Attr](https://pkg.go.dev/log/slog#Int64)
    * [func String(key, value string) Attr](https://pkg.go.dev/log/slog#String)
    * [func Time(key string, v time.Time) Attr](https://pkg.go.dev/log/slog#Time)
    * [func Uint64(key string, v uint64) Attr](https://pkg.go.dev/log/slog#Uint64)
  *     * [func (a Attr) Equal(b Attr) bool](https://pkg.go.dev/log/slog#Attr.Equal)
    * [func (a Attr) String() string](https://pkg.go.dev/log/slog#Attr.String)
  * [type Handler](https://pkg.go.dev/log/slog#Handler)
  * [type HandlerOptions](https://pkg.go.dev/log/slog#HandlerOptions)
  * [type JSONHandler](https://pkg.go.dev/log/slog#JSONHandler)
  *     * [func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler](https://pkg.go.dev/log/slog#NewJSONHandler)
  *     * [func (h *JSONHandler) Enabled(_ context.Context, level Level) bool](https://pkg.go.dev/log/slog#JSONHandler.Enabled)
    * [func (h *JSONHandler) Handle(_ context.Context, r Record) error](https://pkg.go.dev/log/slog#JSONHandler.Handle)
    * [func (h *JSONHandler) WithAttrs(attrs []Attr) Handler](https://pkg.go.dev/log/slog#JSONHandler.WithAttrs)
    * [func (h *JSONHandler) WithGroup(name string) Handler](https://pkg.go.dev/log/slog#JSONHandler.WithGroup)
  * [type Kind](https://pkg.go.dev/log/slog#Kind)
  *     * [func (k Kind) String() string](https://pkg.go.dev/log/slog#Kind.String)
  * [type Level](https://pkg.go.dev/log/slog#Level)
  *     * [func SetLogLoggerLevel(level Level) (oldLevel Level)](https://pkg.go.dev/log/slog#SetLogLoggerLevel)
  *     * [func (l Level) AppendText(b []byte) ([]byte, error)](https://pkg.go.dev/log/slog#Level.AppendText)
    * [func (l Level) Level() Level](https://pkg.go.dev/log/slog#Level.Level)
    * [func (l Level) MarshalJSON() ([]byte, error)](https://pkg.go.dev/log/slog#Level.MarshalJSON)
    * [func (l Level) MarshalText() ([]byte, error)](https://pkg.go.dev/log/slog#Level.MarshalText)
    * [func (l Level) String() string](https://pkg.go.dev/log/slog#Level.String)
    * [func (l *Level) UnmarshalJSON(data []byte) error](https://pkg.go.dev/log/slog#Level.UnmarshalJSON)
    * [func (l *Level) UnmarshalText(data []byte) error](https://pkg.go.dev/log/slog#Level.UnmarshalText)
  * [type LevelVar](https://pkg.go.dev/log/slog#LevelVar)
  *     * [func (v *LevelVar) AppendText(b []byte) ([]byte, error)](https://pkg.go.dev/log/slog#LevelVar.AppendText)
    * [func (v *LevelVar) Level() Level](https://pkg.go.dev/log/slog#LevelVar.Level)
    * [func (v *LevelVar) MarshalText() ([]byte, error)](https://pkg.go.dev/log/slog#LevelVar.MarshalText)
    * [func (v *LevelVar) Set(l Level)](https://pkg.go.dev/log/slog#LevelVar.Set)
    * [func (v *LevelVar) String() string](https://pkg.go.dev/log/slog#LevelVar.String)
    * [func (v *LevelVar) UnmarshalText(data []byte) error](https://pkg.go.dev/log/slog#LevelVar.UnmarshalText)
  * [type Leveler](https://pkg.go.dev/log/slog#Leveler)
  * [type LogValuer](https://pkg.go.dev/log/slog#LogValuer)
  * [type Logger](https://pkg.go.dev/log/slog#Logger)
  *     * [func Default() *Logger](https://pkg.go.dev/log/slog#Default)
    * [func New(h Handler) *Logger](https://pkg.go.dev/log/slog#New)
    * [func With(args ...any) *Logger](https://pkg.go.dev/log/slog#With)
  *     * [func (l *Logger) Debug(msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.Debug)
    * [func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.DebugContext)
    * [func (l *Logger) Enabled(ctx context.Context, level Level) bool](https://pkg.go.dev/log/slog#Logger.Enabled)
    * [func (l *Logger) Error(msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.Error)
    * [func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.ErrorContext)
    * [func (l *Logger) Handler() Handler](https://pkg.go.dev/log/slog#Logger.Handler)
    * [func (l *Logger) Info(msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.Info)
    * [func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.InfoContext)
    * [func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.Log)
    * [func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)](https://pkg.go.dev/log/slog#Logger.LogAttrs)
    * [func (l *Logger) Warn(msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.Warn)
    * [func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)](https://pkg.go.dev/log/slog#Logger.WarnContext)
    * [func (l *Logger) With(args ...any) *Logger](https://pkg.go.dev/log/slog#Logger.With)
    * [func (l *Logger) WithGroup(name string) *Logger](https://pkg.go.dev/log/slog#Logger.WithGroup)
  * [type Record](https://pkg.go.dev/log/slog#Record)
  *     * [func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record](https://pkg.go.dev/log/slog#NewRecord)
  *     * [func (r *Record) Add(args ...any)](https://pkg.go.dev/log/slog#Record.Add)
    * [func (r *Record) AddAttrs(attrs ...Attr)](https://pkg.go.dev/log/slog#Record.AddAttrs)
    * [func (r Record) Attrs(f func(Attr) bool)](https://pkg.go.dev/log/slog#Record.Attrs)
    * [func (r Record) Clone() Record](https://pkg.go.dev/log/slog#Record.Clone)
    * [func (r Record) NumAttrs() int](https://pkg.go.dev/log/slog#Record.NumAttrs)
    * [func (r Record) Source() *Source](https://pkg.go.dev/log/slog#Record.Source)
  * [type Source](https://pkg.go.dev/log/slog#Source)
  * [type TextHandler](https://pkg.go.dev/log/slog#TextHandler)
  *     * [func NewTextHandler(w io.Writer, opts *HandlerOptions) *TextHandler](https://pkg.go.dev/log/slog#NewTextHandler)
  *     * [func (h *TextHandler) Enabled(_ context.Context, level Level) bool](https://pkg.go.dev/log/slog#TextHandler.Enabled)
    * [func (h *TextHandler) Handle(_ context.Context, r Record) error](https://pkg.go.dev/log/slog#TextHandler.Handle)
    * [func (h *TextHandler) WithAttrs(attrs []Attr) Handler](https://pkg.go.dev/log/slog#TextHandler.WithAttrs)
    * [func (h *TextHandler) WithGroup(name string) Handler](https://pkg.go.dev/log/slog#TextHandler.WithGroup)
  * [type Value](https://pkg.go.dev/log/slog#Value)
  *     * [func AnyValue(v any) Value](https://pkg.go.dev/log/slog#AnyValue)
    * [func BoolValue(v bool) Value](https://pkg.go.dev/log/slog#BoolValue)
    * [func DurationValue(v time.Duration) Value](https://pkg.go.dev/log/slog#DurationValue)
    * [func Float64Value(v float64) Value](https://pkg.go.dev/log/slog#Float64Value)
    * [func GroupValue(as ...Attr) Value](https://pkg.go.dev/log/slog#GroupValue)
    * [func Int64Value(v int64) Value](https://pkg.go.dev/log/slog#Int64Value)
    * [func IntValue(v int) Value](https://pkg.go.dev/log/slog#IntValue)
    * [func StringValue(value string) Value](https://pkg.go.dev/log/slog#StringValue)
    * [func TimeValue(v time.Time) Value](https://pkg.go.dev/log/slog#TimeValue)
    * [func Uint64Value(v uint64) Value](https://pkg.go.dev/log/slog#Uint64Value)
  *     * [func (v Value) Any() any](https://pkg.go.dev/log/slog#Value.Any)
    * [func (v Value) Bool() bool](https://pkg.go.dev/log/slog#Value.Bool)
    * [func (v Value) Duration() time.Duration](https://pkg.go.dev/log/slog#Value.Duration)
    * [func (v Value) Equal(w Value) bool](https://pkg.go.dev/log/slog#Value.Equal)
    * [func (v Value) Float64() float64](https://pkg.go.dev/log/slog#Value.Float64)
    * [func (v Value) Group() []Attr](https://pkg.go.dev/log/slog#Value.Group)
    * [func (v Value) Int64() int64](https://pkg.go.dev/log/slog#Value.Int64)
    * [func (v Value) Kind() Kind](https://pkg.go.dev/log/slog#Value.Kind)
    * [func (v Value) LogValuer() LogValuer](https://pkg.go.dev/log/slog#Value.LogValuer)
    * [func (v Value) Resolve() (rv Value)](https://pkg.go.dev/log/slog#Value.Resolve)
    * [func (v Value) String() string](https://pkg.go.dev/log/slog#Value.String)
    * [func (v Value) Time() time.Time](https://pkg.go.dev/log/slog#Value.Time)
    * [func (v Value) Uint64() uint64](https://pkg.go.dev/log/slog#Value.Uint64)


#### Examples [¶](https://pkg.go.dev/log/slog#pkg-examples "Go to Examples")
  * [Package (DiscardHandler)](https://pkg.go.dev/log/slog#example-package-DiscardHandler)
  * [Package (Wrapping)](https://pkg.go.dev/log/slog#example-package-Wrapping)
  * [Group](https://pkg.go.dev/log/slog#example-Group)
  * [GroupAttrs](https://pkg.go.dev/log/slog#example-GroupAttrs)
  * [Handler (LevelHandler)](https://pkg.go.dev/log/slog#example-Handler-LevelHandler)
  * [HandlerOptions (CustomLevels)](https://pkg.go.dev/log/slog#example-HandlerOptions-CustomLevels)
  * [LogValuer (Group)](https://pkg.go.dev/log/slog#example-LogValuer-Group)
  * [LogValuer (Secret)](https://pkg.go.dev/log/slog#example-LogValuer-Secret)
  * [SetLogLoggerLevel (Log)](https://pkg.go.dev/log/slog#example-SetLogLoggerLevel-Log)
  * [SetLogLoggerLevel (Slog)](https://pkg.go.dev/log/slog#example-SetLogLoggerLevel-Slog)


### Constants [¶](https://pkg.go.dev/log/slog#pkg-constants "Go to Constants")
[View Source](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/handler.go;l=176) ```
const (
	// TimeKey is the key used by the built-in handlers for the time
	// when the log method is called. The associated Value is a [time.Time].
	TimeKey = "time"
	// LevelKey is the key used by the built-in handlers for the level
	// of the log call. The associated value is a [Level].
	LevelKey = "level"
	// MessageKey is the key used by the built-in handlers for the
	// message of the log call. The associated value is a string.
	MessageKey = "msg"
	// SourceKey is the key used by the built-in handlers for the source file
	// and line of the log call. The associated value is a *[Source].
	SourceKey = "source"
)
```

Keys for "built-in" attributes. 
### Variables [¶](https://pkg.go.dev/log/slog#pkg-variables "Go to Variables")
This section is empty.
### Functions [¶](https://pkg.go.dev/log/slog#pkg-functions "Go to Functions")
####  func [Debug](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=280) [¶](https://pkg.go.dev/log/slog#Debug "Go to Debug")
```
func Debug(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Debug calls [Logger.Debug](https://pkg.go.dev/log/slog#Logger.Debug) on the default logger. 
####  func [DebugContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=285) [¶](https://pkg.go.dev/log/slog#DebugContext "Go to DebugContext")
```
func DebugContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

DebugContext calls [Logger.DebugContext](https://pkg.go.dev/log/slog#Logger.DebugContext) on the default logger. 
####  func [Error](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=310) [¶](https://pkg.go.dev/log/slog#Error "Go to Error")
```
func Error(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Error calls [Logger.Error](https://pkg.go.dev/log/slog#Logger.Error) on the default logger. 
####  func [ErrorContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=315) [¶](https://pkg.go.dev/log/slog#ErrorContext "Go to ErrorContext")
```
func ErrorContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

ErrorContext calls [Logger.ErrorContext](https://pkg.go.dev/log/slog#Logger.ErrorContext) on the default logger. 
####  func [Info](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=290) [¶](https://pkg.go.dev/log/slog#Info "Go to Info")
```
func Info(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Info calls [Logger.Info](https://pkg.go.dev/log/slog#Logger.Info) on the default logger. 
####  func [InfoContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=295) [¶](https://pkg.go.dev/log/slog#InfoContext "Go to InfoContext")
```
func InfoContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

InfoContext calls [Logger.InfoContext](https://pkg.go.dev/log/slog#Logger.InfoContext) on the default logger. 
####  func [Log](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=320) [¶](https://pkg.go.dev/log/slog#Log "Go to Log")
```
func Log(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Log calls [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log) on the default logger. 
####  func [LogAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=325) [¶](https://pkg.go.dev/log/slog#LogAttrs "Go to LogAttrs")
```
func LogAttrs(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level), msg string[](https://pkg.go.dev/builtin#string), attrs ...Attr[](https://pkg.go.dev/log/slog#Attr))
```

LogAttrs calls [Logger.LogAttrs](https://pkg.go.dev/log/slog#Logger.LogAttrs) on the default logger. 
####  func [NewLogLogger](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=174) [¶](https://pkg.go.dev/log/slog#NewLogLogger "Go to NewLogLogger")
```
func NewLogLogger(h Handler[](https://pkg.go.dev/log/slog#Handler), level Level[](https://pkg.go.dev/log/slog#Level)) *log[](https://pkg.go.dev/log).Logger[](https://pkg.go.dev/log#Logger)
```

NewLogLogger returns a new [log.Logger](https://pkg.go.dev/log#Logger) such that each call to its Output method dispatches a Record to the specified handler. The logger acts as a bridge from the older log API to newer structured logging handlers. 
####  func [SetDefault](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=62) [¶](https://pkg.go.dev/log/slog#SetDefault "Go to SetDefault")
```
func SetDefault(l *Logger[](https://pkg.go.dev/log/slog#Logger))
```

SetDefault makes l the default [Logger](https://pkg.go.dev/log/slog#Logger), which is used by the top-level functions [Info](https://pkg.go.dev/log/slog#Info), [Debug](https://pkg.go.dev/log/slog#Debug) and so on. After this call, output from the log package's default Logger (as with [log.Print](https://pkg.go.dev/log#Print), etc.) will be logged using l's Handler, at a level controlled by [SetLogLoggerLevel](https://pkg.go.dev/log/slog#SetLogLoggerLevel). 
####  func [Warn](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=300) [¶](https://pkg.go.dev/log/slog#Warn "Go to Warn")
```
func Warn(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Warn calls [Logger.Warn](https://pkg.go.dev/log/slog#Logger.Warn) on the default logger. 
####  func [WarnContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=305) [¶](https://pkg.go.dev/log/slog#WarnContext "Go to WarnContext")
```
func WarnContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

WarnContext calls [Logger.WarnContext](https://pkg.go.dev/log/slog#Logger.WarnContext) on the default logger. 
### Types [¶](https://pkg.go.dev/log/slog#pkg-types "Go to Types")
####  type [Attr](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=12) [¶](https://pkg.go.dev/log/slog#Attr "Go to Attr")
```
type Attr struct {
	Key  string[](https://pkg.go.dev/builtin#string)
	Value Value[](https://pkg.go.dev/log/slog#Value)
}
```

An Attr is a key-value pair. 
####  func [Any](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=93) [¶](https://pkg.go.dev/log/slog#Any "Go to Any")
```
func Any(key string[](https://pkg.go.dev/builtin#string), value any[](https://pkg.go.dev/builtin#any)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Any returns an Attr for the supplied value. See [AnyValue](https://pkg.go.dev/log/slog#AnyValue) for how values are treated. 
####  func [Bool](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=44) [¶](https://pkg.go.dev/log/slog#Bool "Go to Bool")
```
func Bool(key string[](https://pkg.go.dev/builtin#string), v bool[](https://pkg.go.dev/builtin#bool)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Bool returns an Attr for a bool. 
####  func [Duration](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=55) [¶](https://pkg.go.dev/log/slog#Duration "Go to Duration")
```
func Duration(key string[](https://pkg.go.dev/builtin#string), v time[](https://pkg.go.dev/time).Duration[](https://pkg.go.dev/time#Duration)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Duration returns an Attr for a [time.Duration](https://pkg.go.dev/time#Duration). 
####  func [Float64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=39) [¶](https://pkg.go.dev/log/slog#Float64 "Go to Float64")
```
func Float64(key string[](https://pkg.go.dev/builtin#string), v float64[](https://pkg.go.dev/builtin#float64)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Float64 returns an Attr for a floating-point number. 
####  func [Group](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=66) [¶](https://pkg.go.dev/log/slog#Group "Go to Group")
```
func Group(key string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Group returns an Attr for a Group [Value](https://pkg.go.dev/log/slog#Value). The first argument is the key; the remaining arguments are converted to Attrs as in [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log). 
Use Group to collect several key-value pairs under a single key on a log line, or as the result of LogValue in order to log a single value as multiple Attrs. 
Example [¶](https://pkg.go.dev/log/slog#example-Group "Go to Example")
```
Output:
level=INFO msg=finished req.method=GET req.url=localhost status=200 duration=1s

```

Share Format Run
####  func [GroupAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=75) [¶](https://pkg.go.dev/log/slog#GroupAttrs "Go to GroupAttrs") added in go1.25.0
```
func GroupAttrs(key string[](https://pkg.go.dev/builtin#string), attrs ...Attr[](https://pkg.go.dev/log/slog#Attr)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

GroupAttrs returns an Attr for a Group [Value](https://pkg.go.dev/log/slog#Value) consisting of the given Attrs. 
GroupAttrs is a more efficient version of [Group](https://pkg.go.dev/log/slog#Group) that accepts only [Attr](https://pkg.go.dev/log/slog#Attr) values. 
Example [¶](https://pkg.go.dev/log/slog#example-GroupAttrs "Go to Example")
```
Output:
level=INFO msg=finished status=200 req.method=POST req.url=localhost req.content-length=0
level=INFO msg=finished status=200 method=POST url=localhost content-length=0

```

Share Format Run
####  func [Int](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=29) [¶](https://pkg.go.dev/log/slog#Int "Go to Int")
```
func Int(key string[](https://pkg.go.dev/builtin#string), value int[](https://pkg.go.dev/builtin#int)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Int converts an int to an int64 and returns an Attr with that value. 
####  func [Int64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=23) [¶](https://pkg.go.dev/log/slog#Int64 "Go to Int64")
```
func Int64(key string[](https://pkg.go.dev/builtin#string), value int64[](https://pkg.go.dev/builtin#int64)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Int64 returns an Attr for an int64. 
####  func [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=18) [¶](https://pkg.go.dev/log/slog#String "Go to String")
```
func String(key, value string[](https://pkg.go.dev/builtin#string)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

String returns an Attr for a string value. 
####  func [Time](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=50) [¶](https://pkg.go.dev/log/slog#Time "Go to Time")
```
func Time(key string[](https://pkg.go.dev/builtin#string), v time[](https://pkg.go.dev/time).Time[](https://pkg.go.dev/time#Time)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Time returns an Attr for a [time.Time](https://pkg.go.dev/time#Time). It discards the monotonic portion. 
####  func [Uint64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=34) [¶](https://pkg.go.dev/log/slog#Uint64 "Go to Uint64")
```
func Uint64(key string[](https://pkg.go.dev/builtin#string), v uint64[](https://pkg.go.dev/builtin#uint64)) Attr[](https://pkg.go.dev/log/slog#Attr)
```

Uint64 returns an Attr for a uint64. 
####  func (Attr) [Equal](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=98) [¶](https://pkg.go.dev/log/slog#Attr.Equal "Go to Attr.Equal")
```
func (a Attr[](https://pkg.go.dev/log/slog#Attr)) Equal(b Attr[](https://pkg.go.dev/log/slog#Attr)) bool[](https://pkg.go.dev/builtin#bool)
```

Equal reports whether a and b have equal keys and values. 
####  func (Attr) [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go;l=102) [¶](https://pkg.go.dev/log/slog#Attr.String "Go to Attr.String")
```
func (a Attr[](https://pkg.go.dev/log/slog#Attr)) String() string[](https://pkg.go.dev/builtin#string)
```

####  type [Handler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/handler.go;l=33) [¶](https://pkg.go.dev/log/slog#Handler "Go to Handler")
```
type Handler interface {
	// Enabled reports whether the handler handles records at the given level.
	// The handler ignores records whose level is lower.
	// It is called early, before any arguments are processed,
	// to save effort if the log event should be discarded.
	// If called from a Logger method, the first argument is the context
	// passed to that method, or context.Background() if nil was passed
	// or the method does not take a context.
	// The context is passed so Enabled can use its values
	// to make a decision.
	Enabled(context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), Level[](https://pkg.go.dev/log/slog#Level)) bool[](https://pkg.go.dev/builtin#bool)
	// Handle handles the Record.
	// It will only be called when Enabled returns true.
	// The Context argument is as for Enabled.
	// It is present solely to provide Handlers access to the context's values.
	// Canceling the context should not affect record processing.
	// (Among other things, log messages may be necessary to debug a
	// cancellation-related problem.)
	//
	// Handle methods that produce output should observe the following rules:
	//  - If r.Time is the zero time, ignore the time.
	//  - If r.PC is zero, ignore it.
	//  - Attr's values should be resolved.
	//  - If an Attr's key and value are both the zero value, ignore the Attr.
	//   This can be tested with attr.Equal(Attr{}).
	//  - If a group's key is empty, inline the group's Attrs.
	//  - If a group has no Attrs (even if it has a non-empty key),
	//   ignore it.
	//
	// [Logger] discards any errors from Handle. Wrap the Handle method to
	// process any errors from Handlers.
	Handle(context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), Record[](https://pkg.go.dev/log/slog#Record)) error[](https://pkg.go.dev/builtin#error)
	// WithAttrs returns a new Handler whose attributes consist of
	// both the receiver's attributes and the arguments.
	// The Handler owns the slice: it may retain, modify or discard it.
	WithAttrs(attrs []Attr[](https://pkg.go.dev/log/slog#Attr)) Handler[](https://pkg.go.dev/log/slog#Handler)
	// WithGroup returns a new Handler with the given group appended to
	// the receiver's existing groups.
	// The keys of all subsequent attributes, whether added by With or in a
	// Record, should be qualified by the sequence of group names.
	//
	// How this qualification happens is up to the Handler, so long as
	// this Handler's attribute keys differ from those of another Handler
	// with a different sequence of group names.
	//
	// A Handler should treat WithGroup as starting a Group of Attrs that ends
	// at the end of the log event. That is,
	//
	//   logger.WithGroup("s").LogAttrs(ctx, level, msg, slog.Int("a", 1), slog.Int("b", 2))
	//
	// should behave like
	//
	//   logger.LogAttrs(ctx, level, msg, slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))
	//
	// If the name is empty, WithGroup returns the receiver.
	WithGroup(name string[](https://pkg.go.dev/builtin#string)) Handler[](https://pkg.go.dev/log/slog#Handler)
}
```

A Handler handles log records produced by a Logger. 
A typical handler may print log records to standard error, or write them to a file or database, or perhaps augment them with additional attributes and pass them on to another handler. 
Any of the Handler's methods may be called concurrently with itself or with other methods. It is the responsibility of the Handler to manage this concurrency. 
Users of the slog package should not invoke Handler methods directly. They should use the methods of [Logger](https://pkg.go.dev/log/slog#Logger) instead. 
Before implementing your own handler, consult <https://go.dev/s/slog-handler-guide>. 
Example (LevelHandler) [¶](https://pkg.go.dev/log/slog#example-Handler-LevelHandler "Go to Example \(LevelHandler\)")
This example shows how to Use a LevelHandler to change the level of an existing Handler while preserving its other behavior. 
This example demonstrates increasing the log level to reduce a logger's output. 
Another typical use would be to decrease the log level (to LevelDebug, say) during a part of the program that was suspected of containing a bug. 
```
Output:
level=WARN msg=printed

```

Share Format Run
```
var DiscardHandler Handler[](https://pkg.go.dev/log/slog#Handler) = discardHandler{}
```

DiscardHandler discards all log output. DiscardHandler.Enabled returns false for all Levels. 
####  type [HandlerOptions](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/handler.go;l=135) [¶](https://pkg.go.dev/log/slog#HandlerOptions "Go to HandlerOptions")
```
type HandlerOptions struct {
	// AddSource causes the handler to compute the source code position
	// of the log statement and add a SourceKey attribute to the output.
	AddSource bool[](https://pkg.go.dev/builtin#bool)
	// Level reports the minimum record level that will be logged.
	// The handler discards records with lower levels.
	// If Level is nil, the handler assumes LevelInfo.
	// The handler calls Level.Level for each record processed;
	// to adjust the minimum level dynamically, use a LevelVar.
	Level Leveler[](https://pkg.go.dev/log/slog#Leveler)
	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// The attribute's value has been resolved (see [Value.Resolve]).
	// If ReplaceAttr returns a zero Attr, the attribute is discarded.
	//
	// The built-in attributes with keys "time", "level", "source", and "msg"
	// are passed to this function, except that time is omitted
	// if zero, and source is omitted if AddSource is false.
	//
	// The first argument is a list of currently open groups that contain the
	// Attr. It must not be retained or modified. ReplaceAttr is never called
	// for Group attributes, only their contents. For example, the attribute
	// list
	//
	//   Int("a", 1), Group("g", Int("b", 2)), Int("c", 3)
	//
	// results in consecutive calls to ReplaceAttr with the following arguments:
	//
	//   nil, Int("a", 1)
	//   []string{"g"}, Int("b", 2)
	//   nil, Int("c", 3)
	//
	// ReplaceAttr can be used to change the default keys of the built-in
	// attributes, convert types (for example, to replace a `time.Time` with the
	// integer seconds since the Unix epoch), sanitize personal information, or
	// remove attributes from the output.
	ReplaceAttr func(groups []string[](https://pkg.go.dev/builtin#string), a Attr[](https://pkg.go.dev/log/slog#Attr)) Attr[](https://pkg.go.dev/log/slog#Attr)
}
```

HandlerOptions are options for a [TextHandler](https://pkg.go.dev/log/slog#TextHandler) or [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler). A zero HandlerOptions consists entirely of default values. 
Example (CustomLevels) [¶](https://pkg.go.dev/log/slog#example-HandlerOptions-CustomLevels "Go to Example \(CustomLevels\)")
This example demonstrates using custom log levels and custom log level names. In addition to the default log levels, it introduces Trace, Notice, and Emergency levels. The ReplaceAttr changes the way levels are printed for both the standard log levels and the custom log levels. 
```
Output:
sev=EMERGENCY msg="missing pilots"
sev=ERROR msg="failed to start engines" err="missing fuel"
sev=WARNING msg="falling back to default value"
sev=NOTICE msg="all systems are running"
sev=INFO msg="initiating launch"
sev=DEBUG msg="starting background job"
sev=TRACE msg="button clicked"

```

Share Format Run
####  type [JSONHandler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=23) [¶](https://pkg.go.dev/log/slog#JSONHandler "Go to JSONHandler")
```
type JSONHandler struct {
	// contains filtered or unexported fields
}
```

JSONHandler is a [Handler](https://pkg.go.dev/log/slog#Handler) that writes Records to an [io.Writer](https://pkg.go.dev/io#Writer) as line-delimited JSON objects. 
####  func [NewJSONHandler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=30) [¶](https://pkg.go.dev/log/slog#NewJSONHandler "Go to NewJSONHandler")
```
func NewJSONHandler(w io[](https://pkg.go.dev/io).Writer[](https://pkg.go.dev/io#Writer), opts *HandlerOptions[](https://pkg.go.dev/log/slog#HandlerOptions)) *JSONHandler[](https://pkg.go.dev/log/slog#JSONHandler)
```

NewJSONHandler creates a [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler) that writes to w, using the given options. If opts is nil, the default options are used. 
####  func (*JSONHandler) [Enabled](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=46) [¶](https://pkg.go.dev/log/slog#JSONHandler.Enabled "Go to JSONHandler.Enabled")
```
func (h *JSONHandler[](https://pkg.go.dev/log/slog#JSONHandler)) Enabled(_ context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level)) bool[](https://pkg.go.dev/builtin#bool)
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower. 
####  func (*JSONHandler) [Handle](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=88) [¶](https://pkg.go.dev/log/slog#JSONHandler.Handle "Go to JSONHandler.Handle")
```
func (h *JSONHandler[](https://pkg.go.dev/log/slog#JSONHandler)) Handle(_ context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), r Record[](https://pkg.go.dev/log/slog#Record)) error[](https://pkg.go.dev/builtin#error)
```

Handle formats its argument [Record](https://pkg.go.dev/log/slog#Record) as a JSON object on a single line. 
If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output as with json.Marshal. 
The level's key is "level" and its value is the result of calling [Level.String](https://pkg.go.dev/log/slog#Level.String). 
If the AddSource option is set and source information is available, the key is "source", and the value is a record of type [Source](https://pkg.go.dev/log/slog#Source). 
The message's key is "msg". 
To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr]. 
Values are formatted as with an [encoding/json.Encoder](https://pkg.go.dev/encoding/json#Encoder) with SetEscapeHTML(false), with two exceptions. 
First, an Attr whose Value is of type error is formatted as a string, by calling its Error method. Only errors in Attrs receive this special treatment, not errors embedded in structs, slices, maps or other data structures that are processed by the [encoding/json](https://pkg.go.dev/encoding/json) package. 
Second, an encoding failure does not cause Handle to return an error. Instead, the error message is formatted as a string. 
Each call to Handle results in a single serialized call to io.Writer.Write. 
####  func (*JSONHandler) [WithAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=52) [¶](https://pkg.go.dev/log/slog#JSONHandler.WithAttrs "Go to JSONHandler.WithAttrs")
```
func (h *JSONHandler[](https://pkg.go.dev/log/slog#JSONHandler)) WithAttrs(attrs []Attr[](https://pkg.go.dev/log/slog#Attr)) Handler[](https://pkg.go.dev/log/slog#Handler)
```

WithAttrs returns a new [JSONHandler](https://pkg.go.dev/log/slog#JSONHandler) whose attributes consists of h's attributes followed by attrs. 
####  func (*JSONHandler) [WithGroup](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go;l=56) [¶](https://pkg.go.dev/log/slog#JSONHandler.WithGroup "Go to JSONHandler.WithGroup")
```
func (h *JSONHandler[](https://pkg.go.dev/log/slog#JSONHandler)) WithGroup(name string[](https://pkg.go.dev/builtin#string)) Handler[](https://pkg.go.dev/log/slog#Handler)
```

####  type [Kind](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=44) [¶](https://pkg.go.dev/log/slog#Kind "Go to Kind")
```
type Kind int[](https://pkg.go.dev/builtin#int)
```

Kind is the kind of a [Value](https://pkg.go.dev/log/slog#Value). 
```
const (
	KindAny Kind[](https://pkg.go.dev/log/slog#Kind) = iota[](https://pkg.go.dev/builtin#iota)
	KindBool
	KindDuration
	KindFloat64
	KindInt64
	KindString
	KindTime
	KindUint64
	KindGroup
	KindLogValuer
)
```

####  func (Kind) [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=75) [¶](https://pkg.go.dev/log/slog#Kind.String "Go to Kind.String")
```
func (k Kind[](https://pkg.go.dev/log/slog#Kind)) String() string[](https://pkg.go.dev/builtin#string)
```

####  type [Level](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=17) [¶](https://pkg.go.dev/log/slog#Level "Go to Level")
```
type Level int[](https://pkg.go.dev/builtin#int)
```

A Level is the importance or severity of a log event. The higher the level, the more important or severe the event. 
```
const (
	LevelDebug Level[](https://pkg.go.dev/log/slog#Level) = -4
	LevelInfo Level[](https://pkg.go.dev/log/slog#Level) = 0
	LevelWarn Level[](https://pkg.go.dev/log/slog#Level) = 4
	LevelError Level[](https://pkg.go.dev/log/slog#Level) = 8
)
```

Names for common levels. 
Level numbers are inherently arbitrary, but we picked them to satisfy three constraints. Any system can map them to another numbering scheme if it wishes. 
First, we wanted the default level to be Info, Since Levels are ints, Info is the default value for int, zero. 
Second, we wanted to make it easy to use levels to specify logger verbosity. Since a larger level means a more severe event, a logger that accepts events with smaller (or more negative) level means a more verbose logger. Logger verbosity is thus the negation of event severity, and the default verbosity of 0 accepts all events at least as severe as INFO. 
Third, we wanted some room between levels to accommodate schemes with named levels between ours. For example, Google Cloud Logging defines a Notice level between Info and Warn. Since there are only a few of these intermediate levels, the gap between the numbers need not be large. Our gap of 4 matches OpenTelemetry's mapping. Subtracting 9 from an OpenTelemetry level in the DEBUG, INFO, WARN and ERROR ranges converts it to the corresponding slog Level range. OpenTelemetry also has the names TRACE and FATAL, which slog does not. But those OpenTelemetry levels can still be represented as slog Levels by using the appropriate integers. 
####  func [SetLogLoggerLevel](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=44) [¶](https://pkg.go.dev/log/slog#SetLogLoggerLevel "Go to SetLogLoggerLevel") added in go1.22.0
```
func SetLogLoggerLevel(level Level[](https://pkg.go.dev/log/slog#Level)) (oldLevel Level[](https://pkg.go.dev/log/slog#Level))
```

SetLogLoggerLevel controls the level for the bridge to the [log](https://pkg.go.dev/log) package. 
Before [SetDefault](https://pkg.go.dev/log/slog#SetDefault) is called, slog top-level logging functions call the default [log.Logger](https://pkg.go.dev/log#Logger). In that mode, SetLogLoggerLevel sets the minimum level for those calls. By default, the minimum level is Info, so calls to [Debug](https://pkg.go.dev/log/slog#Debug) (as well as top-level logging calls at lower levels) will not be passed to the log.Logger. After calling 
```
slog.SetLogLoggerLevel(slog.LevelDebug)

```

calls to [Debug](https://pkg.go.dev/log/slog#Debug) will be passed to the log.Logger. 
After [SetDefault](https://pkg.go.dev/log/slog#SetDefault) is called, calls to the default [log.Logger](https://pkg.go.dev/log#Logger) are passed to the slog default handler. In that mode, SetLogLoggerLevel sets the level at which those calls are logged. That is, after calling 
```
slog.SetLogLoggerLevel(slog.LevelDebug)

```

A call to [log.Printf](https://pkg.go.dev/log#Printf) will result in output at level [LevelDebug](https://pkg.go.dev/log/slog#LevelDebug). 
SetLogLoggerLevel returns the previous value. 
Example (Log) [¶](https://pkg.go.dev/log/slog#example-SetLogLoggerLevel-Log "Go to Example \(Log\)")
This example shows how to use slog.SetLogLoggerLevel to change the minimal level of the internal default handler for slog package before calling slog.SetDefault. 
```
Output:
log debug
INFO info
log debug
DEBUG debug
INFO info

```

Share Format Run
Example (Slog) [¶](https://pkg.go.dev/log/slog#example-SetLogLoggerLevel-Slog "Go to Example \(Slog\)")
This example shows how to use slog.SetLogLoggerLevel to change the minimal level of the internal writer that uses the custom handler for log package after calling slog.SetDefault. 
```
Output:
level=ERROR msg=error

```

Share Format Run
####  func (Level) [AppendText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=103) [¶](https://pkg.go.dev/log/slog#Level.AppendText "Go to Level.AppendText") added in go1.24.0
```
func (l Level[](https://pkg.go.dev/log/slog#Level)) AppendText(b []byte[](https://pkg.go.dev/builtin#byte)) ([]byte[](https://pkg.go.dev/builtin#byte), error[](https://pkg.go.dev/builtin#error))
```

AppendText implements [encoding.TextAppender](https://pkg.go.dev/encoding#TextAppender) by calling [Level.String](https://pkg.go.dev/log/slog#Level.String). 
####  func (Level) [Level](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=156) [¶](https://pkg.go.dev/log/slog#Level.Level "Go to Level.Level")
```
func (l Level[](https://pkg.go.dev/log/slog#Level)) Level() Level[](https://pkg.go.dev/log/slog#Level)
```

Level returns the receiver. It implements [Leveler](https://pkg.go.dev/log/slog#Leveler). 
####  func (Level) [MarshalJSON](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=81) [¶](https://pkg.go.dev/log/slog#Level.MarshalJSON "Go to Level.MarshalJSON")
```
func (l Level[](https://pkg.go.dev/log/slog#Level)) MarshalJSON() ([]byte[](https://pkg.go.dev/builtin#byte), error[](https://pkg.go.dev/builtin#error))
```

MarshalJSON implements [encoding/json.Marshaler](https://pkg.go.dev/encoding/json#Marshaler) by quoting the output of [Level.String](https://pkg.go.dev/log/slog#Level.String). 
####  func (Level) [MarshalText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=109) [¶](https://pkg.go.dev/log/slog#Level.MarshalText "Go to Level.MarshalText")
```
func (l Level[](https://pkg.go.dev/log/slog#Level)) MarshalText() ([]byte[](https://pkg.go.dev/builtin#byte), error[](https://pkg.go.dev/builtin#error))
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [Level.AppendText](https://pkg.go.dev/log/slog#Level.AppendText). 
####  func (Level) [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=59) [¶](https://pkg.go.dev/log/slog#Level.String "Go to Level.String")
```
func (l Level[](https://pkg.go.dev/log/slog#Level)) String() string[](https://pkg.go.dev/builtin#string)
```

String returns a name for the level. If the level has a name, then that name in uppercase is returned. If the level is between named values, then an integer is appended to the uppercased name. Examples: 
```
LevelWarn.String() => "WARN"
(LevelInfo+2).String() => "INFO+2"

```

####  func (*Level) [UnmarshalJSON](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=93) [¶](https://pkg.go.dev/log/slog#Level.UnmarshalJSON "Go to Level.UnmarshalJSON")
```
func (l *Level[](https://pkg.go.dev/log/slog#Level)) UnmarshalJSON(data []byte[](https://pkg.go.dev/builtin#byte)) error[](https://pkg.go.dev/builtin#error)
```

UnmarshalJSON implements [encoding/json.Unmarshaler](https://pkg.go.dev/encoding/json#Unmarshaler) It accepts any string produced by [Level.MarshalJSON](https://pkg.go.dev/log/slog#Level.MarshalJSON), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO". 
####  func (*Level) [UnmarshalText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=118) [¶](https://pkg.go.dev/log/slog#Level.UnmarshalText "Go to Level.UnmarshalText")
```
func (l *Level[](https://pkg.go.dev/log/slog#Level)) UnmarshalText(data []byte[](https://pkg.go.dev/builtin#byte)) error[](https://pkg.go.dev/builtin#error)
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler). It accepts any string produced by [Level.MarshalText](https://pkg.go.dev/log/slog#Level.MarshalText), ignoring case. It also accepts numeric offsets that would result in a different string on output. For example, "Error-8" would marshal as "INFO". 
####  type [LevelVar](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=163) [¶](https://pkg.go.dev/log/slog#LevelVar "Go to LevelVar")
```
type LevelVar struct {
	// contains filtered or unexported fields
}
```

A LevelVar is a [Level](https://pkg.go.dev/log/slog#Level) variable, to allow a [Handler](https://pkg.go.dev/log/slog#Handler) level to change dynamically. It implements [Leveler](https://pkg.go.dev/log/slog#Leveler) as well as a Set method, and it is safe for use by multiple goroutines. The zero LevelVar corresponds to [LevelInfo](https://pkg.go.dev/log/slog#LevelInfo). 
####  func (*LevelVar) [AppendText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=183) [¶](https://pkg.go.dev/log/slog#LevelVar.AppendText "Go to LevelVar.AppendText") added in go1.24.0
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) AppendText(b []byte[](https://pkg.go.dev/builtin#byte)) ([]byte[](https://pkg.go.dev/builtin#byte), error[](https://pkg.go.dev/builtin#error))
```

AppendText implements [encoding.TextAppender](https://pkg.go.dev/encoding#TextAppender) by calling [Level.AppendText](https://pkg.go.dev/log/slog#Level.AppendText). 
####  func (*LevelVar) [Level](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=168) [¶](https://pkg.go.dev/log/slog#LevelVar.Level "Go to LevelVar.Level")
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) Level() Level[](https://pkg.go.dev/log/slog#Level)
```

Level returns v's level. 
####  func (*LevelVar) [MarshalText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=189) [¶](https://pkg.go.dev/log/slog#LevelVar.MarshalText "Go to LevelVar.MarshalText")
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) MarshalText() ([]byte[](https://pkg.go.dev/builtin#byte), error[](https://pkg.go.dev/builtin#error))
```

MarshalText implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler) by calling [LevelVar.AppendText](https://pkg.go.dev/log/slog#LevelVar.AppendText). 
####  func (*LevelVar) [Set](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=173) [¶](https://pkg.go.dev/log/slog#LevelVar.Set "Go to LevelVar.Set")
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) Set(l Level[](https://pkg.go.dev/log/slog#Level))
```

Set sets v's level to l. 
####  func (*LevelVar) [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=177) [¶](https://pkg.go.dev/log/slog#LevelVar.String "Go to LevelVar.String")
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) String() string[](https://pkg.go.dev/builtin#string)
```

####  func (*LevelVar) [UnmarshalText](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=195) [¶](https://pkg.go.dev/log/slog#LevelVar.UnmarshalText "Go to LevelVar.UnmarshalText")
```
func (v *LevelVar[](https://pkg.go.dev/log/slog#LevelVar)) UnmarshalText(data []byte[](https://pkg.go.dev/builtin#byte)) error[](https://pkg.go.dev/builtin#error)
```

UnmarshalText implements [encoding.TextUnmarshaler](https://pkg.go.dev/encoding#TextUnmarshaler) by calling [Level.UnmarshalText](https://pkg.go.dev/log/slog#Level.UnmarshalText). 
####  type [Leveler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go;l=210) [¶](https://pkg.go.dev/log/slog#Leveler "Go to Leveler")
```
type Leveler interface {
	Level() Level[](https://pkg.go.dev/log/slog#Level)
}
```

A Leveler provides a [Level](https://pkg.go.dev/log/slog#Level) value. 
As Level itself implements Leveler, clients typically supply a Level value wherever a Leveler is needed, such as in [HandlerOptions](https://pkg.go.dev/log/slog#HandlerOptions). Clients who need to vary the level dynamically can provide a more complex Leveler implementation such as *[LevelVar](https://pkg.go.dev/log/slog#LevelVar). 
####  type [LogValuer](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=487) [¶](https://pkg.go.dev/log/slog#LogValuer "Go to LogValuer")
```
type LogValuer interface {
	LogValue() Value[](https://pkg.go.dev/log/slog#Value)
}
```

A LogValuer is any Go value that can convert itself into a Value for logging. 
This mechanism may be used to defer expensive operations until they are needed, or to expand a single value into a sequence of components. 
Example (Group) [¶](https://pkg.go.dev/log/slog#example-LogValuer-Group "Go to Example \(Group\)")
Share Format Run
Example (Secret) [¶](https://pkg.go.dev/log/slog#example-LogValuer-Secret "Go to Example \(Secret\)")
This example demonstrates a Value that replaces itself with an alternative representation to avoid revealing secrets. 
```
Output:
level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN

```

Share Format Run
####  type [Logger](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=111) [¶](https://pkg.go.dev/log/slog#Logger "Go to Logger")
```
type Logger struct {
	// contains filtered or unexported fields
}
```

A Logger records structured information about each call to its Log, Debug, Info, Warn, and Error methods. For each call, it creates a [Record](https://pkg.go.dev/log/slog#Record) and passes it to a [Handler](https://pkg.go.dev/log/slog#Handler). 
To create a new Logger, call [New](https://pkg.go.dev/log/slog#New) or a Logger method that begins "With". 
####  func [Default](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=55) [¶](https://pkg.go.dev/log/slog#Default "Go to Default")
```
func Default() *Logger[](https://pkg.go.dev/log/slog#Logger)
```

Default returns the default [Logger](https://pkg.go.dev/log/slog#Logger). 
####  func [New](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=151) [¶](https://pkg.go.dev/log/slog#New "Go to New")
```
func New(h Handler[](https://pkg.go.dev/log/slog#Handler)) *Logger[](https://pkg.go.dev/log/slog#Logger)
```

New creates a new Logger with the given non-nil Handler. 
####  func [With](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=159) [¶](https://pkg.go.dev/log/slog#With "Go to With")
```
func With(args ...any[](https://pkg.go.dev/builtin#any)) *Logger[](https://pkg.go.dev/log/slog#Logger)
```

With calls [Logger.With](https://pkg.go.dev/log/slog#Logger.With) on the default logger. 
####  func (*Logger) [Debug](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=198) [¶](https://pkg.go.dev/log/slog#Logger.Debug "Go to Logger.Debug")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Debug(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Debug logs at [LevelDebug](https://pkg.go.dev/log/slog#LevelDebug). 
####  func (*Logger) [DebugContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=203) [¶](https://pkg.go.dev/log/slog#Logger.DebugContext "Go to Logger.DebugContext")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) DebugContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

DebugContext logs at [LevelDebug](https://pkg.go.dev/log/slog#LevelDebug) with the given context. 
####  func (*Logger) [Enabled](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=164) [¶](https://pkg.go.dev/log/slog#Logger.Enabled "Go to Logger.Enabled")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Enabled(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level)) bool[](https://pkg.go.dev/builtin#bool)
```

Enabled reports whether l emits log records at the given context and level. 
####  func (*Logger) [Error](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=228) [¶](https://pkg.go.dev/log/slog#Logger.Error "Go to Logger.Error")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Error(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Error logs at [LevelError](https://pkg.go.dev/log/slog#LevelError). 
####  func (*Logger) [ErrorContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=233) [¶](https://pkg.go.dev/log/slog#Logger.ErrorContext "Go to Logger.ErrorContext")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) ErrorContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

ErrorContext logs at [LevelError](https://pkg.go.dev/log/slog#LevelError) with the given context. 
####  func (*Logger) [Handler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=121) [¶](https://pkg.go.dev/log/slog#Logger.Handler "Go to Logger.Handler")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Handler() Handler[](https://pkg.go.dev/log/slog#Handler)
```

Handler returns l's Handler. 
####  func (*Logger) [Info](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=208) [¶](https://pkg.go.dev/log/slog#Logger.Info "Go to Logger.Info")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Info(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Info logs at [LevelInfo](https://pkg.go.dev/log/slog#LevelInfo). 
####  func (*Logger) [InfoContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=213) [¶](https://pkg.go.dev/log/slog#Logger.InfoContext "Go to Logger.InfoContext")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) InfoContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

InfoContext logs at [LevelInfo](https://pkg.go.dev/log/slog#LevelInfo) with the given context. 
####  func (*Logger) [Log](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=188) [¶](https://pkg.go.dev/log/slog#Logger.Log "Go to Logger.Log")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Log(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Log emits a log record with the current time and the given level and message. The Record's Attrs consist of the Logger's attributes followed by the Attrs specified by args. 
The attribute arguments are processed as follows: 
  * If an argument is an Attr, it is used as is.
  * If an argument is a string and this is not the last argument, the following argument is treated as the value and the two are combined into an Attr.
  * Otherwise, the argument is treated as a value with key "!BADKEY".


####  func (*Logger) [LogAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=193) [¶](https://pkg.go.dev/log/slog#Logger.LogAttrs "Go to Logger.LogAttrs")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) LogAttrs(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level), msg string[](https://pkg.go.dev/builtin#string), attrs ...Attr[](https://pkg.go.dev/log/slog#Attr))
```

LogAttrs is a more efficient version of [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log) that accepts only Attrs. 
####  func (*Logger) [Warn](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=218) [¶](https://pkg.go.dev/log/slog#Logger.Warn "Go to Logger.Warn")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) Warn(msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

Warn logs at [LevelWarn](https://pkg.go.dev/log/slog#LevelWarn). 
####  func (*Logger) [WarnContext](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=223) [¶](https://pkg.go.dev/log/slog#Logger.WarnContext "Go to Logger.WarnContext")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) WarnContext(ctx context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), msg string[](https://pkg.go.dev/builtin#string), args ...any[](https://pkg.go.dev/builtin#any))
```

WarnContext logs at [LevelWarn](https://pkg.go.dev/log/slog#LevelWarn) with the given context. 
####  func (*Logger) [With](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=126) [¶](https://pkg.go.dev/log/slog#Logger.With "Go to Logger.With")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) With(args ...any[](https://pkg.go.dev/builtin#any)) *Logger[](https://pkg.go.dev/log/slog#Logger)
```

With returns a Logger that includes the given attributes in each output operation. Arguments are converted to attributes as if by [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log). 
####  func (*Logger) [WithGroup](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go;l=141) [¶](https://pkg.go.dev/log/slog#Logger.WithGroup "Go to Logger.WithGroup")
```
func (l *Logger[](https://pkg.go.dev/log/slog#Logger)) WithGroup(name string[](https://pkg.go.dev/builtin#string)) *Logger[](https://pkg.go.dev/log/slog#Logger)
```

WithGroup returns a Logger that starts a group, if name is non-empty. The keys of all attributes added to the Logger will be qualified by the given name. (How that qualification happens depends on the [Handler.WithGroup] method of the Logger's Handler.) 
If name is empty, WithGroup returns the receiver. 
####  type [Record](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=20) [¶](https://pkg.go.dev/log/slog#Record "Go to Record")
```
type Record struct {
	// The time at which the output method (Log, Info, etc.) was called.
	Time time[](https://pkg.go.dev/time).Time[](https://pkg.go.dev/time#Time)
	// The log message.
	Message string[](https://pkg.go.dev/builtin#string)
	// The level of the event.
	Level Level[](https://pkg.go.dev/log/slog#Level)
	// The program counter at the time the record was constructed, as determined
	// by runtime.Callers. If zero, no program counter is available.
	//
	// The only valid use for this value is as an argument to
	// [runtime.CallersFrames]. In particular, it must not be passed to
	// [runtime.FuncForPC].
	PC uintptr[](https://pkg.go.dev/builtin#uintptr)
	// contains filtered or unexported fields
}
```

A Record holds information about a log event. Copies of a Record share state. Do not modify a Record after handing out a copy to it. Call [NewRecord](https://pkg.go.dev/log/slog#NewRecord) to create a new Record. Use [Record.Clone](https://pkg.go.dev/log/slog#Record.Clone) to create a copy with no shared state. 
####  func [NewRecord](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=58) [¶](https://pkg.go.dev/log/slog#NewRecord "Go to NewRecord")
```
func NewRecord(t time[](https://pkg.go.dev/time).Time[](https://pkg.go.dev/time#Time), level Level[](https://pkg.go.dev/log/slog#Level), msg string[](https://pkg.go.dev/builtin#string), pc uintptr[](https://pkg.go.dev/builtin#uintptr)) Record[](https://pkg.go.dev/log/slog#Record)
```

NewRecord creates a [Record](https://pkg.go.dev/log/slog#Record) from the given arguments. Use [Record.AddAttrs](https://pkg.go.dev/log/slog#Record.AddAttrs) to add attributes to the Record. 
NewRecord is intended for logging APIs that want to support a [Handler](https://pkg.go.dev/log/slog#Handler) as a backend. 
####  func (*Record) [Add](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=129) [¶](https://pkg.go.dev/log/slog#Record.Add "Go to Record.Add")
```
func (r *Record[](https://pkg.go.dev/log/slog#Record)) Add(args ...any[](https://pkg.go.dev/builtin#any))
```

Add converts the args to Attrs as described in [Logger.Log](https://pkg.go.dev/log/slog#Logger.Log), then appends the Attrs to the [Record](https://pkg.go.dev/log/slog#Record)'s list of Attrs. It omits empty groups. 
####  func (*Record) [AddAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=97) [¶](https://pkg.go.dev/log/slog#Record.AddAttrs "Go to Record.AddAttrs")
```
func (r *Record[](https://pkg.go.dev/log/slog#Record)) AddAttrs(attrs ...Attr[](https://pkg.go.dev/log/slog#Attr))
```

AddAttrs appends the given Attrs to the [Record](https://pkg.go.dev/log/slog#Record)'s list of Attrs. It omits empty groups. 
####  func (Record) [Attrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=82) [¶](https://pkg.go.dev/log/slog#Record.Attrs "Go to Record.Attrs")
```
func (r Record[](https://pkg.go.dev/log/slog#Record)) Attrs(f func(Attr[](https://pkg.go.dev/log/slog#Attr)) bool[](https://pkg.go.dev/builtin#bool))
```

Attrs calls f on each Attr in the [Record](https://pkg.go.dev/log/slog#Record). Iteration stops if f returns false. 
####  func (Record) [Clone](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=70) [¶](https://pkg.go.dev/log/slog#Record.Clone "Go to Record.Clone")
```
func (r Record[](https://pkg.go.dev/log/slog#Record)) Clone() Record[](https://pkg.go.dev/log/slog#Record)
```

Clone returns a copy of the record with no shared state. The original record and the clone can both be modified without interfering with each other. 
####  func (Record) [NumAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=76) [¶](https://pkg.go.dev/log/slog#Record.NumAttrs "Go to Record.NumAttrs")
```
func (r Record[](https://pkg.go.dev/log/slog#Record)) NumAttrs() int[](https://pkg.go.dev/builtin#int)
```

NumAttrs returns the number of attributes in the [Record](https://pkg.go.dev/log/slog#Record). 
####  func (Record) [Source](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=220) [¶](https://pkg.go.dev/log/slog#Record.Source "Go to Record.Source") added in go1.25.0
```
func (r Record[](https://pkg.go.dev/log/slog#Record)) Source() *Source[](https://pkg.go.dev/log/slog#Source)
```

Source returns a new Source for the log event using r's PC. If the PC field is zero, meaning the Record was created without the necessary information or the location is unavailable, then nil is returned. 
####  type [Source](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go;l=185) [¶](https://pkg.go.dev/log/slog#Source "Go to Source")
```
type Source struct {
	// Function is the package path-qualified function name containing the
	// source line. If non-empty, this string uniquely identifies a single
	// function in the program. This may be the empty string if not known.
	Function string[](https://pkg.go.dev/builtin#string) `json:"function"`
	// File and Line are the file name and line number (1-based) of the source
	// line. These may be the empty string and zero, respectively, if not known.
	File string[](https://pkg.go.dev/builtin#string) `json:"file"`
	Line int[](https://pkg.go.dev/builtin#int)  `json:"line"`
}
```

Source describes the location of a line of source code. 
####  type [TextHandler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=21) [¶](https://pkg.go.dev/log/slog#TextHandler "Go to TextHandler")
```
type TextHandler struct {
	// contains filtered or unexported fields
}
```

TextHandler is a [Handler](https://pkg.go.dev/log/slog#Handler) that writes Records to an [io.Writer](https://pkg.go.dev/io#Writer) as a sequence of key=value pairs separated by spaces and followed by a newline. 
####  func [NewTextHandler](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=28) [¶](https://pkg.go.dev/log/slog#NewTextHandler "Go to NewTextHandler")
```
func NewTextHandler(w io[](https://pkg.go.dev/io).Writer[](https://pkg.go.dev/io#Writer), opts *HandlerOptions[](https://pkg.go.dev/log/slog#HandlerOptions)) *TextHandler[](https://pkg.go.dev/log/slog#TextHandler)
```

NewTextHandler creates a [TextHandler](https://pkg.go.dev/log/slog#TextHandler) that writes to w, using the given options. If opts is nil, the default options are used. 
####  func (*TextHandler) [Enabled](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=44) [¶](https://pkg.go.dev/log/slog#TextHandler.Enabled "Go to TextHandler.Enabled")
```
func (h *TextHandler[](https://pkg.go.dev/log/slog#TextHandler)) Enabled(_ context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), level Level[](https://pkg.go.dev/log/slog#Level)) bool[](https://pkg.go.dev/builtin#bool)
```

Enabled reports whether the handler handles records at the given level. The handler ignores records whose level is lower. 
####  func (*TextHandler) [Handle](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=92) [¶](https://pkg.go.dev/log/slog#TextHandler.Handle "Go to TextHandler.Handle")
```
func (h *TextHandler[](https://pkg.go.dev/log/slog#TextHandler)) Handle(_ context[](https://pkg.go.dev/context).Context[](https://pkg.go.dev/context#Context), r Record[](https://pkg.go.dev/log/slog#Record)) error[](https://pkg.go.dev/builtin#error)
```

Handle formats its argument [Record](https://pkg.go.dev/log/slog#Record) as a single line of space-separated key=value items. 
If the Record's time is zero, the time is omitted. Otherwise, the key is "time" and the value is output in RFC3339 format with millisecond precision. 
The level's key is "level" and its value is the result of calling [Level.String](https://pkg.go.dev/log/slog#Level.String). 
If the AddSource option is set and source information is available, the key is "source" and the value is output as FILE:LINE. 
The message's key is "msg". 
To modify these or other attributes, or remove them from the output, use [HandlerOptions.ReplaceAttr]. 
If a value implements [encoding.TextMarshaler](https://pkg.go.dev/encoding#TextMarshaler), the result of MarshalText is written. Otherwise, the result of [fmt.Sprint](https://pkg.go.dev/fmt#Sprint) is written. 
Keys and values are quoted with [strconv.Quote](https://pkg.go.dev/strconv#Quote) if they contain Unicode space characters, non-printing characters, '"' or '='. 
Keys inside groups consist of components (keys or group names) separated by dots. No further escaping is performed. Thus there is no way to determine from the key "a.b.c" whether there are two groups "a" and "b" and a key "c", or a single group "a.b" and a key "c", or single group "a" and a key "b.c". If it is necessary to reconstruct the group structure of a key even in the presence of dots inside components, use [HandlerOptions.ReplaceAttr] to encode that information in the key. 
Each call to Handle results in a single serialized call to io.Writer.Write. 
####  func (*TextHandler) [WithAttrs](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=50) [¶](https://pkg.go.dev/log/slog#TextHandler.WithAttrs "Go to TextHandler.WithAttrs")
```
func (h *TextHandler[](https://pkg.go.dev/log/slog#TextHandler)) WithAttrs(attrs []Attr[](https://pkg.go.dev/log/slog#Attr)) Handler[](https://pkg.go.dev/log/slog#Handler)
```

WithAttrs returns a new [TextHandler](https://pkg.go.dev/log/slog#TextHandler) whose attributes consists of h's attributes followed by attrs. 
####  func (*TextHandler) [WithGroup](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go;l=54) [¶](https://pkg.go.dev/log/slog#TextHandler.WithGroup "Go to TextHandler.WithGroup")
```
func (h *TextHandler[](https://pkg.go.dev/log/slog#TextHandler)) WithGroup(name string[](https://pkg.go.dev/builtin#string)) Handler[](https://pkg.go.dev/log/slog#Handler)
```

####  type [Value](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=21) [¶](https://pkg.go.dev/log/slog#Value "Go to Value")
```
type Value struct {
	// contains filtered or unexported fields
}
```

A Value can represent any Go value, but unlike type any, it can represent most small values without an allocation. The zero Value corresponds to nil. 
####  func [AnyValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=221) [¶](https://pkg.go.dev/log/slog#AnyValue "Go to AnyValue")
```
func AnyValue(v any[](https://pkg.go.dev/builtin#any)) Value[](https://pkg.go.dev/log/slog#Value)
```

AnyValue returns a [Value](https://pkg.go.dev/log/slog#Value) for the supplied value. 
If the supplied value is of type Value, it is returned unmodified. 
Given a value of one of Go's predeclared string, bool, or (non-complex) numeric types, AnyValue returns a Value of kind [KindString](https://pkg.go.dev/log/slog#KindString), [KindBool](https://pkg.go.dev/log/slog#KindBool), [KindUint64](https://pkg.go.dev/log/slog#KindUint64), [KindInt64](https://pkg.go.dev/log/slog#KindInt64), or [KindFloat64](https://pkg.go.dev/log/slog#KindFloat64). The width of the original numeric type is not preserved. 
Given a [time.Time](https://pkg.go.dev/time#Time) or [time.Duration](https://pkg.go.dev/time#Duration) value, AnyValue returns a Value of kind [KindTime](https://pkg.go.dev/log/slog#KindTime) or [KindDuration](https://pkg.go.dev/log/slog#KindDuration). The monotonic time is not preserved. 
For nil, or values of all other types, including named types whose underlying type is numeric, AnyValue returns a value of kind [KindAny](https://pkg.go.dev/log/slog#KindAny). 
####  func [BoolValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=134) [¶](https://pkg.go.dev/log/slog#BoolValue "Go to BoolValue")
```
func BoolValue(v bool[](https://pkg.go.dev/builtin#bool)) Value[](https://pkg.go.dev/log/slog#Value)
```

BoolValue returns a [Value](https://pkg.go.dev/log/slog#Value) for a bool. 
####  func [DurationValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=173) [¶](https://pkg.go.dev/log/slog#DurationValue "Go to DurationValue")
```
func DurationValue(v time[](https://pkg.go.dev/time).Duration[](https://pkg.go.dev/time#Duration)) Value[](https://pkg.go.dev/log/slog#Value)
```

DurationValue returns a [Value](https://pkg.go.dev/log/slog#Value) for a [time.Duration](https://pkg.go.dev/time#Duration). 
####  func [Float64Value](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=129) [¶](https://pkg.go.dev/log/slog#Float64Value "Go to Float64Value")
```
func Float64Value(v float64[](https://pkg.go.dev/builtin#float64)) Value[](https://pkg.go.dev/log/slog#Value)
```

Float64Value returns a [Value](https://pkg.go.dev/log/slog#Value) for a floating-point number. 
####  func [GroupValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=179) [¶](https://pkg.go.dev/log/slog#GroupValue "Go to GroupValue")
```
func GroupValue(as ...Attr[](https://pkg.go.dev/log/slog#Attr)) Value[](https://pkg.go.dev/log/slog#Value)
```

GroupValue returns a new [Value](https://pkg.go.dev/log/slog#Value) for a list of Attrs. The caller must not subsequently mutate the argument slice. 
####  func [Int64Value](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=119) [¶](https://pkg.go.dev/log/slog#Int64Value "Go to Int64Value")
```
func Int64Value(v int64[](https://pkg.go.dev/builtin#int64)) Value[](https://pkg.go.dev/log/slog#Value)
```

Int64Value returns a [Value](https://pkg.go.dev/log/slog#Value) for an int64. 
####  func [IntValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=114) [¶](https://pkg.go.dev/log/slog#IntValue "Go to IntValue")
```
func IntValue(v int[](https://pkg.go.dev/builtin#int)) Value[](https://pkg.go.dev/log/slog#Value)
```

IntValue returns a [Value](https://pkg.go.dev/log/slog#Value) for an int. 
####  func [StringValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=109) [¶](https://pkg.go.dev/log/slog#StringValue "Go to StringValue")
```
func StringValue(value string[](https://pkg.go.dev/builtin#string)) Value[](https://pkg.go.dev/log/slog#Value)
```

StringValue returns a new [Value](https://pkg.go.dev/log/slog#Value) for a string. 
####  func [TimeValue](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=153) [¶](https://pkg.go.dev/log/slog#TimeValue "Go to TimeValue")
```
func TimeValue(v time[](https://pkg.go.dev/time).Time[](https://pkg.go.dev/time#Time)) Value[](https://pkg.go.dev/log/slog#Value)
```

TimeValue returns a [Value](https://pkg.go.dev/log/slog#Value) for a [time.Time](https://pkg.go.dev/time#Time). It discards the monotonic portion. 
####  func [Uint64Value](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=124) [¶](https://pkg.go.dev/log/slog#Uint64Value "Go to Uint64Value")
```
func Uint64Value(v uint64[](https://pkg.go.dev/builtin#uint64)) Value[](https://pkg.go.dev/log/slog#Value)
```

Uint64Value returns a [Value](https://pkg.go.dev/log/slog#Value) for a uint64. 
####  func (Value) [Any](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=271) [¶](https://pkg.go.dev/log/slog#Value.Any "Go to Value.Any")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Any() any[](https://pkg.go.dev/builtin#any)
```

Any returns v's value as an any. 
####  func (Value) [Bool](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=336) [¶](https://pkg.go.dev/log/slog#Value.Bool "Go to Value.Bool")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Bool() bool[](https://pkg.go.dev/builtin#bool)
```

Bool returns v's value as a bool. It panics if v is not a bool. 
####  func (Value) [Duration](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=349) [¶](https://pkg.go.dev/log/slog#Value.Duration "Go to Value.Duration")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Duration() time[](https://pkg.go.dev/time).Duration[](https://pkg.go.dev/time#Duration)
```

Duration returns v's value as a [time.Duration](https://pkg.go.dev/time#Duration). It panics if v is not a time.Duration. 
####  func (Value) [Equal](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=421) [¶](https://pkg.go.dev/log/slog#Value.Equal "Go to Value.Equal")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Equal(w Value[](https://pkg.go.dev/log/slog#Value)) bool[](https://pkg.go.dev/builtin#bool)
```

Equal reports whether v and w represent the same Go value. 
####  func (Value) [Float64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=363) [¶](https://pkg.go.dev/log/slog#Value.Float64 "Go to Value.Float64")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Float64() float64[](https://pkg.go.dev/builtin#float64)
```

Float64 returns v's value as a float64. It panics if v is not a float64. 
####  func (Value) [Group](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=407) [¶](https://pkg.go.dev/log/slog#Value.Group "Go to Value.Group")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Group() []Attr[](https://pkg.go.dev/log/slog#Attr)
```

Group returns v's value as a []Attr. It panics if v's [Kind](https://pkg.go.dev/log/slog#Kind) is not [KindGroup](https://pkg.go.dev/log/slog#KindGroup). 
####  func (Value) [Int64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=318) [¶](https://pkg.go.dev/log/slog#Value.Int64 "Go to Value.Int64")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Int64() int64[](https://pkg.go.dev/builtin#int64)
```

Int64 returns v's value as an int64. It panics if v is not a signed integer. 
####  func (Value) [Kind](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=87) [¶](https://pkg.go.dev/log/slog#Value.Kind "Go to Value.Kind")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Kind() Kind[](https://pkg.go.dev/log/slog#Kind)
```

Kind returns v's Kind. 
####  func (Value) [LogValuer](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=401) [¶](https://pkg.go.dev/log/slog#Value.LogValuer "Go to Value.LogValuer")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) LogValuer() LogValuer[](https://pkg.go.dev/log/slog#LogValuer)
```

LogValuer returns v's value as a LogValuer. It panics if v is not a LogValuer. 
####  func (Value) [Resolve](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=500) [¶](https://pkg.go.dev/log/slog#Value.Resolve "Go to Value.Resolve")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Resolve() (rv Value[](https://pkg.go.dev/log/slog#Value))
```

Resolve repeatedly calls LogValue on v while it implements [LogValuer](https://pkg.go.dev/log/slog#LogValuer), and returns the result. If v resolves to a group, the group's attributes' values are not recursively resolved. If the number of LogValue calls exceeds a threshold, a Value containing an error is returned. Resolve's return value is guaranteed not to be of Kind [KindLogValuer](https://pkg.go.dev/log/slog#KindLogValuer). 
####  func (Value) [String](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=304) [¶](https://pkg.go.dev/log/slog#Value.String "Go to Value.String")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) String() string[](https://pkg.go.dev/builtin#string)
```

String returns Value's value as a string, formatted like [fmt.Sprint](https://pkg.go.dev/fmt#Sprint). Unlike the methods Int64, Float64, and so on, which panic if v is of the wrong kind, String never panics. 
####  func (Value) [Time](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=377) [¶](https://pkg.go.dev/log/slog#Value.Time "Go to Value.Time")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Time() time[](https://pkg.go.dev/time).Time[](https://pkg.go.dev/time#Time)
```

Time returns v's value as a [time.Time](https://pkg.go.dev/time#Time). It panics if v is not a time.Time. 
####  func (Value) [Uint64](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go;l=327) [¶](https://pkg.go.dev/log/slog#Value.Uint64 "Go to Value.Uint64")
```
func (v Value[](https://pkg.go.dev/log/slog#Value)) Uint64() uint64[](https://pkg.go.dev/builtin#uint64)
```

Uint64 returns v's value as a uint64. It panics if v is not an unsigned integer. 
##  ![](https://pkg.go.dev/static/shared/icon/insert_drive_file_gm_grey_24dp.svg) Source Files [¶](https://pkg.go.dev/log/slog#section-sourcefiles "Go to Source Files")
[View all Source files](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog)
  * [attr.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/attr.go "attr.go")
  * [doc.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/doc.go "doc.go")
  * [handler.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/handler.go "handler.go")
  * [json_handler.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/json_handler.go "json_handler.go")
  * [level.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/level.go "level.go")
  * [logger.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/logger.go "logger.go")
  * [record.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/record.go "record.go")
  * [text_handler.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/text_handler.go "text_handler.go")
  * [value.go](https://cs.opensource.google/go/go/+/go1.25.0:src/log/slog/value.go "value.go")


##  ![](https://pkg.go.dev/static/shared/icon/folder_gm_grey_24dp.svg) Directories [¶](https://pkg.go.dev/log/slog#section-directories "Go to Directories")
Show internal  Expand all
Path | Synopsis  
---|---  
![](https://pkg.go.dev/static/shared/icon/arrow_right_gm_grey_24dp.svg) [internal](https://pkg.go.dev/log/slog/internal@go1.25.0) |   
[benchmarks](https://pkg.go.dev/log/slog/internal/benchmarks@go1.25.0) Package benchmarks contains benchmarks for slog. | Package benchmarks contains benchmarks for slog.  
[buffer](https://pkg.go.dev/log/slog/internal/buffer@go1.25.0) Package buffer provides a pool-allocated byte buffer. | Package buffer provides a pool-allocated byte buffer.  
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
- URL: https://pkg.go.dev/log/slog
- Harvested: 2025-08-20T00:35:38.846164
- Profile: deep_research
- Agent: architect
