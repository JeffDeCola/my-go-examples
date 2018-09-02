# logging-error-handling example

`logging-error-handling` _is an example of logging and error handling._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## LOGGING

Log package `import "log"`

Setup,

`var logger = log.New(os.Stderr, "JeffLogging:", log.Lshortfile)`

Use

`logger.Println("Print Logging info")`

You could add levels to logger, but that is not shown in this example.

## ERROR HANDLING

Go has a built in error interface type.

### Example of error

```go
data, err := ioutil.ReadFile("readthis.txt")
if err != nil {
    panic(err)
}
```

A nil value means there was no error.

### Make your own errors with `errors.New`

```go
var err error
if sum > 4 {
    err = errors.New("This is my error")
    return sum, err
}
```

Print err,

`logger.Fatalf("Error: %s", err)`

## RUN

```bash
go run logging-error-handling.go
```
