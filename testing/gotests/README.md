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

## THE FUNCTION & METHOD

They will simple just calculate the area of a rectangle.

```go
// Rectangle area using function
func rectangleAreaFunction(r Rectangle) float64 {
    return r.width * r.height
}

// Rectangle area using method
func (r Rectangle) rectangleAreaMethod() float64 {
    return r.width * r.height
}
```

## RUN

Normal run is

```go
go run gotests.go
```

## GENERATE TEST FILE & RUN YOUR TEST (0% COVERED)

This will generate a test file with the hooks in place to perform unit testing,

```bash
cd rectangle
gotests -w -all rectangle.go
```

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

## ADD TETS TO FUNCTION (50% COVERED)

Lets add unit testing to your function.
Open `gotests_test.go` and notice your testing function
`Test_rectangleAreaFunction(t *testing.T)`.  This function wraps around
you function.

Lets simple check that width 6 and height 2 is 12.

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
coverage: 12.5% of statements
```

## ADD TESTS TO METHOD (100% COVERED)

Lets add unit testing to your method,

```go
        // TODO: Add test cases.
        {"testMethod1",
            fields{
                width:  6.0,
                height: 3.0,
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
coverage: 25% of statements
```
