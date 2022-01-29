# io-reader

_Read data (a stream of bytes) from a buffer, a file, stdin and a pipe to a buffer
using the standard `io` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#test)

Documentation and references,

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input/output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Reader is an interface from which you can
read a stream of bytes into a buffer.

![IMAGE - buffered-io.jpg - IMAGE](../../docs/pics/input-output/buffered-io.jpg)

The io.Reader interface looks like,

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

Hence, we use the Method Read(),

```go
n, err := r.Read(buffer)
```

### FROM A STRING (*strings.Reader)

```go
sourceString := "This data is being put into a string reader"
s := strings.NewReader(sourceString)
```

### FROM A BUFFER (*bytes.Reader)

```go
sourceBuffer := []byte("This data is being put into a byte reader")
b := bytes.NewReader(sourceBuffer)
```

### FROM A FILE (*os.File)

```go
f, err := os.Open("test.txt")
```

### FROM A USER (os.Stdin)

Just use os.Stdin.

### FROM A PIPE (*io.PipeReader)

Use the pipe output rpipe,

```go
rpipe, wpipe := io.Pipe()
```

## RUN

Run,

```bash
go run io-reader.go
```

## TEST

To create _test files,

```bash
gotests -w -all io-reader.go
```

To unit test the code,

```bash
go test -cover ./... 
```
