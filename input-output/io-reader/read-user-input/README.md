# read-user-input

_Read user input (os.Stdin) to a buffer (using Read method)
and string (using Fscan)._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/read-user-input#test)

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Read user input until "stop".

## RUN

User one word. This is just an example, not for production.

```bash
go run read-user-input.go
```

## TEST

To create `_test` files,

```bash
gotests -w -all read-user-input.go
```

To unit test the code,

```bash
go test -cover ./...
```
