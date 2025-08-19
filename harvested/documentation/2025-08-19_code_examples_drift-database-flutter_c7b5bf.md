---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:20:51.783888'
profile: code_examples
source: https://drift.simonbinder.eu/docs/
topic: drift-database-flutter
---

# drift-database-flutter

[ Skip to content ](https://drift.simonbinder.eu/#welcome-to-drift)
[ ![logo](https://drift.simonbinder.eu/images/icon.png) ](https://drift.simonbinder.eu/ "Drift")
Drift 
Home 
Type to start searching
[ GitHub 
  * drift-2.28.1
  * 3k
  * 408

](https://github.com/simolus3/drift "Go to repository")
  * [ Home ](https://drift.simonbinder.eu/)
  * [ Documentation ](https://drift.simonbinder.eu/setup/)
  * [ Examples ](https://drift.simonbinder.eu/Examples/)
  * [ Tools ](https://drift.simonbinder.eu/Tools/)
  * [ Guides ](https://drift.simonbinder.eu/guides/datetime-migrations/)
  * [ Pub ](https://pub.dev/packages/drift)


[ ![logo](https://drift.simonbinder.eu/images/icon.png) ](https://drift.simonbinder.eu/ "Drift") Drift 
[ GitHub 
  * drift-2.28.1
  * 3k
  * 408

](https://github.com/simolus3/drift "Go to repository")
  * Home  [ Home  ](https://drift.simonbinder.eu/) Table of contents 
    * [ Welcome to drift  ](https://drift.simonbinder.eu/#welcome-to-drift)
    * [ Getting started  ](https://drift.simonbinder.eu/#getting-started)
  * Documentation  Documentation 
    * [ Getting Started  ](https://drift.simonbinder.eu/setup/)
    * [ Tables  ](https://drift.simonbinder.eu/dart_api/tables/)
    * [ Generated table rows  ](https://drift.simonbinder.eu/dart_api/rows/)
    * [ Manager  ](https://drift.simonbinder.eu/dart_api/manager/)
    * [ Transactions  ](https://drift.simonbinder.eu/dart_api/transactions/)
    * Core API  Core API 
      * [ Selects  ](https://drift.simonbinder.eu/dart_api/select/)
      * [ Writes (update, insert, delete)  ](https://drift.simonbinder.eu/dart_api/writes/)
      * [ Expressions  ](https://drift.simonbinder.eu/dart_api/expressions/)
      * [ Runtime schema inspection  ](https://drift.simonbinder.eu/dart_api/schema_inspection/)
      * [ Views  ](https://drift.simonbinder.eu/dart_api/views/)
      * [ DAOs  ](https://drift.simonbinder.eu/dart_api/daos/)
      * [ Stream queries  ](https://drift.simonbinder.eu/dart_api/streams/)
    * SQL API  SQL API 
      * [ Getting Started  ](https://drift.simonbinder.eu/sql_api/)
      * [ Drift files  ](https://drift.simonbinder.eu/sql_api/drift_files/)
      * [ Custom SQL types  ](https://drift.simonbinder.eu/sql_api/types/)
      * [ Supported sqlite extensions  ](https://drift.simonbinder.eu/sql_api/extensions/)
      * [ Custom queries  ](https://drift.simonbinder.eu/sql_api/custom_queries/)
      * [ Experimental IDE  ](https://drift.simonbinder.eu/sql_api/sql_ide/)
    * Migrations  Migrations 
      * [ Getting Started  ](https://drift.simonbinder.eu/Migrations/)
      * [ Exporting schemas  ](https://drift.simonbinder.eu/Migrations/exports/)
      * [ Schema migration helpers  ](https://drift.simonbinder.eu/Migrations/step_by_step/)
      * [ Testing migrations  ](https://drift.simonbinder.eu/Migrations/tests/)
      * [ The migrator API  ](https://drift.simonbinder.eu/Migrations/api/)
    * [ Type converters  ](https://drift.simonbinder.eu/type_converters/)
    * Code Generation  Code Generation 
      * [ Options overview  ](https://drift.simonbinder.eu/generation_options/)
      * [ Modular code generation  ](https://drift.simonbinder.eu/generation_options/modular/)
      * [ Using drift classes in other builders  ](https://drift.simonbinder.eu/generation_options/in_other_builders/)
    * Platforms  Platforms 
      * [ Platforms overview  ](https://drift.simonbinder.eu/Platforms/)
      * [ Native Drift (cross-platform)  ](https://drift.simonbinder.eu/Platforms/vm/)
      * [ Web  ](https://drift.simonbinder.eu/Platforms/web/)
      * [ PostgreSQL support  ](https://drift.simonbinder.eu/Platforms/postgres/)
      * [ Encryption  ](https://drift.simonbinder.eu/Platforms/encryption/)
      * [ Remote sqld & Turso  ](https://drift.simonbinder.eu/Platforms/libsql/)
    * [ Isolates  ](https://drift.simonbinder.eu/isolates/)
    * [ Testing  ](https://drift.simonbinder.eu/testing/)
    * [ FAQ  ](https://drift.simonbinder.eu/faq/)
    * [ Community  ](https://drift.simonbinder.eu/community_tools/)
  * Examples  Examples 
    * [ Examples  ](https://drift.simonbinder.eu/Examples/)
    * [ Flutter  ](https://drift.simonbinder.eu/Examples/flutter/)
    * [ Many-to-many relationships  ](https://drift.simonbinder.eu/Examples/relationships/)
    * [ Tracing database operations  ](https://drift.simonbinder.eu/Examples/tracing/)
    * [ Backend synchronization  ](https://drift.simonbinder.eu/Examples/server_sync/)
    * [ Importing and exporting databases  ](https://drift.simonbinder.eu/Examples/existing_databases/)
  * Tools  Tools 
    * [ CLI  ](https://drift.simonbinder.eu/Tools/)
    * [ DevTools extension  ](https://drift.simonbinder.eu/Tools/devtools/)
  * Guides  Guides 
    * [ DateTime Storage  ](https://drift.simonbinder.eu/guides/datetime-migrations/)
    * [ Upgrading Drift  ](https://drift.simonbinder.eu/upgrading/)
    * [ Migrate to Drift  ](https://drift.simonbinder.eu/migrating_to_drift/)
    * [ Install from GitHub  ](https://drift.simonbinder.eu/Internals/)
  * [ Pub  ](https://pub.dev/packages/drift)


Table of contents 
  * [ Welcome to drift  ](https://drift.simonbinder.eu/#welcome-to-drift)
  * [ Getting started  ](https://drift.simonbinder.eu/#getting-started)


[ ](https://github.com/simolus3/drift/edit/develop/docs/docs/index.md "Edit this page") [ ](https://github.com/simolus3/drift/issues/new?template=docs.md&title=Documentation issue: Home "Create documentation issue")
# Home
## Welcome to drift[¶](https://drift.simonbinder.eu/#welcome-to-drift "Permanent link")
Drift is a reactive persistence library for Dart and Flutter applications. It's built on top of database libraries like [the sqlite3 package](https://pub.dev/packages/sqlite3), [sqflite](https://pub.dev/packages/sqflite) and others. Adding to these libraries, drift provides additional features, like:
  * **Type safety** : Instead of writing SQL queries manually and parsing the `List<Map<String, dynamic>>` that they return, drift turns rows into objects of your choice.
  * **Stream queries** : Drift lets you "watch" your queries with zero additional effort. Any query can be turned into an auto-updating stream that emits new items when the underlying data changes.
  * **Fluent queries** : Drift generates a Dart api that you can use to write queries and automatically get their results. Keep an updated list of all users with `select(users).watch()`. That's it! No SQL to write, no rows to parse.
  * **Type-safe SQL** : If you prefer to write SQL, that's fine! Drift has an SQL parser and analyzer built in. It can parse your queries at compile time, figure out what columns they're going to return and generate Dart code to represent your rows.
  * **Migration utils** : Drift makes writing migrations easier thanks to utility functions like `.createAllTables()`. You don't need to manually write your `CREATE TABLE` statements and keep them updated.


And much more! Drift validates data before inserting it, so you can get helpful error messages instead of just an SQL error code. Of course, it supports transactions. And DAOs. And efficient batched insert statements. The list goes on.
## Getting started[¶](https://drift.simonbinder.eu/#getting-started "Permanent link")
To get started with drift, follow the [setup guide](https://drift.simonbinder.eu/setup/). It explains everything from setting up the dependencies to writing database classes and generating code.
It also links a few pages intended for developers getting started with drift, so that you can explore the areas you're most interested in first.
Back to top  [ Next  Getting Started  ](https://drift.simonbinder.eu/setup/)
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)
[ ](https://github.com/simolus3/drift "Project on GitHub")


## Source Information
- URL: https://drift.simonbinder.eu/docs/
- Harvested: 2025-08-19T01:20:51.783888
- Profile: code_examples
- Agent: unknown
