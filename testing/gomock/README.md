# gomock example

`gomock` _is a example of
using gomock on an interface and using the go tool `gotests` for unit testing._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WHAT IS GOMOCK

GoMock is a mock framework for Go.
**Its useful for mocking an interface**
It enjoys a somewhat official status as
part of the github.com/golang organization.
`mockgen` is the code generation tool.

Install,

```bash
go get -u -v github.com/golang/mock/gomock
go get -u -v github.com/golang/mock/mockgen
```

Verify you have it,

```bash
$GOPATH/bin/mockgen
```

## HOW TO USE GOMOCK

Four basic steps:

* Use mockgen to generate a mock for the interface you wish to mock
* In your test, create an instance of gomock.Controller and pass it
  to your mock object’s constructor to obtain a mock object
* Call EXPECT() on your mocks to set up their expectations and return values
* Call Finish() on the mock controller to assert the mock’s expectations

More on this in a bit. First lets look at the packages.

## CREATURES PACKAGE

Lets describe a werewolf and a vampire in our creatures package
`/creatures/creatures.go`,

```go
type Werewolf struct {
    TimeofDay string
}
type Vampire struct {
    Age int
}
```

Where werewolf has the following methods,

* Kind() - during day is a human, otherwise werewolf.
* Fly()  - no
* Sound() - during days says "hell0" otherwise "howls".

And vampire has the following methods,

* Kind() - older than 100, he's a vampire.
* Fly()  - yes
* Sound() - "I want to drink your blood"

## LABORATORY PACKAGE

I created another package laboratory where I keep my creatures.

```go
type Creatures interface {
    Kind() string
    Fly() bool
    Sound() string
}
```

Where I have two functions `Greet` and `FlyAway` using the interface.

## RUN

```bash
go run helloween.go
```

## GENERATE TEST FILE USING GOTESTS

The following command will generate a test file with the hooks
in place to perform unit testing,

```bash
cd laboratory
gotests -w -all .
cd creatures
gotests -w -all .
```

## RUN YOUR TEST & COVERAGE (0% COVERED)

Now if you run go test,

```bash
cd ..
go test -v -cover laboratory/*
go test -v -cover creatures/*
```

Optional commands,

```bash
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```

You should get 0% covered,

```bash
PASS
coverage: 0.0% of statements
```

## ADD TESTS TO CREATURES PACKAGE (100% COVERED)

There are 6 methods in the creatures package.

As an example lets add tests to the function `TestWerewolf_Kind(t *testing.T)`,

```go
        // TODO: Add test cases.
        {
            "Test1",
            fields{
                TimeofDay: "day",
            },
            "human",
        },
        {
            "Test2",
            fields{
                TimeofDay: "night",
            },
            "werewolf",
        },
```

Do this for all six methods in the creatures package. Now when you run your test,

```bash
cd creatures
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```

Your coverage shall now be,

```bash
PASS
coverage: 100% of statements
```

## ADD TESTS TO LABORATORY PACKAGE (100% COVERED)

We have two functions to test; `Greet()` and `FlyAway()` that each use an interface.
So lets mock that interface.

Using mockgen, we create a file `mockcreature.go`,

```bash
cd laboratory
mockgen -source=laboratory.go -package=laboratory -destination=mockcreature.go
```

To use add the following lines to the top of your test function
`func TestGreet(t *testing.T) {`

```go
func TestGreet(t *testing.T) {
    var ctrl = gomock.NewController(t)
    defer ctrl.Finish()
    var mockcreature = NewMockCreatures(ctrl)
    mockcreature.EXPECT().Sound().Times(1).Return("poo")
```

Where you test is,

```go
        // TODO: Add test cases.
        {
            "Test1",
            args{
                c: mockcreature,
            },
            "poo",
        },
```

```bash
cd laboratory
go test -v -cover .
go test -coverprofile coverage.out
go tool cover -html=coverage.out -o coverage.html
```
