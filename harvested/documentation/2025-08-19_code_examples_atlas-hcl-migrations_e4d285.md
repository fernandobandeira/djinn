---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:20:33.588528'
profile: code_examples
source: https://atlasgo.io/getting-started
topic: atlas-hcl-migrations
---

# atlas-hcl-migrations

[Skip to main content](https://atlasgo.io/getting-started#__docusaurus_skipToContent_fallback)
**New:** Snowflake, Spanner, and PG Partitions [Read now](https://atlasgo.io/blog/2025/07/21/v036-snowflake-postgres-partitions-and-azure-devops)
[![Atlas](https://atlasgo.io/u/logo-atlas.svg)![Atlas](https://atlasgo.io/u/logo-atlas.svg)](https://atlasgo.io/)[Docs](https://atlasgo.io/docs)
[Guides](https://atlasgo.io/getting-started)
[GuidesStep-by-step guides to various Atlas use cases](https://atlasgo.io/guides)[DatabasesDatabase-specific guides for using Atlas](https://atlasgo.io/databases)[ORMs and FrameworksIntegrate Atlas with the ORM of your choice](https://atlasgo.io/orms)

[Use Cases](https://atlasgo.io/getting-started)
[Modernize Schema MigrationsApply modern CI/CD to schema changes](https://atlasgo.io/use-cases/modernize-database-migrations)[Standardize Schema MigrationsOne migration tool to rule them all](https://atlasgo.io/use-cases/standardize-schema-migrations)[Database per tenantManage thousands of databases as one](https://atlasgo.io/use-cases/database-per-tenant)[Database GovernanceComing soonEnd-to-end control and compliance](https://atlasgo.io/getting-started)

[Blog](https://atlasgo.io/blog)[Case Studies](https://atlasgo.io/case-studies)[Pricing](https://atlasgo.io/pricing)
Ask AI
[](https://github.com/ariga/atlas)[](https://discord.gg/zZ6sWVg6NT)[](https://twitter.com/atlasgo_io)[](https://atlasnewsletter.substack.com/)
Search
  * [Home](https://atlasgo.io/docs)
  * [Getting Started](https://atlasgo.io/getting-started)
  * [Atlas vs Others](https://atlasgo.io/atlas-vs-others)
  * [Declarative Workflows](https://atlasgo.io/declarative/inspect)
    * [Schema Inspection](https://atlasgo.io/declarative/inspect)
    * [Applying Changes](https://atlasgo.io/declarative/apply)
    * [Pre-planning Changes](https://atlasgo.io/declarative/plan)
    * [Comparing Schemas](https://atlasgo.io/declarative/diff)
  * [Versioned Workflows](https://atlasgo.io/versioned/intro)
    * [Quick Introduction](https://atlasgo.io/versioned/intro)
    * [Generating Migrations](https://atlasgo.io/versioned/diff)
    * [Migration Linting](https://atlasgo.io/versioned/lint)
    * [Applying Migrations](https://atlasgo.io/versioned/apply)
    * [Down Migrations](https://atlasgo.io/versioned/down)
    * [Pre-migration Checks](https://atlasgo.io/versioned/checks)
    * [Manual Migrations](https://atlasgo.io/versioned/new)
    * [Migration Troubleshooting](https://atlasgo.io/versioned/troubleshoot)
    * [Import Existing Migrations](https://atlasgo.io/versioned/import)
    * [Checkpoints](https://atlasgo.io/versioned/checkpoint)
  * [Schema as Code](https://atlasgo.io/atlas-schema)
    * [HCL Syntax](https://atlasgo.io/atlas-schema/hcl)
    * [SQL Syntax](https://atlasgo.io/atlas-schema/sql)
    * [ORMs and Frameworks](https://atlasgo.io/orms)
    * [External Integrations](https://atlasgo.io/atlas-schema/external)
    * [Atlas Configuration](https://atlasgo.io/atlas-schema/projects)
  * [Testing Framework](https://atlasgo.io/testing/schema)
  * [Concepts](https://atlasgo.io/concepts/url)
  * [Cloud](https://atlasgo.io/cloud/getting-started)
  * [Integrations](https://atlasgo.io/integrations)
  * [HCL Docs](https://atlasgo.io/hcl/docs)
  * [CLI Reference](https://atlasgo.io/cli-reference)
  * [Contributing](https://atlasgo.io/contributing)
  * [Getting Support](https://atlasgo.io/support)
  * [Trust Center](https://atlasgo.io/trust)


  * [](https://atlasgo.io/)
  * Getting Started


On this page
# Quick Introduction
Atlas is a language-independent tool for managing and migrating database schemas using modern DevOps principles. It offers two workflows:
  * **Declarative** : Similar to Terraform, Atlas compares the current state of the database to the desired state, as defined in an [HCL](https://atlasgo.io/atlas-schema/hcl), [SQL](https://atlasgo.io/atlas-schema/sql), or [ORM](https://atlasgo.io/atlas-schema/external) schema. Based on this comparison, it generates and executes a migration plan to transition the database to its desired state.
  * **Versioned** : Unlike other tools, Atlas automatically plans schema migrations for you. Users can describe their desired database schema in [HCL](https://atlasgo.io/atlas-schema/hcl), [SQL](https://atlasgo.io/atlas-schema/sql), or their chosen [ORM](https://atlasgo.io/atlas-schema/external), and by utilizing Atlas, they can plan, lint, and apply the necessary migrations to the database.


### Installation[​](https://atlasgo.io/getting-started#installation "Direct link to Installation")
  * macOS + Linux
  * Homebrew
  * Docker
  * Windows
  * Manual Installation


To download and install the latest release of the Atlas CLI, simply run the following in your terminal:
```
curl-sSf https://atlasgo.sh |sh
```

Get the latest release with [Homebrew](https://brew.sh/):
```
brew install ariga/tap/atlas
```

To pull the Atlas image and run it as a Docker container:
```
docker pull arigaio/atlasdocker run --rm arigaio/atlas --help
```

If the container needs access to the host network or a local directory, use the `--net=host` flag and mount the desired directory:
```
docker run --rm--net=host \-v$(pwd)/migrations:/migrations \ arigaio/atlas migrate apply--url"mysql://root:pass@:3306/test"
```

Download the [latest release](https://release.ariga.io/atlas/atlas-windows-amd64-latest.exe) and move the atlas binary to a file location on your system PATH.
If you want to manually install the Atlas CLI, pick one of the below builds suitable for your system.
  * MacOS 
    * [amd64](https://release.ariga.io/atlas/atlas-darwin-amd64-latest) ([md5](https://release.ariga.io/atlas/atlas-darwin-amd64-latest.md5) / [sha256](https://release.ariga.io/atlas/atlas-darwin-amd64-latest.sha256))
    * [arm64](https://release.ariga.io/atlas/atlas-darwin-arm64-latest) ([md5](https://release.ariga.io/atlas/atlas-darwin-arm64-latest.md5) / [sha256](https://release.ariga.io/atlas/atlas-darwin-arm64-latest.sha256))
  * Windows 
    * [amd64](https://release.ariga.io/atlas/atlas-windows-amd64-latest.exe) ([md5](https://release.ariga.io/atlas/atlas-windows-amd64-latest.exe.md5) / [sha256](https://release.ariga.io/atlas/atlas-windows-amd64-latest.exe.sha256))
  * Linux 
    * [amd64](https://release.ariga.io/atlas/atlas-linux-amd64-latest) ([md5](https://release.ariga.io/atlas/atlas-linux-amd64-latest.md5) / [sha256](https://release.ariga.io/atlas/atlas-linux-amd64-latest.sha256))
    * [arm64](https://release.ariga.io/atlas/atlas-linux-arm64-latest) ([md5](https://release.ariga.io/atlas/atlas-linux-arm64-latest.md5) / [sha256](https://release.ariga.io/atlas/atlas-linux-arm64-latest.sha256))


The default binaries distributed in official releases are released under the [Atlas EULA](https://ariga.io/legal/atlas/eula). If you would like obtain a copy of Atlas Community Edition (under an Apache 2 license) follow the instructions [here](https://atlasgo.io/community-edition).
### Start a local database container[​](https://atlasgo.io/getting-started#start-a-local-database-container "Direct link to Start a local database container")
For the purpose of this guide, we will start a local Docker container running MySQL.
```
docker run --rm-d--name atlas-demo -p3306:3306 -eMYSQL_ROOT_PASSWORD=pass -eMYSQL_DATABASE=example mysql
```

For this example, we will start with a schema that represents a `users` table, in which each user has an ID and a name:
```
CREATEtable users ( id intPRIMARYKEY, name varchar(100));
```

To create the table above on our local database, we can run the following command:
```
dockerexec atlas-demo mysql -ppass-e'CREATE table example.users(id int PRIMARY KEY, name varchar(100))'
```

### Inspecting our database[​](https://atlasgo.io/getting-started#inspecting-our-database "Direct link to Inspecting our database")
The `atlas schema inspect` command supports reading the database description provided by a URL and outputting it in three different formats: [Atlas DDL](https://atlasgo.io/atlas-schema/hcl) (default), SQL, and JSON. In this guide, we will demonstrate the flow using both the Atlas DDL and SQL formats, as the JSON format is often used for processing the output using `jq`.
  * Atlas DDL (HCL)
  * SQL


To inspect our locally-running MySQL instance, use the `-u` flag and write the output to a file named `schema.hcl`:
```
atlas schema inspect -u"mysql://root:pass@localhost:3306/example"> schema.hcl
```

Open the `schema.hcl` file to view the Atlas schema that describes our database.
schema.hcl
```
table "users"{schema= schema.example column "id"{null=falsetype= int} column "name"{null=truetype= varchar(100)}primary_key{columns=[column.id]}}
```

This block represents a [table](https://atlasgo.io/atlas-schema/hcl#table) resource with `id`, and `name` columns. The `schema` field references the `example` schema that is defined elsewhere in this document. In addition, the `primary_key` sub-block defines the `id` column as the primary key for the table. Atlas strives to mimic the syntax of the database that the user is working against. In this case, the type for the `id` column is `int`, and `varchar(100)` for the `name` column.
To inspect our locally-running MySQL instance, use the `-u` flag and write the output to a file named `schema.sql`:
```
atlas schema inspect -u"mysql://root:pass@localhost:3306/example"--format'{{ sql . }}'> schema.sql
```

Open the `schema.sql` file to view the inspected SQL schema that describes our database.
schema.sql
```
-- create "users" tableCREATETABLE`users`(`id`intNOTNULL,`name`varchar(100)NULL,PRIMARYKEY(`id`))CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
```

Now, consider we want to add a `blog_posts` table and have our schema represent a simplified blogging system.
[![Blog ERD](https://atlasgo.io/u/cloud/images/explore-page.png)](https://gh.atlasgo.cloud/explore/9717d499)
Let's add the following to our inspected schema, and use Atlas to plan and apply the changes to our database.
  * Atlas DDL (HCL)
  * SQL


Edit the `schema.hcl` file and add the following `table` block:
schema.hcl
```
table "blog_posts"{schema= schema.example column "id"{null=falsetype= int} column "title"{null=truetype= varchar(100)} column "body"{null=truetype= text} column "author_id"{null=truetype= int}primary_key{columns=[column.id]} foreign_key "author_fk"{columns=[column.author_id]ref_columns=[table.users.column.id]}}
```

In addition to the elements we saw in the `users` table, here we can find a [foreign key](https://atlasgo.io/atlas-schema/hcl#foreign-key) block, declaring that the `author_id` column references the `id` column on the `users` table.
Edit the `schema.sql` file and add the following `CREATE TABLE` statement:
schema.sql
```
-- create "blog_posts" tableCREATETABLE`blog_posts`(`id`intNOTNULL,`title`varchar(100)NULL,`body`textNULL,`author_id`intNULL,PRIMARYKEY(`id`),CONSTRAINT`author_fk`FOREIGNKEY(`author_id`)REFERENCES`example`.`users`(`id`));
```

Now, let's apply these changes by running a migration. In Atlas, migrations can be applied in two types of workflows: _declarative_ and _versioned_.
### Declarative Migrations[​](https://atlasgo.io/getting-started#declarative-migrations "Direct link to Declarative Migrations")
The declarative approach requires the user to define the _desired_ end schema, and Atlas provides a safe way to alter the database to get there. Let's see this in action.
Continuing the example, in order to apply the changes to our database we will run the `apply` command:
  * Atlas DDL (HCL)
  * SQL


```
atlas schema apply \-u"mysql://root:pass@localhost:3306/example"\--to file://schema.hcl
```

```
atlas schema apply \-u"mysql://root:pass@localhost:3306/example"\--to file://schema.sql \ --dev-url "docker://mysql/8/example"
```

Atlas presents the plan it created by displaying the SQL statements. For example, for a MySQL database we will see the following:
```
-- Planned Changes:-- Create "blog_posts" tableCREATE TABLE `example`.`blog_posts` (`id` int NOT NULL, `title` varchar(100) NULL, `body` text NULL, `author_id` int NULL, PRIMARY KEY (`id`), INDEX `author_id` (`author_id`), CONSTRAINT `author_fk` FOREIGN KEY (`author_id`) REFERENCES `example`.`users` (`id`))Use the arrow keys to navigate: ↓ ↑ → ←? Are you sure?: ▸ Apply  Abort
```

Apply the changes, and that's it! You have successfully run a declarative migration.
To ensure that the changes have been made to the schema, you can run the `inspect` command again. This time, we use the `--web`/`-w` flag to open the Atlas Web UI and view the schema.
```
atlas schema inspect \-u"mysql://root:pass@localhost:3306/example"\--web
```

note
If you are using an old version of Atlas, you may need to replace the `--web` flag with `--visualize`.
### Versioned Migrations[​](https://atlasgo.io/getting-started#versioned-migrations "Direct link to Versioned Migrations")
Alternatively, the versioned migration workflow, sometimes called "change-based migrations", allows each change to the database schema to be checked-in to source control and reviewed during code-review. Users can still benefit from Atlas intelligently planning migrations for them, however they are not automatically applied.
To start, we will calculate the difference between the _desired_ and _current_ state of the database by running the `atlas migrate diff` command.
To run this command, we need to provide the necessary parameters:
  * `--dir` the URL to the migration directory, by default it is `file://migrations`.
  * `--to` the URL of the desired state. A state can be specified using a database URL, HCL or SQL schema, or another migration directory.
  * `--dev-url` a URL to a [Dev Database](https://atlasgo.io/concepts/dev-database) that will be used to compute the diff.


  * Atlas DDL (HCL)
  * SQL


```
atlas migrate diff create_blog_posts \--dir"file://migrations"\--to"file://schema.hcl"\ --dev-url "docker://mysql/8/example"
```

```
atlas migrate diff create_blog_posts \--dir"file://migrations"\--to"file://schema.sql"\ --dev-url "docker://mysql/8/example"
```

Run `ls migrations`, and you will notice that Atlas has created two files:
  * 20220811074144_create_blog_posts.sql
  * atlas.sum


```
-- create "blog_posts" tableCREATETABLE`example`.`blog_posts`(`id`intNOTNULL,`title`varchar(100)NULL,`body`textNULL,`author_id`intNULL,PRIMARYKEY(`id`),INDEX`author_id`(`author_id`),CONSTRAINT`author_fk`FOREIGNKEY(`author_id`)REFERENCES`example`.`users`(`id`))
```

In addition to the migration directory, Atlas maintains a file name `atlas.sum` which is used to ensure the integrity of the migration directory and force developers to deal with situations where migration order or contents was modified after the fact.
```
h1:t1fEP1rSsGf1gYrYCjsGyEyuM0cnhATlq93B7h8uXxY=20220811074144_create_blog_posts.sql h1:liZcCBbAn/HyBTqBAEVar9fJNKPTb2Eq+rEKZeCFC9M=
```

Now that we have our migration files ready, you can use the `migrate apply` command to apply the changes to the database. To learn more about this process, check out the [Versioned Migrations Quickstart Guide](https://atlasgo.io/versioned/intro)
### Next Steps[​](https://atlasgo.io/getting-started#next-steps "Direct link to Next Steps")
In this short tutorial we learned how to use Atlas to inspect databases, as well as use declarative and versioned migrations. Read more about the use-cases for the two approaches [here](https://atlasgo.io/concepts/declarative-vs-versioned) to help you decide which workflow works best for you.
Need help getting started?
We have a super friendly [#getting-started](https://discord.gg/8mvDUG22) channel on our community chat on Discord.
For web-based, free, and fun (GIFs included) support:
[Join our Discord server](https://discord.gg/zZ6sWVg6NT)
![Newsletter planet icon](https://atlasgo.io/pages/newsletter-planet.svg)
Get the latest Atlas tips and updates in our newsletter.
![Newsletter ufo icon](https://atlasgo.io/pages/newsletter-ufo.svg)
**Tags:**
  * [quickstart](https://atlasgo.io/tags/quickstart)
  * [getting started](https://atlasgo.io/tags/getting-started)
  * [installation](https://atlasgo.io/tags/installation)
  * [overview](https://atlasgo.io/tags/overview)


[PreviousHome](https://atlasgo.io/docs)[NextAtlas vs Others](https://atlasgo.io/atlas-vs-others)
  * [Installation](https://atlasgo.io/getting-started#installation)
  * [Start a local database container](https://atlasgo.io/getting-started#start-a-local-database-container)
  * [Inspecting our database](https://atlasgo.io/getting-started#inspecting-our-database)
  * [Declarative Migrations](https://atlasgo.io/getting-started#declarative-migrations)
  * [Versioned Migrations](https://atlasgo.io/getting-started#versioned-migrations)
  * [Next Steps](https://atlasgo.io/getting-started#next-steps)


Docs
  * [Home](https://atlasgo.io/docs)
  * [Atlas vs Others ](https://atlasgo.io/atlas-vs-others)
  * [CLI Reference](https://atlasgo.io/cli-reference)
  * [Blog](https://atlasgo.io/blog)
  * [Guides](https://atlasgo.io/guides)
  * [FAQs](https://atlasgo.io/faq)
  * [GoDoc](https://pkg.go.dev/ariga.io/atlas)


Integrations
  * [GitHub Actions](https://atlasgo.io/integrations/github-actions)
  * [Kubernetes Operator](https://atlasgo.io/integrations/kubernetes)
  * [Terraform](https://atlasgo.io/integrations/terraform-provider)
  * [Go API](https://atlasgo.io/integrations/go-sdk)


Legal
  * [Privacy Policy](https://ariga.io/legal/privacy)
  * [Terms of Service](https://ariga.io/legal/tos)
  * [Trust Center](https://atlasgo.io/trust)


Community
  * [GitHub](https://github.com/ariga/atlas)
  * [Discord](https://discord.gg/zZ6sWVg6NT)
  * [Twitter / X](https://twitter.com/atlasgo_io)
  * [YouTube](https://youtube.com/@ariga_io)


Copyright © 2025 The Atlas Authors. To manage your cookie settings [click here.](javascript:void\(0\);)


## Source Information
- URL: https://atlasgo.io/getting-started
- Harvested: 2025-08-19T01:20:33.588528
- Profile: code_examples
- Agent: unknown
