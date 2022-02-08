# goroutines-waitgroup

_Concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

Other examples using,

* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)
  **<- You are here**
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

Table of Contents,

* _tbd._

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

A WaitGroup allows to wait for a collection of goroutines to finish.

First make a waitgroup,

```go
var wg sync.WaitGroup
```

* wg.Add(1) - Add the goroutines to the waitgroup
* wg.Done() - Call done when goroutine finished
* wg.Wait() - Wait for all the the waitgroups to finish

This illustration may help,

![IMAGE - goroutines-waitgroup.jpg - IMAGE](../../../../docs/pics/goroutines/goroutines-waitgroup.jpg)

## RUN

Run,

```bash
go run goroutines-waitgroup.go
```

## TEST

To create _test files,

```bash
gotests -w -all goroutines-waitgroup.go
```

To unit test the code,

```bash
go test -cover ./... 
```
