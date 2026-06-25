# SLICES EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_A flexible dynamically-sized array._

tl;dr

```go
// DECLARE
var a []float64                   // var name []type  (nil slice, no length)

// ASSIGN (append to grow)
a = append(a, 1.1)                // name = append(name, value...)
a = append(a, 2.0)

// DECLARE & INITIALIZE
var b = []float32{1.1, 2.0}       // var name = []type{values...}
c := []float32{3.4, 4.5}          // short form (declare & initialize)

// MAKE (preallocate length & capacity)
d := make([]int, 2, 25)           // name := make([]type, length, capacity)
d[0] = 6                          // indexable up to length
d[1] = 7
```

Most widely used,

```go
// SLICE: var (nil), then append
    var nums []int
    for _, n := range data { nums = append(nums, n) }

// MAP: make, then fill
    counts := make(map[string]int)
    for _, w := range words { counts[w]++ }
```

Examples

* **BUILT-IN**
  * [arrays](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays)
  * [slices](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices)
    **YOU ARE HERE**
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices#run)

Documentation and Reference

* [the go programming language specification - slice types](https://go.dev/ref/spec#Slice_types)
* [a tour of go - slices](https://go.dev/tour/moretypes/7)
* [go by example - slices](https://gobyexample.com/slices)

## OVERVIEW

Slices are,

* A data structure (holds data)
* Dynamically-sized (grows with append)
* A view into an underlying array
* Sequence of elements of a single type
* A list/collection identified by an index
* Zero-based indexes
* The length is NOT part of the type — `[]float32` holds any length
* Slices are references — copying shares the backing array (change one, change both)

Length is how many elements exist; capacity is how many fit before Go grows the
backing array. Slices are used far more than arrays.

Declare then append,

```go
var a []float64               // var name []type  (nil slice, no length)
a = append(a, 1.1)            // name = append(name, value...)
a = append(a, 2.0)
```

Declare and initialize,

```go
var b = []float32{1.1, 2.0}   // var name = []type{values...}
c := []float32{3.4, 4.5}      // short form (declare & initialize)
```

Make (preallocate length & capacity),

```go
d := make([]int, 2, 25)       // name := make([]type, length, capacity)
d[0] = 6                      // index within length
d[1] = 7
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Declare:               []
Assign:                [1.1 2]
Declare & Initialize:  [1.1 2]
Declare & Initialize:  [4.4 5]
Make:                  [0 0]  len=2 cap=25
Make:                  [6 7]  len=2 cap=25
```
