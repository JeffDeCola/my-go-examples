# CONSTRUCTOR SIMPLE EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Constructors are functions that build and return a new instance of a struct,
often with defaults._

Examples

* [constructor-simple](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-simple)
  **YOU ARE HERE**
* [constructor-with-error](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-with-error)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-simple#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/constructors/constructor-simple#run)

Documentation and Reference

* [go spec - composite literals](https://go.dev/ref/spec#Composite_literals)
* [effective go - constructors](https://go.dev/doc/effective_go#composite_literals)

## OVERVIEW

A constructor is basically "I want a thing".

```go
t := newThing(4)
```

You want a thing and you don't want to worry about setting it up correctly.
The constructor owns the setup so the caller can't get it wrong.
Defaults are filled in, anything computed gets computed, and
(in the with-error variant) bad inputs get rejected at the door.
The caller's side stays that one clean line.

Our example follows this structure,

```go
type thing struct {...}            // the blueprint - what a thing IS
func newThing(a int) *thing {...}  // the factory  - how a thing gets BUILT
t := newThing(4)                   // the caller   - "I want a thing"
```

The constructor builds in three steps - BUILD the values,
ASSEMBLE the struct, and RETURNS a pointer to it.

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
The thing is {a:4 s:default}
The thing is {a:5 s:happy}
```
