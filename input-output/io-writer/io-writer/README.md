# IO.WRITER EXAMPLE

_Write data (a stream of bytes) to a buffer, file, stdout and a pipe
from a buffer using the standard `io` package._

Other Examples

* IO.READER
  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader)
  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple)
* IO.WRITER
  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)
    **<- YOU ARE HERE**
  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer#test)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer#an-illustration-that-may-help)

Documentation and Reference

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input/output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Writer is an interface from which you can
write a stream of bytes from a buffer.

The io.Writer interface looks like,

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

Hence, we use the method Write(),

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
go run main.go
```

## TEST

To create _test files,

```bash
gotests -w -all main.go
```

To unit test the code,

```bash
go test -cover ./... 
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - buffered-io.jpg - IMAGE](../../../docs/pics/input-output/buffered-io.jpg)
