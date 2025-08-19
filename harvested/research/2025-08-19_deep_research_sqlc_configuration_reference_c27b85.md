---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:18:43.026349'
profile: deep_research
source: https://docs.sqlc.dev/en/latest/reference/config.html
topic: sqlc Configuration Reference
---

# sqlc Configuration Reference

[ ![Logo](https://docs.sqlc.dev/en/latest/_static/logo.png) ](https://docs.sqlc.dev/en/latest/index.html)
latest  stable  v1.29.0  v1.28.0  v1.27.0  v1.26.0  v1.25.0  v1.24.0  v1.23.0  v1.22.0  v1.21.0  v1.20.0  v1.19.1  v1.19.0  v1.18.0  v1.17.2  v1.17.1  v1.17.0  v1.16.0  v1.15.0  v1.14.0  v1.13.0  v1.12.0  v1.11.0  v1.10.0  v1.9.0  v1.8.0  v1.7.0 
Overview
  * [Installing sqlc](https://docs.sqlc.dev/en/latest/overview/install.html)


Tutorials
  * [Getting started with MySQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-mysql.html)
  * [Getting started with PostgreSQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)
  * [Getting started with SQLite](https://docs.sqlc.dev/en/latest/tutorials/getting-started-sqlite.html)


Commands
  * [`generate` - Generating code](https://docs.sqlc.dev/en/latest/howto/generate.html)
  * [`push` - Uploading projects](https://docs.sqlc.dev/en/latest/howto/push.html)
  * [`verify` - Verifying schema changes](https://docs.sqlc.dev/en/latest/howto/verify.html)
  * [`vet` - Linting queries](https://docs.sqlc.dev/en/latest/howto/vet.html)


How-to Guides
  * [Retrieving rows](https://docs.sqlc.dev/en/latest/howto/select.html)
  * [Counting rows](https://docs.sqlc.dev/en/latest/howto/query_count.html)
  * [Inserting rows](https://docs.sqlc.dev/en/latest/howto/insert.html)
  * [Updating rows](https://docs.sqlc.dev/en/latest/howto/update.html)
  * [Deleting rows](https://docs.sqlc.dev/en/latest/howto/delete.html)
  * [Preparing queries](https://docs.sqlc.dev/en/latest/howto/prepared_query.html)
  * [Using transactions](https://docs.sqlc.dev/en/latest/howto/transactions.html)
  * [Naming parameters](https://docs.sqlc.dev/en/latest/howto/named_parameters.html)
  * [Modifying the database schema](https://docs.sqlc.dev/en/latest/howto/ddl.html)
  * [Configuring generated structs](https://docs.sqlc.dev/en/latest/howto/structs.html)
  * [Embedding structs](https://docs.sqlc.dev/en/latest/howto/embedding.html)
  * [Overriding types](https://docs.sqlc.dev/en/latest/howto/overrides.html)
  * [Renaming fields](https://docs.sqlc.dev/en/latest/howto/rename.html)


sqlc Cloud
  * [Managed databases](https://docs.sqlc.dev/en/latest/howto/managed-databases.html)


Reference
  * [Changelog](https://docs.sqlc.dev/en/latest/reference/changelog.html)
  * [CLI](https://docs.sqlc.dev/en/latest/reference/cli.html)
  * [Configuration](https://docs.sqlc.dev/en/latest/reference/config.html)
    * [Version 2](https://docs.sqlc.dev/en/latest/reference/config.html#version-2)
      * [sql](https://docs.sqlc.dev/en/latest/reference/config.html#sql)
      * [codegen](https://docs.sqlc.dev/en/latest/reference/config.html#codegen)
      * [database](https://docs.sqlc.dev/en/latest/reference/config.html#database)
      * [analyzer](https://docs.sqlc.dev/en/latest/reference/config.html#analyzer)
      * [gen](https://docs.sqlc.dev/en/latest/reference/config.html#gen)
        * [go](https://docs.sqlc.dev/en/latest/reference/config.html#go)
        * [kotlin](https://docs.sqlc.dev/en/latest/reference/config.html#kotlin)
        * [python](https://docs.sqlc.dev/en/latest/reference/config.html#python)
        * [json](https://docs.sqlc.dev/en/latest/reference/config.html#json)
      * [plugins](https://docs.sqlc.dev/en/latest/reference/config.html#plugins)
      * [rules](https://docs.sqlc.dev/en/latest/reference/config.html#rules)
      * [Global overrides](https://docs.sqlc.dev/en/latest/reference/config.html#global-overrides)
    * [Version 1](https://docs.sqlc.dev/en/latest/reference/config.html#version-1)
      * [packages](https://docs.sqlc.dev/en/latest/reference/config.html#packages)
      * [overrides](https://docs.sqlc.dev/en/latest/reference/config.html#id1)
      * [rename](https://docs.sqlc.dev/en/latest/reference/config.html#rename)
  * [Datatypes](https://docs.sqlc.dev/en/latest/reference/datatypes.html)
  * [Environment variables](https://docs.sqlc.dev/en/latest/reference/environment-variables.html)
  * [Database and language support](https://docs.sqlc.dev/en/latest/reference/language-support.html)
  * [Macros](https://docs.sqlc.dev/en/latest/reference/macros.html)
  * [Query annotations](https://docs.sqlc.dev/en/latest/reference/query-annotations.html)


Conceptual Guides
  * [Using sqlc in CI/CD](https://docs.sqlc.dev/en/latest/howto/ci-cd.html)
  * [Using Go and pgx](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html)
  * [Using plugins](https://docs.sqlc.dev/en/latest/guides/plugins.html)
  * [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html)
  * [Privacy and data collection](https://docs.sqlc.dev/en/latest/guides/privacy.html)


Sponsored By
[ ![Riza logo](https://sqlc.dev/sponsors/riza-readme.png) ](https://riza.io?utm_source=sqlc+docs)
[sqlc](https://docs.sqlc.dev/en/latest/index.html)
  * [](https://docs.sqlc.dev/en/latest/index.html)
  * Configuration
  * [ View page source](https://docs.sqlc.dev/en/latest/_sources/reference/config.md.txt)


# Configuration[](https://docs.sqlc.dev/en/latest/reference/config.html#configuration "Link to this heading")
The `sqlc` tool is configured via a `sqlc.(yaml|yml)` or `sqlc.json` file. This file must be in the directory where the `sqlc` command is run.
## Version 2[](https://docs.sqlc.dev/en/latest/reference/config.html#version-2 "Link to this heading")
```
version:"2"
cloud:
project:"<PROJECT_ID>"
sql:
-schema:"postgresql/schema.sql"
queries:"postgresql/query.sql"
engine:"postgresql"
gen:
go:
package:"authors"
out:"postgresql"
database:
managed:true
rules:
-sqlc/db-prepare
-schema:"mysql/schema.sql"
queries:"mysql/query.sql"
engine:"mysql"
gen:
go:
package:"authors"
out:"mysql"

```

### sql[](https://docs.sqlc.dev/en/latest/reference/config.html#sql "Link to this heading")
Each mapping in the `sql` collection has the following keys:
  * `name`:
    * An human-friendly identifier for this query set. Optional.
  * `engine`:
    * One of `postgresql`, `mysql` or `sqlite`.
  * `schema`:
    * Directory of SQL migrations or path to single SQL file; or a list of paths.
  * `queries`:
    * Directory of SQL queries or path to single SQL file; or a list of paths.
  * `codegen`:
    * A collection of mappings to configure code generators. See [codegen](https://docs.sqlc.dev/en/latest/reference/config.html#codegen) for the supported keys.
  * `gen`:
    * A mapping to configure built-in code generators. See [gen](https://docs.sqlc.dev/en/latest/reference/config.html#gen) for the supported keys.
  * `database`:
    * A mapping to configure database connections. See [database](https://docs.sqlc.dev/en/latest/reference/config.html#database) for the supported keys.
  * `rules`:
    * A collection of rule names to run via `sqlc vet`. See [rules](https://docs.sqlc.dev/en/latest/reference/config.html#rules) for configuration options.
  * `analyzer`:
    * A mapping to configure query analysis. See [analyzer](https://docs.sqlc.dev/en/latest/reference/config.html#analyzer) for the supported keys.
  * `strict_function_checks`
    * If true, return an error if a called SQL function does not exist. Defaults to `false`.
  * `strict_order_by`
    * If true, return an error if a order by column is ambiguous. Defaults to `true`.


### codegen[](https://docs.sqlc.dev/en/latest/reference/config.html#codegen "Link to this heading")
The `codegen` mapping supports the following keys:
  * `out`:
    * Output directory for generated code.
  * `plugin`:
    * The name of the plugin. Must be defined in the `plugins` collection.
  * `options`:
    * A mapping of plugin-specific options.


```
version:'2'
plugins:
-name:py
wasm:
url:https://github.com/sqlc-dev/sqlc-gen-python/releases/download/v0.16.0-alpha/sqlc-gen-python.wasm
sha256:428476c7408fd4c032da4ec74e8a7344f4fa75e0f98a5a3302f238283b9b95f2
sql:
-schema:"schema.sql"
queries:"query.sql"
engine:postgresql
codegen:
-out:src/authors
plugin:py
options:
package:authors
emit_sync_querier:true
emit_async_querier:true
query_parameter_limit:5

```

### database[](https://docs.sqlc.dev/en/latest/reference/config.html#database "Link to this heading")
The `database` mapping supports the following keys:
  * `managed`:
    * If true, connect to a [managed database](https://docs.sqlc.dev/en/latest/howto/managed-databases.html). Defaults to `false`.
  * `uri`:
    * Database connection URI


The `uri` string can contain references to environment variables using the `${...}` syntax. In the following example, the connection string will have the value of the `PG_PASSWORD` environment variable set as its password.
```
version:'2'
sql:
-schema:schema.sql
queries:query.sql
engine:postgresql
database:
uri:postgresql://postgres:${PG_PASSWORD}@localhost:5432/authors
gen:
go:
package:authors
out:postgresql

```

### analyzer[](https://docs.sqlc.dev/en/latest/reference/config.html#analyzer "Link to this heading")
The `analyzer` mapping supports the following keys:
  * `database`:
    * If false, do not use the configured database for query analysis. Defaults to `true`.


### gen[](https://docs.sqlc.dev/en/latest/reference/config.html#gen "Link to this heading")
The `gen` mapping supports the following keys:
#### go[](https://docs.sqlc.dev/en/latest/reference/config.html#go "Link to this heading")
  * `package`:
    * The package name to use for the generated code. Defaults to `out` basename.
  * `out`:
    * Output directory for generated code.
  * `sql_package`:
    * Either `pgx/v4`, `pgx/v5` or `database/sql`. Defaults to `database/sql`.
  * `sql_driver`:
    * Either `github.com/jackc/pgx/v4`, `github.com/jackc/pgx/v5`, `github.com/lib/pq` or `github.com/go-sql-driver/mysql`. No defaults. Required if query annotation `:copyfrom` is used.
  * `emit_db_tags`:
    * If true, add DB tags to generated structs. Defaults to `false`.
  * `emit_prepared_queries`:
    * If true, include support for prepared queries. Defaults to `false`.
  * `emit_interface`:
    * If true, output a `Querier` interface in the generated package. Defaults to `false`.
  * `emit_exact_table_names`:
    * If true, struct names will mirror table names. Otherwise, sqlc attempts to singularize plural table names. Defaults to `false`.
  * `emit_empty_slices`:
    * If true, slices returned by `:many` queries will be empty instead of `nil`. Defaults to `false`.
  * `emit_exported_queries`:
    * If true, autogenerated SQL statement can be exported to be accessed by another package.
  * `emit_json_tags`:
    * If true, add JSON tags to generated structs. Defaults to `false`.
  * `emit_result_struct_pointers`:
    * If true, query results are returned as pointers to structs. Queries returning multiple results are returned as slices of pointers. Defaults to `false`.
  * `emit_params_struct_pointers`:
    * If true, parameters are passed as pointers to structs. Defaults to `false`.
  * `emit_methods_with_db_argument`:
    * If true, generated methods will accept a DBTX argument instead of storing a DBTX on the `*Queries` struct. Defaults to `false`.
  * `emit_pointers_for_null_types`:
    * If true, generated types for nullable columns are emitted as pointers (ie. `*string`) instead of `database/sql` null types (ie. `NullString`). Currently only supported for PostgreSQL if `sql_package` is `pgx/v4` or `pgx/v5`, and for SQLite. Defaults to `false`.
  * `emit_enum_valid_method`:
    * If true, generate a Valid method on enum types, indicating whether a string is a valid enum value.
  * `emit_all_enum_values`:
    * If true, emit a function per enum type that returns all valid enum values.
  * `emit_sql_as_comment`:
    * If true, emits the SQL statement as a code-block comment above the generated function, appending to any existing comments. Defaults to `false`.
  * `build_tags`:
    * If set, add a `//go:build <build_tags>` directive at the beginning of each generated Go file.
  * `initialisms`:
    * An array of [initialisms](https://google.github.io/styleguide/go/decisions.html#initialisms) to upper-case. For example, `app_id` becomes `AppID`. Defaults to `["id"]`.
  * `json_tags_id_uppercase`:
    * If true, “Id” in json tags will be uppercase. If false, will be camelcase. Defaults to `false`
  * `json_tags_case_style`:
    * `camel` for camelCase, `pascal` for PascalCase, `snake` for snake_case or `none` to use the column name in the DB. Defaults to `none`.
  * `omit_unused_structs`:
    * If `true`, sqlc won’t generate table and enum structs that aren’t used in queries for a given package. Defaults to `false`.
  * `output_batch_file_name`:
    * Customize the name of the batch file. Defaults to `batch.go`.
  * `output_db_file_name`:
    * Customize the name of the db file. Defaults to `db.go`.
  * `output_models_file_name`:
    * Customize the name of the models file. Defaults to `models.go`.
  * `output_querier_file_name`:
    * Customize the name of the querier file. Defaults to `querier.go`.
  * `output_copyfrom_file_name`:
    * Customize the name of the copyfrom file. Defaults to `copyfrom.go`.
  * `output_files_suffix`:
    * If specified the suffix will be added to the name of the generated files.
  * `query_parameter_limit`:
    * The number of positional arguments that will be generated for Go functions. To always emit a parameter struct, set this to `0`. Defaults to `1`.
  * `rename`:
    * Customize the name of generated struct fields. See [Renaming fields](https://docs.sqlc.dev/en/latest/howto/rename.html) for usage information.
  * `overrides`:
    * A collection of configurations to override sqlc’s default Go type choices. See [Overriding types](https://docs.sqlc.dev/en/latest/howto/overrides.html) for usage information.


##### overrides[](https://docs.sqlc.dev/en/latest/reference/config.html#overrides "Link to this heading")
See [Overriding types](https://docs.sqlc.dev/en/latest/howto/overrides.html) for an in-depth guide to using type overrides.
#### kotlin[](https://docs.sqlc.dev/en/latest/reference/config.html#kotlin "Link to this heading")
> Removed in v1.17.0 and replaced by the [sqlc-gen-kotlin](https://github.com/sqlc-dev/sqlc-gen-kotlin) plugin. Follow the [migration guide](https://docs.sqlc.dev/en/latest/guides/migrating-to-sqlc-gen-kotlin.html) to switch.
  * `package`:
    * The package name to use for the generated code.
  * `out`:
    * Output directory for generated code.
  * `emit_exact_table_names`:
    * If true, use the exact table name for generated models. Otherwise, guess a singular form. Defaults to `false`.


#### python[](https://docs.sqlc.dev/en/latest/reference/config.html#python "Link to this heading")
> Removed in v1.17.0 and replaced by the [sqlc-gen-python](https://github.com/sqlc-dev/sqlc-gen-python) plugin. Follow the [migration guide](https://docs.sqlc.dev/en/latest/guides/migrating-to-sqlc-gen-python.html) to switch.
  * `package`:
    * The package name to use for the generated code.
  * `out`:
    * Output directory for generated code.
  * `emit_exact_table_names`:
    * If true, use the exact table name for generated models. Otherwise, guess a singular form. Defaults to `false`.
  * `emit_sync_querier`:
    * If true, generate a class with synchronous methods. Defaults to `false`.
  * `emit_async_querier`:
    * If true, generate a class with asynchronous methods. Defaults to `false`.
  * `emit_pydantic_models`:
    * If true, generate classes that inherit from `pydantic.BaseModel`. Otherwise, define classes using the `dataclass` decorator. Defaults to `false`.


#### json[](https://docs.sqlc.dev/en/latest/reference/config.html#json "Link to this heading")
  * `out`:
    * Output directory for the generated JSON.
  * `filename`:
    * Filename for the generated JSON document. Defaults to `codegen_request.json`.
  * `indent`:
    * Indent string to use in the JSON document. Defaults to .


### plugins[](https://docs.sqlc.dev/en/latest/reference/config.html#plugins "Link to this heading")
Each mapping in the `plugins` collection has the following keys:
  * `name`:
    * The name of this plugin. Required
  * `env`
    * A list of environment variables to pass to the plugin. By default, no environment variables are passed.
  * `process`: A mapping with a single `cmd` key
    * `cmd`:
      * The executable to call when using this plugin
    * `format`:
      * The format expected. Supports `json` and `protobuf` formats. Defaults to `protobuf`.
  * `wasm`: A mapping with a two keys `url` and `sha256`
    * `url`:
      * The URL to fetch the WASM file. Supports the `https://` or `file://` schemes.
    * `sha256`
      * The SHA256 checksum for the downloaded file.


```
version:"2"
plugins:
-name:"py"
wasm:
url:"https://github.com/sqlc-dev/sqlc-gen-python/releases/download/v0.16.0-alpha/sqlc-gen-python.wasm"
sha256:"428476c7408fd4c032da4ec74e8a7344f4fa75e0f98a5a3302f238283b9b95f2"
-name:"js"
env:
-PATH
process:
cmd:"sqlc-gen-json"

```

### rules[](https://docs.sqlc.dev/en/latest/reference/config.html#rules "Link to this heading")
Each mapping in the `rules` collection has the following keys:
  * `name`:
    * The name of this rule. Required
  * `rule`:
    * A [Common Expression Language (CEL)](https://github.com/google/cel-spec) expression. Required.
  * `message`:
    * An optional message shown when this rule evaluates to `true`.


See the [vet](https://docs.sqlc.dev/en/latest/howto/vet.html) documentation for a list of built-in rules and help writing custom rules.
```
version:"2"
sql:
-schema:"query.sql"
queries:"query.sql"
engine:"postgresql"
gen:
go:
package:"authors"
out:"db"
rules:
-no-pg
-no-delete
-only-one-param
-no-exec
rules:
-name:no-pg
message:"invalidengine:postgresql"
rule:|
config.engine == "postgresql"
-name:no-delete
message:"don'tusedeletestatements"
rule:|
query.sql.contains("DELETE")
-name:only-one-param
message:"toomanyparameters"
rule:|
query.params.size() > 1
-name:no-exec
message:"don'tuseexec"
rule:|
query.cmd == "exec"

```

### Global overrides[](https://docs.sqlc.dev/en/latest/reference/config.html#global-overrides "Link to this heading")
Sometimes, the same configuration must be done across various specifications of code generation. Then a global definition for type overriding and field renaming can be done using the `overrides` mapping the following manner:
```
version:"2"
overrides:
go:
rename:
id:"Identifier"
overrides:
-db_type:"pg_catalog.timestamptz"
nullable:true
engine:"postgresql"
go_type:
import:"gopkg.in/guregu/null.v4"
package:"null"
type:"Time"
sql:
-schema:"postgresql/schema.sql"
queries:"postgresql/query.sql"
engine:"postgresql"
gen:
go:
package:"authors"
out:"postgresql"
-schema:"mysql/schema.sql"
queries:"mysql/query.sql"
engine:"mysql"
gen:
go:
package:"authors"
out:"mysql"

```

With the previous configuration, whenever a struct field is generated from a table column that is called `id`, it will generated as `Identifier`.
Also, whenever there is a nullable `timestamp with time zone` column in a Postgres table, it will be generated as `null.Time`. Note that the mapping for global type overrides has a field called `engine` that is absent in the regular type overrides. This field is only used when there are multiple definitions using multiple engines. Otherwise, the value of the `engine` key defaults to the engine that is currently being used.
Currently, type overrides and field renaming, both global and regular, are only fully supported in Go.
## Version 1[](https://docs.sqlc.dev/en/latest/reference/config.html#version-1 "Link to this heading")
```
version:"1"
packages:
-name:"db"
path:"internal/db"
queries:"./sql/query/"
schema:"./sql/schema/"
engine:"postgresql"
emit_db_tags:false
emit_prepared_queries:true
emit_interface:false
emit_exact_table_names:false
emit_empty_slices:false
emit_exported_queries:false
emit_json_tags:true
emit_result_struct_pointers:false
emit_params_struct_pointers:false
emit_methods_with_db_argument:false
emit_pointers_for_null_types:false
emit_enum_valid_method:false
emit_all_enum_values:false
build_tags:"some_tag"
json_tags_case_style:"camel"
omit_unused_structs:false
output_batch_file_name:"batch.go"
output_db_file_name:"db.go"
output_models_file_name:"models.go"
output_querier_file_name:"querier.go"
output_copyfrom_file_name:"copyfrom.go"
query_parameter_limit:1

```

### packages[](https://docs.sqlc.dev/en/latest/reference/config.html#packages "Link to this heading")
Each mapping in the `packages` collection has the following keys:
  * `name`:
    * The package name to use for the generated code. Defaults to `path` basename.
  * `path`:
    * Output directory for generated code.
  * `queries`:
    * Directory of SQL queries or path to single SQL file; or a list of paths.
  * `schema`:
    * Directory of SQL migrations or path to single SQL file; or a list of paths.
  * `engine`:
    * Either `postgresql` or `mysql`. Defaults to `postgresql`.
  * `sql_package`:
    * Either `pgx/v4`, `pgx/v5` or `database/sql`. Defaults to `database/sql`.
  * `overrides`:
    * A list of type override configurations. See the [Overriding types](https://docs.sqlc.dev/en/latest/howto/overrides.html) guide for details.
  * `emit_db_tags`:
    * If true, add DB tags to generated structs. Defaults to `false`.
  * `emit_prepared_queries`:
    * If true, include support for prepared queries. Defaults to `false`.
  * `emit_interface`:
    * If true, output a `Querier` interface in the generated package. Defaults to `false`.
  * `emit_exact_table_names`:
    * If true, struct names will mirror table names. Otherwise, sqlc attempts to singularize plural table names. Defaults to `false`.
  * `emit_empty_slices`:
    * If true, slices returned by `:many` queries will be empty instead of `nil`. Defaults to `false`.
  * `emit_exported_queries`:
    * If true, autogenerated SQL statement can be exported to be accessed by another package.
  * `emit_json_tags`:
    * If true, add JSON tags to generated structs. Defaults to `false`.
  * `emit_result_struct_pointers`:
    * If true, query results are returned as pointers to structs. Queries returning multiple results are returned as slices of pointers. Defaults to `false`.
  * `emit_params_struct_pointers`:
    * If true, parameters are passed as pointers to structs. Defaults to `false`.
  * `emit_methods_with_db_argument`:
    * If true, generated methods will accept a DBTX argument instead of storing a DBTX on the `*Queries` struct. Defaults to `false`.
  * `emit_pointers_for_null_types`:
    * If true and `sql_package` is set to `pgx/v4` or `pgx/v5`, generated types for nullable columns are emitted as pointers (ie. `*string`) instead of `database/sql` null types (ie. `NullString`). Defaults to `false`.
  * `emit_enum_valid_method`:
    * If true, generate a Valid method on enum types, indicating whether a string is a valid enum value.
  * `emit_all_enum_values`:
    * If true, emit a function per enum type that returns all valid enum values.
  * `build_tags`:
    * If set, add a `//go:build <build_tags>` directive at the beginning of each generated Go file.
  * `json_tags_case_style`:
    * `camel` for camelCase, `pascal` for PascalCase, `snake` for snake_case or `none` to use the column name in the DB. Defaults to `none`.
  * `omit_unused_structs`:
    * If `true`, sqlc won’t generate table and enum structs that aren’t used in queries for a given package. Defaults to `false`.
  * `output_batch_file_name`:
    * Customize the name of the batch file. Defaults to `batch.go`.
  * `output_db_file_name`:
    * Customize the name of the db file. Defaults to `db.go`.
  * `output_models_file_name`:
    * Customize the name of the models file. Defaults to `models.go`.
  * `output_querier_file_name`:
    * Customize the name of the querier file. Defaults to `querier.go`.
  * `output_copyfrom_file_name`:
    * Customize the name of the copyfrom file. Defaults to `copyfrom.go`.
  * `output_files_suffix`:
    * If specified the suffix will be added to the name of the generated files.
  * `query_parameter_limit`:
    * Positional arguments that will be generated in Go functions (`>= 0`). To always emit a parameter struct, you would need to set it to `0`. Defaults to `1`.


### overrides[](https://docs.sqlc.dev/en/latest/reference/config.html#id1 "Link to this heading")
See the version 1 configuration section of the [Overriding types](https://docs.sqlc.dev/en/latest/howto/overrides.html#version-1-configuration) guide for details.
### rename[](https://docs.sqlc.dev/en/latest/reference/config.html#rename "Link to this heading")
Struct field names are generated from column names using a simple algorithm: split the column name on underscores and capitalize the first letter of each part.
```
account   -> Account
spotify_url -> SpotifyUrl
app_id   -> AppID

```

If you’re not happy with a field’s generated name, use the `rename` mapping to pick a new name. The keys are column names and the values are the struct field name to use.
```
version:"1"
packages:[...]
rename:
spotify_url:"SpotifyURL"

```

[ Previous](https://docs.sqlc.dev/en/latest/reference/cli.html "CLI") [Next ](https://docs.sqlc.dev/en/latest/reference/datatypes.html "Datatypes")
© Copyright 2024, Riza, Inc..
Built with [Sphinx](https://www.sphinx-doc.org/) using a [theme](https://github.com/readthedocs/sphinx_rtd_theme) provided by [Read the Docs](https://readthedocs.org). 


## Source Information
- URL: https://docs.sqlc.dev/en/latest/reference/config.html
- Harvested: 2025-08-19T15:18:43.026349
- Profile: deep_research
- Agent: architect
