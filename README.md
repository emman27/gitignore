# Gitignore generator

This is a CLI tool for generating gitignores

## Why I wrote this

Most commonly, when I start my project, my workflow looks like this

```bash
mkdir -p ~/dev/my-project
cd ~/dev/my-project
git init
touch .gitignore
# search github for gitignore file for the language/framework I'm using
# paste gitignore contents into file
```

This project aims to replicate those 2 manual steps with a simple command - for a Go project, for example:

```bash
mkdir -p ~/dev/my-project
cd ~/dev/my-project
git init
gitignore generate go
```

## Installation

The easiest way to install this package is to download a versioned release at the [releases page](https://github.com/emman27/gitignore/releases)

### Building from source

Requires a working `go` installation. This package was developed in `go 1.12`, but should work with any decently recent version of `go`.

```bash
go get -u github.com/emman27/gitignore
```

## Supported keywords

The list of supported gitignores will grow based on usage. As of now, only the following keywords are supported for generation

```text
go
golang
Go
```

## Testing

Testing is done using the `testify` framework as well as `mockery` for mock generation. You will need to generate the mock code for the tests to run. If you don't have an internet connection, run the tests in short mode to avoid integration tests failing

```bash
go generate ./...
go test -short ./...
go test ./...
```
