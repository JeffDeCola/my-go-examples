# GENERICS TYPE EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A generic type uses a type parameter so one definition serves every element
type, instead of writing one type per element._

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
* [generics-type](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-type)
  **YOU ARE HERE**

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-type#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/structs-and-types/generics/generics-type#run)

Documentation and Reference

* [go spec - type parameters](https://go.dev/ref/spec#Type_parameter_declarations)
* [tutorial - getting started with generics](https://go.dev/doc/tutorial/generics)

## OVERVIEW

Without generics, a container that holds different element types has to be
written once per type - same logic, different element.

```go
type intStack struct    { items []int }
type stringStack struct { items []string }
```

A generic type replaces them all with one definition, using a _type
parameter_ `T`.

```go
type stack[T any] struct {
    items []T
}
```

You pick the concrete type when you create an instance, and the compiler
enforces it from there.

```go
var si stack[int]      // holds ints
var ss stack[string]   // holds strings - same type, different T
```

Methods on a generic type repeat the type parameter on the receiver.

```go
func (s *stack[T]) push(item T) {...}
func (s *stack[T]) pop() T      {...}
```

You never _need_ generics - you could write a type per element, or use a
slice of `any`. A generic type just lets you do it once, type-safely, so a
`stack[int]` can never accidentally hold a string.

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```text
STACK OF INTS
    stack is  [10 20]
    pop       20
    stack is  [10]
STACK OF STRINGS
    stack is  [a b]
    pop       b
    stack is  [a]
```
