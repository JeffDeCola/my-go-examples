# MAPS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_An unordered collection of key-value pairs (a hash table)._

tl;dr

```go
// DECLARE (nil map - cannot assign yet)
var a map[string]int                  // var name map[keytype]valuetype

// MAKE
a = make(map[string]int)              // name = make(map[keytype]valuetype)

// ASSIGN
a["go"] = 2009                        // name[key] = value
a["rust"] = 2010

// DECLARE & INITIALIZE
var b = map[string]int{"go": 2009}    // var name = map[k]v{key: value...}
c := map[string]int{"c": 1972}        // short form (declare & initialize)

// READ (comma-ok)
v, ok := a["go"]                      // value, present? (ok=false if absent)

// DELETE
delete(a, "rust")                     // delete(name, key)
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
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
    **YOU ARE HERE**
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps#run)

Documentation and Reference

* [the go programming language specification - map types](https://go.dev/ref/spec#Map_types)
* [a tour of go - maps](https://go.dev/tour/moretypes/19)
* [go by example - maps](https://gobyexample.com/maps)

## OVERVIEW

Maps are,

* A data structure (holds data)
* Unordered key-value pairs (a hash table)
* Keys are unique; each key maps to one value
* Indexed by key, not position — `m["go"]`, not `m[0]`
* Dynamically-sized (grows as you add keys)
* A nil map can be read but NOT written — `make`
  (or a literal) before assigning, or it panics
* Reading a missing key returns the zero value.
  Use `v, ok := m[key]` to tell absent from zero
* Range order is random — never rely on iteration order
* Maps are references — like slices, copying shares the same data

Maps are used heavily in Go. (`fmt` prints them in sorted key order,
so output is stable even though iteration is not.)

Declare then make,

```go
var a map[string]int          // nil map - not really needed
a = make(map[string]int)      // initialize before assigning
```

Assign,

```go
a["go"] = 2009                // name[key] = value
a["rust"] = 2010
```

Declare and initialize,

```go
var b = map[string]int{"go": 2009, "rust": 2010}   // var name = map[k]v{key: value...}
c := map[string]int{"c": 1972, "python": 1991}     // short form (declare & initialize)
```

Read (comma-ok),

```go
v, ok := a["go"]              // 2009, true
w, ok := a["java"]            // 0, false  (absent key -> zero value)
```

Delete,

```go
delete(a, "rust")            // delete(name, key)
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Declare:               map[] true
Make:                  map[] false
Assign:                map[go:2009 rust:2010]
Declare & Initialize:  map[go:2009 rust:2010]
Declare & Initialize:  map[c:1972 python:1991]
Read go:               2009 true
Read java:             0 false
Delete rust:           map[go:2009]
```
