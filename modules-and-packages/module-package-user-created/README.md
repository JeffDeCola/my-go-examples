# module-package-user-created

_This module contains a user created package._

Other examples using,

* [module-package-user-created](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-package-user-created)
  **<- YOU ARE HERE**
* [module-package-public](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/module-package-public)

Table of Contents,

* tbd

Documentation and reference,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

A go **module** is used to manage various versions of dependencies/packages.

In go, each directory is considered its own **package**. A package is a collection
of related go source files in the same directory.

## STRUCTURE

For this example,

* **module-package-user-created** _Module_
  * **something.go** _Must use package main with func main()_
  * go.mod
  * **arithmetic** _User Created Package arithmetic_
    * **sum.go** _Must use package arithmetic_

## CREATE A MODULE (go.mod)

```cmd
go mod init module-package-user-create
```

Will look like,

```go
module module-package-user-create

go 1.19
```

## RUN

To run,

```bash
go run something.go
```

## TEST

To create _test files,

```bash
gotests -w -all something.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## VS CODE AND MULTI MODULE REPO (go.work)

If you are using VS Code and have a multi module repo like this,
you will need to add a go.work file to the root of your workspace.

```txt
cd [workspace root]
go work init
go work use my-go-examples/modules-and-packages/module-package-user-created
go work use my-go-examples/modules-and-packages/module-package-public
```

It will look like,

```txt
go 1.19

use (
    ./my-go-examples/modules-and-packages/module-package-user-created
    ./my-go-examples/modules-and-packages/module-package-public
)
```
