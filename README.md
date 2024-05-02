# How to run the code

# How to run the tests

# Thought process and choices

## Components

Services

You will see that almost all packages implement a Service. These services are the main drivers behind the functionality.

They are all named the same, since in Go the package is more important.

This way we get nice initialization calls like `csv.NewService()`

Contracts

You will also notice that each package has a contracts file. This is where we define the interfaces that are implemented
in each of the packages.

This is very important for mocking.

Mocks

You will also find a mocks file in each of the packages. This is where we have a separate mock implementation of the contracts.

These mock implementations can provide a way to override the functionality, while still implementing the interface.

The mocks can be passed anywhere that an interface is expected and the mocks match that interface.

## Authentication

There are two ways to authenticate with the app:

Using `api.go`

You can set your API ID and token in /entities/api.go.
This file is included in the .gitignore file so it won't be checked in.

Below is the format for this file:

```
const (
  ID = "REPLACE_ME"
  Token = "REPLACE_ME"
)
```

If the file is not present, you can also pass the ID and token through the flags on startup:

`empora -apiID "myID" -apiToken "myToken"`

## Smarty SDK

I've chosen to go with the Go SDK provided for interacting with their API. Since I am doing simple lookups, this was
the most direct path to achieving that.

If more complex requests were needed in the future, that could definitely be changed and we could use a standard HTTP client.

## Dependency Injection

Everything here is testable except for the line where we generate the client for calling the API.

This is achieved through dependency injection, where we inject structs that implement interfaces.

Because we accept interfaces for our functions, we can create mocked implementations and pass those to our functions.

With the mocked implementations, we have full control over the behavior of the application and can test multiple scenarios.

## Packages

### Naming

Why are your packages named like that? The answer to this is that this is the Golang standard :)

Go prioritizes concise package names that are all one word, not separated by underscores and not camel cased. It makes for
consistent code bases and is a familiar sight to those who write code in Go. This helps reduce time to ramp up on projects.

There is a lot of debate about camel case, snake case, etc. for naming in a lot of areas, but consistency is more important than
which one is "right". I'm a big believer in consistency. Let's decide on the best option, and then once its decided lets stay consistent.

### Division

I've created an address client that can perform address-related actions. This felt like a logical separation to me, it could be extended
in the future to perform additional address-related tasks.

I've created an entities package to hold all of the objects. This gives a single location to access each of the entities for the application.
In the future, this could be broken out into smaller files under the entities package, for example `addressentities.go`, `titleentities.go`, etc.

The entities could be placed into the package that uses them (i.e. the `Address` struct goes in the address client), but some entities are used by
multiple actors. Then entities start to get placed in specific packages and universal locations which can get messy. This provides one location to
access entities.
