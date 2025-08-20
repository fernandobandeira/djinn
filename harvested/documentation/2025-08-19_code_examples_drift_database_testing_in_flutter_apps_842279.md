---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:20:55.809757'
profile: code_examples
source: https://drift.simonbinder.eu/docs/testing/
topic: Drift Database Testing in Flutter Apps
---

# Drift Database Testing in Flutter Apps

[ Skip to content ](https://drift.simonbinder.eu/testing/#setup)
[ ![logo](https://drift.simonbinder.eu/images/icon.png) ](https://drift.simonbinder.eu/ "Drift")
Drift 
Testing 
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
    * Testing  [ Testing  ](https://drift.simonbinder.eu/testing/) Table of contents 
      * [ Setup  ](https://drift.simonbinder.eu/testing/#setup)
      * [ Writing tests  ](https://drift.simonbinder.eu/testing/#writing-tests)
      * [ Testing migrations  ](https://drift.simonbinder.eu/testing/#testing-migrations)
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
  * [ Setup  ](https://drift.simonbinder.eu/testing/#setup)
  * [ Writing tests  ](https://drift.simonbinder.eu/testing/#writing-tests)
  * [ Testing migrations  ](https://drift.simonbinder.eu/testing/#testing-migrations)


[ ](https://github.com/simolus3/drift/edit/develop/docs/docs/testing.md "Edit this page") [ ](https://github.com/simolus3/drift/issues/new?template=docs.md&title=Documentation issue: Testing "Create documentation issue")
# Testing
Flutter apps using drift can always be tested with [integration tests](https://flutter.dev/docs/cookbook/testing/integration/introduction) running on a real device. This guide focuses on writing unit tests for a database written in drift. Those tests can be run and debugged on your computer without additional setup, you don't need a physical device to run them.
## Setup[¶](https://drift.simonbinder.eu/testing/#setup "Permanent link")
For this guide, we're going to test a very simple database that stores user names. The only important change from a regular drift database is the constructor: We make the `QueryExecutor` argument explicit instead of having a no-args constructor that passes a fixed executor (like a `FlutterQueryExecutor` or a `NativeDatabase`) to the superclass:
```
[](https://drift.simonbinder.eu/testing/#__codelineno-0-1)import'package:drift/drift.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-0-2)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-3)part'database.g.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-0-4)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-5)classUsersextendsTable{
[](https://drift.simonbinder.eu/testing/#__codelineno-0-6)IntColumngetid=>integer().autoIncrement()();
[](https://drift.simonbinder.eu/testing/#__codelineno-0-7)TextColumngetname=>text()();
[](https://drift.simonbinder.eu/testing/#__codelineno-0-8)}
[](https://drift.simonbinder.eu/testing/#__codelineno-0-9)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-10)@DriftDatabase(tables:[Users])
[](https://drift.simonbinder.eu/testing/#__codelineno-0-11)classMyDatabaseextends_$MyDatabase{
[](https://drift.simonbinder.eu/testing/#__codelineno-0-12)MyDatabase(QueryExecutore):super(e);
[](https://drift.simonbinder.eu/testing/#__codelineno-0-13)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-14)@override
[](https://drift.simonbinder.eu/testing/#__codelineno-0-15)intgetschemaVersion=>1;
[](https://drift.simonbinder.eu/testing/#__codelineno-0-16)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-17)/// Creates a user and returns their id
[](https://drift.simonbinder.eu/testing/#__codelineno-0-18)Future<int>createUser(Stringname){
[](https://drift.simonbinder.eu/testing/#__codelineno-0-19)returninto(users).insert(UsersCompanion.insert(name:name));
[](https://drift.simonbinder.eu/testing/#__codelineno-0-20)}
[](https://drift.simonbinder.eu/testing/#__codelineno-0-21)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-22)/// Changes the name of a user with the [id] to the [newName].
[](https://drift.simonbinder.eu/testing/#__codelineno-0-23)Future<void>updateName(intid,StringnewName){
[](https://drift.simonbinder.eu/testing/#__codelineno-0-24)returnupdate(users).replace(User(id:id,name:newName));
[](https://drift.simonbinder.eu/testing/#__codelineno-0-25)}
[](https://drift.simonbinder.eu/testing/#__codelineno-0-26)
[](https://drift.simonbinder.eu/testing/#__codelineno-0-27)Stream<User>watchUserWithId(intid){
[](https://drift.simonbinder.eu/testing/#__codelineno-0-28)return(select(users)..where((u)=>u.id.equals(id))).watchSingle();
[](https://drift.simonbinder.eu/testing/#__codelineno-0-29)}
[](https://drift.simonbinder.eu/testing/#__codelineno-0-30)}

```

Installing sqlite
We can't distribute an sqlite installation as a pub package (at least not as something that works outside of a Flutter build), so you need to ensure that you have the sqlite3 shared library installed on your system.
On macOS, it's installed by default.
On Linux, you can use the `libsqlite3-dev` package on Ubuntu and the `sqlite3` package on Arch (other distros will have similar packages).
On Windows, you can [download 'Precompiled Binaries for Windows'](https://www.sqlite.org/download.html) and extract `sqlite3.dll` into a folder that's in your `PATH` environment variable. Then restart your device to ensure that all apps will run with this `PATH` change.
As `sqlite3_flutter_libs` bundles the latest sqlite3 version with your app, using a recent sqlite3 version is recommended to avoid differences in how tests behave from your app. The minimum sqlite version tested with drift is 3.29.0, but many drift features like `returning` or generated columns will require more recent versions.
## Writing tests[¶](https://drift.simonbinder.eu/testing/#writing-tests "Permanent link")
We can create an in-memory version of the database by using a `NativeDatabase.memory()` instead of a `FlutterQueryExecutor` or other implementations. A good place to open the database is the `setUp` and `tearDown` methods from `package:test`:
```
[](https://drift.simonbinder.eu/testing/#__codelineno-1-1)import'package:drift/drift.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-1-2)import'package:drift/native.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-1-3)import'package:test/test.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-1-4)// the file defined above, you can test any drift database of course
[](https://drift.simonbinder.eu/testing/#__codelineno-1-5)import'database.dart';
[](https://drift.simonbinder.eu/testing/#__codelineno-1-6)
[](https://drift.simonbinder.eu/testing/#__codelineno-1-7)voidmain(){
[](https://drift.simonbinder.eu/testing/#__codelineno-1-8)lateMyDatabasedatabase;
[](https://drift.simonbinder.eu/testing/#__codelineno-1-9)
[](https://drift.simonbinder.eu/testing/#__codelineno-1-10)setUp((){
[](https://drift.simonbinder.eu/testing/#__codelineno-1-11)database=MyDatabase(DatabaseConnection(
[](https://drift.simonbinder.eu/testing/#__codelineno-1-12)NativeDatabase.memory(),
[](https://drift.simonbinder.eu/testing/#__codelineno-1-13)closeStreamsSynchronously:true,[](https://drift.simonbinder.eu/testing/#__code_1_annotation_1)
[](https://drift.simonbinder.eu/testing/#__codelineno-1-14)));
[](https://drift.simonbinder.eu/testing/#__codelineno-1-15)});
[](https://drift.simonbinder.eu/testing/#__codelineno-1-16)tearDown(()async{
[](https://drift.simonbinder.eu/testing/#__codelineno-1-17)awaitdatabase.close();
[](https://drift.simonbinder.eu/testing/#__codelineno-1-18)});
[](https://drift.simonbinder.eu/testing/#__codelineno-1-19)}

```

With that setup in place, we can finally write some tests: 
```
[](https://drift.simonbinder.eu/testing/#__codelineno-2-1)test('users can be created',()async{
[](https://drift.simonbinder.eu/testing/#__codelineno-2-2)finalid=awaitdatabase.createUser('some user');
[](https://drift.simonbinder.eu/testing/#__codelineno-2-3)finaluser=awaitdatabase.watchUserWithId(id).first;
[](https://drift.simonbinder.eu/testing/#__codelineno-2-4)
[](https://drift.simonbinder.eu/testing/#__codelineno-2-5)expect(user.name,'some user');
[](https://drift.simonbinder.eu/testing/#__codelineno-2-6)});
[](https://drift.simonbinder.eu/testing/#__codelineno-2-7)
[](https://drift.simonbinder.eu/testing/#__codelineno-2-8)test('stream emits a new user when the name updates',()async{
[](https://drift.simonbinder.eu/testing/#__codelineno-2-9)finalid=awaitdatabase.createUser('first name');
[](https://drift.simonbinder.eu/testing/#__codelineno-2-10)
[](https://drift.simonbinder.eu/testing/#__codelineno-2-11)finalexpectation=expectLater(
[](https://drift.simonbinder.eu/testing/#__codelineno-2-12)database.watchUserWithId(id).map((user)=>user.name),
[](https://drift.simonbinder.eu/testing/#__codelineno-2-13)emitsInOrder(['first name','changed name']),
[](https://drift.simonbinder.eu/testing/#__codelineno-2-14));
[](https://drift.simonbinder.eu/testing/#__codelineno-2-15)
[](https://drift.simonbinder.eu/testing/#__codelineno-2-16)awaitdatabase.updateName(id,'changed name');
[](https://drift.simonbinder.eu/testing/#__codelineno-2-17)awaitexpectation;
[](https://drift.simonbinder.eu/testing/#__codelineno-2-18)});

```

## Testing migrations[¶](https://drift.simonbinder.eu/testing/#testing-migrations "Permanent link")
Drift can help you generate code for schema migrations. For more details, see [this guide](https://drift.simonbinder.eu/Migrations/tests/).
Back to top  [ Previous  Isolates  ](https://drift.simonbinder.eu/isolates/) [ Next  FAQ  ](https://drift.simonbinder.eu/faq/)
Made with [ Material for MkDocs ](https://squidfunk.github.io/mkdocs-material/)
[ ](https://github.com/simolus3/drift "Project on GitHub")


## Source Information
- URL: https://drift.simonbinder.eu/docs/testing/
- Harvested: 2025-08-19T23:20:55.809757
- Profile: code_examples
- Agent: architect
