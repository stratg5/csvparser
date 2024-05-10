## Running the app

I have provided a few pre-built binaries that you can run, you just need to cd into the directory where they live and run:

`./parser-darwin-amd64` on Mac (or `./parser-darwin-arm64` for M architecture), or `.\parser.exe` on Windows.

If you are running Windows or Mac, these binaries should be natively compiled and work out of the box, you do not need a runtime.

You may also build and run the app yourself:

- With Go installed (https://go.dev/doc/install), cd into the project directory
- Run `go build`
- A binary is output, run the binary as above, depending on your OS
- Note: you can also run `go run main.go`, which will run the app directly from the code

## Running the tests

With go installed and a terminal open to the top level directory, run `go test ./...`. This will run all tests.

## Flags

You can pass a few different flags to make the program easier to use:

`-inputPath` is the path to the input csv file. It defaults to `./input.csv'
`-outputPath`is the path to the output csv file. It defaults to`./output.csv'
`-id` can be used to pass your smarty ID to the app if its not set in the api.go file (more on that below)
`-token` can be used to pass your smarty token to the app if its not set in the api.go file (more on that below)

Example: `./parser-darwin-amd64 -inputPath ./input.csv -outputPath ./output.csv -id YOUR_SMARTY_ID -token YOUR_SMARTY_TOKEN`

## Components

**Services**

You will see that almost all packages implement a Service. These services are the main drivers behind the functionality.

They are all named the same, since in Go the package is more important.

This way we get nice initialization calls like `csv.NewService()`

**Contracts**

You will also notice that each package has a contracts file. This is where we define the interfaces that are implemented
in each of the packages.

This is very important for mocking.

**Mocks**

You will also find a mocks file in each of the packages. This is where we have a separate mock implementation of the contracts.

These mock implementations can provide a way to override the functionality, while still implementing the interface.

The mocks can be passed anywhere that an interface is expected and the mocks match that interface.

## Authentication

There are two ways to authenticate with the app:

Using `api.go`

You can set your API ID and token in /entities/api.go, just replace the ID and token with your ID and token.

You can also pass the ID and token through the flags on startup:

`parser -id "myID" -token "myToken"`

## Smarty SDK

I've chosen to go with the Go SDK provided for interacting with their API. Since I am doing simple lookups, this was
the most direct path to achieving that.

If more complex requests were needed in the future, that could definitely be changed and we could use a standard HTTP client.

## Dependency Injection

Everything here is testable except for the line where we generate the client for calling the API.

This is achieved through dependency injection, where we inject structs that implement interfaces.

Because we accept interfaces for our functions, we can create mocked implementations and pass those to our functions.

With the mocked implementations, we have full control over the behavior of the application and can test multiple scenarios.

## Package Naming

Why are your packages named like that? The answer to this is that this is the Golang standard :)

Go prioritizes concise package names that are all one word, not separated by underscores and not camel cased. It makes for
consistent code bases and is a familiar sight to those who write code in Go. This helps reduce time to ramp up on projects.

There is a lot of debate about camel case, snake case, etc. for naming in a lot of areas, but consistency is more important than
which one is "right". I'm a big believer in consistency. Let's decide on the best option, and then once its decided lets stay consistent.

## Package Structure

I've created an address client that can perform address-related actions. This felt like a logical separation to me, and it could be extended
in the future to perform additional address-related tasks.

I've created an entities package to hold all of the objects. This gives a single location to access each of the entities for the application.
In the future, this could be broken out into smaller files under the entities package, for example `addressentities.go`, `titleentities.go`, etc.

The entities could be placed into the package that uses them (i.e. the `Address` struct goes in the address client), but some entities are used by
multiple actors. Then entities start to get placed in specific packages and universal locations which can get messy. This provides one location to
access entities.

`main.go` is responsible for initializing all the dependencies and kicking off the main flow. This is where all of the real implementations get created
and injected into the services responsible for performing the logic.

## Output

Output is one part that I would definitely talk to someone about if I were actually implementing this for a project. The desired output is a bit odd
to fit into a csv, it would be better as a text file or to change the output so it is better suited for csv.

The `->` with address data on either side especially felt like a format that was not well suited to csv, as well as some of the spacing before fields.
Ultimately I went with csv output since the input was csv, but its something I would want to discuss with the people using the output of the report to
try and determine if maybe we can better suit their needs.
