---
agent_context: architect
confidence: 0.95
harvested_at: '2025-08-19T23:09:34.063548'
profile: deep_research
source: https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/
topic: Testcontainers for Go Integration Testing
---

# Testcontainers for Go Integration Testing

[![Testcontainers Logo](https://testcontainers.com/images/testcontainers-logo.svg) ](https://testcontainers.com/) Menu
  * [Desktop](https://testcontainers.com/desktop/)
  * [Cloud](https://testcontainers.com/cloud/)
  * [Getting Started](https://testcontainers.com/getting-started/)
  * [Guides](https://testcontainers.com/guides/)
  * [Modules](https://testcontainers.com/modules/)
  * Docs
    * [![](https://testcontainers.com/images/language-logos/java.svg) Java](https://java.testcontainers.org/)
    * [![](https://testcontainers.com/images/language-logos/go.svg) Go](https://golang.testcontainers.org/)
    * [![](https://testcontainers.com/images/language-logos/dotnet.svg) .NET](https://dotnet.testcontainers.org/)
    * [![](https://testcontainers.com/images/language-logos/nodejs.svg) Node.js](https://node.testcontainers.org/)
    * [![](https://testcontainers.com/images/language-logos/clojure.svg) Clojure](https://cljdoc.org/d/clj-test-containers/clj-test-containers/)
    * [![](https://testcontainers.com/images/language-logos/elixir.svg) Elixir](https://github.com/testcontainers/testcontainers-elixir)
    * [![](https://testcontainers.com/images/language-logos/haskell.svg) Haskell](https://github.com/testcontainers/testcontainers-hs)
    * [![](https://testcontainers.com/images/language-logos/python.svg) Python](https://testcontainers-python.readthedocs.io/en/latest/)
    * [![](https://testcontainers.com/images/language-logos/ruby.svg) Ruby](https://github.com/testcontainers/testcontainers-ruby)
    * [![](https://testcontainers.com/images/language-logos/rust.svg) Rust](https://rust.testcontainers.org/)
    * [![](https://testcontainers.com/images/language-logos/php.svg) PHP](https://github.com/testcontainers/testcontainers-php)
    * [![](https://testcontainers.com/images/language-logos/c.svg) Native](https://github.com/testcontainers/testcontainers-native)
  * [Blog](https://atomicjar.com/category/testcontainers/)
  * [ Slack](https://slack.testcontainers.org/)
  * [ GitHub](https://github.com/testcontainers/)


[ Guides](https://testcontainers.com/guides)
# Getting started with Testcontainers for Go
  * Go
  * PostgreSQL


[ Get the code](https://github.com/testcontainers/tc-guide-getting-started-with-testcontainers-for-go)
  * [Prerequisites](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_prerequisites)
  * [What we are going to achieve in this guide](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_what_we_are_going_to_achieve_in_this_guide)
  * [Getting Started](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_getting_started)
  * [Create Customer struct](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_create_customer_struct)
  * [Create Repository](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_create_repository)
  * [Write test for Repository using testcontainers-go](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_write_test_for_repository_using_testcontainers_go)
  * [Reusing the containers and running multiple tests as a suite](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_reusing_the_containers_and_running_multiple_tests_as_a_suite)
  * [Run tests](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_run_tests)
  * [Summary](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_summary)
  * [Further Reading](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_further_reading)


In this guide, you will learn how to
  * Create a Go application with modules support.
  * Implement a Repository to manage customers data in PostgreSQL database using pgx driver.
  * Test the database interactions using testcontainers-go.


## Prerequisites[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_prerequisites)
  * Go 1.19+
  * Your favorite IDE(VS Code, GoLand)
  * A Docker environment supported by Testcontainers <https://golang.testcontainers.org/system_requirements/>


## What we are going to achieve in this guide[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_what_we_are_going_to_achieve_in_this_guide)
We are going to create a Go project and implement a Repository to save and retrieve the customer details from a PostgreSQL database. Then we will test this repository using the testcontainers-go postgres module.
## Getting Started[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_getting_started)
Let’s start with creating a Go project.
```
$ mkdir testcontainers-go-demo
$ cd testcontainers-go-demo
$ go mod init github.com/testcontainers/testcontainers-go-demoCopy
```

We are going to use the [jackc/pgx](https://github.com/jackc/pgx) PostgreSQL Driver to interact with the Postgres database and the **testcontainers-go** [postgres module](https://golang.testcontainers.org/modules/postgres/) to spin up a Postgres docker instance for testing. Also, we are going to use [testify](https://github.com/stretchr/testify) for running multiple tests as a suite and for writing assertions.
![Note](https://testcontainers.com/images/icons/note.svg)| If you are new to Testcontainers, then please visit [Testcontainers Getting Started](https://testcontainers.com/getting-started/) page to learn more about Testcontainers and the benefits of using it.  
---|---  
Let’s install these dependencies.
```
$ go get github.com/jackc/pgx/v5
$ go get github.com/testcontainers/testcontainers-go
$ go get github.com/testcontainers/testcontainers-go/modules/postgres
$ go get github.com/stretchr/testifyCopy
```

After installing these dependencies, your `go.mod` file should look like this:
```
module github.com/testcontainers/testcontainers-go-demo
go 1.19
require (
  github.com/jackc/pgx/v5 v5.3.1
  github.com/stretchr/testify v1.8.3
  github.com/testcontainers/testcontainers-go v0.20.1
  github.com/testcontainers/testcontainers-go/modules/postgres v0.20.1
)
require (
  // indirect dependencies here
)Copy
```

## Create Customer struct[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_create_customer_struct)
First, let us start with creating a `types.go` file in `customer` package and define the `Customer` struct to model the customer details as follows:
```
package customer
type Customer struct {
	Id  int
	Name string
	Email string
}Copy
```

## Create Repository[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_create_repository)
Next, create `customer/repo.go` file, define the `Repository` struct and then add methods to create a new customer and get a customer by email as follows:
```
package customer
import (
	"context"
	"fmt"
	"os"
	"github.com/jackc/pgx/v5"
)
type Repository struct {
	conn *pgx.Conn
}
func NewRepository(ctx context.Context, connStr string) (*Repository, error) {
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return &Repository{
		conn: conn,
	}, nil
}
func (r Repository) CreateCustomer(ctx context.Context, customer Customer) (Customer, error) {
	err := r.conn.QueryRow(ctx,
		"INSERT INTO customers (name, email) VALUES ($1, $2) RETURNING id",
		customer.Name, customer.Email).Scan(&customer.Id)
	return customer, err
}
func (r Repository) GetCustomerByEmail(ctx context.Context, email string) (Customer, error) {
	var customer Customer
	query := "SELECT id, name, email FROM customers WHERE email = $1"
	err := r.conn.QueryRow(ctx, query, email).
		Scan(&customer.Id, &customer.Name, &customer.Email)
	if err != nil {
		return Customer{}, err
	}
	return customer, nil
}Copy
```

Let’s understand what is going on here:
  * We have defined a `Repository` struct with a field of type `*pgc.Conn` which will be used for performing database operations.
  * We have defined a helper function `NewRepository(connStr)` that takes a database connection string and initializes `Repository`.
  * Then we have implemented `CreateCustomer()` and `GetCustomerByEmail()` methods on the `Repository` receiver.


## Write test for Repository using testcontainers-go[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_write_test_for_repository_using_testcontainers_go)
We have our `Repository` implementation ready, but for testing we need a PostgreSQL database. We can use **testcontainers-go** to spin up a Postgres database in a Docker container and run our tests connecting to that database.
In real applications we might use some database migration tool, but for this guide let us use a simple script to initialize our database.
Create a `testdata/init-db.sql` file to create CUSTOMERS table and insert the sample data as follows:
```
CREATE TABLE IF NOT EXISTS customers (id serial, name varchar(255), email varchar(255));
INSERT INTO customers(name, email) VALUES ('John', 'john@gmail.com');Copy
```

The **testcontainers-go** library provides the generic **Container** abstraction that can be used to run any containerised service. To further simplify, testcontainers-go provides technology specific modules that will reduce the boilerplate and also provides a functional options pattern to easily construct the container instance.
For example, PostgresContainer provides `WithImage()`, `WithDatabase()`, `WithUsername()`, `WithPassword()` etc functions to set various properties of Postgres containers easily.
Now, let’s create the `customer/repo_test.go` file and implement the test as follows:
```
package customer
import (
	"context"
	"path/filepath"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)
func TestCustomerRepository(t *testing.T) {
	ctx := context.Background()
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	assert.NoError(t, err)
	customerRepo, err := NewRepository(ctx, connStr)
	assert.NoError(t, err)
	c, err := customerRepo.CreateCustomer(ctx, Customer{
		Name: "Henry",
		Email: "henry@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, c)
	customer, err := customerRepo.GetCustomerByEmail(ctx, "henry@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "Henry", customer.Name)
	assert.Equal(t, "henry@gmail.com", customer.Email)
}Copy
```

Let’s understand what is going on here:
  * We have created an instance of `PostgresContainer` by specifying the Docker image `postgres:15.3-alpine`, from which the container needs to be created.
  * We have configured the initialization scripts using `WithInitScripts(…​)` so that after the database starts, the CUSTOMERS table will be created and sample data will be inserted.
  * Next, we have specified the username, password and database name for the Postgres container.
  * We have configured the `WaitStrategy` that will help to determine whether the Postgres container is fully ready to use or not.
  * Then, we have defined the test cleanup function using `t.Cleanup(…​)` so that at the end of the test the Postgres container will be removed.
  * Next, we obtained the database `ConnectionString` from `PostgresContainer` and initialized `Repository`.
  * Then, we have created a new customer with the email `henry@gmail.com` and verified that a customer with the email `henry@gmail.com` exists in our database.


## Reusing the containers and running multiple tests as a suite[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_reusing_the_containers_and_running_multiple_tests_as_a_suite)
In the previous section, we saw how to spin up a Postgres Docker container for a single test. But usually we might have multiple tests in a single file, and we may want to reuse the same Postgres Docker container for all the tests in that file.
We can use [testify suite](https://pkg.go.dev/github.com/stretchr/testify/suite) package to implement common test setup and teardown actions.
First, let us extract `PostgresContainer` creation logic into a separate file called `testhelpers/containers.go`.
```
package testhelpers
import (
	"context"
	"path/filepath"
	"time"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)
type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}
func CreatePostgresContainer(ctx context.Context) (*PostgresContainer, error) {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithInitScripts(filepath.Join("..", "testdata", "init-db.sql")),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString: connStr,
	}, nil
}Copy
```

In `containers.go`, we have defined `PostgresContainer` struct which extends testcontainers-go `PostgresContainer` struct to provide easy access to `ConnectionString` and created `CreatePostgresContainer()` function to instantiate `PostgresContainer`.
Now, let’s create `customer/repo_suite_test.go` file and implement tests for creating a new customer and getting customer by email by using testify suite package as follows:
```
package customer
import (
	"context"
	"log"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go-demo/testhelpers"
)
type CustomerRepoTestSuite struct {
	suite.Suite
	pgContainer *testhelpers.PostgresContainer
	repository *Repository
	ctx     context.Context
}
func (suite *CustomerRepoTestSuite) SetupSuite() {
	suite.ctx = context.Background()
	pgContainer, err := testhelpers.CreatePostgresContainer(suite.ctx)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer
	repository, err := NewRepository(suite.ctx, suite.pgContainer.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}
	suite.repository = repository
}
func (suite *CustomerRepoTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("error terminating postgres container: %s", err)
	}
}
func (suite *CustomerRepoTestSuite) TestCreateCustomer() {
	t := suite.T()
	customer, err := suite.repository.CreateCustomer(suite.ctx, Customer{
		Name: "Henry",
		Email: "henry@gmail.com",
	})
	assert.NoError(t, err)
	assert.NotNil(t, customer.Id)
}
func (suite *CustomerRepoTestSuite) TestGetCustomerByEmail() {
	t := suite.T()
	customer, err := suite.repository.GetCustomerByEmail(suite.ctx, "john@gmail.com")
	assert.NoError(t, err)
	assert.NotNil(t, customer)
	assert.Equal(t, "John", customer.Name)
	assert.Equal(t, "john@gmail.com", customer.Email)
}
func TestCustomerRepoTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerRepoTestSuite))
}Copy
```

Let’s understand what is going on here:
  * We have created `CustomerRepoTestSuite` by extending `suite.Suite` struct and added fields which will be used across multiple tests in that suite.
  * In the `SetupSuite()` function which will be executed only once before executing the tests, we have created `PostgresContainer` and initialized `Repository`.
  * In `TearDownSuite()` function which will be executed only once after all the tests in that suite are executed, we are terminating the container which will destroy the Postgres Docker container.
  * Next, we have created the tests `TestCreateCustomer()` and `TestGetCustomerByEmail()` as receiver functions on the suite.
  * Finally, we have created the test function `TestCustomerRepoTestSuite(t *testing.T)` which will run the test suite when we execute the tests using `go test`.


![Tip](https://testcontainers.com/images/icons/tip.svg)| For the purpose of this guide, we are not resetting the data in the database. But it is a good practice to reset the database in a known state before running any test.  
---|---  
## Run tests[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_run_tests)
You can run all the tests using `go test ./…​` and optionally add the flag "-v" for displaying verbose output.
```
$ go test -v ./...
=== RUN  TestCustomerRepoTestSuite
...
...
2023/06/13 09:27:11 🐳 Creating container for image docker.io/testcontainers/ryuk:0.4.0
2023/06/13 09:27:11 ✅ Container created: 2881f4e311a2
2023/06/13 09:27:11 🐳 Starting container: 2881f4e311a2
2023/06/13 09:27:12 🚧 Waiting for container id 2881f4e311a2 image: docker.io/testcontainers/ryuk:0.4.0
2023/06/13 09:27:12 ✅ Container started: 2881f4e311a2
2023/06/13 09:27:12 🐳 Creating container for image postgres:15.3-alpine
2023/06/13 09:27:12 ✅ Container created: a98029633d02
2023/06/13 09:27:12 🐳 Starting container: a98029633d02
2023/06/13 09:27:13 🚧 Waiting for container id a98029633d02 image: postgres:15.3-alpine
2023/06/13 09:27:14 ✅ Container started: a98029633d02
=== RUN  TestCustomerRepoTestSuite/TestCreateCustomer
=== RUN  TestCustomerRepoTestSuite/TestGetCustomerByEmail
2023/06/13 09:27:14 🐳 Terminating container: a98029633d02
2023/06/13 09:27:15 🚫 Container terminated: a98029633d02
--- PASS: TestCustomerRepoTestSuite (3.66s)
  --- PASS: TestCustomerRepoTestSuite/TestCreateCustomer (0.00s)
  --- PASS: TestCustomerRepoTestSuite/TestGetCustomerByEmail (0.00s)
=== RUN  TestCustomerRepository
2023/06/13 09:27:15 🐳 Creating container for image postgres:15.3-alpine
2023/06/13 09:27:15 ✅ Container created: fcf4241a61ab
2023/06/13 09:27:15 🐳 Starting container: fcf4241a61ab
2023/06/13 09:27:15 🚧 Waiting for container id fcf4241a61ab image: postgres:15.3-alpine
2023/06/13 09:27:16 ✅ Container started: fcf4241a61ab
2023/06/13 09:27:16 🐳 Terminating container: fcf4241a61ab
2023/06/13 09:27:17 🚫 Container terminated: fcf4241a61ab
--- PASS: TestCustomerRepository (1.94s)
PASS
ok 	github.com/testcontainers/testcontainers-go-demo/customer	6.177s
?  	github.com/testcontainers/testcontainers-go-demo/testhelpers	[no test files]Copy
```

You should see two Postgres docker containers automatically started: one for the suite and its two tests, and the other for the initial test we created, and all those tests should PASS. You can also notice that after tests are executed, the containers are stopped and removed automatically.
## Summary[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_summary)
The Testcontainers for Go library helped us to write integration tests by using the same type of database, Postgres, that we use in production as opposed to using mocks. As we are not using mocks and talking to the real services, we are free to do any code refactoring and still ensure that the application is working as expected.
To learn more about Testcontainers visit <http://testcontainers.com>
## Further Reading[](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/#_further_reading)
  * <https://golang.testcontainers.org/>
  * <https://golang.testcontainers.org/quickstart/>
  * <https://golang.testcontainers.org/modules/postgres/>
  * <https://testcontainers.com/modules/>


## Want to stay up to date?
[Read the Newsletter](https://newsletter.testcontainers.com/)
## Links
  * [Community Champions](https://testcontainers.com/community-champions/)
  * [Testcontainers for Java](https://java.testcontainers.org/)
  * [Testcontainers for Go](https://golang.testcontainers.org/)
  * [Testcontainers for .NET](https://dotnet.testcontainers.org/)
  * [Testcontainers for Node.js](https://node.testcontainers.org/)
  * [Testcontainers for Clojure](https://cljdoc.org/d/clj-test-containers/clj-test-containers/)
  * [Testcontainers for Elixir](https://github.com/testcontainers/testcontainers-elixir)
  * [Testcontainers for Haskell](https://github.com/testcontainers/testcontainers-hs)
  * [Testcontainers for Python](https://testcontainers-python.readthedocs.io/en/latest/)
  * [Testcontainers for Ruby](https://github.com/testcontainers/testcontainers-ruby)
  * [Testcontainers for Rust](https://rust.testcontainers.org/)
  * [Testcontainers for PHP](https://github.com/testcontainers/testcontainers-php)
  * [Testcontainers for Native](https://github.com/testcontainers/testcontainers-native)


## Join the community
We hope that you find Testcontainers reliable and intuitive to use. However sometimes things don't go the way we'd expect and we'd like to try and help out if we can.
  * [![Slack](https://testcontainers.com/images/slack.svg)](https://slack.testcontainers.org/)
  * [![GitHub](https://testcontainers.com/images/github.svg)](https://github.com/testcontainers)
  * [![StackOverflow](https://testcontainers.com/images/stackoverflow.svg)](https://stackoverflow.com/questions/tagged/testcontainers)
  * [![Twitter](https://testcontainers.com/images/twitter.svg)](https://twitter.com/testcontainers)


[Privacy Policy](https://www.iubenda.com/privacy-policy/58807048 "Privacy Policy") [Cookie Policy](https://www.iubenda.com/privacy-policy/58807048/cookie-policy "Cookie Policy ") [![California Consumer Privacy Act \(CCPA\) Opt-Out Icon](data:image/svg+xml;charset=UTF-8,%3csvg version='1.1' id='Layer_1' xmlns='http://www.w3.org/2000/svg' xmlns:xlink='http://www.w3.org/1999/xlink' x='0px' y='0px' viewBox='0 0 30 14' style='enable-background:new 0 0 30 14;' xml:space='preserve'%3e%3cstyle type='text/css'%3e .st0%7bfill-rule:evenodd;clip-rule:evenodd;fill:%23FFFFFF;%7d .st1%7bfill-rule:evenodd;clip-rule:evenodd;fill:%230066FF;%7d .st2%7bfill:%23FFFFFF;%7d .st3%7bfill:%230066FF;%7d %3c/style%3e%3cg%3e%3cg id='final---dec.11-2020_1_'%3e%3cg id='_x30_208-our-toggle_2_' transform='translate\(-1275.000000, -200.000000\)'%3e%3cg id='Final-Copy-2_2_' transform='translate\(1275.000000, 200.000000\)'%3e%3cpath class='st0' d='M7.4,12.8h6.8l3.1-11.6H7.4C4.2,1.2,1.6,3.8,1.6,7S4.2,12.8,7.4,12.8z'/%3e%3c/g%3e%3c/g%3e%3c/g%3e%3cg id='final---dec.11-2020'%3e%3cg id='_x30_208-our-toggle' transform='translate\(-1275.000000, -200.000000\)'%3e%3cg id='Final-Copy-2' transform='translate\(1275.000000, 200.000000\)'%3e%3cpath class='st1' d='M22.6,0H7.4c-3.9,0-7,3.1-7,7s3.1,7,7,7h15.2c3.9,0,7-3.1,7-7S26.4,0,22.6,0z M1.6,7c0-3.2,2.6-5.8,5.8-5.8 h9.9l-3.1,11.6H7.4C4.2,12.8,1.6,10.2,1.6,7z'/%3e%3cpath id='x' class='st2' d='M24.6,4c0.2,0.2,0.2,0.6,0,0.8l0,0L22.5,7l2.2,2.2c0.2,0.2,0.2,0.6,0,0.8c-0.2,0.2-0.6,0.2-0.8,0 l0,0l-2.2-2.2L19.5,10c-0.2,0.2-0.6,0.2-0.8,0c-0.2-0.2-0.2-0.6,0-0.8l0,0L20.8,7l-2.2-2.2c-0.2-0.2-0.2-0.6,0-0.8 c0.2-0.2,0.6-0.2,0.8,0l0,0l2.2,2.2L23.8,4C24,3.8,24.4,3.8,24.6,4z'/%3e%3cpath id='y' class='st3' d='M12.7,4.1c0.2,0.2,0.3,0.6,0.1,0.8l0,0L8.6,9.8C8.5,9.9,8.4,10,8.3,10c-0.2,0.1-0.5,0.1-0.7-0.1l0,0 L5.4,7.7c-0.2-0.2-0.2-0.6,0-0.8c0.2-0.2,0.6-0.2,0.8,0l0,0L8,8.6l3.8-4.5C12,3.9,12.4,3.9,12.7,4.1z'/%3e%3c/g%3e%3c/g%3e%3c/g%3e%3c/g%3e%3c/svg%3e) Your Privacy Choices ](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/)[Notice at Collection](https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/)


## Source Information
- URL: https://testcontainers.com/guides/getting-started-with-testcontainers-for-go/
- Harvested: 2025-08-19T23:09:34.063548
- Profile: deep_research
- Agent: architect
