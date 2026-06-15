# STRUCTS BASIC EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Declaring a struct, creating instances and accessing fields._

tl;dr

```go
// BASIC - declare a type, make instances, use the fields
    type thing struct {                 // declare the struct
        a int
        s string
    }

    t1 := thing{a: 4, s: "whatever"}    // field-name literal
    var t2 thing                        // zero values - {a:0 s:}
    t3 := &thing{a: 7, s: "happy"}      // pointer to a struct

    total := t1.a + 5                   // access a field
    t1.s = "smile"                      // modify a field

// NESTING - a struct as a NAMED field, reached through that name
    type bigNesting struct {
        b int
        t thing                         // named field of type thing
    }

    n.t.a                               // must go through field name 't'

// EMBEDDING - a struct as a type with NO field name, fields promoted
    type bigEmbedding struct {
        b int
        thing                           // embedded - no field name
    }

    e.a                                 // promoted - reach thing's field directly
```

Examples

* [structs-basic](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-basic)
  **YOU ARE HERE**
* [structs-nesting-embedding](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-basic#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-basic#run)

Documentation and Reference

* [go spec - struct types](https://go.dev/ref/spec#Struct_types)
* [effective go - structs](https://go.dev/doc/effective_go#composite_literals)

## OVERVIEW

A struct groups related fields under one type. You declare the type once,
then create instances of it several ways.

```go
t1 := thing{a: 4, s: "whatever"}   // field-name literal
var t2 thing                       // zero value - each field gets its type's zero
t3 := &thing{a: 7, s: "happy"}     // pointer to a struct
```

Fields are accessed and modified with dot notation.

```go
total := t1.a + 5   // access
t1.s = "smile"      // modify
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
PRINT THE STRUCTS
    t1 is {a:4 s:whatever}
    t2 is {a:0 s:}
    t3 is &{a:7 s:happy}
ACCESS A FIELD
    t1.a 4 plus 5 is 9
MODIFY A FIELD
    t1 modified is {a:4 s:smile}
```
