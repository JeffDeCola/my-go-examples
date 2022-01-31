# io-writer-simple

_Write data (a stream of bytes) to a buffer
from a buffer using the standard `io` package._

Other examples using,

* IO.READER
  * [io-reader](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader-/io-reader)
  * [io-reader-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader-/io-reader-simple)
* IO.WRITER
  * [io-writer](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer)
  * [io-writer-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple)
    **<- YOU ARE HERE**

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-writer/io-writer-simple#test)

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

Therefore,

```go
// CREATE THE IO WRITER INTERFACE
var w io.Writer

// CREATE BUFFER TO READ FROM
buffer := []byte("This data is being put into a byte reader\n")

// CREATE THE BUFFER b TO WRITE TO
b := new(bytes.Buffer)
fmt.Printf("Buffer in:  %s\n", b.String())

// ASSIGN B TO WRITER
w = b

// WRITE METHOD (USING io.Writer)
_, err := w.Write(buffer)
if err != nil {
    fmt.Printf("Error with io.Writer: %s", err)
}

// PRINT b
fmt.Printf("Buffer out: %s\n", b.String())
```

## RUN

Run,

```bash
go run io-writer-simple.go
```

## TEST

To create _test files,

```bash
gotests -w -all io-writer-simple.go
```

To unit test the code,

```bash
go test -cover ./... 
```
