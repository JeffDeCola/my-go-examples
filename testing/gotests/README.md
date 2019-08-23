# gotests-complex-function example

`gotests` _is an example of
using the go tool gotests to test a function and a method._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WHAT IS GOTESTS

`gotests` makes writing go tests easy. It's a golang cli that generates
table driven tests based on its target source files' function
and method signatures.

Install,

```bash
go get -u github.com/cweill/gotests/...
```

## UNIT TEST ON A FUNCTION & METHOD

I created a package
[rectangle.go](https://github.com/JeffDeCola/my-go-examples/blob/master/testing/gotests/rectangle/rectangle.go)
that calculates the area of a rectangle
using either a function or a method,

```go
// RectangleAreaFunction area using function
func RectangleAreaFunction(r Rectangle) float64 {
    return r.Width * r.Height
}

// RectangleAreaMethod area using method
func (r Rectangle) RectangleAreaMethod() float64 {
    return r.Width * r.Height
}
```

## RUN

Normal run,

```go
go run gotests.go
```

## GENERATE TEST FILE

The following command will generate a test file with the hooks in place
to perform unit testing,

```bash
cd rectangle
gotests -w -all rectangle.go
```

## RUN YOUR TEST & COVERAGE (0% COVERED)

Now if you run `go test`,

```bash
cd rectangle
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```

You should get 0% covered,

```bash
PASS
coverage: 0.0% of statements
```

## ADD TESTS TO FUNCTION (50% COVERED)

Lets add unit testing to your function.
Open `gotests_test.go` and notice your testing function
`Test_rectangleAreaFunction(t *testing.T)`.  This function wraps around
you function.

Lets simple check that Width 6.0 and Height 3.0 is 18.0.

```go
// TODO: Add some test cases.
        {"testFunction1",
            args{
                r: Rectangle{
                    Width:  6.0,
                    Height: 3.0,
                },
            },
            18.0,
        },
```

Now when you run your test,

```bash
cd rectangle
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Your coverage shall now be,

```bash
PASS
coverage: 50% of statements
```

## ADD TESTS TO METHOD (100% COVERED)

Lets add unit testing to your method,

```go
        // TODO: Add test cases.
        {"testMethod1",
            fields{
                Width:  6.0,
                Height: 3.0,
            },
            18.0,
        },
```

Now when you run your test,

```bash
cd rectangle
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Your coverage shall now be,

```bash
PASS
coverage: 100% of statements
```
