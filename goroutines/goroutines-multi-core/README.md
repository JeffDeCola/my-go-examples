# goroutines-multi-core example

`goroutines-multi-core`  _is an example of
concurrency across multi-cores._

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## SETUP

This program will allow you to set the numbers of workers per core and check the
statistics of each process and core.

```go
// GO RUNTIME
const numberCores = 5                       // Number of CPU you want to use
const lockThread = true                     // locked the goroutine to a thread (Done in go runtime)
const lockCore = true                       // locked the thread to a core (Done in C)

// WORKERS
const useGoroutine = true                   // Do you want to use goroutines
const numberWorkers = 3                     // Number of workers
const timeWork = 5                          // Amount of time it takes a worker to finish

// BUFFER CHANNEL
var channelBufferSize = numberWorkers + 1   // How many channel buffers
```

This diagram will help explain what we are trying to do,

![IMAGE - goroutines-multi-core - IMAGE](../../docs/pics/goroutines-multi-core.jpg)

## SOME THINGS TO NOTE

The `go runtime` will schedule goroutines to cores and threads.  And this
can change a lot.  No goroutine is locked to a thread or a core.

### LOCK GOROUTINE TO A THREAD

```go
runtime.LockOSThread()
defer runtime.UnlockOSThread()
```

### LOCK THREAD TO A CPU/CORE

This is outside go and uses c.  So since we now have a locked thread,
lets lock that thread to a CPU.

```go
// Get the cpu your are on
cpuID:= C.sched_getcpu()
// lock the thread to a cpu
C.lock_thread(C.int(cpuID))
```

### GOMAXPROCS

The GOMAXPROCS variable limits the number of threads that can
execute user-level Go code simultaneously.

### PIN A GOROUTINE TO A CPU

This is tricky and needs a bit of C code.



## RUN

```go
go run goroutines-multi-core.go
```
