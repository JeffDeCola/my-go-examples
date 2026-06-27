# STACK EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_- A LIFO (last-in first-out) built from a slice._

Examples

* **BUILT-IN**
  * [arrays](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays)
  * [slices](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices)
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
    **YOU ARE HERE**
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack#run)

Documentation and Reference

* [the go programming language specification - slice types](https://go.dev/ref/spec#Slice_types)
* [go wiki - slicetricks (stack)](https://go.dev/wiki/SliceTricks#stack)

## OVERVIEW

A stack is LIFO (last in, first out).

Go has no stack type, so we build one from a slice. Push and pop both act on
the tail (the end of the slice), which is what makes it last-in, first-out.

A stack is,

* A LIFO collection — last in, first out
* Built from a slice — the tail is the top of the stack
* Push adds to the tail (`append`)
* Pop reads the tail, then chops it off (re-slice)
* Popping an empty stack must be guarded, or the re-slice panics

Stack vs queue is one line. Both append to the tail; the difference is which
end you remove from:

* stack (LIFO) — remove the tail: `s.items[:len(s.items)-1]`
* queue (FIFO) — remove the head: `q.items[1:]`

Push,

```go
func (s *stack) push(x int) {
    s.items = append(s.items, x)
}
```

Pop (read the tail, then remove it),

```go
func (s *stack) pop() (int, bool) {
    if len(s.items) == 0 {
        return 0, false            // empty - nothing to pop
    }
    top := s.items[len(s.items)-1]     // read tail
    s.items = s.items[:len(s.items)-1] // chop tail
    return top, true
}
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Stack: [1 2 3]
Pop:   3
Pop:   2
Pop:   1
Pop empty: false
```
