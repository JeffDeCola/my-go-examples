# goroutines-waitgroup example

`goroutines-waitgroup` _is an example of
concurrency using a waitgroup (waiting for a collection of goroutines to finish)._

These are my 5 main example of using goroutines,

* [goroutines-async-channel-receive-no-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-receive-no-waiting)
* [goroutines-async-channel-send-receive-waiting](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-async-channel-send-receive-waiting)
* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
* **goroutines-waitgroup** <- You are here
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## WAITGROUP

A WaitGroup allows to wait for a collection of goroutines to finish.

First make a waitgroup,

```go
var wg sync.WaitGroup
```

* wg.Add(1) - Add the goroutines to the waitgroup
* wg.Done() - Call done when goroutine finished
* wg.Wait() - Wait for all the the waitgroups are done

![IMAGE - goroutines-waitgroup - IMAGE](../../docs/pics/goroutines-waitgroup.jpg)

## RUN

```bash
go run goroutines_waitgroup.go
```
