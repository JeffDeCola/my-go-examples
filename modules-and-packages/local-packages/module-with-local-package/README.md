# MODULE WITH LOCAL PACKAGE EXAMPLE

_A go module with a local package._

Other Examples

* [module-with-local-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package)
  **<- YOU ARE HERE**
* [module-with-remote-package](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/remote-packages/module-with-remote-package)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#overview)
* [STRUCTURE](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#structure)
* [CREATE A MODULE (go.mod)](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#create-a-module-gomod)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#test)
* [VS CODE AND GO MULTI-MODULES (go.work)](https://github.com/JeffDeCola/my-go-examples/tree/master/modules-and-packages/local-packages/module-with-local-package#vs-code-and-go-multi-modules-gowork)

## OVERVIEW

A go **module** is used to manage various versions of dependencies/packages.

In go, each directory is considered its own **package**. A package is a collection
of related go source files in the same directory.

## STRUCTURE

![IMAGE - go-module-with-local-package - IMAGE](../../../docs/pics/modules-and-packages/go-module-with-local-package.jpg)

## CREATE A MODULE (go.mod)

```cmd
go mod init module-with-local-package
```

Will look like,

```txt
module module-with-local-package

go 1.19
```

## RUN

To run,

```bash
go run main.go
```

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

Now edit test files.

To unit test the code,

```bash
go test -cover ./...
```

## VS CODE AND GO MULTI-MODULES (go.work)

If you are using VS Code and have a multi module repo like this,
you will need to add a `go.work` file to the root of your workspace.

```txt
cd [workspace root]
go work init
go work use my-go-examples/modules-and-packages/module-with-local-package
go work use my-go-examples/modules-and-packages/module-with-remote-package
```

It will look like,

```txt
go 1.19

use (
    ./my-go-examples/modules-and-packages/module-with-local-package
    ./my-go-examples/modules-and-packages/module-with-remote-package
)
```

This diagram may help,

![IMAGE - vs-code-multi-root-workspace-with-go-multi-modules - IMAGE](../../../docs/pics/modules-and-packages/vs-code-multi-root-workspace-with-go-multi-modules.jpg)
