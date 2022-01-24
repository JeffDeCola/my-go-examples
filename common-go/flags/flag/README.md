# flag

_The `flag` standard package makes it easy to implement command-line flag parsing._

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag#run)
* [TEST](https://github.com/JeffDeCola/my-go-examples/tree/master/common-go/flags/flag#test)

Documentation and reference,

* Refer to the
  [flag](https://golang.org/pkg/flag/)
  package for more info
* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

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

## RUN

For integers and strings, the following formats are permitted,

* -flag x
* -flag=x

String,

```bash
go run flag.go -s "happy days"
go run flag.go -s="happy days"
```

Integer,

```bash
go run flag.go -i 56
go run flag.go -i=56
```

Boolean,

```bash
go run flag.go -b
```

## TEST

To create _test files,

```bash
gotests -w -all flag.go
```

To unit test the code,

```bash
go test -cover ./... 
```
