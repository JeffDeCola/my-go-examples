# callback example

`callback` _is an example of
passing a function (as an argument) to a function._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## CALLBACK

Callback simply means call back the function
you passed in. Calling back home.

In this example, you will be
passing the following anonymous functions to
the sum() function,

```go
// Giving an int, determine if even
anonFunc1 := func(n int) bool {
    if n%2 == 0 {
        return true
    }
    return false
}
// Giving an int, determine if odd
anonFunc2 := func(n int) bool {
    if n%2 != 0 {
        return true
    }
    return false
}
```

It just tests if the number is even or odd.

## RUN

```go
go run callback.go
```
