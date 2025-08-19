---
agent_context: unknown
confidence: 0.95
harvested_at: '2025-08-19T01:20:06.913336'
profile: code_examples
source: https://gqlgen.com/getting-started/
topic: gqlgen-graphql-setup
---

# gqlgen-graphql-setup

v0.17.78
v0.17.78 [v0.17.77](https://gqlgen.com/v0.17.77/) [v0.17.76](https://gqlgen.com/v0.17.76/) [v0.17.75](https://gqlgen.com/v0.17.75/) [v0.17.74](https://gqlgen.com/v0.17.74/) [v0.17.73](https://gqlgen.com/v0.17.73/) [v0.17.72](https://gqlgen.com/v0.17.72/) [v0.17.71](https://gqlgen.com/v0.17.71/) [v0.17.70](https://gqlgen.com/v0.17.70/) [v0.17.69](https://gqlgen.com/v0.17.69/) [v0.17.68](https://gqlgen.com/v0.17.68/) [v0.17.67](https://gqlgen.com/v0.17.67/) [v0.17.66](https://gqlgen.com/v0.17.66/) [v0.17.65](https://gqlgen.com/v0.17.65/) [v0.17.64](https://gqlgen.com/v0.17.64/) [v0.17.63](https://gqlgen.com/v0.17.63/) [v0.17.62](https://gqlgen.com/v0.17.62/) [v0.17.61](https://gqlgen.com/v0.17.61/) [v0.17.60](https://gqlgen.com/v0.17.60/) [v0.17.59](https://gqlgen.com/v0.17.59/) [master](https://gqlgen.com/master/)
[ gqlgen ](http://github.com/99designs/gqlgen)
  * [Introduction](https://gqlgen.com/)
  * [Getting Started](https://gqlgen.com/getting-started/)
  * [Configuration](https://gqlgen.com/config/)
  * [Feature Comparison](https://gqlgen.com/feature-comparison/)
  * Reference
    * [ APQ ](https://gqlgen.com/reference/apq/)
    * [ Changesets ](https://gqlgen.com/reference/changesets/)
    * [ Dataloaders ](https://gqlgen.com/reference/dataloaders/)
    * [ Field Collection ](https://gqlgen.com/reference/field-collection/)
    * [ File Upload ](https://gqlgen.com/reference/file-upload/)
    * [ Handling Errors ](https://gqlgen.com/reference/errors/)
    * [ Introspection ](https://gqlgen.com/reference/introspection/)
    * [ Model Generation ](https://gqlgen.com/reference/model-generation/)
    * [ Name Collision ](https://gqlgen.com/reference/name-collision/)
    * [ Plugins ](https://gqlgen.com/reference/plugins/)
    * [ Query Complexity ](https://gqlgen.com/reference/complexity/)
    * [ Resolvers ](https://gqlgen.com/reference/resolvers/)
    * [ Scalars ](https://gqlgen.com/reference/scalars/)
    * [ Schema Directives ](https://gqlgen.com/reference/directives/)
    * [ pkg.go.dev → ](https://pkg.go.dev/github.com/99designs/gqlgen)
  * Recipes
    * [ Apollo Federation ](https://gqlgen.com/recipes/federation/)
    * [ Authentication ](https://gqlgen.com/recipes/authentication/)
    * [ CORS ](https://gqlgen.com/recipes/cors/)
    * [ Enum binding ](https://gqlgen.com/recipes/enum/)
    * [ Generated Model Extra Fields ](https://gqlgen.com/recipes/extra_fields/)
    * [ Gin ](https://gqlgen.com/recipes/gin/)
    * [ Migrating to 0.11 ](https://gqlgen.com/recipes/migration-0.11/)
    * [ Modelgen hook ](https://gqlgen.com/recipes/modelgen-hook/)
    * [ Subscriptions ](https://gqlgen.com/recipes/subscriptions/)


# Getting Started
Building GraphQL servers in golang
[ [edit] ](https://github.com/99designs/gqlgen/blob/master/docs/content/getting-started.md)
This tutorial will take you through the process of building a GraphQL server with gqlgen that can:
  * Return a list of todos
  * Create new todos
  * Mark off todos as they are completed


You can find the finished code for this tutorial [here](https://github.com/vektah/gqlgen-tutorials/tree/master/gettingstarted)
## Set up Project
Create a directory for your project, and [initialise it as a Go Module](https://golang.org/doc/tutorial/create-module):
```
mkdir gqlgen-todos
cd gqlgen-todos
go mod init github.com/[username]/gqlgen-todos

```

Next, create a `tools.go` file and add gqlgen as a [tool dependency for your module](https://go.dev/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module).
```
//go:build tools
package tools
import (
	_ "github.com/99designs/gqlgen"
)

```

To automatically add the dependency to your `go.mod` run
```
go mod tidy

```

By default you’ll be using the latest version of gqlgen, but if you want to specify a particular version you can use `go get` (replacing `VERSION` with the particular version desired)
```
go get -d github.com/99designs/gqlgen@VERSION

```

## Building the server
### Create the project skeleton
```
go run github.com/99designs/gqlgen init

```

This will create our suggested package layout. You can modify these paths in gqlgen.yml if you need to.
```
├── go.mod
├── go.sum
├── gqlgen.yml        - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated      - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model        - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go     - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls   - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go - the resolver implementation for schema.graphql
└── server.go        - The entry point to your app. Customize it however you see fit

```

### Define your schema
gqlgen is a schema-first library — before writing code, you describe your API using the GraphQL [Schema Definition Language](http://graphql.org/learn/schema/). By default this goes into a file called `schema.graphqls` but you can break it up into as many different files as you want.
The schema that was generated for us was:
```
typeTodo{id:ID!text:String!done:Boolean!user:User!}typeUser{id:ID!name:String!}typeQuery{todos:[Todo!]!}inputNewTodo{text:String!userId:String!}typeMutation{createTodo(input:NewTodo!):Todo!}
```

### Implement the resolvers
When executed, gqlgen’s `generate` command compares the schema file (`graph/schema.graphqls`) with the models `graph/model/*`, and, wherever it can, it will bind directly to the model. That was done already when `init` was run. We’ll edit the schema later in the tutorial, but for now, let’s look at what was generated already.
If we take a look in `graph/schema.resolvers.go` we will see all the times that gqlgen couldn’t match them up. For us it was twice:
```
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

```

We just need to implement these two methods to get our server working:
First we need somewhere to track our state, lets put it in `graph/resolver.go`. The `graph/resolver.go` file is where we declare our app’s dependencies, like our database. It gets initialized once in `server.go` when we create the graph.
```
type Resolver struct{
	todos []*model.Todo
}

```

Returning to `graph/schema.resolvers.go`, let’s implement the bodies of those automatically generated resolver functions. For `CreateTodo`, we’ll use the [`crypto.rand` package](https://pkg.go.dev/crypto/rand#Int) to simply return a todo with a randomly generated ID and store that in the in-memory todos list — in a real app, you’re likely to use a database or some other backend service.
```
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text: input.Text,
		ID:  fmt.Sprintf("T%d", randNumber),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

```

### Run the server
We now have a working server, to start it:
```
go run server.go

```

Open http://localhost:8080 in a browser. Here are some queries to try, starting with creating a todo:
```
mutationcreateTodo{createTodo(input:{text:"todo",userId:"1"}){user{id}textdone}}
```

And then querying for it:
```
queryfindTodos{todos{textdoneuser{name}}}
```

### Don’t eagerly fetch the user
This example is great, but in the real world fetching most objects is expensive. We don’t want to load the User on the todo unless the user actually asked for it. So lets replace the generated `Todo` model with something slightly more realistic.
First let’s enable `autobind`, allowing gqlgen to use your custom models if it can find them rather than generating them. We do this by uncommenting the `autobind` config line in `gqlgen.yml`:
```
# gqlgen will search for any type names in the schema in these go packages# if they match it will use them, otherwise it will generate them.autobind:- "github.com/[username]/gqlgen-todos/graph/model"
```

And add `Todo` fields resolver config in `gqlgen.yml` to generate resolver for `user` field
```
# This section declares type mapping between the GraphQL and go type systems## The first line in each type will be used as defaults for resolver arguments and# modelgen, the others will be allowed when binding to fields. Configure them to# your likingmodels:ID:model:- github.com/99designs/gqlgen/graphql.ID- github.com/99designs/gqlgen/graphql.Int- github.com/99designs/gqlgen/graphql.Int64- github.com/99designs/gqlgen/graphql.Int32Int:model:- github.com/99designs/gqlgen/graphql.Int32Todo:fields:user:resolver:true
```

Next, create a new file called `graph/model/todo.go`
```
package model
type Todo struct {
	ID   string `json:"id"`
	Text  string `json:"text"`
	Done  bool  `json:"done"`
	UserID string `json:"userId"`
	User  *User `json:"user"`
}

```

And run `go run github.com/99designs/gqlgen generate`.
> If you run into this error `package github.com/99designs/gqlgen: no Go files` while executing the `generate` command above, follow the instructions in [this](https://github.com/99designs/gqlgen/issues/800#issuecomment-888908950) comment for a possible solution.
Now if we look in `graph/schema.resolvers.go` we can see a new resolver, lets implement it and fix `CreateTodo`.
```
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text:  input.Text,
		ID:   fmt.Sprintf("T%d", randNumber),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}
func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

```

## Finishing touches
At the top of our `resolver.go`, between `package` and `import`, add the following line:
```
//go:generate go run github.com/99designs/gqlgen generate

```

This magic comment tells `go generate` what command to run when we want to regenerate our code. To run go generate recursively over your entire project, use this command:
```
go generate ./...

```



## Source Information
- URL: https://gqlgen.com/getting-started/
- Harvested: 2025-08-19T01:20:06.913336
- Profile: code_examples
- Agent: unknown
