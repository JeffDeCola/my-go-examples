# ARRAYS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A fixed-length collection of elements of the same type._

tl;dr

```go
// DECLARE
var a [2]float32                  // var name [length]type

// ASSIGN
a[0] = 1.1                        // name[index] = value
a[1] = 2.0

// DECLARE & INITIALIZE
var b = [2]float32{1.1, 2.0}      // var name = [length]type{values...}
c := [2]float32{4.4, 5.0}         // short form (declare & initialize)
```

Examples

* **BUILT-IN**
  * [arrays](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays)
    **YOU ARE HERE**
  * [slices](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices)
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays#run)

Documentation and Reference

* [the go programming language specification - array types](https://go.dev/ref/spec#Array_types)
* [a tour of go - arrays](https://go.dev/tour/moretypes/6)
* [go by example - arrays](https://gobyexample.com/arrays)

## OVERVIEW

Arrays are,

* A data structure (holds data)
* Do not change in size (not dynamic)
* Sequence of elements of a single type
* A list/collection identified by an index
* Zero-based indexes
* The length is part of the type — `[2]float32` and `[3]float32` are different types
* Arrays are values — assigning or passing one copies every element

Arrays are really not used that much.

Declare then assign,

```go
var a [2]float32              // var name [length]type
a[0] = 1.1                    // name[index] = value
a[1] = 2.0
```

Declare and initialize,

```go
var b = [2]float32{1.1, 2.0}  // var name = [length]type{values...}
c := [2]float32{4.4, 5.0}     // short form (declare & initialize)
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Declare:               [0 0]
Assign:                [1.1 2]
Declare & Initialize:  [1.1 2]
Declare & Initialize:  [4.4 5]
```
