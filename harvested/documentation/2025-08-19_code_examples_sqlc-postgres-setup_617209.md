---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:20:23.717294'
profile: code_examples
source: https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html
topic: sqlc-postgres-setup
---

# sqlc-postgres-setup

[ ![Logo](https://docs.sqlc.dev/en/stable/_static/logo.png) ](https://docs.sqlc.dev/en/stable/index.html)
latest  stable  v1.29.0  v1.28.0  v1.27.0  v1.26.0  v1.25.0  v1.24.0  v1.23.0  v1.22.0  v1.21.0  v1.20.0  v1.19.1  v1.19.0  v1.18.0  v1.17.2  v1.17.1  v1.17.0  v1.16.0  v1.15.0  v1.14.0  v1.13.0  v1.12.0  v1.11.0  v1.10.0  v1.9.0  v1.8.0  v1.7.0 
Overview
  * [Installing sqlc](https://docs.sqlc.dev/en/stable/overview/install.html)


Tutorials
  * [Getting started with MySQL](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html)
  * [Getting started with PostgreSQL](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html)
    * [Setting up](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#setting-up)
    * [Schema and queries](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#schema-and-queries)
    * [Generating code](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#generating-code)
    * [Using generated code](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#using-generated-code)
    * [Query verification](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#query-verification)
  * [Getting started with SQLite](https://docs.sqlc.dev/en/stable/tutorials/getting-started-sqlite.html)


Commands
  * [`generate` - Generating code](https://docs.sqlc.dev/en/stable/howto/generate.html)
  * [`push` - Uploading projects](https://docs.sqlc.dev/en/stable/howto/push.html)
  * [`verify` - Verifying schema changes](https://docs.sqlc.dev/en/stable/howto/verify.html)
  * [`vet` - Linting queries](https://docs.sqlc.dev/en/stable/howto/vet.html)


How-to Guides
  * [Retrieving rows](https://docs.sqlc.dev/en/stable/howto/select.html)
  * [Counting rows](https://docs.sqlc.dev/en/stable/howto/query_count.html)
  * [Inserting rows](https://docs.sqlc.dev/en/stable/howto/insert.html)
  * [Updating rows](https://docs.sqlc.dev/en/stable/howto/update.html)
  * [Deleting rows](https://docs.sqlc.dev/en/stable/howto/delete.html)
  * [Preparing queries](https://docs.sqlc.dev/en/stable/howto/prepared_query.html)
  * [Using transactions](https://docs.sqlc.dev/en/stable/howto/transactions.html)
  * [Naming parameters](https://docs.sqlc.dev/en/stable/howto/named_parameters.html)
  * [Modifying the database schema](https://docs.sqlc.dev/en/stable/howto/ddl.html)
  * [Configuring generated structs](https://docs.sqlc.dev/en/stable/howto/structs.html)
  * [Embedding structs](https://docs.sqlc.dev/en/stable/howto/embedding.html)
  * [Overriding types](https://docs.sqlc.dev/en/stable/howto/overrides.html)
  * [Renaming fields](https://docs.sqlc.dev/en/stable/howto/rename.html)


sqlc Cloud
  * [Managed databases](https://docs.sqlc.dev/en/stable/howto/managed-databases.html)


Reference
  * [Changelog](https://docs.sqlc.dev/en/stable/reference/changelog.html)
  * [CLI](https://docs.sqlc.dev/en/stable/reference/cli.html)
  * [Configuration](https://docs.sqlc.dev/en/stable/reference/config.html)
  * [Datatypes](https://docs.sqlc.dev/en/stable/reference/datatypes.html)
  * [Environment variables](https://docs.sqlc.dev/en/stable/reference/environment-variables.html)
  * [Database and language support](https://docs.sqlc.dev/en/stable/reference/language-support.html)
  * [Macros](https://docs.sqlc.dev/en/stable/reference/macros.html)
  * [Query annotations](https://docs.sqlc.dev/en/stable/reference/query-annotations.html)


Conceptual Guides
  * [Using sqlc in CI/CD](https://docs.sqlc.dev/en/stable/howto/ci-cd.html)
  * [Using Go and pgx](https://docs.sqlc.dev/en/stable/guides/using-go-and-pgx.html)
  * [Using plugins](https://docs.sqlc.dev/en/stable/guides/plugins.html)
  * [Developing sqlc](https://docs.sqlc.dev/en/stable/guides/development.html)
  * [Privacy and data collection](https://docs.sqlc.dev/en/stable/guides/privacy.html)


Sponsored By
[ ![Riza logo](https://sqlc.dev/sponsors/riza-readme.png) ](https://riza.io?utm_source=sqlc+docs)
[sqlc](https://docs.sqlc.dev/en/stable/index.html)
  * [](https://docs.sqlc.dev/en/stable/index.html)
  * Getting started with PostgreSQL
  * [ View page source](https://docs.sqlc.dev/en/stable/_sources/tutorials/getting-started-postgresql.md.txt)


# Getting started with PostgreSQL[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#getting-started-with-postgresql "Link to this heading")
This tutorial assumes that the latest version of sqlc is [installed](https://docs.sqlc.dev/en/stable/overview/install.html) and ready to use.
We’ll generate Go code here, but other [language plugins](https://docs.sqlc.dev/en/stable/reference/language-support.html) are available. You’ll naturally need the Go toolchain if you want to build and run a program with the code sqlc generates, but sqlc itself has no dependencies.
At the end, you’ll push your SQL queries to [sqlc Cloud](https://dashboard.sqlc.dev/) for further insights and analysis.
## Setting up[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#setting-up "Link to this heading")
Create a new directory called `sqlc-tutorial` and open it up.
Initialize a new Go module named `tutorial.sqlc.dev/app`:
```
gomodinittutorial.sqlc.dev/app

```

sqlc looks for either a `sqlc.(yaml|yml)` or `sqlc.json` file in the current directory. In our new directory, create a file named `sqlc.yaml` with the following contents:
```
version:"2"
sql:
-engine:"postgresql"
queries:"query.sql"
schema:"schema.sql"
gen:
go:
package:"tutorial"
out:"tutorial"
sql_package:"pgx/v5"

```

## Schema and queries[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#schema-and-queries "Link to this heading")
sqlc needs to know your database schema and queries in order to generate code. In the same directory, create a file named `schema.sql` with the following content:
```
CREATETABLEauthors(
idBIGSERIALPRIMARYKEY,
nametextNOTNULL,
biotext
);

```

Next, create a `query.sql` file with the following five queries:
```
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
-- name: UpdateAuthor :exec
UPDATEauthors
setname=$2,
bio=$3
WHEREid=$1;
-- name: DeleteAuthor :exec
DELETEFROMauthors
WHEREid=$1;

```

If you prefer, you can alter the `UpdateAuthor` query to return the updated record:
```
-- name: UpdateAuthor :one
UPDATEauthors
setname=$2,
bio=$3
WHEREid=$1
RETURNING*;

```

## Generating code[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#generating-code "Link to this heading")
You are now ready to generate code. You shouldn’t see any output when you run the `generate` subcommand, unless something goes wrong:
```
sqlcgenerate

```

You should now have a `tutorial` subdirectory with three files containing Go source code. These files comprise a Go package named `tutorial`:
```
├── go.mod
├── query.sql
├── schema.sql
├── sqlc.yaml
└── tutorial
  ├── db.go
  ├── models.go
  └── query.sql.go

```

## Using generated code[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#using-generated-code "Link to this heading")
You can use your newly-generated `tutorial` package from any Go program. Create a file named `tutorial.go` and add the following contents:
```
packagemain
import(
"context"
"log"
"reflect"
"github.com/jackc/pgx/v5"
"github.com/jackc/pgx/v5/pgtype"
"tutorial.sqlc.dev/app/tutorial"
)
funcrun()error{
ctx:=context.Background()
conn,err:=pgx.Connect(ctx,"user=pqgotest dbname=pqgotest sslmode=verify-full")
iferr!=nil{
returnerr
}
deferconn.Close(ctx)
queries:=tutorial.New(conn)
// list all authors
authors,err:=queries.ListAuthors(ctx)
iferr!=nil{
returnerr
}
log.Println(authors)
// create an author
insertedAuthor,err:=queries.CreateAuthor(ctx,tutorial.CreateAuthorParams{
Name:"Brian Kernighan",
Bio:pgtype.Text{String:"Co-author of The C Programming Language and The Go Programming Language",Valid:true},
})
iferr!=nil{
returnerr
}
log.Println(insertedAuthor)
// get the author we just inserted
fetchedAuthor,err:=queries.GetAuthor(ctx,insertedAuthor.ID)
iferr!=nil{
returnerr
}
// prints true
log.Println(reflect.DeepEqual(insertedAuthor,fetchedAuthor))
returnnil
}
funcmain(){
iferr:=run();err!=nil{
log.Fatal(err)
}
}

```

Before this code will compile you’ll need to fetch the relevant PostgreSQL driver. You can use `lib/pq` with the standard library’s `database/sql` package, but in this tutorial we’ve used `pgx/v5`:
```
gogetgithub.com/jackc/pgx/v5
gobuild./...

```

The program should compile without errors. To make that possible, sqlc generates readable, **idiomatic** Go code that you otherwise would’ve had to write yourself. Take a look in `tutorial/query.sql.go`.
Of course for this program to run successfully you’ll need to compile after replacing the database connection parameters in the call to `pgx.Connect()` with the correct parameters for your database. And your database must have the `authors` table as defined in `schema.sql`.
You should now have a working program using sqlc’s generated Go source code, and hopefully can see how you’d use sqlc in your own real-world applications.
## Query verification[](https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html#query-verification "Link to this heading")
[sqlc Cloud](https://dashboard.sqlc.dev) provides additional verification, catching subtle bugs. To get started, create a [dashboard account](https://dashboard.sqlc.dev). Once you’ve signed in, create a project and generate an auth token. Add your project’s ID to the `cloud` block to your sqlc.yaml.
```
version:"2"
cloud:
# Replace <PROJECT_ID> with your project ID from the sqlc Cloud dashboard
project:"<PROJECT_ID>"
sql:
-engine:"postgresql"
queries:"query.sql"
schema:"schema.sql"
gen:
go:
package:"tutorial"
out:"tutorial"
sql_package:"pgx/v5"

```

Replace `<PROJECT_ID>` with your project ID from the sqlc Cloud dashboard. It will look something like `01HA8SZH31HKYE9RR3N3N3TSJM`.
And finally, set the `SQLC_AUTH_TOKEN` environment variable:
```
exportSQLC_AUTH_TOKEN="<your sqlc auth token>"

```

```
$sqlcpush--tagtutorial

```

In the sidebar, go to the “Queries” section to see your published queries. Run `verify` to ensure that previously published queries continue to work against updated database schema.
```
$sqlcverify--againsttutorial

```

[ Previous](https://docs.sqlc.dev/en/stable/tutorials/getting-started-mysql.html "Getting started with MySQL") [Next ](https://docs.sqlc.dev/en/stable/tutorials/getting-started-sqlite.html "Getting started with SQLite")
© Copyright 2024, Riza, Inc..
Built with [Sphinx](https://www.sphinx-doc.org/) using a [theme](https://github.com/readthedocs/sphinx_rtd_theme) provided by [Read the Docs](https://readthedocs.org). 


## Source Information
- URL: https://docs.sqlc.dev/en/stable/tutorials/getting-started-postgresql.html
- Harvested: 2025-08-19T01:20:23.717294
- Profile: code_examples
- Agent: unknown
