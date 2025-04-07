# GOROUTINES WAITGROUP

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_Concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

Other Examples

* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-multi-core)
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-waitgroup)
  **<- You are here**
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-worker-pools)

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-waitgroup#overview)
* [RUN](https://github.com/JeffDeCola/my-go-examples/tree/master/go-runtime/goroutines/goroutines-waitgroup#run)

## OVERVIEW

Go is written for **concurrency**. The go runtime schedules goroutines on threads.
The OS schedules these threads on cpus/cores.

A WaitGroup waits for a collection of goroutines to finish.

First make a waitgroup,

```go
var wg sync.WaitGroup
```

Then,

* **wg.Add(1)** - Add the goroutines to the waitgroup
* **wg.Done()** - Call done when goroutine finished
* **wg.Wait()** - Wait for all the the waitgroups to finish

This illustration may help,

![IMAGE - goroutines-waitgroup.svg - IMAGE](../../../docs/pics/goroutines/goroutines-waitgroup.svg)

## RUN

Run,

```bash
go run main.go
```
