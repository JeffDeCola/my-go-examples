# errors (github.com/pkg/errors) example

`errors` _non-standard package. An example of
error handling using non-standard `github.com/pkg/errors` package._

Refer to the
[github.com/pkg/errors](https://github.com/pkg/errors)
package for more info.

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## HANDLE ERRORS

Go doesn’t have exceptions, so it doesn’t have try, catch or anything similar.
So how do we handle errors? With Multiple return values.

## MULTIPLE RETURN VALUES

As an example,

```go
func a(x int) (int, error) {
    if x == 42 {
        // Make your error
        return -1, errors.New("can't work with 42")
    }
    return x + 3, nil
}

func main() {
    //MULTIPLE RETURN VALUES
    r, err := b(42)
    if err != nil {
        // Handle the error
        fmt.Printf("Error is %+v\n", err)
        log.Fatal("ERROR:", err)
    }
    // Continue
    fmt.Println("Returned", r)
}
```

The beauty of the `github.com/pkg/errors` package is that it
traces the error back to the source. That's why we use
`fmt.Printf("Error is %+v\n", err)`.

Since you check errors a lot, make an error checker,

```go
func checkErr(err error) {
    if err != nil {
        fmt.Printf("Error is %+v\n", err)
        log.Fatal("ERROR:", err)
    }
}

func main() {
    //MULTIPLE RETURN VALUES (Simpler form)
    r, err = b(42)
    checkErr(err)
    fmt.Println("Returned", r)
}
```

## RUN

```go
go run errors.go
```
