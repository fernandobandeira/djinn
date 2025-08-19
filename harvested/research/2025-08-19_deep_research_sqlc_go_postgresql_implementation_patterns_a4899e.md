---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:17:33.832437'
profile: deep_research
source: https://docs.sqlc.dev
topic: sqlc Go PostgreSQL Implementation Patterns
---

# sqlc Go PostgreSQL Implementation Patterns

[ ![Logo](https://docs.sqlc.dev/en/latest/_static/logo.png) ](https://docs.sqlc.dev/en/latest/)
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
[sqlc](https://docs.sqlc.dev/en/latest/)
  * [](https://docs.sqlc.dev/en/latest/)
  * sqlc Documentation
  * [ View page source](https://docs.sqlc.dev/en/latest/_sources/index.rst.txt)


# sqlc Documentation[](https://docs.sqlc.dev/en/latest/#sqlc-documentation "Link to this heading")
> 

And lo, the Great One looked down upon the people and proclaimed:
    
> “SQL is actually pretty great”
sqlc generates **fully type-safe idiomatic Go code** from SQL. Here’s how it works:
  1. You write SQL queries
  2. You run sqlc to generate Go code that presents type-safe interfaces to those queries
  3. You write application code that calls the methods sqlc generated


Seriously, it’s that easy. You don’t have to write any boilerplate SQL querying code ever again.
[Next ](https://docs.sqlc.dev/en/latest/overview/install.html "Installing sqlc")
© Copyright 2024, Riza, Inc..
Built with [Sphinx](https://www.sphinx-doc.org/) using a [theme](https://github.com/readthedocs/sphinx_rtd_theme) provided by [Read the Docs](https://readthedocs.org). 


## Source Information
- URL: https://docs.sqlc.dev
- Harvested: 2025-08-19T15:17:33.832437
- Profile: deep_research
- Agent: architect
