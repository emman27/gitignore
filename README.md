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

```sh
gitignore generate go
```

## Supported keywords

The list of supported gitignores will grow based on usage. As of now, only the following keywords are supported for generation

```text
go
```
