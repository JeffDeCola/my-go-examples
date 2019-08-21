# recursion example

`recursion` _is a example of
a function calling itself to make a fibonacci series._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## FIBONACCI SERIES

Each Fibonacci number is the sum of the two preceding numbers.

Hence, the simplest series is 1,1,2,3,5,8, etc.

### METHOD 1

For this example given n, complete the fibonacci function so it returns
`fibonacci(n)` where,

```go
fibN := fibonacci(n - 2) + fibonacci(n - 1)
```

The problem is the it always calculates backwards and does the same
thing over and over.

### METHOD 2

This is more of a bottom up approach that I like better.
It  more flows what I would do with a pen and piece of paper.

## RUN

```bash
go run recursion1.go
go run recursion2.go
```
