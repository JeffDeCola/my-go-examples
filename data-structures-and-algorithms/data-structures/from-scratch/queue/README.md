# QUEUE EXAMPLE

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_- A FIFO (first-in first-out) built from a slice._

Examples

* **BUILT-IN**
  * [arrays](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/arrays)
  * [slices](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/slices)
  * [maps](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/built-in/maps)
* **FROM SCRATCH**
  * [sets](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/sets)
  * [stack](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/stack)
  * [queue](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue)
    **YOU ARE HERE**
  * [linked-list](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/linked-list)
  * [binary-tree](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/binary-tree)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/data-structures-and-algorithms/data-structures/from-scratch/queue#run)

Documentation and Reference

* [the go programming language specification - slice types](https://go.dev/ref/spec#Slice_types)
* [package container/list](https://pkg.go.dev/container/list)

## OVERVIEW

A queue is FIFO (first in, first out).

Go has no queue type, so we build one from a slice.
enqueue adds to the tail, dequeue removes from the head.

A queue is,

* A FIFO collection — first in, first out
* Built from a slice — tail is the back, head is the front
* Enqueue adds to the tail (`append`) — same as a stack's push
* Dequeue reads the head, then chops it off (`q.items[1:]`)
* Dequeuing an empty queue must be guarded, or the read panics

Stack vs queue is one line. Both append to the tail; the difference is which
end you remove from:

* stack (LIFO) — remove the tail: `s.items[:len(s.items)-1]`
* queue (FIFO) — remove the head: `q.items[1:]`

Enqueue,

```go
func (q *queue) enqueue(i int) {
    q.items = append(q.items, i)   // add to the tail
}
```

Dequeue (read the head, then remove it),

```go
func (q *queue) dequeue() (int, bool) {
    if len(q.items) == 0 {
       return 0, false        // empty - nothing to dequeue
    }
    front := q.items[0]    // read head
    q.items = q.items[1:]  // chop head
    return front, true
}
```

## RUN

To run,

```bash
go run main.go
```

The output should look like,

```bash
Queue:     [1 2 3]
Dequeue:   1
Dequeue:   2
Dequeue:   3
Dequeue empty: false
```
