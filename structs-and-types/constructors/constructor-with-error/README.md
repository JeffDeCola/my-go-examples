# CONSTRUCTOR WITH ERROR EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Expanding on constructor-simple to add a configuration struct
and error handling._

Examples

* [constructor-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-simple)
* [constructor-with-error](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-with-error)
  **YOU ARE HERE**

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-with-error#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-with-error#run)

Documentation and Reference

* [errors](https://pkg.go.dev/errors)
  standard package
* [working with errors in go 1.13](https://go.dev/blog/go1.13-errors)
* [effective go - constructors](https://go.dev/doc/effective_go#composite_literals)

## OVERVIEW

Our example expands on the constructor-simple example to
have the following structure,

```go
type config struct {...}                      // the configuration
type thing struct {...}                       // the blueprint - what a thing IS
func newThing(c config) (*thing, error) {...} // the factory  - BUILD it
cfg := config{...}                            // the user   - fills out config
t, err := newThing(cfg)                       // the caller - "I want a thing"
// if err != nil { handle and return }        // the caller - error handled ONCE
```

Note the three ways an error gets made,

```go
errors.New("s must not be empty")                    // static string
fmt.Errorf("a must not be negative, got %d", c.a)    // has a %verb
fmt.Errorf("hostname failed: %w", err)               // wrapping
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
a is 4, s is "happy", hostname is "TK3-PC"
a is 5, s is "monkey", hostname is "TK3-PC"
newThing failed: a must not be negative, got -3
exit status 1
```
