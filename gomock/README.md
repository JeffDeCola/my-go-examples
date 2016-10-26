# gomock example

`gomock` _is a example of using gomock on an interface
for unit testing._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## METHODS - CREATURES

* Warewolf - TimeofDay string
  * Kind() - during day is a human, otherwise warewolf.
  * Fly()  - no
  * Sound() - during days says "holle" otherwise "howls".

* Vampire - age int
  * Kind() - older than 100, he's a vampire.
  * Fly()  - yes
  * Sound() - "I want to drink your blood"

## INTERFACE & FUNCTIONS - LABORATORY

```go
type Creatures interface {
    Kind() string
    Fly() bool
    Sound() string
}
```

Where I have two functions `Greet` and `FlyAway` using the interface.

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

```bash
go run helloween.go
```

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