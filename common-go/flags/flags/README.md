# FLAGS EXAMPLE

_The
[flag](https://pkg.go.dev/flag)
 standard package makes it easy to implement command-line flag parsing._

tl;dr

```go
// STRING
stringPtr := flag.String("s", "Cloudy Days", "This is the flag for a string")
// INTEGER
integerPtr := flag.Int("i", 1, "This is the flag for an integer")
// BOOLEAN
boolPtr := flag.Bool("b", false, "This is the flag for a boolean")
// Parse the flags
flag.Parse()
```

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags#overview)
* [PREREQUISITES](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags#prerequisites)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flags#test)

Documentation and Reference

* [flag](https://golang.org/pkg/flag/)
  standard package

## OVERVIEW

We will go over the 3 basic flags,

```go
func main() {

    // STRING
    stringPtr := flag.String("s", "Cloudy Days", "This is the flag for a string")

    // INTEGER
    integerPtr := flag.Int("i", 1, "This is the flag for an integer")

    // BOOLEAN
    boolPtr := flag.Bool("b", false, "This is the flag for a boolean")

    // Parse the flags
    flag.Parse()

    // Print out the flags
    fmt.Printf("\n")
    fmt.Printf("String flag is %s\n", *stringPtr)
    fmt.Printf("Integer flag is %d\n", *integerPtr)
    fmt.Printf("Integer flag is %t\n\n", *boolPtr)

}
```

## PREREQUISITES

You will need the following go packages,

```bash
go install github.com/cweill/gotests/...@latest
```

## RUN

For integers and strings, the following formats are permitted,

* -flag x
* -flag=x

String,

```bash
go run main.go -s "happy days"
go run main.go -s="happy days"
```

Integer,

```bash
go run main.go -i 56
go run main.go -i=56
```

Boolean,

```bash
go run main.go -b
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
