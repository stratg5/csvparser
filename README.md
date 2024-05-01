# How to run the code

# How to run the tests

# Thought process and choices

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

## Dependency Injection

Everything here is testable except for the line where we generate the client for calling the API.

This is achieved through dependency injection, where we inject structs that implement interfaces.

Because we accept interfaces for our functions, we can create mocked implementations and pass those to our functions.

With the mocked implementations, we have full control over the behavior of the application and can test multiple scenarios.

## Packages

Why are your packages named like that? The answer to this is that this is the Golang standard :)

Go prioritizes concise package names that are all one word, not separated by underscores and not camel cased. It makes for
consistent code bases and is a familiar sight to those who write code in Go. This helps reduce time to ramp up on projects.

There is a lot of debate about camel case, snake case, etc. for naming in a lot of areas, but consistency is more important than
which one is "right". I'm a big believer in consistency. Let's decide on the best option, and then once its decided lets stay consistent.
