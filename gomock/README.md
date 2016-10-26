# gomock example

`gomock` _is a example of using gomock on an interface
for unit testing._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## METHODS - CREATURES

Warewolf - TimeofDay string
    Kind() - during day is a human, otherwise warewolf.
    Fly()  - no
    Sound() - during days says "holle" otherwise "howls"

Vampire - age int
    Kind() - older than 100, he's a vampire.
    Fly()  - yes
    Sound() - "I want to drink your blood"

## INTERFACE - LABORATORY

```go
type Creatures interface {
	Kind() string
	Fly() bool
	Sound() string
}
```

## RUN TEST

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