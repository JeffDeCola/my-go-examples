# io-writer

_Write data (a stream of bytes) to a buffer, file, stdout
and a pipe from a buffer using the standard `io` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer#test)

Documentation and references,

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input/output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Writer is an interface from which you can
write a stream of bytes from a buffer.

![IMAGE - buffered-io.jpg - IMAGE](../../docs/pics/input-output/buffered-io.jpg)

The io.Writer interface looks like,

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

Hence, we use the Method Write(),

```go
n, err := r.Write(buffer)
```

To a buffer (*bytes.Buffer),

```go
b := new(bytes.Buffer)
```

To a file (*os.File),

```go
f, err := os.Create("filename.txt")
```

To a user (os.Stdout),

```go
o := os.Stdout
```

To a pipe (*io.PipeWriter),

```go
pipeReader, pipeWriter := io.Pipe()
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
