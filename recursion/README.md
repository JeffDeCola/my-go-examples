# recursion example

`recursion` _is a example of a function calling itself.  A fibonacci series is a perfect example._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## FIBONACCI SERIES

Each number ( Fibonacci number ) is the sum of the two preceding numbers.

Hence, the simplest series is 1,1,2,3,5,8, etc.

for this example given n, complete the fibonacci function so it returns `fibonacci(n)` where

`fibonacci(n) = fibonacci(n-1) + fibonacci(n-2)`

## RUN

```bash
go run recursion.go ??????????????
```

## RUN TEST

```bash
go test -v -cover ./...
```

`recursion_test.go` file was created by running:

```bash
gotests -w -all recursion.go
```