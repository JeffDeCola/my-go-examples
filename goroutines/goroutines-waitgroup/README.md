# goroutines-waitgroup example

`goroutines-waitgroup`  _is an example of
concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WAITGROUP

A WaitGroup allows to wait for a collection of goroutines to finish.

First make a waitgroup,

```go
var wg sync.WaitGroup
```

* wg.Add(1) - Add the goroutines to the wg
* wg.Done() - Call done when goroutine finished
* wg.Wait() - Wait to all the wg are done

![IMAGE - goroutines-waitgroup - IMAGE](../../docs/pics/goroutines-waitgroup.jpg)

## RUN

```go
go run goroutines_waitgroup.go
```
