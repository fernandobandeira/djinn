---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T15:18:24.759979'
profile: deep_research
source: https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html
topic: sqlc pgx integration guide
---

# sqlc pgx integration guide

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
  * [Datatypes](https://docs.sqlc.dev/en/latest/reference/datatypes.html)
  * [Environment variables](https://docs.sqlc.dev/en/latest/reference/environment-variables.html)
  * [Database and language support](https://docs.sqlc.dev/en/latest/reference/language-support.html)
  * [Macros](https://docs.sqlc.dev/en/latest/reference/macros.html)
  * [Query annotations](https://docs.sqlc.dev/en/latest/reference/query-annotations.html)


Conceptual Guides
  * [Using sqlc in CI/CD](https://docs.sqlc.dev/en/latest/howto/ci-cd.html)
  * [Using Go and pgx](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html)
    * [Getting started](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html#getting-started)
    * [Generated code walkthrough](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html#generated-code-walkthrough)
  * [Using plugins](https://docs.sqlc.dev/en/latest/guides/plugins.html)
  * [Developing sqlc](https://docs.sqlc.dev/en/latest/guides/development.html)
  * [Privacy and data collection](https://docs.sqlc.dev/en/latest/guides/privacy.html)


Sponsored By
[ ![Riza logo](https://sqlc.dev/sponsors/riza-readme.png) ](https://riza.io?utm_source=sqlc+docs)
[sqlc](https://docs.sqlc.dev/en/latest/index.html)
  * [](https://docs.sqlc.dev/en/latest/index.html)
  * Using Go and pgx
  * [ View page source](https://docs.sqlc.dev/en/latest/_sources/guides/using-go-and-pgx.rst.txt)


# Using Go and pgx[](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html#using-go-and-pgx "Link to this heading")
Note
`pgx/v5` is supported starting from v1.18.0.
pgx is a pure Go driver and toolkit for PostgreSQL. It’s become the default PostgreSQL package for many Gophers since lib/pq was put into maintenance mode.
## Getting started[](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html#getting-started "Link to this heading")
To start generating code that uses pgx, set the `sql_package` field in your `sqlc.yaml` configuration file. Valid options are `pgx/v4` or `pgx/v5`
```
version:"2"
sql:
-engine:"postgresql"
queries:"query.sql"
schema:"query.sql"
gen:
go:
package:"db"
sql_package:"pgx/v5"
out:"db"

```

If you don’t have an existing sqlc project on hand, create a directory with the configuration file above and the following `query.sql` file.
```
CREATETABLEauthors(
idBIGSERIALPRIMARYKEY,
nametextNOTNULL,
biotext
);
-- name: GetAuthor :one
SELECT*FROMauthors
WHEREid=$1LIMIT1;
-- name: ListAuthors :many
SELECT*FROMauthors
ORDERBYname;
-- name: CreateAuthor :one
INSERTINTOauthors(
name,bio
)VALUES(
$1,$2
)
RETURNING*;
-- name: DeleteAuthor :exec
DELETEFROMauthors
WHEREid=$1;

```

Generating the code will now give you pgx-compatible database access methods.
```
sqlcgenerate

```

## Generated code walkthrough[](https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html#generated-code-walkthrough "Link to this heading")
The generated code is very similar to the code generated when using `lib/pq`. However, instead of using `database/sql`, the code uses pgx types directly.
```
packagemain
import(
"context"
"fmt"
"os"
"github.com/jackc/pgx/v5"
"example.com/sqlc-tutorial/db"
)
funcmain(){
// urlExample := "postgres://username:password@localhost:5432/database_name"
conn,err:=pgx.Connect(context.Background(),os.Getenv("DATABASE_URL"))
iferr!=nil{
fmt.Fprintf(os.Stderr,"Unable to connect to database: %v\n",err)
os.Exit(1)
}
deferconn.Close(context.Background())
q:=db.New(conn)
author,err:=q.GetAuthor(context.Background(),1)
iferr!=nil{
fmt.Fprintf(os.Stderr,"GetAuthor failed: %v\n",err)
os.Exit(1)
}
fmt.Println(author.Name)
}

```

Note
For production applications, consider using pgxpool for connection pooling:
```
import(
"github.com/jackc/pgx/v5/pgxpool"
"example.com/sqlc-tutorial/db"
)
funcmain(){
pool,err:=pgxpool.New(context.Background(),os.Getenv("DATABASE_URL"))
iferr!=nil{
fmt.Fprintf(os.Stderr,"Unable to create connection pool: %v\n",err)
os.Exit(1)
}
deferpool.Close()
q:=db.New(pool)
// Use q the same way as with single connections
}

```

[ Previous](https://docs.sqlc.dev/en/latest/howto/ci-cd.html "Using sqlc in CI/CD") [Next ](https://docs.sqlc.dev/en/latest/guides/plugins.html "Using plugins")
© Copyright 2024, Riza, Inc..
Built with [Sphinx](https://www.sphinx-doc.org/) using a [theme](https://github.com/readthedocs/sphinx_rtd_theme) provided by [Read the Docs](https://readthedocs.org). 


## Source Information
- URL: https://docs.sqlc.dev/en/latest/guides/using-go-and-pgx.html
- Harvested: 2025-08-19T15:18:24.759979
- Profile: deep_research
- Agent: architect
