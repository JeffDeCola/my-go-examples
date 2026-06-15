# GENERICS FUNCTION EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A generic function uses a type parameter so a and b are the same type,
and so is the result - one function for many types instead of copy-pasting._

tl;dr

```go
// FUNCTION - one function, many types
    type Number interface {       // a constraint - which types are allowed
        int | float64
    }

    func max[T Number](a T, b T) T {...}  // T is the type parameter

    biggerNumber(3, 5)                     // T inferred as int
    biggerNumber(2.5, 1.5)                 // T inferred as float64

  // Could have done this inline without type interface
  func max[T int | float64](a T, b T) T { ... }

// TYPE - one type definition, many element types
    type stack[T any] struct {    // any = the loosest constraint
        items []T
    }

    var si stack[int]             // a stack of ints
    var ss stack[string]          // the SAME type, holding strings
```

Examples

* [generics-function](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-function)
  **YOU ARE HERE**
* [generics-type](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-type)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-function#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-function#run)

Documentation and Reference

* [go spec - type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
* [tutorial - getting started with generics](https://go.dev/doc/tutorial/generics)

## OVERVIEW

Without generics, a function only works for the type it was written for - so
the same logic gets copy-pasted once per type.

```go
func maxInt(a, b int) int           {...}
func maxFloat(a, b float64) float64 {...}
func maxString(a, b string) string  {...}
```

Same logic, three times. A generic function replaces them all with one, using
a _type parameter_ `T`.

```go
func max[T Number](a T, b T) T {...}

biggerNumber(3, 5)       // T is int
biggerNumber(2.5, 1.5)   // T is float64
```

The three pieces:

* the _type parameter_ `[T ...]` - a placeholder for the type
* the _constraint_ `Number` - which types T is allowed to be (always an interface)
* `any` - the loosest constraint, meaning any type at all

The caller doesn't spell out T - Go infers it from the arguments. The inputs
and the result all share that one type.

You never _need_ generics - you could write a function per type, or use `any`
with type assertions. Generics just let you do it once, type-safely, and even
for types you haven't seen yet.

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```text
MAX WITH DIFFERENT TYPES
    biggerNumber(3, 5)       is  5
    biggerNumber(2.5, 1.5)   is  2.5
    biggerNumber[int](7, 2)  is  7
```
