# STRUCTS NESTING EMBEDDING EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Nesting and embedding a struct inside another struct, and how field
access differs between them._

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
* [structs-nesting-embedding](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding)
  **YOU ARE HERE**

Table of Contents

* [NESTING](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding#nesting)
* [EMBEDDING](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding#embedding)
* [WHY EMBEDDING OVER NESTING](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding#why-embedding-over-nesting)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/structs/structs-nesting-embedding#run)

Documentation and Reference

* [go spec - struct types](https://go.dev/ref/spec#Struct_types)
* [effective go - embedding](https://go.dev/doc/effective_go#embedding)

## NESTING

A nested struct is a field like any other - it just happens to have a struct
type.

```go
type bigNesting struct {
    b int
    t thing      // named field of type thing
}

n.t.a            // must go through the field name 't'
```

## EMBEDDING

Embedding omits the field name - you write just the type. Its fields are
promoted, so you can reach them directly on the outer struct, as if they
belonged to it.

```go
type bigEmbedding struct {
    b int
    thing        // embedded - no field name
}

e.a              // promoted - direct access
e.thing.a        // long way still works, but not needed
```

## WHY EMBEDDING OVER NESTING

For data, nesting and embedding are functionally the same.
Embedding just saves you typing (e.a instead of e.thing.a).
Its real power is with methods and interfaces, which get promoted too.
See the methods and interfaces examples.

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
PRINT BOTH STRUCTS - Same data, but note the field name
    Nesting:       n     is  {b:10 t:{a:5 s:smile}}
    Embedding:     e     is  {b:10 thing:{a:5 s:smile}}
ACCESS FIELD a - It's simpler with embedding
    Nesting:       n.t.a is  5
    Embedding:     e.a   is  5
```
