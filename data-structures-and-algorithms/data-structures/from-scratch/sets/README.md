# SETS EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_- A collection of unique elements built from a map._

Examples

* **BUILT-IN**
  * [arrays](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays)
  * [slices](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices)
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
    **YOU ARE HERE**
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets#run)

Documentation and Reference

* [the go programming language specification - struct types](https://go.dev/ref/spec#Struct_types)
* [the empty struct - dave cheney](https://dave.cheney.net/2014/03/25/the-empty-struct)

## OVERVIEW

A set is a collection of **unique** things.
Adding a duplicate has no effect.
For example, a guest list RSVP. Add the same name twice, there's still one guest.

Go has no set type. So this example is about building one from a
primitive you already have, a map.

Sets are,

* A collection of unique elements
* Built from a map — the keys ARE the set
* Membership-tested, not ordered (range order is random)
* Adding a duplicate has no effect (keys are unique)
* The value is unused — store `struct{}{}`, which takes zero bytes

For a set, we want to be able to

* ADD an element (Guest RSVPs)
* CHECK if element (Guest already is on list)
* DELETE an element (Guest Cancels)

The structure looks like,

```go
rsvp := make(map[string]struct{})

// ADD (RSVP)
// name[key] = struct{}{}  (value stores nothing) Will not take up memory
rsvp["Jeff"] = struct{}{}
rsvp["Larry"] = struct{}{}
rsvp["Sam"] = struct{}{}
rsvp["Jeff"] = struct{}{} // duplicate

// READ - normal map use
_, ok := rsvp["Larry"] // is Amy on the list?
_, ok = rsvp["Bob"] // bob never RSVP'd

// DELETE - normal map use
delete(rsvp, "Sam") // delete(name, key)
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Make map:              map[]
RSVP (Jeff twice):     map[Jeff:{} Larry:{} Sam:{}] - len 3
Is Larry coming:       true
Is Bob coming:         false
DELETE Sam:            map[Jeff:{} Larry:{}] - len 2
```
