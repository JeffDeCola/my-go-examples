# io-writer

_Write data (a stream of bytes) to a buffer, a file, stdout and a pipe from a buffer
using the standard `io` package._

Table of Contents,

* ?????

Documentation and references,

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input and output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Writer is an interface from which you can
write a stream of bytes from a buffer.

### TO A BUFFER

```go
???
```

### TO A FILE

```go
???
```

### TO STDOUT

```go
???
```

### TO A PIPE

```go
???
```

## RUN

Run,

```bash
go run io-writer.go
```

## TEST

To create _test files,

```bash
gotests -w -all io-writer.go
```

To unit test the code,

```bash
go test -cover ./... 
```
