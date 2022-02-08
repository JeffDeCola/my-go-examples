# goroutines-worker-pools

_Concurrency using worker pools._

Other examples using,

* [goroutines-multi-core](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-multi-core)
* [goroutines-waitgroup](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-waitgroup)
* [goroutines-worker-pools](https://github.com/JeffDeCola/my-go-examples/tree/master/goroutines/goroutines-worker-pools)
  **<- You are here**

Table of Contents,

* _tbd._

Documentation and references,

* This repos [github webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

Worker pools are goroutines that do work.
Usually they will pull from a channel buffer to get data on what to do.

### TICK TIME

In this example, doSomething is using tick time.
Meaning, asking the workers to do a bunch of stuff in period of time
(a tick).
In this example the default is to have the workers complete
5 jobs in 10 seconds.
And we only have 2 workers that take 7 seconds to complete its task.
So you can see this will break.

So you will have to fix this by doing one of the following,

* Increase the number of workers
* Reduce the number of jobs per tick
* Speed up the workers

This illustration  may help,

![IMAGE - channels-buffered.jpg - IMAGE](../../../../docs/pics/in-process-communication/channels-buffered.jpg)

## RUN

Run,

```bash
go run goroutines-worker-pools.go
```

**Simply press return to stop.**

These are the defaults, play around with them.  Right now it will break,
so you need to adjust this.

```go
// SET CONSTANTS FOR WORKER doWork()
const numberWorkers = 2                                             // How many workers you want
const workerTime = 7                                                // How long it takes a worker to work

// SET CONSTANTS FOR doSomething()
var jobList = []string{"jobA", "jobB", "jobC", "jobD", "jobE"}      // 5 jobs with jobNames
var ticktimeSeconds = 10                                            // Tick time to send a bunch of jobs workers

// OTHER
var channelBufferSize = 1                                           // How many channel buffers
```

## TEST

To create _test files,

```bash
gotests -w -all goroutines-worker-pools.go
```

To unit test the code,

```bash
go test -cover ./... 
```
