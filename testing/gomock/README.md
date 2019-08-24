# gomock example

`gomock` _is a example of
using gomock on an interface for unit testing._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WHAT IS GOMOCK

GoMock is a mock framework for Go. It enjoys a somewhat official status as
part of the github.com/golang organization. `mockgen` is the code generation tool.

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

I'll keep my creature int he laboratory. Lets create an interface
that describes creatures,

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

## GREET FUNCTION

`func Greet(c Creatures) string {`

Hence the test would be setup as

```go
func TestGreet(t *testing.T) {
    var ctrl = gomock.NewController(t)
    defer ctrl.Finish()
    var mockcreature = NewMockCreatures(ctrl)
    mockcreature.EXPECT().Sound().Times(1).Return("poo")
    type args struct {
        c Creatures
    }
    tests := []struct {
        name string
        args args
        want string
    }{
```

## FLYAWAY FUNCTION

`func FlyAway(canfly Creatures) string {`

Hence the test would be setup as,

```go
func TestFlyAway(t *testing.T) {
    var ctrl = gomock.NewController(t)
    defer ctrl.Finish()
    var mockcreature = NewMockCreatures(ctrl)
    gomock.InOrder(
        mockcreature.EXPECT().Fly().Times(1).Return(true),
        mockcreature.EXPECT().Fly().Times(1).Return(false),
    )
    mockcreature.EXPECT().Kind().AnyTimes().Return("Pig")
    type args struct {
        canfly Creatures
    }
    tests := []struct {
        name string
        args args
        want string
    }{
```

## RUN


## RUN TEST

```bash
go test -v -cover ./...
```

## MOCKGEN

`mockgen` file was created by running:

```bash
mockgen -source=laboratory.go -package=laboratory -destination=mockcreature.go
```

`laboratory_test.go` file was created by running:

```bash
gotests -w -all laboratory.go
```