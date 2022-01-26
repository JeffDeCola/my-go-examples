# io-reader

_Read data (a stream of bytes) from a buffer, a file, stdin and a pipe to a buffer
using the standard `io` package._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#overview)
  * [FROM A BUFFER (STRING)](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#from-a-buffer-string)
  * [FROM A FILE](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#from-a-file)
  * [FROM STDIN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#from-stdin)
  * [FROM A PIPE](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#from-a-pipe)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/input-output/io-reader#test)

Documentation and references,

* Refer to the
  [io](https://pkg.go.dev/io)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Buffered I/O is extremely powerful in go.
Input and output operations model data as streams of bytes that
can be read or written to.

Simply put, the io.Reader is an interface from which you can
read a stream of bytes into a buffer.

![IMAGE - buffered-io.jpg - IMAGE](docs/pics/input-output/buffered-io.jpg)

r is the io.Reader and can handle multiple types of streams,

```go
// r is the io.Reader
buffer := make([]byte, 20)
// THE READER PLACES THE BYTES INTO A BUFFER
n, err := r.Read(buffer)
```

### FROM A BUFFER (strings.Reader)

```go
sourceBuffer := []byte("This data is being put into a buffer")
r := string.NewReader(string(sourceBuffer))
// r strings.Reader put into the reader
```

### FROM A FILE (os.File)

```go
f, err := os.Open("test.txt")
// f os.File put into the reader
```

### FROM A USER (os.Stdin)

```go
// os.Stdin is put into the reader
```

### FROM A PIPE (io.PipeReader)

```go
rpipe, wpipe := io.Pipe()
// rpipe io.PipeReader is put into the reader
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
