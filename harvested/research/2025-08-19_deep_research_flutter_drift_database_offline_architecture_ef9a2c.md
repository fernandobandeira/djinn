---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T11:36:49.796939'
profile: deep_research
source: https://drift.simonbinder.eu/docs/getting-started/
topic: Flutter Drift Database Offline Architecture
---

# Flutter Drift Database Offline Architecture

[ Skip to content ](https://drift.simonbinder.eu/setup/#the-dependencies)
[ ![logo](https://drift.simonbinder.eu/images/icon.png) ](https://drift.simonbinder.eu/ "Drift")
Drift 
Setup 
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
  * [ Home  ](https://drift.simonbinder.eu/)
  * Documentation  Documentation 
    * Getting Started  [ Getting Started  ](https://drift.simonbinder.eu/setup/) Table of contents 
      * [ The dependencies  ](https://drift.simonbinder.eu/setup/#the-dependencies)
      * [ Database class  ](https://drift.simonbinder.eu/setup/#database-class)
      * [ Next steps  ](https://drift.simonbinder.eu/setup/#next-steps)
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
  * [ The dependencies  ](https://drift.simonbinder.eu/setup/#the-dependencies)
  * [ Database class  ](https://drift.simonbinder.eu/setup/#database-class)
  * [ Next steps  ](https://drift.simonbinder.eu/setup/#next-steps)


[ ](https://github.com/simolus3/drift/edit/develop/docs/docs/setup.md "Edit this page") [ ](https://github.com/simolus3/drift/issues/new?template=docs.md&title=Documentation issue: Getting Started "Create documentation issue")
# Getting Started
Drift is a powerful database library for Dart and Flutter applications. To support its advanced capabilities like type-safe SQL queries, verification of your database and migrations, it uses a builder and command-line tooling that runs at compile-time.
This means that the setup involves a little more than just adding a single dependency to your pubspec. This page explains how to add drift to your project and gives pointers to the next steps. If you're stuck adding drift, or have questions or feedback about the project, please share that with the community by [starting a discussion on GitHub](https://github.com/simolus3/drift/discussions). If you want to look at an example app for inspiration, a cross-platform Flutter app using drift is available [as part of the drift repository](https://github.com/simolus3/drift/tree/develop/examples/app).
## The dependencies[¶](https://drift.simonbinder.eu/setup/#the-dependencies "Permanent link")
First, let's add drift to your project's `pubspec.yaml`. In addition to the core drift dependencies (`drift` and `drift_dev` to generate code), we're also adding a package to open database on the respective platform.
[Flutter (sqlite3)](https://drift.simonbinder.eu/setup/#__tabbed_1_1)[Dart (sqlite3)](https://drift.simonbinder.eu/setup/#__tabbed_1_2)[Dart (Postgres)](https://drift.simonbinder.eu/setup/#__tabbed_1_3)
```
[](https://drift.simonbinder.eu/setup/#__codelineno-0-1)dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-0-2)drift:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-0-3)drift_flutter:^0.2.5
[](https://drift.simonbinder.eu/setup/#__codelineno-0-4)path_provider:^2.1.5
[](https://drift.simonbinder.eu/setup/#__codelineno-0-5)
[](https://drift.simonbinder.eu/setup/#__codelineno-0-6)dev_dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-0-7)drift_dev:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-0-8)build_runner:^2.6.0

```

Alternatively, you can achieve the same result using the following command:
```
[](https://drift.simonbinder.eu/setup/#__codelineno-1-1)dart pub add drift drift_flutter path_provider dev:drift_dev dev:build_runner

```

```
[](https://drift.simonbinder.eu/setup/#__codelineno-2-1)dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-2-2)drift:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-2-3)sqlite3:^2.8.0
[](https://drift.simonbinder.eu/setup/#__codelineno-2-4)
[](https://drift.simonbinder.eu/setup/#__codelineno-2-5)dev_dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-2-6)drift_dev:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-2-7)build_runner:^2.6.0

```

Alternatively, you can achieve the same result using the following command:
```
[](https://drift.simonbinder.eu/setup/#__codelineno-3-1)dart pub add drift sqlite3 dev:drift_dev dev:build_runner

```

```
[](https://drift.simonbinder.eu/setup/#__codelineno-4-1)dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-4-2)drift:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-4-3)postgres:^3.5.6
[](https://drift.simonbinder.eu/setup/#__codelineno-4-4)drift_postgres:^1.3.1
[](https://drift.simonbinder.eu/setup/#__codelineno-4-5)
[](https://drift.simonbinder.eu/setup/#__codelineno-4-6)dev_dependencies:
[](https://drift.simonbinder.eu/setup/#__codelineno-4-7)drift_dev:^2.28.1
[](https://drift.simonbinder.eu/setup/#__codelineno-4-8)build_runner:^2.6.0

```

Alternatively, you can achieve the same result using the following command:
```
[](https://drift.simonbinder.eu/setup/#__codelineno-5-1)dart pub add drift postgres drift_postgres dev:drift_dev dev:build_runner

```

Drift only generates code for sqlite3 by default. So, also create a `build.yaml` to [configure](https://drift.simonbinder.eu/generation_options/) `drift_dev`:
```
[](https://drift.simonbinder.eu/setup/#__codelineno-6-1)targets:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-2)$default:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-3)builders:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-4)drift_dev:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-5)options:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-6)sql:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-7)dialects:
[](https://drift.simonbinder.eu/setup/#__codelineno-6-8)-postgres
[](https://drift.simonbinder.eu/setup/#__codelineno-6-9)# Uncomment if you need to support both
[](https://drift.simonbinder.eu/setup/#__codelineno-6-10)#       - sqlite

```

## Database class[¶](https://drift.simonbinder.eu/setup/#database-class "Permanent link")
Every project using drift needs at least one class to access a database. This class references all the tables you want to use and is the central entry point for drift's code generator. In this example, we'll assume that this database class is defined in a file called `database.dart` and somewhere under `lib/`. Of course, you can put this class in any Dart file you like.
To make the database useful, we'll also add a simple table to it. This table, `TodoItems`, can be used to store todo items for a todo list app. Everything there is to know about defining tables in Dart is described on the [Dart tables](https://drift.simonbinder.eu/dart_api/tables/) page. If you prefer using SQL to define your tables, drift supports that too! You can read all about that [here](https://drift.simonbinder.eu/sql_api/).
For now, populate the contents of `database.dart` with these tables which could form the persistence layer of a simple todolist application:
```
import 'package:drift/drift.dart';
part 'database.g.dart';
class TodoItems extends Table {
 IntColumn get id => integer().autoIncrement()();
 TextColumn get title => text().withLength(min: 6, max: 32)();
 TextColumn get content => text().named('body')();
 DateTimeColumn get createdAt => dateTime().nullable()();
}
@DriftDatabase(tables: [TodoItems])
class AppDatabase extends _$AppDatabase {
}
```

You will get an analyzer warning on the `part` statement and on `extends _$AppDatabase`. This is expected because drift's generator did not run yet. You can do that by invoking [build_runner](https://pub.dev/packages/build_runner):
  * `dart run build_runner build` generates all the required code once.
  * `dart run build_runner watch` watches for changes in your sources and generates code with incremental rebuilds. This is suitable for development sessions.


After running either command, the `database.g.dart` file containing the generated `_$AppDatabase` class will have been generated. You will now see errors related to missing overrides and a missing constructor. The constructor is responsible for telling drift how to open the database. The `schemaVersion` getter is relevant for migrations after changing the database, we can leave it at `1` for now. Update `database.dart` so it now looks like this:
[Flutter (sqlite3)](https://drift.simonbinder.eu/setup/#__tabbed_2_1)[Dart (sqlite3)](https://drift.simonbinder.eu/setup/#__tabbed_2_2)[Dart (Postgres)](https://drift.simonbinder.eu/setup/#__tabbed_2_3)
```
import 'package:drift/drift.dart';
import 'package:drift_flutter/drift_flutter.dart';
import 'package:path_provider/path_provider.dart';
part 'database.g.dart';
class TodoItems extends Table {
 IntColumn get id => integer().autoIncrement()();
 TextColumn get title => text().withLength(min: 6, max: 32)();
 TextColumn get content => text().named('body')();
 DateTimeColumn get createdAt => dateTime().nullable()();
}
@DriftDatabase(tables: [TodoItems])
class AppDatabase extends _$AppDatabase {
 // After generating code, this class needs to define a `schemaVersion` getter
 // and a constructor telling drift where the database should be stored.
 // These are described in the getting started guide: https://drift.simonbinder.eu/setup/
 AppDatabase([QueryExecutor? executor]) : super(executor ?? _openConnection());
 @override
 int get schemaVersion => 1;
 static QueryExecutor _openConnection() {
  return driftDatabase(
   name: 'my_database',
   native: const DriftNativeOptions(
    // By default, `driftDatabase` from `package:drift_flutter` stores the
    // database files in `getApplicationDocumentsDirectory()`.
    databaseDirectory: getApplicationSupportDirectory,
   ),
   // If you need web support, see https://drift.simonbinder.eu/platforms/web/
  );
 }
}
```

If you need to customize how databases are opened, you can also set the connection up manually:
Manual database setup
```
[](https://drift.simonbinder.eu/setup/#__codelineno-7-1)import'dart:io';
[](https://drift.simonbinder.eu/setup/#__codelineno-7-2)import'package:drift/native.dart';
[](https://drift.simonbinder.eu/setup/#__codelineno-7-3)import'package:path_provider/path_provider.dart';
[](https://drift.simonbinder.eu/setup/#__codelineno-7-4)import'package:path/path.dart'asp;
[](https://drift.simonbinder.eu/setup/#__codelineno-7-5)import'package:sqlite3/sqlite3.dart';
[](https://drift.simonbinder.eu/setup/#__codelineno-7-6)import'package:sqlite3_flutter_libs/sqlite3_flutter_libs.dart';
[](https://drift.simonbinder.eu/setup/#__codelineno-7-7)
[](https://drift.simonbinder.eu/setup/#__codelineno-7-8)LazyDatabase_openConnection(){
[](https://drift.simonbinder.eu/setup/#__codelineno-7-9)// the LazyDatabase util lets us find the right location for the file async.
[](https://drift.simonbinder.eu/setup/#__codelineno-7-10)returnLazyDatabase(()async{
[](https://drift.simonbinder.eu/setup/#__codelineno-7-11)// put the database file, called db.sqlite here, into the documents folder
[](https://drift.simonbinder.eu/setup/#__codelineno-7-12)// for your app.
[](https://drift.simonbinder.eu/setup/#__codelineno-7-13)finaldbFolder=awaitgetApplicationDocumentsDirectory();
[](https://drift.simonbinder.eu/setup/#__codelineno-7-14)finalfile=File(p.join(dbFolder.path,'db.sqlite'));
[](https://drift.simonbinder.eu/setup/#__codelineno-7-15)
[](https://drift.simonbinder.eu/setup/#__codelineno-7-16)// Also work around limitations on old Android versions
[](https://drift.simonbinder.eu/setup/#__codelineno-7-17)if(Platform.isAndroid){
[](https://drift.simonbinder.eu/setup/#__codelineno-7-18)awaitapplyWorkaroundToOpenSqlite3OnOldAndroidVersions();
[](https://drift.simonbinder.eu/setup/#__codelineno-7-19)}
[](https://drift.simonbinder.eu/setup/#__codelineno-7-20)
[](https://drift.simonbinder.eu/setup/#__codelineno-7-21)// Make sqlite3 pick a more suitable location for temporary files - the
[](https://drift.simonbinder.eu/setup/#__codelineno-7-22)// one from the system may be inaccessible due to sandboxing.
[](https://drift.simonbinder.eu/setup/#__codelineno-7-23)finalcachebase=(awaitgetTemporaryDirectory()).path;
[](https://drift.simonbinder.eu/setup/#__codelineno-7-24)// We can't access /tmp on Android, which sqlite3 would try by default.
[](https://drift.simonbinder.eu/setup/#__codelineno-7-25)// Explicitly tell it about the correct temporary directory.
[](https://drift.simonbinder.eu/setup/#__codelineno-7-26)sqlite3.tempDirectory=cachebase;
[](https://drift.simonbinder.eu/setup/#__codelineno-7-27)
[](https://drift.simonbinder.eu/setup/#__codelineno-7-28)returnNativeDatabase.createInBackground(file);
[](https://drift.simonbinder.eu/setup/#__codelineno-7-29)});
[](https://drift.simonbinder.eu/setup/#__codelineno-7-30)}

```

The Android-specific workarounds are necessary because sqlite3 attempts to use `/tmp` to store private data on unix-like systems, which is forbidden on Android. We also use this opportunity to work around a problem some older Android devices have with loading custom libraries through `dart:ffi`.
```
import 'package:drift/drift.dart';
import 'dart:io';
import 'package:drift/native.dart';
part 'database.g.dart';
class TodoItems extends Table {
 IntColumn get id => integer().autoIncrement()();
 TextColumn get title => text().withLength(min: 6, max: 32)();
 TextColumn get content => text().named('body')();
 DateTimeColumn get createdAt => dateTime().nullable()();
}
@DriftDatabase(tables: [TodoItems])
class AppDatabase extends _$AppDatabase {
 // After generating code, this class needs to define a `schemaVersion` getter
 // and a constructor telling drift where the database should be stored.
 // These are described in the getting started guide: https://drift.simonbinder.eu/setup/
 AppDatabase([QueryExecutor? executor]) : super(executor ?? _openConnection());
 @override
 int get schemaVersion => 1;
 static QueryExecutor _openConnection() {
  return NativeDatabase.createInBackground(File('path/to/your/database'));
 }
}
```

```
import 'package:drift/drift.dart';
import 'package:drift_postgres/drift_postgres.dart';
import 'package:postgres/postgres.dart' as pg;
part 'database.g.dart';
class TodoItems extends Table {
 IntColumn get id => integer().autoIncrement()();
 TextColumn get title => text().withLength(min: 6, max: 32)();
 TextColumn get content => text().named('body')();
 DateTimeColumn get createdAt => dateTime().nullable()();
}
@DriftDatabase(tables: [TodoItems])
class AppDatabase extends _$AppDatabase {
 // After generating code, this class needs to define a `schemaVersion` getter
 // and a constructor telling drift where the database should be stored.
 // These are described in the getting started guide: https://drift.simonbinder.eu/setup/
 AppDatabase([QueryExecutor? executor]) : super(executor ?? _openConnection());
 @override
 int get schemaVersion => 1;
 static QueryExecutor _openConnection() {
  return PgDatabase(
   endpoint: pg.Endpoint(
    host: 'localhost',
    database: 'database',
    username: 'dart',
    password: 'mysecurepassword',
   ),
  );
 }
}
```

## Next steps[¶](https://drift.simonbinder.eu/setup/#next-steps "Permanent link")
Congratulations! With this setup complete, your project is ready to use drift. This short snippet shows how the database can be opened and how to run inserts and selects:
```
void main() async {
 WidgetsFlutterBinding.ensureInitialized();
 final database = AppDatabase();
 await database.into(database.todoItems).insert(TodoItemsCompanion.insert(
    title: 'todo: finish drift setup',
    content: 'We can now write queries and define our own tables.',
   ));
 List<TodoItem> allItems = await database.select(database.todoItems).get();
 print('items in database: $allItems');
}
```

But drift can do so much more! These pages provide more information useful when getting started with drift:
  * [Dart tables](https://drift.simonbinder.eu/dart_api/tables/): This page describes how to define your own tables in Dart. For an overview of the classes drift generates for tables, check out [row classes](https://drift.simonbinder.eu/dart_api/rows/).
  * For new drift users or users not familiar with SQL, the [manager](https://drift.simonbinder.eu/dart_api/manager/) APIs for tables allows writing most queries with a syntax you're likely familiar with from ORMs or other packages.
  * Writing queries: Drift-generated classes support writing the most common SQL statements, like [selects](https://drift.simonbinder.eu/dart_api/select/) or [inserts, updates and deletes](https://drift.simonbinder.eu/dart_api/writes/).
  * Something to keep in mind for later: When changing the database, for instance by adding new columns or tables, you need to write a migration so that existing databases are transformed to the new format. Drift's extensive [migration tools](https://drift.simonbinder.eu/Migrations/) help with that.
  * Take a look at our [FAQ](https://drift.simonbinder.eu/faq/)! It will help you with the most common questions about drift projects.


Once you're familiar with the basics, the [overview here](https://drift.simonbinder.eu/) shows what more drift has to offer. This includes transactions, automated tooling to help with migrations, multi-platform support and more.
Back to top  [ Previous  Home  ](https://drift.simonbinder.eu/) [ Next  Tables  ](https://drift.simonbinder.eu/dart_api/tables/)
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)
[ ](https://github.com/simolus3/drift "Project on GitHub")


## Source Information
- URL: https://drift.simonbinder.eu/docs/getting-started/
- Harvested: 2025-08-19T11:36:49.796939
- Profile: deep_research
- Agent: architect
