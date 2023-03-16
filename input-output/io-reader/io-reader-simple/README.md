# IO.READER SIMPLE EXAMPLE

_Read data (a stream of bytes) from a buffer to a buffer
using the standard `io` package._

Other Examples

* IO.READER
  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader)
  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple)
    **<- YOU ARE HERE**
* IO.WRITER
  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)
  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple#run)
* [AN ILLUSTRATION THAT MAY HELP](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader/io-reader-simple#an-illustration-that-may-help)

Documentation and Reference

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input/output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Reader is an interface from which you can
read a stream of bytes into a buffer.

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

Therefore,

```go
// CREATE THE IO READER INTERFACE
var r io.Reader

// CREATE THE BUFFER b TO READ FROM
b := new(bytes.Buffer)
b.WriteString("This is data in byte.buffer")
fmt.Printf("Buffer in:  %s\n", b.String())

// CREATE BUFFER TO WRITE TO
buffer := make([]byte, 100)

// ASSIGN BUFFER TO READER
r = b

// READ METHOD (USING io.Reader)
n, err := r.Read(buffer)
if err != nil {
    fmt.Printf("Error with io.Reader: %s", err)
}

// PRINT BUFFER
fmt.Printf("Buffer out: %s\n", string(buffer[:n]))
```

## RUN

Run,

```bash
go run main.go
```

## AN ILLUSTRATION THAT MAY HELP

![IMAGE - buffered-io.jpg - IMAGE](../../../docs/pics/input-output/buffered-io.jpg)
